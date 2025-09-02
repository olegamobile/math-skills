package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	data := read_data()
	fmt.Println("Average: ", int(math.Round(average(data))))
	fmt.Println("Median: ", int(math.Round(median(data))))
	fmt.Println("Variance: ", int(math.Round(variance(data))))
	fmt.Println("Standard Deviation: ", int(math.Round(std_dev(data))))

}

func average(numbers []float64) float64 {
	// x̄ = (1 / n) * Σ xi

	if len(numbers) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range numbers {
		sum += v
	}
	return sum / float64(len(numbers))
}

func median(numbers []float64) float64 {
	// Me = x((n + 1) / 2) for odd aount of numbers
	// Me = (x(n / 2) + x(n / 2 + 1)) / 2 for even amount of numbers

	n := len(numbers)
	if n == 0 {
		return 0
	}
	sorted := make([]float64, n)
	copy(sorted, numbers)
	
	for i := 1; i < n; i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j] > key {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}
	if n%2 == 1 {
		return sorted[n/2]
	}
	return (sorted[n/2-1] + sorted[n/2]) / 2
}

func variance(numbers []float64) float64 {
	// σ^2 = (1 / N) * Σ (xi - μ)^2

	if len(numbers) == 0 {
		return 0
	}
	avg := average(numbers)
	sum := 0.0
	for _, v := range numbers {
		diff := v - avg
		sum += diff * diff
	}
	return sum / float64(len(numbers))
}

func std_dev(numbers []float64) float64 {
	// σ = sqrt( (1 / N) * Σ (xi - μ)^2 )

	return math.Sqrt(variance(numbers))
}


func read_data() []float64 {
	if len(os.Args) < 2 {
		panic("No input file provided")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %v", err))
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			panic(fmt.Sprintf("Invalid data in file: %v", err))
		}
		numbers = append(numbers, val)
	}
	if len(numbers) == 0 {
		panic("No valid data found in file")
	}
	return numbers
}
