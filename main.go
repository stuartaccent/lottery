package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Occurrences struct {
	Key   string
	Value int
}

var allNumbers = make([]int, 49)

func drawLottery(rng *rand.Rand) []int {
	rng.Shuffle(len(allNumbers), func(i, j int) {
		allNumbers[i], allNumbers[j] = allNumbers[j], allNumbers[i]
	})

	lotteryNumbers := allNumbers[:6]

	sort.Ints(lotteryNumbers)
	return lotteryNumbers
}

func runSimulation(times, topLimit int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	occurrences := make(map[string]int)

	for range times {
		numbers := drawLottery(rng)
		key := fmt.Sprintf("%v", numbers)
		occurrences[key]++
	}

	var sortedOccurrences []Occurrences
	for k, v := range occurrences {
		sortedOccurrences = append(sortedOccurrences, Occurrences{k, v})
	}

	sort.Slice(sortedOccurrences, func(i, j int) bool {
		return sortedOccurrences[i].Value > sortedOccurrences[j].Value
	})

	if len(sortedOccurrences) < topLimit {
		topLimit = len(sortedOccurrences)
	}

	fmt.Println("Top 10 Most Frequent Lottery Numbers Sets:")
	for i := range topLimit {
		fmt.Printf("%d. %s: %d times\n", i+1, sortedOccurrences[i].Key, sortedOccurrences[i].Value)
	}
}

// Draws 6 numbers from a range of 49 numbers, 1 to 49, and runs the simulation 1,000,000 times
// to find the most frequent lottery numbers sets.
// The top 20 most frequent lottery numbers sets are printed.
func main() {
	for i := range 49 {
		allNumbers[i] = i + 1
	}
	runSimulation(1000000, 20)
}
