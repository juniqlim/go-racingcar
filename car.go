package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

func (cs *Cars) print() {
	for _, car := range cs.cars {
		fmt.Printf("%s : ", car.name)
		for i := 0; i < car.position; i++ {
			fmt.Print("-")
		}
		fmt.Println()
	}
	fmt.Println()
}

func NewCars(s string) Cars {
	names := strings.Split(s, ",")
	var cars []Car
	for _, name := range names {
		cars = append(cars, Car{name: name, position: 0})
	}
	return Cars{cars: cars}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("경주할 자동차 이름을 입력하세요(이름은 쉼표(,)를 기준으로 구분).")
	carNames, _ := reader.ReadString('\n')
	carNames = strings.TrimSpace(carNames)
	cars := NewCars(carNames)
	fmt.Println("시도할 회수는 몇회인가요?")
	tryNumberStr, _ := reader.ReadString('\n')
	tryNumberStr = strings.TrimSpace(tryNumberStr)
	tryNumber, _ := strconv.Atoi(tryNumberStr)

	for i := 0; i < tryNumber; i++ {
		cars.move(randomNumber)
		cars.print()
	}
}
