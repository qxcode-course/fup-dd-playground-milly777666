package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Scanln(&name)
	fmt.Scanln(&age)

	var classification string

	if age < 12 {
		classification = "crianca"
	} else if age < 18 {
		classification = "jovem"
	} else if age < 65 {
		classification = "adulto"
	} else if age < 1000 {
		classification = "idoso"
	} else {
		classification = "mumia"
	}

	fmt.Printf("%s eh %s\n", name, classification)
}
