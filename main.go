package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Input number: ")
	fmt.Scan(&input)

	number, err := oddAndEvenNumbers(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(number)
}

func oddAndEvenNumbers(input string) (string, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return "", fmt.Errorf("Input isn't number")
	}

	if i%2 == 0 {
		return "even number", nil
	} else {
		return "odd number", nil
	}
}
