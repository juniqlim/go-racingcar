package main

import (
	"math/rand"
	"time"
)

type Car struct {
	name     string
	position int
}

func (c *Car) move(numberGenerator func() int) {
	if numberGenerator() > 5 {
		c.position = c.position + 1
	}
}

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}

func sixNumber() int {
	return 6
}

type Cars struct {
	cars []Car
}

func (cs *Cars) move(numberGenerator func() int) {
	for i := range cs.cars {
		cs.cars[i].move(numberGenerator)
	}
}
