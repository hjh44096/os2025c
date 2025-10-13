package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var month string = now.Month().String()
	var day int = now.Day()
	fmt.Println(month, day)

	//univ := "Go$ Inha$"
	changer := strings.NewReplacer("$", "!")
	//changed := changer.Replace("univ")
	changed := changer.Replace("Go$ Inha$")
	fmt.Println(changed)
}
