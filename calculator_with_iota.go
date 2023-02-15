package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	default:
		panic("unhandled operation")
	}
}

func main() {
	add := Operation(Add)
	fmt.Println("7+2:", add.calculate(7, 2))

	subtract := Operation(Subtract)
	fmt.Println("7-2:", subtract.calculate(7, 2))

	multiply := Operation(Multiply)
	fmt.Println("7*2:", multiply.calculate(7, 2))

	divide := Operation(Divide)
	fmt.Println("7/2:", divide.calculate(7, 2))
}
