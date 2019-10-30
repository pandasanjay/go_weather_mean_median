package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	//Geting from the url
	//res, err := http.Get("http://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2015.txt")
	f, err := os.Open("./Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
	defer f.Close()
	//defer res.Body.Close()
	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("Total Records: ", len(rows)-1)
	fmt.Println("Mean Air Temp: ", getMean(rows, 1), getMedian(rows, 1))
	fmt.Println("Mean Barometric: ", getMean(rows, 2), getMedian(rows, 2))
	fmt.Println("Mean Wind Speed: ", getMean(rows, 7), getMedian(rows, 7))
}

//Pass the row
//find the index from
func getMean(rows [][]string, idx int) float64 {
	var total float64
	for ind, row := range rows {
		if ind >= 1 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			total += val
		}
	}
	return total / float64(len(rows)-1)
}

func getMedian(rows [][]string, idx int) float64 {
	var columData []float64
	for ind, row := range rows {
		if ind != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			columData = append(columData, val)
		}
	}
	sort.Float64s(columData)
	if len(columData)%2 == 0 {
		firstValue := columData[len(columData)/2]
		secondValue := columData[len(columData)/2-1]
		return (firstValue + secondValue) / 2
	}
	return columData[(len(columData))/2]
}
