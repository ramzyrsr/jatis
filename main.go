package main

import (
	"fmt"
	"unicode/utf8"
)

// number 3
func fibonaci(x, y int) {
	for n := 2; n < 1000000; n++ {
		if (x+n)%2 == 0 {
			fmt.Println("Sum of even = ", x+n)
		}
		if (y+n)%2 == 1 {
			fmt.Println("Sum of odd = ", y+n)
		}
	}
}

// number 4
func charOccurence(str string) {
	fmt.Println("counts =", utf8.RuneCountInString(str))
}

// number 7
func royalty(spent int, voucher string) {
	var countAward = spent / 1000000
	var totalAward int

	for i := 0; i < countAward; i++ {
		totalAward += 10000
	}
	fmt.Println("total award: ", totalAward)
}

// number 8
func Dice() {
	var probabilty int

	// dice 1
	for i := 0; i < 6; i++ {
		// dice 2
		for j := 0; j < 6; j++ {
			// dice 3
			for k := 0; k < 6; k++ {
				if (i+j+k) > 2 && (i+j+k) < 19 {
					probabilty++
				}
			}
		}
	}
	fmt.Println("probabilty: ", probabilty)
}

func main() {
	fibonaci(2, 3)

	charOccurence("hello")

	royalty(3000000, "test")

	Dice()
}

// number 10
// getting your windows cleaned approximately
// every three to six months will guarantee you a
// beautifully clean and clear view.
