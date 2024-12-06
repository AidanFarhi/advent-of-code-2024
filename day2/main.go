package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := 0
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		nums := strings.Split(sc.Text(), " ")
		safe := true
		increasing := true
		if nums[0] > nums[1] {
			increasing = false
		}
		for i := 0; i < len(nums)-1; i++ {
			x, _ := strconv.Atoi(nums[i])
			y, _ := strconv.Atoi(nums[i+1])
			diff := 0
			if increasing {
				if y <= x {
					safe = false
					break
				}
				diff = y - x

			} else {
				if x <= y {
					safe = false
					break
				}
				diff = x - y
			}
			if diff < 1 || diff > 3 {
				safe = false
				break
			}
		}
		if safe {
			result++
		}
	}
	fmt.Println(result)
}
