package main

import (
	"fmt"
)

func main() {
	//Loops forvever until break condition, iterates i by 1 every cycle.
	for i := 1; true; i++ {
		if isSpecialNumber(i) {
			name := makeName(i)
			fmt.Println(name)
			if name == "FizzBuzzWoofBarkBark!" {
				break //loop break condition
			}
		} else {
			fmt.Println(i)
		}
	}
}

// Checks if number passed is divisible by the array of possible divisors
func isSpecialNumber(n int) bool {
	divisors := [4]int{3, 5, 7, 11}
	for _, d := range divisors {
		if n%d == 0 {
			return true
		}
	}
	return false
}

// Creates the name for our divisible number... i.e 3 is Fizz, 11 is BarkBark.
func makeName(n int) string {
	name := ""
	partsUsed := 0
	if n%3 == 0 {
		name += "Fizz"
		partsUsed++
	}
	if n%5 == 0 {
		name += "Buzz"
		partsUsed++
	}
	if n%7 == 0 {
		name += "Woof"
		partsUsed++
	}
	if n%11 == 0 {
		name += "BarkBark"
		partsUsed++
	}
	if partsUsed == 4 {
		name += "!"
	}
	return name
}
