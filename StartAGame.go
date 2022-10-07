package main

import (
	"fmt"
	hangman "hangman/Functions"
	"strings"
)

func StartPlaying(GameInProgress Game) { // go file in root to use struc "game" freely
	hangman.PrintWord(GameInProgress.Word, GameInProgress.RevealedLettres)

	for GameInProgress.Tries < 10 && !WordCompleted(GameInProgress) {
		fmt.Printf("Choose a letter : ")
		guess := ChooseLetter(GameInProgress)
		GameInProgress.RevealedLettres = append(GameInProgress.RevealedLettres, strings.ToUpper(guess))
		if IsGoodAnswer(GameInProgress, guess) {
			hangman.PrintWord(GameInProgress.Word, GameInProgress.RevealedLettres)
		} else {
			PrintJose(GameInProgress)
			GameInProgress.Tries++
		}
	}

	println("ggwp...")
}
