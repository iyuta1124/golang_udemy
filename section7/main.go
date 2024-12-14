package main

import "fmt"

// func Plus(x, y int) int {
// 	return x + y
// }

// func Div(x, y int) (int, int) {
// 	q := x / 2
// 	r := x % y
// 	return q, r
// }

// func Double(price int) (result int) {
// 	result = price * 2
// 	return
// }

func ReturnFunc() func() {
	return func() {
		fmt.Println("aaa")
	}
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func Integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// i := Plus(1, 2)
	// fmt.Println(i)

	// p := Double(2)

	// i2, i3 := Div(9, 3)
	// fmt.Println(i2, i3, p)

	// 無名関数
	// f2 := func(x, y int) int {
	// 	return x + y
	// }(1, 1)
	// fmt.Println(f2)

	// f3 := ReturnFunc()
	// f3()

	// クロージャー
	// f := Later()
	// fmt.Println(f("1"))
	// fmt.Println(f("2"))
	// fmt.Println(f("3"))
	// fmt.Println(f("4"))
	// fmt.Println(f("5"))

	// ジェネレーター
	// goにはPHPなどのように機能として実装されてはいないがクロージャを利用することで再現可能

	f := Integers()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}
