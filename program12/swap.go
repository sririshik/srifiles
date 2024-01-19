package main

import "fmt"

func Swap_num(x, y string) (string, string) {
	return y, x

}
func main() {

	a, b := Swap_num("hello", "world")
	fmt.Println(a, b)

}
