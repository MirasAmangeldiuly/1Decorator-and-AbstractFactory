package main

import "fmt"

type Burger interface {
	GetDescription() string
	GetCost() int
}

type PlainBurger struct{}

func (pb PlainBurger) GetDescription() string {
	return "Обычный бургер"
}

func (pb PlainBurger) GetCost() int {
	return 100
}

type BurgerDecorator interface {
	GetDescription() string
	GetCost() int
}

type CheeseDecorator struct {
	burger Burger
}

func (cd CheeseDecorator) GetDescription() string {
	return cd.burger.GetDescription() + ", с сыром"
}

func (cd CheeseDecorator) GetCost() int {
	return cd.burger.GetCost() + 20
}

type BaconDecorator struct {
	burger Burger
}

func (bd BaconDecorator) GetDescription() string {
	return bd.burger.GetDescription() + ", с беконом"
}

func (bd BaconDecorator) GetCost() int {
	return bd.burger.GetCost() + 30
}

func main() {

	burger := PlainBurger{}
	fmt.Println("Описание:", burger.GetDescription())
	fmt.Println("Цена:", burger.GetCost())

	burgerWithCheese := CheeseDecorator{burger}
	fmt.Println("\nОписание:", burgerWithCheese.GetDescription())
	fmt.Println("Цена:", burgerWithCheese.GetCost())

	burgerWithBacon := BaconDecorator{burger}
	fmt.Println("\nОписание:", burgerWithBacon.GetDescription())
	fmt.Println("Цена:", burgerWithBacon.GetCost())

	burgerWithCheeseAndBacon := BaconDecorator{burgerWithCheese}
	fmt.Println("\nОписание:", burgerWithCheeseAndBacon.GetDescription())
	fmt.Println("Цена:", burgerWithCheeseAndBacon.GetCost())
}
