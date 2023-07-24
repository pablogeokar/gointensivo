package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pablogeokar/gointensivo/internal/infra/database"
	"github.com/pablogeokar/gointensivo/internal/usecase"
)

type Car struct {
	Model string
	Color string
}

// método
func (c Car) Start() {
	println(c.Model + " has been started")
}

// função
func soma(x, y int) int {
	return x + y
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)

	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
