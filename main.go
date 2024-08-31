package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	numberToGuess := generateNumber()

	fmt.Println("Number is guessed. Try to guess it. Write 'exit' if you want to quit")

	for {
		fmt.Println("Your number is ")
		var userNumber string
		fmt.Scan(&userNumber)

		if userNumber == "exit" {
			break
		}

		if len(userNumber) != 4 {
			fmt.Println("Your number is too long. Enter 4-digits number")
			continue
		}

		cows, bulls := calculateAnimals(numberToGuess, userNumber)
		fmt.Printf("Cows are %d, bulls are %d\n", cows, bulls)

		if bulls == 4 {
			fmt.Println("Congrats! Number to guess is ", numberToGuess)
			break
		}

	}
}

func generateNumber() string {
	digits := rand.Perm(10)[:4]
	if digits[0] == 0 {
		digits[0], digits[1] = digits[1], digits[0]
	}
	number := ""
	for _, digit := range digits {
		number += strconv.Itoa(digit)
	}

	return number
}

func calculateAnimals(pcNumber, userNumber string) (cows, bulls int) {
	for i, pcDigit := range pcNumber {
		for j, userDigit := range userNumber {
			if pcDigit == userDigit {
				if i == j {
					bulls++
				} else {
					cows++
				}
			}
		}
	}
	return
}
