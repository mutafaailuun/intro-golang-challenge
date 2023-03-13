package main

import "fmt"

func main() {
	i := 21
	j := true
	k := 15
	l := "\u042F"
	russianUnicode := 1071
	var float float64 = 123.456
	
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Println(j)
	fmt.Println(l)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", k)
	fmt.Printf("%X \n", k)
	fmt.Printf("%U \n", russianUnicode)
	fmt.Printf("%f\n", float)
	fmt.Printf("%e\n", float)

}