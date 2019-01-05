// Ninja Level 12 - Hands-On Exercise from Udemy course on Learning Go Programming,
// by Todd McLeod

//Provide basic statistics on car sales for traditional U.S. muscle cars
//I felt like doing something different than the dog years example in the course syllabus :-)
package main

import (
	"encoding/json"
	"fmt"
)

//	{"carModelName": "Firebird", "carModelYear": "1996", "carModelSales": "32622"},
//	{"carModelName": "Firebird", "carModelYear": "1997", "carModelSales": "32524"},
//	{"carModelName": "Mustang", "carModelYear": "1996", "carModelSales": "122674"},
//	{"carModelName": "Mustang", "carModelYear": "1997", "carModelSales": "116610"}
//	]`)

func main() {
	var carModel, fullName string
	var modelYear int
	//fmt.Println(jsonBlob)
	populateDb()

	//Input one of four car types, Camaro, Firebird, Mustang
	fmt.Print("Enter Car Model Name (Camaro | Firebird | Mustang): ")
	_, err := fmt.Scan(&carModel)
	if err != nil {
		panic(err)
	}

	//Input a specific year, or enter 0 to see all years
	fmt.Print("Enter a specific Model Year between 1996-1997. Enter 0 for all years: ")
	_, err = fmt.Scan(&modelYear)
	if err != nil {
		panic(err)
	}

	//Print out what the user input was
	switch carModel {
	case "Camaro":
		fullName = "Chevrolet Camaro"
	case "Firebird":
		fullName = "Pontiac Firebird"
	case "Mustang":
		fullName = "Ford Mustang"
	}
	fmt.Println("\t", fullName, "\t\t\t", modelYear)

	fmt.Println("")
}

func populateDb() {

	var jsonBlob = []byte(`[
	{"CarName": "Camaro", "CarYear": "1996", "CarSales": "66866"},
	{"CarName": "Camaro", "CarYear": "1997", "CarSales": "55973"},
	{"CarName": "Firebird", "CarYear": "1996", "CarSales": "32622"},
	{"CarName": "Firebird", "CarYear": "1997", "CarSales": "32524"},
	{"CarName": "Mustang", "CarYear": "1996", "CarSales": "122674"},
	{"CarName": "Mustang", "CarYear": "1997", "CarSales": "116610"}
	]`)

	type carModelSales struct {
		CarName  string
		CarYear  string
		CarSales string
	}

	var allCars []carModelSales

	err := json.Unmarshal(jsonBlob, &allCars)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("%+v\n", allCars)
}
