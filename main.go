package main

import (
	"fmt"
	"math/rand"
	"os"
)

type level struct {
	name     string
	attempts int
}

func main() {
	fmt.Println("Welcome to the the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	levels := map[int]level{
		1: {name: "Easy", attempts: 10},
		2: {name: "Medium", attempts: 5},
		3: {name: "Hard", attempts: 3},
	}

	fmt.Println("Please select the difficulty level:")
	for key, value := range levels {
		fmt.Printf("%d. %s (%d chances)\n", key, value.name, value.attempts)
	}

	choice := 0
	var levelChoice level
	inputCheck := false

	for !inputCheck {
		if levelChoice, inputCheck = levels[choice]; !inputCheck {
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice)
		}
	}

	fmt.Printf("\nGreat! You have selected the %s difficulty level.\n", levelChoice.name)

	number := rand.Intn(100)
	pick := -1
	incorrect_prefix := "Incorrect! The number is "
	attempts := 1

	for attempts <= levelChoice.attempts {
		fmt.Printf("\nEnter your guess: ")
		fmt.Scanln(&pick)
		var incorrect_suffix string
		if pick == number {
			fmt.Printf("\nCongratulations! You guessed the correct number in %d attempts\n", attempts)
			os.Exit(0)
		} else if pick > number {
			incorrect_suffix = "less"
		} else {
			incorrect_suffix = "greater"
		}
		fmt.Printf("%s %s %d\n", incorrect_prefix, incorrect_suffix, pick)
		attempts++
	}
	fmt.Printf("\nSorry! you failed to find the number in %d attempts\n", levelChoice.attempts)
}
