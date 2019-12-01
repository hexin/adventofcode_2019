package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbs := readInput()
	var result int64 = 0
	for _, v := range numbs {
		result += int64(calcFuel(v))
	}
	fmt.Println(result)
}

func calcFuel(mass int) int {
	result := (mass / 3) - 2
	if result > 0 {
		return result + calcFuel(result)
	} else {
		return 0
	}
}

func readInput() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result []int
	for scanner.Scan() {
		asInt, _ := strconv.Atoi(scanner.Text())
		result = append(result, asInt)
	}
	return result
}
