package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarMove(t *testing.T) {
	car := Car{name: "pobi", position: 0}
	car.move(sixNumber)
	assert.Equal(t, 1, car.position)
}

func TestCarsMove(t *testing.T) {
	cars := Cars{cars: []Car{
		{name: "pobi", position: 0},
		{name: "anna", position: 0},
		{name: "show", position: 0},
	}}
	cars.move(sixNumber)
	assert.Equal(t, 1, cars.cars[0].position)
	assert.Equal(t, "anna", cars.cars[1].name)

	cars.move(randomNumber)
	assert.Contains(t, []int{1, 2}, cars.cars[0].position)
}

func TestStringToCars(t *testing.T) {
	cars := NewCars("pobi,anna,show")
	assert.Equal(t, 3, len(cars.cars))
	assert.Equal(t, 0, cars.cars[0].position)
	assert.Equal(t, "anna", cars.cars[1].name)
}

func TestWinner(t *testing.T) {
	cars := Cars{cars: []Car{
		{name: "pobi", position: 0},
		{name: "anna", position: 1},
		{name: "show", position: 1},
	}}
	assert.Equal(t, 2, len(cars.winner()))
	assert.Equal(t, "anna", cars.winner()[0].name)
}
