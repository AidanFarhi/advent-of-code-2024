package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	l1, l2 := []int{}, []int{}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		nums := strings.Split(s.Text(), "   ")
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}

	slices.Sort(l1)
	slices.Sort(l2)

	result := 0
	for i := 0; i < len(l1); i++ {
		result += int(math.Abs(float64(l1[i]) - float64(l2[i])))
	}

	fmt.Println(result)
}
