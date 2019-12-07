package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/0daryo/monkey/repl"
)

func Exec() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
