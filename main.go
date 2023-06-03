package main

import (
	"fmt"
	"goBootcamp/doctor"
)

func main() {
	fmt.Println("Hello, World!")

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)
}