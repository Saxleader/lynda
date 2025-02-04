package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func main() {
	res, err := http.Get("http://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	rdr := csv.NewReader(res.Body)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true
	defer res.Body.Close()
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	// Don't count the header row in len(rows)
	fmt.Println("Total Records: ", len(rows)-1)
	fmt.Println("Air Temp:\t", mean(rows, 1), median(rows, 1))
	fmt.Println("Barometric:\t", mean(rows, 2), median(rows, 2))
	fmt.Println("Wind Speed:\t", mean(rows, 7), median(rows, 7))
}

func median(rows [][]string, idx int) float64 {
	var sorted []float64
	for i, row := range rows {
		if i != 0 {
			val, err := strconv.ParseFloat(row[idx], 64)
			if err != nil {
				panic(err)
			}
			sorted = append(sorted, val)
		}
	}
	sort.Float64s(sorted)
	if len(sorted)%2 == 0 {
		// even number of items
		// for example: 3, 5, 8, 9
		// median is (5 + 8) / 2 = 6.5
		middle := len(sorted) / 2
		higher := sorted[middle]
		lower := sorted[middle-1]
		return (higher + lower) / 2
	}
	// odd number of items
	// for example: 3, 5, 8
	// median is 5
	middle := len(sorted) / 2
	return sorted[middle]
}

func mean(rows [][]string, idx int) float64 {
	var total float64
	for i, row := range rows {
		if i != 0 {
			val, err := strconv.ParseFloat(row[idx], 64)
			if err != nil {
				panic(err)
			}
			total += val
		}
	}
	return total / float64(len(rows)-1)
}
