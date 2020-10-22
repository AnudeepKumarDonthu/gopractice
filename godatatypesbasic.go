package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func printUserName() {
	var myName string = "Anudeep Kumar Donthu"
	var a, _ = fmt.Println(myName)
	fmt.Print(a)
}

func readUserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Full Name :: ")
	myFullName, _ := reader.ReadString('\n')
	fmt.Println(myFullName)
}

func readUserInputAndConvertToFloat() {
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Rating ::")
	rating, _ := reader1.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	fmt.Print(myNumRating + 2)
}

func restaurantRating() {
	var name string
	var rating string

	//Ask for Input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pleae Enter Your Full Name :: ")
	name, _ = reader.ReadString('\n')
	trimmedName := strings.TrimSpace(name)
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please provide User Rating and conver into Intege")
	rating, _ = reader.ReadString('\n')
	floatRatingValue, _ := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	//Print User Input Data
	fmt.Printf("Hello \"%v\", Thank you for rating us with \"%v\" at Time %v\n", trimmedName, floatRatingValue, time.Now().Format(time.Stamp))

	if floatRatingValue == 5 {
		fmt.Println("We people Did a Awesome Work")
	} else if floatRatingValue == 4 || floatRatingValue == 3 {
		fmt.Println("We did cool work and continues to improve")
	} else if floatRatingValue < 3 {
		fmt.Println("We need to work on to improve our quality")
	}
}

func Printmessage() {
	fmt.Println("Anyone Can Call me as I am publicly Available")
}

func pointerPractice() {
	var score float32 = 100
	scoreAddress := &score
	score = score * 0.1

	if scoreAddress != nil {
		fmt.Println("Pointer Address value is", scoreAddress)
		fmt.Println("Pointer value is", *scoreAddress)
	} else {
		fmt.Println("I am nil :(")
	}
}

func arrayPractice() {
	var number [3]string
	number[0] = "Donthu"
	number[1] = "Anudeep"
	number[2] = "Kumar"
	fmt.Println(number)
	fmt.Println("Numbers Array Size :: ", len(number))

	var colours = [4]string{"Green", "Red", "Black", "White"}
	fmt.Println(colours)
	fmt.Println("Colours Size :: ", len(colours))
}

func slicePractice() {
	var names = []string{"Donthu", "Anudeep"}
	fmt.Println(names)

	names = append(names, "Kumar")
	fmt.Println(names)

	names = append(names[1:])
	fmt.Println(names)

	sampleNames := make([]string, 2, 2)
	sampleNames[0] = "Simple"
	sampleNames[1] = "Sample"
	fmt.Println(sampleNames)

	sampleNames = append(sampleNames, "Append")
	fmt.Println(sampleNames)

	myints := []int{6, 4, 9, 1, 5}
	sort.Ints(myints)
	isSorted := sort.IntsAreSorted(myints)
	fmt.Println(myints)
	fmt.Println(isSorted)
}

func mapPractice() {
	score := make(map[string]int) //similar to new which just allocate but no initialization of Memory
	score["cricket"] = 100
	score["carrom"] = 40
	score["tabletenies"] = 20
	score["sncooker"] = 50
	fmt.Println(score)

	crickertScore := score["cricket"]
	fmt.Println(crickertScore)

	delete(score, "sncooker")
	fmt.Println(score)

	//Iterate Map Using for loop
	for k, v := range score {
		fmt.Println(k, " :: ", v)
	}
}

// User  Struct
type User struct {
	Name  string
	Email string
	Age   int32
}

func structPractice() {
	//Type1
	anudeep := User{"Anudeep Kumar", "donthu04@gmail.com", 29}
	fmt.Printf("%+v\n", anudeep)

	//Type2
	var umaUserDetails = new(User)
	umaUserDetails.Name = "Uma"
	umaUserDetails.Email = "akurathi15@gmail.com"
	umaUserDetails.Age = 28
	fmt.Printf("%+v\n", umaUserDetails)

	//Type3
	var kumarDetails = &User{"Kumar", "kumar@gmail.com", 29}
	fmt.Printf("%+v", kumarDetails)
}

func forLoopPractice() {
	//start := 1
	names := []string{"Anudeep", "Kumar", "Donthu", "Uma", "Akurathi"}
	fmt.Println(names)

	//Type1
	for i := 0; i < len(names); i++ {
		fmt.Print(names[i], ",")
	}

	fmt.Println("\n---------------------------")
	//Type2
	for i := range names {
		fmt.Print(names[i], ",")
	}
}

func main() {
	//printUserName()
	//readUserInput()
	//readUserInputAndConvertToFloat()
	//restaurantRating()
	//pointerPractice()
	//arrayPractice()
	//slicePractice()
	//mapPractice()
	//structPractice()
	forLoopPractice()

}
