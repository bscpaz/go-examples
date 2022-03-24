package main

import "fmt"

// The instruction "~" tells to make a approximations of values
type Number interface {
	~int | ~float64
}

//"MyNumber" is compatible with "Number" interface as long as this has "~" instruction.
type MyNumber int

func addVerboseGenerics[T int | float64](numberA T, numberB T) T {
	return numberA + numberB
}

func addGenerics[T Number](numberA T, numberB T) T {
	return numberA + numberB
}

func main() {
	//Test #1
	resultInt := addVerboseGenerics(5, 2)
	fmt.Printf("addVerboseGenerics with int = %v\n\n", resultInt)

	resultFloat := addVerboseGenerics(5.8, 2.9)
	fmt.Printf("addVerboseGenerics with float = %v\n\n", resultFloat)

	//Test #2
	resultInt = addGenerics(5, 2)
	fmt.Printf("addGenerics with int = %v\n\n", resultInt)

	resultFloat = addGenerics(5.8, 2.9)
	fmt.Printf("addGenerics with float = %v\n\n", resultFloat)

	//Test #3
	var myNumberA, myNumberB, myResult MyNumber
	myNumberA = 10
	myNumberB = 15
	myResult = addGenerics(myNumberA, myNumberB)
	fmt.Printf("addGenerics with my 'MyNumber'('~' instruction) = %v\n\n", myResult)

}
