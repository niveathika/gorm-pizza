package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Name     string     `gorm:"size:50;not null"`
	Workers  []*Worker  `gorm:"many2many:worker_recipes"`
	Toppings []*Topping `gorm:"many2many:recipe_toppings"`
}

type Topping struct {
	gorm.Model
	Name   string    `gorm:"size:20;not null"`
	Recipe []*Recipe `gorm:"many2many:recipe_toppings"`
}

type Size struct {
	gorm.Model
	Name string `gorm:"size:20;not null"`
}

type Pizza struct {
	gorm.Model
	RecipeID uint `gorm:"not null;unique_indez:pizzas"`
	Recipe   Recipe
	SizeID   uint `gorm:"not null;unique_indez:pizzas"`
	Size     Size
	Price    float64 `gorm:"not null;type:decimal(10,2)"`
}

func getPizza(db *gorm.DB) Pizza {

	pizza := Pizza{}
	db.Debug().Preload("Recipe").Preload("Size").Preload("Recipe.Workers").Preload("Recipe.Toppings").Find(&pizza, 1)
	fmt.Printf("Pizza (id = 1):\n%+v\n", pizza)

	return pizza
}

func getRecipe(db *gorm.DB) Recipe {
	recipe := Recipe{}
	db.Debug().Preload("Workers").Preload("Toppings").Find(&recipe, `name=?`, "Mozzarella")
	fmt.Printf("Receipe (Name = Mozzarella):\n%+v\n", recipe)
	for _, topping := range recipe.Toppings {
		fmt.Printf("Toppings : \n%+v\n", topping)
	}
	return recipe
}

func getRecipesByTopping(db *gorm.DB) []Recipe {

	onionTopping := Topping{}
	db.Debug().Find(&onionTopping, `name=?`, "Onions")
	fmt.Printf("Toppings : \n%+v\n\n\n", onionTopping)

	recipes := []Recipe{}
	db.Debug().Where(`id IN (SELECT recipe_id FROM recipe_toppings WHERE topping_id = ?)`, onionTopping.ID).Preload("Toppings").Find(&recipes)
	for _, r := range recipes {
		fmt.Printf("Receipe: \n%+v\n", r)
		for _, topping := range r.Toppings {
			fmt.Printf("Toppings : \n%+v\n", topping)
		}
	}
	return recipes
}
