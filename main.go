package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	gameLoop(input)
}

func gameLoop(input *bufio.Reader) {

	var candidates []candidate = make([]candidate, 0, ((10 - 1) * 9 * 8 * 7))
	candidates = permutations(make([]uint8, 4), 3, candidates)
	computersNumber := candidates[rand.Intn(len(candidates))]

	for {

		for {
			fmt.Print("Ask the computer for its number: ")

			parsedInput, err := parseInput(input)
			if err != nil {
				fmt.Println(fmt.Sprintf("Input error: %v; Try again!", err))
			} else {
				answer := computersNumber.Compare(parsedInput)
				fmt.Println(fmt.Sprintf("Bulls: %d, Cows: %d", answer.bulls, answer.cows))

				if answer.bulls == 4 {
					fmt.Println()
					fmt.Println("You won!!!")
					return
				}

				break
			}
		}

		computersQuery := candidates[rand.Intn(len(candidates))]

		for {
			fmt.Println(fmt.Sprintf("The computer asks for your number: %s", computersQuery.Show()))
			fmt.Print("How many bulls: ")
			parsedAnswerBulls, err := parseAnswer(input)

			if parsedAnswerBulls == 4 {
				fmt.Println()
				fmt.Println("The computer won!!!")
				return
			}

			if err != nil {
				fmt.Println(fmt.Sprintf("Input error: %v; Try again!", err))
			} else {

				fmt.Print("How many cows: ")
				parsedAnswerCows, err := parseAnswer(input)

				if err != nil {
					fmt.Println(fmt.Sprintf("Input error: %v; Try again!", err))
				} else {

					if parsedAnswerBulls+parsedAnswerCows > 4 {
						fmt.Println("The sum of cows and bulls can not be greater than 4")
					} else {
						candidates = filterNonMatchingCandidates(candidates, computersQuery, answer{bulls: parsedAnswerBulls, cows: parsedAnswerCows})

						if len(candidates) == 0 {
							fmt.Println()
							fmt.Println("There are no possible numbers matching your answers! The computer is sad :(")
							return
						}
						break
					}
				}
			}
		}

	}
}

func parseInput(input *bufio.Reader) (candidate, error) {
	line, isPrefix, err := input.ReadLine()
	if isPrefix {
		// error - too long input
		return candidate{}, fmt.Errorf("Too long input")
	}
	if err != nil {
		return candidate{}, err
	}
	if len(line) != 4 {
		return candidate{}, fmt.Errorf("Input length must be 4")
	}
	for i := 0; i < len(line); i++ {
		line[i] -= 0x30
		if line[i] < 0 || line[i] > 9 {
			// error
			return candidate{}, fmt.Errorf("Input must contain only digits 0-9")
		}
	}

	c := (candidate)(line)
	if !c.IsValid() {
		return candidate{}, fmt.Errorf("Invalid input %s. There should be no repeating digits and the first digit must not be zero!", c.Show())
	}

	return c, nil
}

func parseAnswer(input *bufio.Reader) (uint8, error) {
	line, isPrefix, err := input.ReadLine()
	if isPrefix {
		// error - too long input
		return 0, fmt.Errorf("Too long input")
	}
	if err != nil {
		return 0, err
	}
	if len(line) != 1 {
		return 0, fmt.Errorf("Input length must be single digit")
	}
	line[0] -= 0x30
	if line[0] < 0 || line[0] > 4 {
		// error
		return 0, fmt.Errorf("Input must contain only digits 0-4")
	}

	return line[0], nil
}
