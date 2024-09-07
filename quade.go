package piscine

import "fmt"

func QuadE(x, y int) {
    if x <= 0 || y <= 0 {
        return
    }

    for i := 0; i < y; i++ {
        for j := 0; j < x; j++ {
            if i == 0 && j == 0 {
                // Upper left corner
                fmt.Print("A")
            } else if i == 0 && j == x-1 {
                // Upper right corner
                fmt.Print("C")
            } else if i == y-1 && j == 0 {
                // Lower left corner
                fmt.Print("C")
            } else if i == y-1 && j == x-1 {
                // Lower right corner
                fmt.Print("A")
            } else {
                // The rest of the characters in between
                fmt.Print("B")
            }
        }
        fmt.Println()
    }
}
