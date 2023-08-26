package main

import (
	"fmt"
	//"strings"
)

func main() {
	var number, val, ch int
	number = 50
	var list []string

	//Global loop
	for {

		var firstname, lastname string
		fmt.Printf("Enter first name \n")
		fmt.Scan(&firstname)
		fmt.Printf("enter last  name \n")
		fmt.Scan(&lastname)
		fmt.Printf("Enter number of ticket \n")
		fmt.Scan(&val)
		number = cal(number, val)
		list = append(list, firstname+" "+lastname)
		fmt.Printf("Thank you %v %v for booking %v ticket\n", firstname, lastname, val)
		fmt.Printf("%v remaining ticket for %v\n", number, val)
		if number <= 0 {
			fmt.Printf("all tickets are sold\n")
			continue
		}
		fmt.Println("if you want to show the whole list press 1 ,0 to continue")
		fmt.Scan(&ch)
		if ch == 1 {
			show(list)
		}

	}

}

func cal(n int, v int) int {
	n = n - v
	return n
}
func show(list []string) {
	fmt.Printf("these are all our booking:%v  \n", list)
}
