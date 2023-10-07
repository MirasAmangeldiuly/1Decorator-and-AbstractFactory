package main

import "fmt"

type RestaurantFactory interface {
	CreateFood() Food
	CreateDrink() Drink
}

type Food interface {
	Eat()
}

type Drink interface {
	Drink()
}

type ItalianRestaurantFactory struct{}

func (irf ItalianRestaurantFactory) CreateFood() Food {
	return ItalianFood{}
}

func (irf ItalianRestaurantFactory) CreateDrink() Drink {
	return ItalianDrink{}
}

type JapaneseRestaurantFactory struct{}

func (jrf JapaneseRestaurantFactory) CreateFood() Food {
	return JapaneseFood{}
}

func (jrf JapaneseRestaurantFactory) CreateDrink() Drink {
	return JapaneseDrink{}
}

type ItalianFood struct{}

func (ifd ItalianFood) Eat() {
	fmt.Println("Итальянская еда")
}

type ItalianDrink struct{}

func (idr ItalianDrink) Drink() {
	fmt.Println("Итальянский напиток")
}

type JapaneseFood struct{}

func (jfd JapaneseFood) Eat() {
	fmt.Println("Японская еда")
}

type JapaneseDrink struct{}

func (jdr JapaneseDrink) Drink() {
	fmt.Println("Японский напиток")
}

func main() {

	italianFactory := ItalianRestaurantFactory{}
	italianFood := italianFactory.CreateFood()
	italianDrink := italianFactory.CreateDrink()

	fmt.Println("Итальянский ресторан:")
	italianFood.Eat()
	italianDrink.Drink()

	japaneseFactory := JapaneseRestaurantFactory{}
	japaneseFood := japaneseFactory.CreateFood()
	japaneseDrink := japaneseFactory.CreateDrink()

	fmt.Println("\nЯпонский ресторан:")
	japaneseFood.Eat()
	japaneseDrink.Drink()
}
