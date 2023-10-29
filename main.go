package main

import (
	"fmt"
)

type Command interface {
	Execute() float64
	Undo() float64
}

type Calculator struct {
	result float64
}

type AddCommand struct {
	calculator *Calculator
	value      float64
}

func (c *AddCommand) Execute() float64 {
	c.calculator.result += c.value
	return c.calculator.result
}

func (c *AddCommand) Undo() float64 {
	c.calculator.result -= c.value
	return c.calculator.result
}

type SubtractCommand struct {
	calculator *Calculator
	value      float64
}

func (c *SubtractCommand) Execute() float64 {
	c.calculator.result -= c.value
	return c.calculator.result
}

func (c *SubtractCommand) Undo() float64 {
	c.calculator.result += c.value
	return c.calculator.result
}

type CalculatorCommandInvoker struct {
	commands   []Command
	current    int
	calculator *Calculator
}

func (ci *CalculatorCommandInvoker) Execute(cmd Command, value int) float64 {
	result := cmd.Execute()
	ci.commands = append(ci.commands[:ci.current], cmd)
	ci.current++
	return result
}

func (ci *CalculatorCommandInvoker) Undo() float64 {
	if ci.current > 0 {
		cmd := ci.commands[ci.current-1]
		result := cmd.Undo()
		ci.current--
		return result
	}
	return ci.calculator.result
}

func (ci *CalculatorCommandInvoker) Redo() float64 {
	if ci.current < len(ci.commands) {
		cmd := ci.commands[ci.current]
		result := cmd.Execute()
		ci.current++
		return result
	}
	return ci.calculator.result
}

func main() {
	calculator := &Calculator{}
	addCommand := &AddCommand{calculator, 5}
	subtractCommand := &SubtractCommand{calculator, 3}

	commandInvoker := &CalculatorCommandInvoker{calculator: calculator}

	fmt.Println("Initial Result:", calculator.result)
	result := commandInvoker.Execute(addCommand, 5)
	fmt.Println("After Adding 5:", result)
	result = commandInvoker.Execute(subtractCommand, 3)
	fmt.Println("After Subtracting 3:", result)
	fmt.Println("Undo:", commandInvoker.Undo())
	fmt.Println("Redo:", commandInvoker.Redo())
}
