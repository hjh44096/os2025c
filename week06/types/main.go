package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(math.Round(2.21))
	fmt.Println(math.Ceil(2.21))
	fmt.Println(strings.Title("go developer!"))
	fmt.Println("kim\ninha\t\"\\")
	fmt.Println("2", "가")
	fmt.Println(reflect.TypeOf(2.31))
	fmt.Println(reflect.TypeOf("Kim inha"))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf(18))
}
