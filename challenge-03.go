package main

import "fmt"

func main() {
	
	kalimat := "Selamat malam! Wahai Rakyat Manusia"

	// Hitung kemunculan setiap huruf
	mapHuruf := make(map[string]int)
	for _, h := range kalimat {
		mapHuruf[string(h)]++
		fmt.Println(string(h))
	}

	fmt.Println()
	fmt.Println(mapHuruf)
}