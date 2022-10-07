package main

import (
	"fmt"
)

func ChooseLetter(GameInProgress Game) Game {
	var letter string
	fmt.Scanln(&letter)
	if !IsAlpha(letter) {
		fmt.Println("Please, select a letter, try again. ")
		fmt.Println("")
		fmt.Println("___________________________________")
	} else if len(letter) > 1 {
		fmt.Println("Please, select only one letter, try again. ")
		fmt.Println("")
		fmt.Println("___________________________________")
	} else {
		println("letter: ", letter)
	}
	return GameInProgress
}

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 96 && s[i] < 123 {
			return true
		}
		if s[i] > 64 && s[i] < 91 {
			return true
		}
	}
	return false
}
