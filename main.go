package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	choices := []string{"stone", "paper", "Scissors"}
	fmt.Println("Welcome to the game of stone, paper, scissors!")
	fmt.Println("Enter your choice (stone, paper, scissors):")

	var userChoice string
	fmt.Print("Your choice: ")
	fmt.Scan(&userChoice)

	// انتخاب تصادفی کامپیوتر
	computerChoice := choices[rand.Intn(len(choices))]
	fmt.Printf("The computer has selected: %s\n", computerChoice)

	// تعیین برنده
	if userChoice == computerChoice {
		fmt.Println("A draw!")
	} else if (userChoice == "Stone" && computerChoice == "Scissors") ||
		(userChoice == "paper" && computerChoice == "stone") ||
		(userChoice == "Scissors" && computerChoice == "paper") {
		fmt.Println("You Won!")
	} else {
		fmt.Println("The computer won!")
	}
}
