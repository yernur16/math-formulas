package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		log.Print("Error, no argument!")
		return
	} else if len(arg) > 1 {
		log.Print("Error, more arguments than needed!")
		return
	}

	dataByte, err := ioutil.ReadFile(arg[0])
	if err != nil {
		log.Fatalf("Error, missed argument %v ", err.Error())
	}
	var numbers []float64

	strFromBytes := strings.Fields(string(dataByte))
	for _, j := range strFromBytes {
		conversion, err := strconv.ParseFloat(j, 64)
		if err != nil {
			log.Print("Error, in the appending data into array")
			return
		}
		numbers = append(numbers, conversion)

	}

	fmt.Println("Average:", int64(math.Round(Average(numbers))))
	fmt.Println("Median:", int64(math.Round(Median(numbers))))
	fmt.Println("Variance:", int64(math.Round(Variance(numbers))))
	fmt.Println("Standard Deviation:", int64(math.Round(StDev(Variance(numbers)))))
}

func Average(n []float64) float64 {
	total := 0.0
	for i := 0; i < len(n); i++ {
		total += n[i]
	}
	return total / float64(len(n))
}

func Median(n []float64) float64 {
	sort.Float64s(n)
	if len(n)%2 == 1 {
		return (n[len(n)/2])
	} else {
		return (0.5 * (n[len(n)/2-1] + n[len(n)/2]))
	}
}

func Variance(n []float64) float64 {
	xMean := Average(n)

	array := []float64{}
	for i := 0; i < len(n); i++ {
		array = append(array, n[i]-xMean)
	}
	for i := 0; i < len(array); i++ {
		array[i] *= array[i]
	}
	return (Average(array))
}

func StDev(n float64) float64 {
	return math.Sqrt((n))
}
