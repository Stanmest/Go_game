package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the number!")
	fmt.Println("You have only seven chances to get the correct number")

	//generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(100)

	//variable declarations
	var guess int
	var limit int = 0
	const even string = "It's an even number!"
	const odd string = "It's an odd number!"

	for {
		fmt.Println("Please pick a number from 1 to 100")
		fmt.Scan(&guess)
		//add a hint if the secret number is even or odd
		if secretNumber%2 == 0 {
			if guess > secretNumber {
				fmt.Println("Hint: Too Big, and ", even)
			} else if guess < secretNumber {
				fmt.Println("Hint: Too Small, and ", even)
			} else {
				fmt.Println("YaY!...You guessed correctly,You win! ")
				break
			}
		} else {
			if guess > secretNumber {
				fmt.Println("Hint: Too Big, and ", odd)
			} else if guess < secretNumber {
				fmt.Println("Hint: Too Small, and ", odd)
			} else {
				fmt.Println("YaY!...You guessed correctly,You win! ")
				break
			}
		}
		limit++
		//end game if number of guesses is 7
		if limit == 7 {
			fmt.Println("Sorry, you've exceeded your trial, you can try again! :) ")
			break
		}
	}

}