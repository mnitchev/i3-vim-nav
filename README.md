# i3-vim-tmux-nav
Seamlessly change focus between i3 windows, Vim and Tmux splits using the same hotkey.

# Installation

## Dependencies

This depends on you having a couple packages installed. Most notably, `xdotool/libxdo`, the second of which should be installed as a dependency of the first. If you want to build the binary from source and you're on Fedora, you'll also need the `libxdo-devel` package. We also currently require you to be using a version of vim with either python or python3 support.
To get seamless switching between vim and tmux you need the [vim-tmux-navigator](https://github.com/christoomey/vim-tmux-navigator).

## Vim plugin

First, install the Vim plugin. The `vim-tmux-navigator` plugin already maps `<C-hjkl>` to switch between vim and tmux splits. In addition to this we have to add `<M-hjkl>` to switch between vim and i3. These mappings won't be used directly, but will rather be called by the the binary installed in the next step. **This may however interfere with mappings that you may already have in your vim config!**

### Using vim-plug
In your .vimrc (vim) or .config/nvim/init.vim (neovim):

```vim
Plug 'termhn/i3-vim-nav'
Pluig 'christoomey/vim-tmux-navigator'

" i3 integration
nnoremap <silent> <M-l> :call Focus('right', 'l')<CR>
nnoremap <silent> <M-h> :call Focus('left', 'h')<CR>
nnoremap <silent> <M-k> :call Focus('up', 'k')<CR>
nnoremap <silent> <M-j> :call Focus('down', 'j')<CR>
```

### Using Pathogen
1. cd ~/.vim/bundle
1. git clone https://github.com/termhn/i3-vim-nav
1. git clone https://github.com/christoomey/vim-tmux-navigator.git
1. add the following to your .vimrc

```vim
" i3 integration
nnoremap <silent> <M-l> :call Focus('right', 'l')<CR>
nnoremap <silent> <M-h> :call Focus('left', 'h')<CR>
nnoremap <silent> <M-k> :call Focus('up', 'k')<CR>
nnoremap <silent> <M-j> :call Focus('down', 'j')<CR>
```

## Binary

Next, install the binary on your PATH. If you have go installed, this can be done simply by

```
go get -u github.com/mnitchev/i3-vim-tmux-nav
```
If not, simply donwload [the latest release]() and put it in PATH.
**NOTE:** make sure it's in the PATH set in `~/.profile`.

Then, in your i3 config (adjust the path to the executable as necessary if you installed it differently). Feel free to change the key bind as you please.

```
bindsym --release Control+h exec --no-startup-id "i3-vim-tmux-nav h"
bindsym --release Control+j exec --no-startup-id "i3-vim-tmux-nav j"
bindsym --release Control+k exec --no-startup-id "i3-vim-tmux-nav k"
bindsym --release Control+l exec --no-startup-id "i3-vim-tmux-nav l"
```

Note: I've gotten a bug where in some installations of i3, it seems to not respect user $PATH additions, even though it seems to recognize them in the variable. If it doesn't work when placed in a user $PATH directory, try hard-coding the path to the binary in the exec commands.
