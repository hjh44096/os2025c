package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Round(2.21))
	fmt.Println(math.Ceil(2.21))
	fmt.Println(strings.Title("go developer!"))
	fmt.Println("kim\ninha\t\"\\")
	fmt.Println("2","가")
}

import (
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOF(2.31))
	fmt.Println(reflect.TypeOF("Kim inha"))
	fmt.Println(reflect.TypeOF(true))
	fmt.Println(reflect.TypeOF('A'))
	fmt.Println(reflect.TypeOF(19))
}