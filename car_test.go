package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarMove(t *testing.T) {
	car := Car{name: "pobi", position: 0}
	car.move(sixNumber)
	assert.Equal(t, car.position, 1)
}
