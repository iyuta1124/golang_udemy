package main

import "fmt"

const Pi = 99999999999999

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(c0, c1, c2)
	fmt.Printf("%T\n", Pi)
}
