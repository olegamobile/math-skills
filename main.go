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

func average(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}

func median(nums []float64) float64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sorted := make([]float64, n)
	copy(sorted, nums)
	// Simple insertion sort
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

func variance(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	avg := average(nums)
	sum := 0.0
	for _, v := range nums {
		diff := v - avg
		sum += diff * diff
	}
	return sum / float64(len(nums))
}

func std_dev(nums []float64) float64 {
	return sqrt(variance(nums))
}

func sqrt(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// read_data reads float64 numbers from a file whose name is given as the first command-line argument.
// It returns a slice of float64 values.
func read_data() []float64 {
	if len(os.Args) < 2 {
		panic("No input file provided")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %v", err))
	}
	defer file.Close()

	var nums []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			panic(fmt.Sprintf("Invalid data in file: %v", err))
		}
		nums = append(nums, val)
	}
	if len(nums) == 0 {
		panic("No valid data found in file")
	}
	return nums
}
