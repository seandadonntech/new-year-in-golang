package main

import "fmt"

func main() {
	var year string
	fmt.Println("Enter the year you are in:")
	fmt.Scan(&year)
	if year == "2023" {
		fmt.Println("nice you the entered the right year")
	} else {
		fmt.Println("wth do you live under a rock?")
	}
}
