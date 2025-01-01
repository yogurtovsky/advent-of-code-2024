package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

)

func readInput(file string, tokens int) [][]string {
    f, err := os.Open(file)

    data := make([][]string, tokens)
    if err != nil {
        log.Fatalf("Failed to open file %v", err) 
    } else {
        scanner := bufio.NewScanner(f)
        
        input := make([]string, 0)
        lineCount := 0
        
        // get input data from  file
        for scanner.Scan() {
            line := scanner.Text()
            input = append(input, strings.Split(line, " ")...)
            lineCount++
        }

        // allocate slices for left and right data
        for idx := range data {
            data[idx] = make([]string, 0)
        }

        step := 0
        // fill data
        for idx := range input {
            step++
            data[step-1] = append(data[step-1], input[idx]) 
    
            if step == tokens {
                step = 0
            }
        }
    }
    return data
}

func main()  {
    data := readInput("input", 2)    

    l_int := make([]int, 0)
    r_int := make([]int, 0)

    for idx := range data[0] {
        l, err_l := strconv.Atoi(data[0][idx])
        r, err_r := strconv.Atoi(data[1][idx])
    
        if err_l != nil {
            log.Fatalf("Failed to convert left value")
            return
        }
        if err_r != nil {
            log.Fatalf("Failed to convert right value")
            return
        }
    
        l_int = append(l_int, l)
        r_int = append(r_int, r)
    }

    sort.Ints(l_int)
    sort.Ints(r_int)

    answer_1 := 0.0
    for idx := range l_int {
        answer_1 = answer_1 + math.Abs((float64)(l_int[idx] - r_int[idx]))
    }

    fmt.Printf("Answer part 1: %v\n", answer_1)

    // part 2

    answer_2 := 0.0

    for idx := range l_int {
        count := 0.0 

        // fuck it we ball it's really not performant but i don't give a shit ¯\_(ツ)_/¯
        for jdx := range r_int {
            if l_int[idx] == r_int[jdx] {
                count++
            }
        }

        answer_2 = answer_2 + (float64)(l_int[idx]) * count
    } 
    fmt.Printf("Answer part 2: %v\n", answer_2)
}

