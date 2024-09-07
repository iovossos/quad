package piscine

import (
	"github.com/01-edu/z01"
)

/*
	func QuadA(x, y int) {
		if x <= 0 || y <= 0 {
			return
		}

		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
					fmt.Print("o")
				} else if i == 0 || i == y-1 {
					fmt.Print("-")
				} else if j == 0 || j == x-1 {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
*/
func QuadB(q, w int) {
	if q <= 0 || w <= 0 {
		return
	}
	for i := 0; i < w; i++ {
		for j := 0; j < q; j++ {
			if (i == 0 || i == w-1) && (j == 0 || j == q-1) {
				z01.PrintRune('0')
			} else if i == 0 || i == w-1 {
				z01.PrintRune('-')
			} else if j == 0 || j == w-1 {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
