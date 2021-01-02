package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"

	"github.com/proxypoke/i3ipc"
	"github.com/vbrown608/xdo-go"
)

type Directions struct {
	i3Focus         string
	tmuxSelectPane  string
	tmuxBorderCheck string
}

func main() {
	directions := map[string]Directions{
		"h": {i3Focus: "left", tmuxSelectPane: "L", tmuxBorderCheck: "left"},
		"j": {i3Focus: "down", tmuxSelectPane: "D", tmuxBorderCheck: "bottom"},
		"k": {i3Focus: "up", tmuxSelectPane: "U", tmuxBorderCheck: "top"},
		"l": {i3Focus: "right", tmuxSelectPane: "R", tmuxBorderCheck: "right"},
	}

	dir := string(os.Args[1])
	if match, _ := regexp.MatchString(`h|j|k|l`, dir); !match {
		log.Fatal("must have an argument h j k or l")
	}

	if windowIsTmux() {
		atBorder, name := getCurrentPane(directions[dir].tmuxBorderCheck)
		if isVim(name) {
			// If vim is on the border of the terminal and another
			// i3 window, then send the M-<dir> keys. This should be
			// mapped to the Focus command from the i3-vim-nav plugin
			// and will let i3 handle navigation if you're moving out
			// of a vim window
			if atBorder {
				executeTmuxCommand("send-keys", "M-"+dir)
				return
			}

			// Otheriwse send C-<dir> which will let tmux handle
			// navigation if you're moving out of a vim window.
			// This requires the vim-tmux-navigator plugin.
			executeTmuxCommand("send-keys", "C-"+dir)
			return
		}
		if !atBorder {
			executeTmuxCommand(
				"select-pane",
				"-"+directions[dir].tmuxSelectPane,
			)
			return
		}
	}

	conn, err := i3ipc.GetIPCSocket()
	if err != nil {
		log.Fatal("could not connect to i3:", err)
	}

	_, err = conn.Command("focus " + directions[dir].i3Focus)
	if err != nil {
		log.Fatalf("failed to focus i3 %s: %s", directions[dir].i3Focus, err)
	}
}

// getCurrentPane will return whether the active pane in
// tmux is on the border (relative to the direction we
// are trying to move in) and the current command executed
// in the pane.
func getCurrentPane(dir string) (bool, string) {
	// See https://man7.org/linux/man-pages/man1/tmux.1.html#FORMATS
	// and https://man7.org/linux/man-pages/man1/tmux.1.html#STATUS_LINE
	// for more details
	out := executeTmuxCommand(
		"display-message",
		"-p",
		"#{pane_at_"+dir+"} #{pane_current_command}",
	)
	s := strings.Split(string(out), " ")
	isBorder, paneCommand := s[0], s[1]

	return isBorder == "1", paneCommand
}

func executeTmuxCommand(args ...string) string {
	tmuxCmd := exec.Command(
		"tmux",
		args...)
	out, err := tmuxCmd.Output()
	if err != nil {
		log.Fatalf("failed executing tmux with arguments: %s, output: %s, error: %s", args, out, err)
	}
	return string(out)
}

func isVim(name string) bool {
	sanitizedName := strings.ToLower(name)
	r, _ := regexp.Compile(`\bn?vim\b`)
	return r.MatchString(sanitizedName)
}

func windowIsTmux() bool {
	xdot := xdo.NewXdo()
	window, err := xdot.GetActiveWindow()
	if err != nil {
		log.Fatal("failed to get i3 window", err)
	}

	name := strings.ToLower(window.GetName())
	user, err := user.Current()
	if err != nil {
		log.Fatal("failed to get username", err)
	}

	// The regex assumes that terminal windows are named
	// "<username>: <command>", which, as far as I can tell,
	// is true for long running processes like tmux.
	// This needs to be done, so that we don't confuse,
	// for example, a browser window which has tmux in it's
	// title and an actual terminal window running tmux.
	r, _ := regexp.Compile(`^` + user.Username + `: \btmux\b`)
	return r.MatchString(name)
}
