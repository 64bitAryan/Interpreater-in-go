package main

import (
	"fmt"
	"github/64bitaryan/interpreater/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s this is monkey programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
