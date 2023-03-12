package main

import "fmt"

func main() {
	i := 21
	j := true
	k := 15
	l := "\u042F"
	m := 1071
	var n float64 = 123.456
	
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Println(j)
	fmt.Println(l)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", k)
	fmt.Printf("%X \n", k)
	fmt.Printf("%U \n", m)
	fmt.Printf("%f\n", n)
	fmt.Printf("%e\n", n)

}