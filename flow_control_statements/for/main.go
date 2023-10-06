package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	anotherexample()
}

func anotherexample() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

/*
for {
}
if you omit the loop condition it loops forever, so an infinite loop is compactly expressed
*/
