package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrInvalidChars = errors.New("invalid characters")
)

func Solution(S string) (string, error) {
	totalWinsByX, totalWinsByO := 0, 0
	consecutiveWinsByX, consecutiveWinsByO := 0, 0

	for _, win := range S {
		switch win {
		case 'X':
			totalWinsByX++
			consecutiveWinsByX++
			consecutiveWinsByO = 0

			if consecutiveWinsByX == 3 {
				consecutiveWinsByX = 0
			}
		case 'O':
			totalWinsByO++
			consecutiveWinsByO++
			consecutiveWinsByX = 0

			if consecutiveWinsByO == 3 {
				consecutiveWinsByO = 0
			}
		default:
			return "", ErrInvalidChars
		}
	}

	if totalWinsByX > totalWinsByO {
		return "X", nil
	} else if totalWinsByO > totalWinsByX {
		return "O", nil
	} else {
		return "tie", nil
	}
}

func main() {
	S := "XOXOXXXOOOXXXOOO"

	result, err := Solution(S)
	if err != nil {
		log.Fatalf("could not compute the solution: %s", err)
	}
	fmt.Println(result)
}
