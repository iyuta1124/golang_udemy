package main

import "fmt"

const Pi = 99999999999999

// iotaは連番を指定するときに使う
// 下記だと0,1,2となる
const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(c0, c1, c2)
	fmt.Printf("%T\n", Pi)
}
