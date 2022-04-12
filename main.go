package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(mysql.Open("root:Test@123@/pizzas?parseTime=true"), &gorm.Config{})
}

func BoxString(x string) *string {
	return &x
}

func main() {
	if db, err := Connect(); err != nil {
		fmt.Printf("Dude! I could not connect to the database. This happened: %s. Please fix everything and try again", err)
	} else {
		fmt.Println("Database connection was successful. Enjoy.")
		// ClearEverything(db)
		// Migrate(db)
		// Seed(db)
		// ListEverything(db)
		//_ = getPizza(db)
		//_ = getRecipe(db)
		_ = getRecipesByTopping(db)
	}
}

func Migrate(db *gorm.DB) {
	workplacePrototype := &Workplace{}
	workerPrototype := &Worker{}
	recipePrototype := &Recipe{}
	sizePrototype := &Size{}
	pizzaPrototype := &Pizza{}
	toppingsPrototype := &Topping{}
	db.AutoMigrate(workplacePrototype, workerPrototype, recipePrototype, sizePrototype, pizzaPrototype, toppingsPrototype)
}

func Seed(db *gorm.DB) {
	cheese := &Topping{
		Name: "Cheese",
	}
	tomatoeSauce := &Topping{
		Name: "Tomatoe Sauce",
	}
	onions := &Topping{
		Name: "Onions",
	}
	tomatoeSlices := &Topping{
		Name: "Tomatoe Slices",
	}
	hamSlices := &Topping{
		Name: "Ham Slices",
	}
	pepperoni := &Topping{
		Name: "Pepperoni",
	}
	recipe1 := &Recipe{
		Name: "Mozzarella",
		Toppings: []*Topping{
			tomatoeSauce,
			cheese,
		},
	}
	recipe2 := &Recipe{
		Name: "Onions",
		Toppings: []*Topping{
			onions,
			cheese,
		},
	}
	recipe3 := &Recipe{
		Name: "Napolitan",
		Toppings: []*Topping{
			tomatoeSauce,
			tomatoeSlices,
			cheese,
			hamSlices,
		},
	}
	recipe4 := &Recipe{
		Name: "Pepperoni",
		Toppings: []*Topping{
			tomatoeSauce,
			pepperoni,
			cheese,
		},
	}
	db.Debug().Save(recipe1)
	db.Save(recipe2)
	sizePersonal := Size{
		Name: "Personal",
	}
	sizeSmall := Size{
		Name: "Small",
	}
	sizeMedium := Size{
		Name: "Medium",
	}
	sizeBig := Size{
		Name: "Big",
	}
	sizeExtraBig := Size{
		Name: "Extra Big",
	}
	db.Save(&sizePersonal)
	db.Save(&sizeSmall)
	db.Save(&sizeMedium)
	db.Save(&sizeBig)
	db.Save(&sizeExtraBig)
	workplace1 := &Workplace{
		Name:    "Workplace One",
		Address: "Fake st. 123rd",
	}
	workplace2 := &Workplace{
		Name:    "Workplace Two",
		Address: "Evergreen Terrace 742nd",
		Phone:   BoxString("(56) 123-4789"),
	}

	worker1 := Worker{
		Name:     "Mauricio Macri",
		Birthday: time.Date(1959, 2, 8, 12, 0, 0, 0, time.UTC),
		Recipes: []*Recipe{
			recipe1, recipe2, recipe3,
		},
		Workplace: workplace2,
	}

	worker2 := Worker{
		Name:     "Donald Trump",
		Birthday: time.Date(1946, 6, 14, 12, 0, 0, 0, time.UTC),
		Recipes: []*Recipe{
			recipe1, recipe2, recipe4,
		},
		Workplace: workplace2,
	}

	for sizeIndex, size := range []Size{sizePersonal, sizeSmall, sizeMedium, sizeBig, sizeExtraBig} {
		for recipeIndex, recipe := range []*Recipe{recipe1, recipe2, recipe3, recipe4} {
			db.Save(&Pizza{
				Size:   size,
				Recipe: *recipe,
				Price:  (float64(0.1)*float64(recipeIndex+1) + 1) + float64(sizeIndex)*5,
			})
		}
	}
	db.Save(workplace1)
	db.Save(workplace2)
	db.Save(worker1)
	db.Save(worker2)

	fmt.Printf("Workplaces created:\n%v\n%v\n", workplace1, workplace2)

	fmt.Printf("Recipes created:\n%v\n", []*Recipe{recipe1, recipe2, recipe3, recipe4})
}
