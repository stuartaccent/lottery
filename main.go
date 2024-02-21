package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type (
	SetOccurrences struct {
		Key   string
		Count int
	}
	NumberOccurrences struct {
		Number int
		Count  int
	}
)

var allNumbers = make([]int, 59)

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
	setOccurrences := make(map[string]int)
	numberOccurrences := make(map[int]int)

	for range times {
		numbers := drawLottery(rng)
		key := fmt.Sprintf("%v", numbers)
		setOccurrences[key]++
		for _, number := range numbers {
			numberOccurrences[number]++
		}
	}

	var sortedSetOccurrences []SetOccurrences
	for k, v := range setOccurrences {
		sortedSetOccurrences = append(sortedSetOccurrences, SetOccurrences{k, v})
	}

	sort.Slice(sortedSetOccurrences, func(i, j int) bool {
		return sortedSetOccurrences[i].Count > sortedSetOccurrences[j].Count
	})

	topSetLimit := min(topLimit, len(sortedSetOccurrences))
	fmt.Printf("Top %d Most Frequent Lottery Numbers:\n", topSetLimit)
	for i := range topSetLimit {
		fmt.Printf("%d. %s: %d times\n", i+1, sortedSetOccurrences[i].Key, sortedSetOccurrences[i].Count)
	}

	var sortedNumberOccurrences []NumberOccurrences
	for number, count := range numberOccurrences {
		sortedNumberOccurrences = append(sortedNumberOccurrences, NumberOccurrences{Number: number, Count: count})
	}

	sort.Slice(sortedNumberOccurrences, func(i, j int) bool {
		return sortedNumberOccurrences[i].Count > sortedNumberOccurrences[j].Count
	})

	topNumberLimit := min(topLimit, len(sortedNumberOccurrences))
	fmt.Printf("\nTop %d Most Frequently Drawn Individual Numbers:\n", topNumberLimit)
	for i := range topNumberLimit {
		fmt.Printf("%d. %d: %d times\n", i+1, sortedNumberOccurrences[i].Number, sortedNumberOccurrences[i].Count)
	}
}

func timeTaken(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

// Draws 6 numbers from a range of 59 numbers, 1 to 59, and runs the simulation 1,000,000 times
// to find the most frequent lottery numbers sets.
// The top 20 most frequent lottery numbers sets are printed.
// The top 20 most frequent drawn individual numbers are printed.
func main() {
	defer timeTaken(time.Now(), "runSimulation")
	for i := range 59 {
		allNumbers[i] = i + 1
	}
	runSimulation(1000000, 20)
}
