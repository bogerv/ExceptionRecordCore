package main

import (
	"fmt"
	"reflect"
)

// zengjia
func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func main() {
	f := Adder()
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(f(1))
	fmt.Println(f(10))
	fmt.Println(f(100))
}
