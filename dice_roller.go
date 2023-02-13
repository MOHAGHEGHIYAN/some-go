package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	dice, sides, rolls := 2, 6, 2
	for r := 1; r <= rolls; r++ {
		sum := 0
		fmt.Println("roll:", r)
		for d := 1; d <= dice; d++ {
			rolled := rand.Intn(sides) + 1
			sum += rolled
			fmt.Println("Dice #", d, "rolled with:", rolled)
		}

		fmt.Println("sum is: ", sum)
		switch sum := sum; {
		case sum%2 == 0:
			fmt.Println("Even\n")
		default:
			fmt.Println("Odd\n")
		}
	}
}
