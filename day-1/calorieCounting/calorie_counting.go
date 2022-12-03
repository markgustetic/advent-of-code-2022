package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	TotalCalories int
}

func main() {
	elves := getElves()
	maxCalories := calculateCalories(elves)

	fmt.Println(maxCalories)
}

func getElves() []Elf {
	var elves []Elf

	f, err := os.Open("../input-test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var currentElf Elf

	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, currentElf)
			currentElf = Elf{}
		} else {
			currentCalories, err := strconv.Atoi(scanner.Text())

			if err != nil {
				log.Fatal(err)
			}

			currentElf.TotalCalories += currentCalories
		}
	}

	elves = append(elves, currentElf)

	return elves
}

func calculateCalories(elves []Elf) int {
	var maxCalories int

	for _, elf := range elves {
		if elf.TotalCalories > maxCalories {
			maxCalories = elf.TotalCalories
		}
	}

	return maxCalories
}
