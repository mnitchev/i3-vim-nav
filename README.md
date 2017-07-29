# i3-vim-nav
Seamlessly navigate between i3 windows and vim splits.
(most credit for code goes to [termhn](https://github.com/termhn/i3-vim-nav); this fork was made for +python3 support)

## Installation

Using pathogen (recommended)
1. cd ~/.vim/bundle
2. git clone https://github.com/isaaclo123/i3-vim-nav
3. add the following to your .vimrc

	```vim
	" i3 integration
	nnoremap <c-l> :call Focus('right', 'l')<CR>
	nnoremap <c-h> :call Focus('left', 'h')<CR>
	nnoremap <c-k> :call Focus('up', 'k')<CR>
	nnoremap <c-j> :call Focus('down', 'j')<CR>
	```

You can also compile the fo binary yourself using

```
go get -u github.com/isaaclo123/i3-vim-nav
```
if you have go installed.

Then, in your i3 config (adjust the path to the executable as necessary if you installed it differently):

```
bindsym --release $mod+h exec --no-startup-id "~/.vim/bundle/i3-vim-nav/i3-vim-nav h"
bindsym --release $mod+j exec --no-startup-id "~/.vim/bundle/i3-vim-nav/i3-vim-nav j"
bindsym --release $mod+k exec --no-startup-id "~/.vim/bundle/i3-vim-nav/i3-vim-nav k"
bindsym --release $mod+l exec --no-startup-id "~/.vim/bundle/i3-vim-nav/i3-vim-nav l"
```

