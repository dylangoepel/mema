# mema
*mema* is a program made to train your mental calculation skills.
## Installation
mema was built using go and its gtk3 binding, gotk3, which may be installed via
```
go get -u github.com/gotk3/gotk3/gtk
'''
In order to build the binary, just run
```
go build
'''
inside the directory you cloned the repository into.
## Usage
mema was made to be fully to be exclusively used with the keyboard, and therefore adapt to the UI philosophy of my [dotfiles](https://github.com/dylangoepel/dotfiles)
and my [window manager](https://github.com/dylangoepel/dwm).
You can start a challenge / view its solution by hitting the Space key, raise / lower the delay between the numbers using the + / - keys and raise / lower the count of numbers using the Up and Down arrow keys.
