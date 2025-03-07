package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DeliveryRequest struct{
	TotalRest float64  `json:"total_rest"`
	Distance float64 `json:"distance_km"`
	Weight float64  `json:"weight_kg"`
	Priority string  `json:"priority"`
}

func main() {
	fmt.Println("Server is ruunning on port 8080")

	s := echo.New()

	s.POST("/calculate", Handler)

	err := s.Start(":8080")

	if err != nil{
		log.Fatal(err)
	}
}



func Handler(ctx echo.Context) error{
	var req DeliveryRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	price := calculatePrice(req)
	return ctx.JSON(http.StatusOK, map[string]float64{"price": price})
}


func calculatePrice(req DeliveryRequest) float64 {
// Get total price from the restaurant
    distance := 20.0 // RUB per km
    weight := 10.0 // RUB per kg
    priorityMultiplier := 1.0

    if req.Priority == "fast" {
        priorityMultiplier = 1.5
    }

    // Apply priority multiplier to the entire total
    price := (req.TotalRest + (req.Distance * distance) + (req.Weight * weight)) * priorityMultiplier

    return price
}