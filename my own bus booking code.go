package main

import (
	"fmt"
)

func search(city string, v string) {

	date := make(map[string]int)
	date["chennai"] = 12
	date["bangalore"] = 13
	date["chennai"] = 14
	date["bangalore"] = 15
	date["tirupathi"] = 12
	date["tirupathi"] = 13
	date["tirupathi"] = 14
	date["hyderabad"] = 15
	trip := make(map[string]string)
	trip["tirupathi"] = "tirupathi"
	trip["hyderabad"] = "chennai"
	trip["bangalore"] = "hyderabad"
	for i, v := range trip {
		if v == city {

			for j := 0; j < 1; j++ {
				if date[i] == date[v] {
					fmt.Printf(" the bus is available from %s to %s with dates %d\n", i, v, date[i])
				} else {
					fmt.Println("dates not available")
				}
			}

		} else {
			fmt.Println("not available")
		}
	}
}

func main() {

	var v string
	fmt.Println("enter the from city")
	fmt.Scanln(&v)
	var t string
	fmt.Println("Enter the to city")
	fmt.Scanln(&t)
	search(t, v)
}
