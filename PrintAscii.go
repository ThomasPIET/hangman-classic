package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func PrintWord() {
	display := ""
	var LettersOfWord []string
	for _, l := range GameInProgress.Word {
		LettersOfWord = append(LettersOfWord, string(l))
	}
	for _, l := range LettersOfWord {
		if DoesContain(GameInProgress.RevealedLettres, l) == true {
			if utf8.RuneCountInString(GameInProgress.Word)*12 > w.colones*65/100-4 || w.ligns/2 < 11 {
				display += " " + l
			} else {
				display = display + l
			}
		} else {
			if utf8.RuneCountInString(GameInProgress.Word)*12 > w.colones*65/100-4 || w.ligns/2 < 11 {
				display += " _"
			} else {
				display += "_"
			}
		}
	}
	PrintAscii(w.ligns/4-3, w.colones*65/200-(utf8.RuneCountInString(GameInProgress.Word)*12)/2, display)
	fmt.Print("\x1B[0m")
}

func PrintAscii(y, x int, word string) {
	if utf8.RuneCountInString(word)*12 > w.colones*65/100-4 || w.ligns/2 < 11 {
		fmt.Print(MoveTo(y+3, x+utf8.RuneCountInString(word)*3-utf8.RuneCountInString(word)/2), "\x1B[1m"+word)
		fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[0m")
		return
	}
	fmt.Print("\x1B[7m")
	data, err := os.ReadFile("FilesAndLists/fancyLetters.txt")
	if err != nil {
		panic(err)
	}
	var CharLigns [][]string
	sep := []byte{13, 10, 13, 10}
	char := strings.Split(string(data), string(sep))
	for _, l := range char {
		CharLigns = append(CharLigns, strings.Split(l, string(sep[2:])))
	}
	characters := [...]string{"<", ">", " ", "-", "(", ")", ";", ",", ".", "_", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "А", "Б", "В", "Г", "Д", "Е", "Ё", "Ж", "З", "И", "Й", "К", "Л", "М", "Н", "О", "П", "Р", "С", "Т", "У", "Ф", "Х", "Ц", "Ч", "Ш", "Щ", "Ъ", "Ы", "Ь", "Э", "Ю", "Я"}
	var LettersOfWord []string
	for _, l := range word {
		LettersOfWord = append(LettersOfWord, string(l))
	}

	for floor := 0; floor < 9; floor++ {
		for k, l := range LettersOfWord {
			for i, c := range characters {
				if l == c {
					fmt.Print(MoveTo(y+floor, x+k*12), CharLigns[i][floor])
				}
			}
		}
	}
	fmt.Print(MoveTo(w.ligns/2+2, 2), "\x1B[0m")
}
