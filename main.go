package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func rollDice() int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func newHand(numDice int) []int {
	hand := make([]int, numDice)

	for i := 0; i < 6; i++ {
		hand[i] = rollDice()
	}
	return hand

}

func countValues(hand []int, value int) int {

	count := 0
	for _, v := range hand {
		if v == value {
			count++
		}
	}
	return count
}

func valueResult(valueDice int, numValue int) (int, bool) {
	baseValue := 0
	switch {
	case valueDice == 1:
		baseValue = 1000
	case valueDice == 2:
		baseValue = 200
	case valueDice == 3:
		valueDice = 300
	case valueDice == 4:
		baseValue = 400
	case valueDice == 5:
		baseValue = 500
	case valueDice == 6:
		baseValue = 600
	}

	if valueDice == 1 || valueDice == 5 {
		switch {
		case numValue < 3:
			if valueDice == 1 {
				return numValue * 100, false
			} else {
				return numValue * 50, false
			}
		case numValue == 3:
			return baseValue, false
		case numValue == 4:
			return baseValue * 2, false
		case numValue == 5:
			return baseValue * 4, false
		case numValue == 6:
			return 0, true

		default:
			return 0, false
		}
	} else {
		switch {
		case numValue == 3:
			return baseValue, false
		case numValue == 4:
			return baseValue * 2, false
		case numValue == 5:
			return baseValue * 4, false
		case numValue == 6:
			return 0, true

		default:
			return 0, false
		}

	}

}

/*
func valueOnes(numOnes int) (int, bool) {

	switch {
	case numOnes < 3:
		return numOnes * 100, false
	case numOnes == 3:
		return 1000, false
	case numOnes == 4:
		return 2000, false
	case numOnes == 5:
		return 4000, false
	case numOnes == 6:
		return 0, true

	default:
		return 0, false
	}

}

func valueTwos(numTwos int) (int, bool) {
	switch {
	case numTwos == 3:
		return 200, false
	case numTwos == 4:
		return 400, false
	case numTwos == 5:
		return 800, false
	case numTwos == 6:
		return 0, true

	default:
		return 0, false
	}

}

func valueThrees(numThrees int) (int, bool) {
	switch {
	case numThrees == 3:
		return 300, false
	case numThrees == 4:
		return 600, false
	case numThrees == 5:
		return 900, false
	case numThrees == 6:
		return 0, true

	default:
		return 0, false
	}

}

func valueFours(numFours int) (int, bool) {
	switch {
	case numFours == 3:
		return 400, false
	case numFours == 4:
		return 800, false
	case numFours == 5:
		return 1600, false
	case numFours == 6:
		return 0, true

	default:
		return 0, false
	}

}

func valueFives(numFives int) (int, bool) {
	switch {
	case numFives < 3:
		return numFives * 100, false
	case numFives == 3:
		return 400, false
	case numFives == 4:
		return 800, false
	case numFives == 5:
		return 1600, false
	case numFives == 6:
		return 0, true

	default:
		return 0, false
	}

}

func valueSixes(numSixes int) (int, bool) {
	switch {

	case numSixes == 3:
		return 400, false
	case numSixes == 4:
		return 800, false
	case numSixes == 5:
		return 1600, false
	case numSixes == 6:
		return 0, true

	default:
		return 0, false
	}

} */

func printResult(hand []int) {

	v := 0
	w := false
	sum := 0

	for i := 1; i < 7; i++ {

		v, w = valueResult(i, countValues(hand, i))
		sum += v
		fmt.Println(i, ": ", v, ", win: ", w)
	}
	fmt.Println("Hand Sum: ", sum)

}

func main() {
	hand := newHand(6)
	fmt.Println(hand)
	printResult(hand)

	result := make(map[string]int)
	var index string
	for i := 0; i < len(hand); i++ {
		index = strconv.Itoa(i + 1)
		result[index] = hand[i]

	}

	fmt.Println(result)

}
