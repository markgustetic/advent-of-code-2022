package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	A             = "Rock"
	X             = "Rock"
	B             = "Paper"
	Y             = "Paper"
	C             = "Sissors"
	Z             = "Sissors"
	RockScore     = 1
	PaperScore    = 2
	ScissorsScore = 3
	Lose          = 0
	Draw          = 3
	Win           = 6
)

type Move struct {
	PlayerMove   string
	OpponentMove string
}

func (m Move) Score() int {
	var totalScore int

	switch m.PlayerMove {
	case "X":
		totalScore += RockScore
	case "Y":
		totalScore += PaperScore
	case "Z":
		totalScore += ScissorsScore
	}

	totalScore += m.ScoreMoves()

	return totalScore
}

func (m Move) ConvertOpponentMove() string {
	switch m.OpponentMove {
	case "A":
		return A
	case "B":
		return B
	default:
		return C
	}
}

func (m Move) ConvertPlayerMove() string {
	switch m.PlayerMove {
	case "Y":
		return Y
	case "X":
		return X
	default:
		return Z
	}
}

func (m Move) ScoreMoves() int {
	playerMove := m.ConvertPlayerMove()
	opponentMove := m.ConvertOpponentMove()

	if playerMove == opponentMove {
		return Draw
	}

	if playerMove == "Rock" {
		if opponentMove == "Scissors" {
			return Win
		} else {
			return Lose
		}
	}

	if playerMove == "Paper" {
		if opponentMove == "Rock" {
			return Win
		} else {
			return Lose
		}
	}

	if opponentMove == "Paper" {
		return Win
	} else {
		return Lose
	}
}

func main() {
	allMoves := getMoves()

	totalScore := calcScores(allMoves)

	fmt.Println(totalScore)
}

func getMoves() []Move {
	var allMoves []Move

	f, err := os.Open("../input-test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")

		allMoves = append(allMoves, Move{moves[1], moves[0]})
	}

	return allMoves
}

func calcScores(moves []Move) int {
	var totalScore int

	for _, move := range moves {
		totalScore += move.Score()
	}

	return totalScore
}
