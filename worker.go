package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Workplace struct {
	gorm.Model
	Name    string  `gorm:"size:50;not null"`
	Address string  `gorm:"size:255;not null"`
	Phone   *string `gorm:"size:20"`
}

type Worker struct {
	gorm.Model
	WorkplaceID uint `gorm:"not null"`
	Workplace   *Workplace
	Name        string    `gorm:"size:61;not null"`
	Birthday    time.Time `gorm:"type:datetime"`
	Phone       *string   `gorm:"size:20"`
	Recipes     []*Recipe `gorm:"many2many:worker_recipes"`
}

func insertWorker(db *gorm.DB) {
	workplace := Workplace{}
	db.Find(&workplace, `name = ?`, "Workplace Two")

	worker := Worker{
		Name:      "Donald Trump",
		Birthday:  time.Date(1946, 6, 14, 12, 0, 0, 0, time.UTC),
		Recipes:   []*Recipe{},
		Workplace: &workplace,
	}
	db.Save(worker)
}

func insertWorkplace(db *gorm.DB) {
	workplace := &Workplace{
		Name:    "Workplace Two",
		Address: "Evergreen Terrace 742nd",
		Phone:   BoxString("(56) 123-4789"),
	}
	db.Save(workplace)
}

func ListEverything(db *gorm.DB) {
	workplaces := []Workplace{}
	db.Find(&workplaces)
	for _, workplace := range workplaces {
		fmt.Printf("Workplace data: %v\n", workplace)

		workers := []Worker{}
		db.Preload("Recipes").Preload("Recipes.Toppings").Find(&workers)

		for _, worker := range workers {
			fmt.Printf("Worker data: %v\n", worker)
			for _, recipe := range worker.Recipes {
				fmt.Printf("Recipe data: %v\n", recipe)
				for _, topping := range recipe.Toppings {
					fmt.Printf("Topping data: %v\n", topping)
				}
			}
		}
	}
}

func deleteAll(db *gorm.DB) {
	err1 := db.Delete(&Workplace{}).Error
	fmt.Printf("Deleting the records:\n%v\n", err1)
}

func deleteWorker(db *gorm.DB) {
	err := db.Delete(&Worker{}, `name = ? and age = ?`, "john", 34).Error
	fmt.Printf("Deleting the records:\n%v\n", err)
}
