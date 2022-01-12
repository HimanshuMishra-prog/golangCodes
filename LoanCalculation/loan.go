package main

import (
	"fmt"
	simple "hello/Assignment/github.com/SimpleInterest"
)

func main() {
	var p float32
	var r int
	var t float32
	fmt.Print("Enter the principal")
	fmt.Scanf("%f", &p)
	fmt.Print("Enter the rate ")
	fmt.Scanf("%d", &r)
	fmt.Println("Enter the number of years")
	fmt.Scanf("%f", &t)
	interest := simple.IntersetCalculation(p, r, t)

	fmt.Println("Your interest is : ", interest)

	if interest > 1000 {
		fmt.Println("Your loan is approved")
	} else {
		fmt.Println("Your loan is rejected")
	}
}
