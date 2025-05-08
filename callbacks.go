package main

import (
	"fmt"
	"os"
)

func callbackHelp() error {
	println("Here are your available commands:")
	availableCommands := getCommand()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	println("")
	return nil
}

func callbackExit() error {
	os.Exit(0)
	return nil
}
