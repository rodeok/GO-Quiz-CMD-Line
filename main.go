package main

import "fmt" // package fmt provides basic formatted I/O primitives.

func main() {
	fmt.Println("Welcome to my Quiz Game")
	fmt.Printf("Enter\nyour name: ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, welcome to the game! \n", name)
	fmt.Printf("\n Enter your age: ")
	var age uint
	fmt.Scan(&age)

	// fmt.Println(age >= 10)
	if age >= 10 {
		fmt.Println("You are eligible to play the game")
	} else {
		fmt.Println("You are not eligible to play the game")
		return
	}
	score := 0
	num_questions := 2

	fmt.Println("What is Better, the RTX 3080 or the RTX 3090")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	fmt.Println(answer, answer2)
	if answer+" "+answer2 == "the RTX 3090" {
		fmt.Println("Correct")
		score++
	} else if answer+" "+answer2 == "the rtx 3090" {
		fmt.Println("Correct")
		score++
	} else if answer+" "+answer2 == "the Rtx 3090" {
		fmt.Println("Correct")
		score++
	} else if answer+" "+answer2 == "the rTx 3090" {
		fmt.Println("Correct")
		score++
	} else if answer+" "+answer2 == "the rtX 3090" {
		fmt.Println("Correct")
		score++
	} else {
		// fmt.Println("Wrong")
	}
	fmt.Printf("How Many Cores does Ryzen 9 3900X have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct")
		score++
	} else {
		fmt.Println("Wrong")
	}
	fmt.Printf("You Scored %v out of %v.", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf(" \n You scored: %v%%.", percent)
}
