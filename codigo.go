package main

import (
	"errors"
	"fmt"
)

func withdrawNotes(amount int) ([]int, error) {
	notes := []int{}
	balance := 500

	if amount <= 0 || amount > balance {
		return notes, errors.New("Não é possível fazer o saque desse valor.")
	}

	for amount > 0 {
		if amount >= 100 {
			notes = append(notes, 100)
			amount -= 100
		} else if amount >= 50 {
			notes = append(notes, 50)
			amount -= 50
		} else if amount >= 20 {
			notes = append(notes, 20)
			amount -= 20
		} else if amount >= 10 {
			notes = append(notes, 10)
			amount -= 10
		} else if amount >= 5 {
			notes = append(notes, 5)
			amount -= 5
		} else if amount >= 2 {
			notes = append(notes, 2)
			amount -= 2
		} else {
			return notes, errors.New("Notas não disponiveis para sacar este valor")
		}
	}

	return notes, nil
}

func main() {
	amount := 400
	notes, err := withdrawNotes(amount)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Notas para o número %d: %v\n", amount, notes)
	}
}
