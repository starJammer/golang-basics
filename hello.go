package main

//Setup $GOROOT - /usr/local/Cellar/go/1.7/libexec
//Setup $GOPATH - /Users/j/Documents/code/go
//In your GOPATH dir, `mkdir -p {src,bin,pkg}`
import (
	"fmt"
)

//go run hello.go
//go build -- creates a binary file named golang-basics
//go build -o hello -- creates a binary file named hello
//go install . -- installs the binary into the your $GOPATH/bin
func main() {
	fmt.Println("I hate hello worlds")
}
