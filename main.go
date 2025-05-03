package main

import ("fmt"
"strings"
"bufio"
"os"
)


type cliCommand struct {
	name string
	description string
	callback func() error
}


func cleanInput(text string) []string {
	
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	
	fmt.Println("Closing the Pokedex... Goodbye!")
	 os.Exit(0)
	 return nil  // This line won't be reached, but satifises the compiler 
}

func commandHelp( commands map[string]cliCommand) func() error {

	// uses a closure to satisfy the struct call back type

	return func () error {

		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println()

		for _, cmd := range commands {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}

		return nil
	}
}

func main() {
	// makes a new instance of the scanner struct, so I can now listen for user input.
	scanner := bufio.NewScanner(os.Stdin)

	// Creates the command registry
	commands := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}

	/* 
	I've had to add the help function after the initialisation of the map, this is due to
	the callback needing the commands map itself and it would create a circular dependency
	*/

	commands["help"] = cliCommand{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp(commands),
		}


	// makes an infinite loop for listening to user inputs

	for {
		fmt.Print("Pokedex > ")

		// waits for the user to input something and press enter
		scanner.Scan()

		// Get what the user has typed as a string value
		input := scanner.Text()

		cleanedInput := cleanInput(input)

		// make sure there is at least value in the input, if not loop will skip to next iteration

		if len(cleanedInput) == 0 {
			continue
		}

		// for Now, I'm only working with the first word of every user input.
		commandName := cleanedInput[0]

		command, exists := commands[commandName]

		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	
	}
}