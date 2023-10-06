package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// switch function
	switchexample()

	// weekDays function
	weekDays()

}

func switchexample() {
	fmt.Print("your local machine: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux ")
	default:
		fmt.Printf("%s. \n", os)
	}
}

func weekDays() {
	fmt.Println("when's saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
