package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	secretNumber := r.Intn(100) + 1
	var guess int

	for {
		fmt.Print("Guess a number between 1 and 100: ")
		fmt.Scan(&guess)
		if guess > secretNumber {
			fmt.Println("Too high! Try again.")
			continue
		} else if guess < secretNumber {
			fmt.Println("Too low! Try again.")
			continue
		} else {
			fmt.Println("Congratulations! You guessed it right.")
			break
		}
	}
}
