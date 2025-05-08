package main

func getCommand() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off cli",
			callback:    callbackExit,
		},
	}
}
