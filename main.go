package main

import (
	"fmt"
	"math/rand"
)

func main() {

	const rounds = 3

	fmt.Println("Rock-Paper-Scissors")

	for i := 0; i < rounds; i++ {
		computerChoiceNum := rand.Intn(rounds)
		var computerChoice string
		switch computerChoiceNum{
		case 0:
			computerChoice = "Rock"
		case 1:
			computerChoice = "Paper"
		case 2:
			computerChoice = "Scissors"
		}
		
	
		//User choice
		var playersChoice string
		fmt.Println("Enter your choice ('Rock', 'Paper', 'Choice': )")
		fmt.Scanln(&playersChoice)

		//game logic
		fmt.Printf("Computer chose: %s\n", computerChoice)

		switch {
		case playersChoice == computerChoice:
			fmt.Println("Its a tie!")
		case playersChoice == "Rock" && computerChoice == "Paper":
			fmt.Println("Paper wraps the rock. Computer wins!")
		case playersChoice == "Paper" && computerChoice == "Scissors":
			fmt.Println("Scissors cuts the paper. Computer wins!")
		case playersChoice == "Scissors" && computerChoice == "Rock":
			fmt.Println("Rocks slams the scissors. Computer wins")
		case playersChoice == "Rock" && computerChoice == "Scissors":
			fmt.Println("Rock slams the scissors. You wins!")
		case playersChoice == "Paper" && computerChoice == "Rock":
			fmt.Println("Paper wraps the rock. You wins!")
		case playersChoice == "Scissors" && computerChoice == "Paper":
			fmt.Println("Scissors cuts the paper. You wins")
		}
	}

	fmt.Println("Game over!")
}
