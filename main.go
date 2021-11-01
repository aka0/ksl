package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/aka0/ksl/repl"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Current User: %s\n", user.Username)
	fmt.Println("Starting repl:")
	repl.Start(os.Stdin, os.Stdout)
}
