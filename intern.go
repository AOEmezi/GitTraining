package main

import "fmt"

type UserData struct {
	a int
	b int
	c int
	f float64
}

func main() {
//List of options for user to choose from via input.
	fmt.Printf("SUM = A")
	fmt.Printf("FACTORIAL = B")
	fmt.Printf("CONCATENATES STRINGS = C")

	fmt.Scanf("%v")

	if {
		fmt.Scanf("%v") = A
		sum()
	} else {
		fmt.Scanf("%v") = B
		factorial()
	} else {
		fmt.Scanf("%v") = C
		concatenates()
	}
		
	for {
		a, b := userInput()
		f:= factorial()

		if a && b = int {
			c:=sum()
			fmt.Printf("%v + %v = %v\n", a, b, c)
			break
		}
	} else _{
		if f != float64 {
			fmt.Printf("factorial(%v)" ,f)
		}
	}	
}

func userInput(a, b int) {
	var userData = UserData{
		a: a,
		b: b,
		f: f,
	}

}

//Create a function in Go that takes two integers as parameters and returns their sum.
func sum() int {
	fmt.Println("Enter your first number: ") //asks for first number
	fmt.Scan(&a)                             //collects input from user

	fmt.Println("Enter your second number: ") //asks for second number
	fmt.Scan(&b)                              //collects input from user

	c := a + b
	return c

}

//Create a function to calculate the factorial of a number
func factorial() float64 {
	for i := 1; i < n; i++ {
		f := n * (n-1)
	}
	return f
}

//Create a function that accepts a variable number of strings and concatenates them into a single string.
func concatenates(firstName string, lastName string, otherNames string) string {
    fullName := fmt.Sprintf("%v %v %v", firstName, lastName, otherNames)
    return fullName
}
