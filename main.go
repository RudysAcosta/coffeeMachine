package main

import "fmt"

var coffeeIngredients = map[string]map[string]map[string]float64{
	"latte": {
		"small": {
			"water": 200.00,
			"milk":  50.0,
			"beans": 15,
		},
	},
}

func main() {
	var countCoffee int

	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&countCoffee)

	fmt.Printf("For %d cups of coffee you will need:\n", countCoffee)

	fmt.Printf("%d ml of water\n", int(coffeeIngredients["latte"]["small"]["water"]*float64(countCoffee)))
	fmt.Printf("%d ml of milk\n", int(coffeeIngredients["latte"]["small"]["milk"]*float64(countCoffee)))
	fmt.Printf("%d g of coffee beans", int(coffeeIngredients["latte"]["small"]["beans"]*float64(countCoffee)))
}
