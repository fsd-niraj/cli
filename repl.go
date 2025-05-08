package main

import (
	"bufio"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("fsd-niraj@cli > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput((text))

		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		availableCommands := getCommand()
		command, ok := availableCommands[commandName]

		if !ok {
			println("cli: command not found")
			continue
		}
		command.callback()
	}

}

type cliCommands struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(str string) []string {
	lowercase := strings.ToLower(str)
	words := strings.Fields(lowercase)
	return words
}
