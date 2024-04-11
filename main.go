package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Rock-Paper-Scissors")

	//possible moves
	moves := [3]string{
		"Rock",
		"Paper",
		"Scissors",
	}

	for {
		//get the players move
	var playerMove string
	fmt.Println("Rock, Paper or Scissors?")
	fmt.Scanln(&playerMove)

	//Computer move
	computerMove := moves[rand.Intn(3)]

	//print the results
	fmt.Printf("You played %s and computer played %s\n", playerMove, computerMove)

	//logic
	switch {
	case playerMove == computerMove:
		fmt.Println("Its a tie!")
	case playerMove == "Rock" && computerMove == "Scissors" || playerMove == "Paper" && computerMove == "Rock" || playerMove == "Scissors" && computerMove == "Paper":
		fmt.Println("You win")
	default:
		fmt.Println("Computer win")
	}
	}
	

}
