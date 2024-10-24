package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var coffeeIngredients = map[int]map[string]map[string]float64{
	1: {
		"small": {
			"water": 250,
			"milk":  0,
			"beans": 16,
			"count": 4,
		},
	},
	2: {
		"small": {
			"water": 350,
			"milk":  75,
			"beans": 20,
			"count": 7,
		},
	},
	3: {
		"small": {
			"water": 200,
			"milk":  100,
			"beans": 12,
			"count": 6,
		},
	},
}

var inventarioCafe = map[string]float64{
	"water":    400,
	"milk":     540,
	"beans":    120,
	"cups":     9,
	"cash_box": 550,
}

func main() {
	for {
		action := menu()

		if action == "buy" {
			buy()
		} else if action == "fill" {
			fill()
		} else if action == "take" {
			take()
		} else if action == "remaining" {
			remaining()
		} else {
			exit()
		}
	}
}

func calculatePossibleCups() int {

	var possibleCups int = -1

	for item, count := range inventarioCafe {
		if item != "cups" {
			cupsCoffee := int(math.Floor(count / coffeeIngredients[1]["small"][item]))

			if cupsCoffee < possibleCups || possibleCups == -1 {
				possibleCups = cupsCoffee
			}
		}
	}

	return possibleCups
}

func buy() {
	fmt.Println()
	var choose string

	for {
		fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
		fmt.Scan(&choose)

		if choose == "back" {
			return
		}

		if num, err := strconv.Atoi(choose); err == nil && num >= 1 && num <= 3 {
			break
		}
	}

	num, _ := strconv.Atoi(choose)
	isMakedCoffee, item := makeCoffee(num)

	if !isMakedCoffee {
		fmt.Printf("Sorry, not enough %s!\n\n", item)
	} else {
		fmt.Printf("I have enough resources, making you a coffee!\n\n")
	}
}

func makeCoffee(typeCoffee int) (bool, string) {
	coffee := coffeeIngredients[typeCoffee]["small"]

	if inventarioCafe["water"] < coffee["water"] {
		return false, "water"
	}

	if inventarioCafe["milk"] < coffee["milk"] {
		return false, "milk"
	}

	if inventarioCafe["beans"] < coffee["beans"] {
		return false, "beans"
	}

	if inventarioCafe["cups"] == 0 {
		return false, "cups"
	}

	inventarioCafe["water"] -= coffee["water"]
	inventarioCafe["milk"] -= coffee["milk"]
	inventarioCafe["beans"] -= coffee["beans"]
	inventarioCafe["cups"] -= 1
	inventarioCafe["cash_box"] += coffee["count"]

	return true, ""
}

func fill() {
	var water, milk, beans, cups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&water)

	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milk)

	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&beans)

	fmt.Printf("Write how many disposable cups you want to add:\n\n")
	fmt.Scan(&cups)

	inventarioCafe["water"] += float64(water)
	inventarioCafe["milk"] += float64(milk)
	inventarioCafe["beans"] += float64(beans)
	inventarioCafe["cups"] += float64(cups)
}

func take() {
	fmt.Printf("\nI gave you $%d\n\n", int(inventarioCafe["cash_box"]))
	inventarioCafe["cash_box"] = 0
}

func remaining() {
	fmt.Println()
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", int(inventarioCafe["water"]))
	fmt.Printf("%d ml of milk\n", int(inventarioCafe["milk"]))
	fmt.Printf("%d g of coffee beans\n", int(inventarioCafe["beans"]))
	fmt.Printf("%d disposable cups\n", int(inventarioCafe["cups"]))
	fmt.Printf("$%d of money\n\n", int(inventarioCafe["cash_box"]))
}

func exit() {
	os.Exit(0)
}

func menu() string {
	var action string

	for {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scan(&action)

		if action == "buy" || action == "fill" || action == "take" || action == "remaining" || action == "exit" {
			break
		}
	}

	return action
}
