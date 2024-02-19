package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("\nWelcome to the Pokedex")
	fmt.Println("Here are the available commands:")
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")

	return nil
}
