package main

import (
	"fmt"
)

func main() {
	
	for i := 0; i <= 4; i++ {
		fmt.Printf("Nilai i = %d \n", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
						for index, char := range "САШАРВО" {
				fmt.Printf("character %#U starts at byte position %d\n", char, index)
			}
			continue
		}
		fmt.Printf("Nilai j = %d \n", j)
	}
}
