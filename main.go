package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	var my_slice []string
	// var group []string
	// var slice_groups [][]string

	f, err := os.Open("names.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		my_slice = append(my_slice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(my_slice), func(i, j int) { my_slice[i], my_slice[j] = my_slice[j], my_slice[i] })

	fmt.Println("Welcome to our team generator" + "\n\n\n" + "How big should the teams be?")

	var teamSize int

	fmt.Scanln(&teamSize)

	var divided [][]string

	for i := 0; i < len(my_slice); i += teamSize {
		end := i + teamSize

		if end > len(my_slice) {
			end = len(my_slice)
		}

		divided = append(divided, my_slice[i:end])
	}

	for i := 0; i < len(divided); i++ {
		fmt.Println(divided[i])
		fmt.Println(" \n")
	}

	// fmt.Printf("%v", divided)

	// _ = group

	// for i := 0; i < len(my_slice); i++ {

	// 	for j := i; j < teamSize; j++ {
	// 		group = append(group, my_slice[i])
	// 		fmt.Println(group)

	// 	}

	// 	slice_groups = append(slice_groups, group)

	// }

}
