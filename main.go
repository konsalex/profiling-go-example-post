package main

import (
	_ "embed"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func createRuleMap(data *FormattedData) map[int][]int {
	ruleMap := make(map[int][]int)

	// Create a map of rules
	for _, rule := range data.Rules {
		if ruleMap[rule[0]] == nil {
			ruleMap[rule[0]] = []int{rule[1]}
		} else {
			ruleMap[rule[0]] = append(ruleMap[rule[0]], rule[1])
		}
	}
	return ruleMap
}

// To avoid using slices.Index, we can create our own index finder
// to avoid the overhead of the new slice allocation
func indexFinder(slice []int, value, maxIndex int) int {
	for i, v := range slice {
		if i >= maxIndex {
			return -1
		}
		if v == value {
			return i
		}
	}
	return -1
}

// Part2Fixes will avoid using so much indexFinder, as its has a slow runtime
func Part2Fixes(data FormattedData) (sum int) {
	ruleMap := createRuleMap(&data)

	for _, update := range data.Updates {
		faulty := false
		i := 0

	OUTER:
		for i < len(update) {
			character := update[i]

			if ruleMap[character] == nil {
				i++
				continue
			}

			// To avoid finding all the time the index of the character
			// we can store the indexes of the characters, and swap them when we swap the letters
			charIndexes := map[int]int{}
			for idx, char := range update[:i] {
				charIndexes[char] = idx
			}

			// Check if the rule is valid
			for _, r := range ruleMap[character] {
				// Rule is invalid
				//if idx := slices.Index(update[:i], r); idx != -1 {
				if idx, ok := charIndexes[r]; ok {
					// We need to be faulty at least once
					faulty = true
					// Swap the character positions and start again
					update[i], update[idx] = update[idx], update[i]
					// Update the indexes
					charIndexes[update[i]], charIndexes[update[idx]] = idx, i
					i = 0
					continue OUTER
				}
			}
			i++
		}

		if faulty {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func Part2(data FormattedData) (sum int) {
	ruleMap := createRuleMap(&data)

	for _, update := range data.Updates {
		faulty := false
		i := 0

		for i < len(update) {
			character := update[i]
			if ruleMap[character] == nil {
				i++
				continue
			}

			// Check if the rule is valid
			for _, r := range ruleMap[character] {
				// Rule is invalid
				//if idx := slices.Index(update[:i], r); idx != -1 {
				if idx := indexFinder(update, r, i); idx != -1 {
					// We need to be faulty at least once
					faulty = true
					// Swap the character positions and start again
					update[i], update[idx] = update[idx], update[i]
					i = 0
					continue
				}
			}
			i++
		}

		if faulty {
			sum += update[len(update)/2]
		}
	}

	return sum
}

type Rule[T any] [2]T

type FormattedData struct {
	Rules   []Rule[int]
	Updates [][]int
}

func StringToFormatted(input string) FormattedData {
	formattedData := FormattedData{
		Rules:   []Rule[int]{},
		Updates: [][]int{},
	}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// Check if line has "," or "|" or nothing
		if strings.Contains(line, "|") {
			// Rule
			rule := Rule[int]{}
			for splitIdx, r := range strings.Split(line, "|") {
				// Convert string to int
				ruleInt, err := strconv.Atoi(r)
				if err != nil {
					panic(err)
				}
				rule[splitIdx] = ruleInt
			}
			formattedData.Rules = append(formattedData.Rules, rule)
		} else if strings.Contains(line, ",") {
			// Update
			update := []int{}
			for _, u := range strings.Split(line, ",") {
				// Convert string to int
				updateInt, err := strconv.Atoi(u)
				if err != nil {
					panic(err)
				}
				update = append(update, updateInt)
			}
			formattedData.Updates = append(formattedData.Updates, update)
		}
	}

	return formattedData
}

func main() {

	f, perr := os.Create("cpu.pprof")
	if perr != nil {
		panic(perr)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 1000; i++ {
		formattedData := StringToFormatted(input)
		Part2(formattedData)
		Part2Fixes(formattedData)
	}
}
