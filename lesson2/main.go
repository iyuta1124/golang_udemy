package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println((i))

	var s string = "hello go"
	fmt.Println((s))

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "hello go2"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s4 string
	fmt.Println(i3, s4)

	i4 := 400
	fmt.Println(i4)
}
