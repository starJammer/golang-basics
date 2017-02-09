package main

import (
	"fmt"
)

//Setup your $GOROOT -- /usr/local/Cellar/go/1.7/libexec
//export GOROOT=/usr/local/Cellar/go/1.7/libexec
//Setup your $GOPATH -- /Users/j/Documents/code/go
//export GOPATH=/Users/j/Documents/code/go
//mkdir -p {$GOPATH/bin,$GOPATH/src/github.com/starJammer,$GOPATH/pkg}
//Optional: Add bin to your path to mak binaries easier to execute
//export PATH=$GOPATH/bin:$PATH

//Move this code into the GOPATH
//mv $GOPATH/src/github.com/starJammer

//Things to try:
//go run hello.go -- compiles and runs in place
//go build -- creates a binary called golang-basics
//go build -o hello -- creates a binary named hello
//go install -- creates the binary and installs it into your $GOPATH/bin directory
func main() {
	fmt.Println("I hate hello worlds")
}
