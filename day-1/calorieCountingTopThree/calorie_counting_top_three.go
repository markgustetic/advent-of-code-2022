package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	TotalCalories int
}

func main() {
	elves := getElves()

	topThreeTotal := calculateTopThree(elves)

	fmt.Println(topThreeTotal)
}

func getElves() []Elf {
	var elves []Elf

	f, err := os.Open("./input-test.txt")

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

func calculateTopThree(elves []Elf) int {
	var topThreeTotal int
	sortedElves := make([]Elf, len(elves))

	copy(sortedElves, elves)

	sort.Slice(sortedElves, func(i, j int) bool {
		return sortedElves[i].TotalCalories > sortedElves[j].TotalCalories
	})

	for i := 0; i < 3; i++ {
		topThreeTotal += sortedElves[i].TotalCalories
	}

	return topThreeTotal
}
