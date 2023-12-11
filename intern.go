package main

import "fmt"

type UserData struct {
	a int
	b int
	c int
	f float64
}

func main() {
	// List of options for the user to choose from via input.
	fmt.Println("Select an option:")
	fmt.Println("SUM = A | FACTORIAL = B | CONCATENATES STRINGS = C")

	var choice string
	fmt.Scan(&choice)

	switch choice {
	case "A":
		var a, b int

		fmt.Println("Enter your first number: ")
		fmt.Scan(&a)
		fmt.Println("Enter your second number: ")
		fmt.Scan(&b)

		c := sum(a, b)
		fmt.Printf("%v + %v = %v\n", a, b, c)

	case "B":
		var n int
		fmt.Println("Enter a number to calculate its factorial: ")
		fmt.Scan(&n)

		f := factorial(n)
		fmt.Printf("Factorial of %v: %v\n", n, f)

	case "C":
		var firstName, lastName, otherNames string
		fmt.Println("Enter first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter other names: ")
		fmt.Scan(&otherNames)
		result := concatenates(firstName, otherNames, lastName)
		fmt.Println("Concatenated String:", result)
	default:
		fmt.Println("Invalid choice")
	}

}
func sum(a, b int) int {
	return a + b
}
func factorial(n int) float64 {
	f := 1.0
	for i := 1; i <= n; i++ {
		f *= float64(i)
	}
	return f
}
func concatenates(firstName, lastName, otherNames string) string {
	fullName := fmt.Sprintf("%v %v %v", firstName, lastName, otherNames)
	return fullName
}
