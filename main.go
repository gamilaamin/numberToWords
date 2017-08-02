package main

import (
	"fmt"
)

var   (
	ones = []string  {"واحد","اثنان","ثلاثه","اربعه","خمسه","سته","سبعه","ثمانيه","تسعه","عشره","احد عشر","اثنا عشر"}
	tens = []string {"عشر","عشرون","ثلاثون","اربعوب","خمسون","ستون","سبعون","ثمانون","تسعون"}
	hundreds = []string  {"مئه","مئتان","ثلاثمائه","اربعمائة","خمسمائة","ستمائة","سبعمائه","ثمانمائه","تسعمائه"}
	thousands = []string  {"الف","الفان"}
	millions = [] string {"مليون","مليار","بليون"}
)
func numToArabic(number int) string  {
	var str string
	//fmt.Println("hello from num to arabic ")
	//fmt.Println("fnumber  ", number)

	//for get index of thousands

	num1000 := number/ 1000
	//fmt.Println("thousand ", num1000)

	// for get index of hundreds

	num100 := (number % 1000) / 100
	//fmt.Println("handreds ", num100)

	//for get index of tens

	num10 := ((number % 1000) % 100) / 10
	//fmt.Println("tens ", num10)

	// for get index of ones

	num :=  (((number % 1000) % 100) % 10)
	//fmt.Println("ones ", num)

	// millions
	//case larger than million and less than 2 million
	if (num1000 == 1000 && num1000 < 2000) {
		y := num1000/1000
		return "مليون و " + numToArabic(number-(y*1000000))
	}
	//case larger than 2 million and less than Billion
	if (num1000 >= 2000 && num1000 < 1000000 ){
		y := num1000/1000
		str += " مليون و "
		return  numToArabic(y) + str + numToArabic(number-(y*1000000))
	}

	//thousands

	//الف او الفان

	if (num1000 < 3 && number >= 1000) {
		str += thousands[num1000 - 1] + " و"
		return str + numToArabic(number - (num1000 * 1000))
	}

	// thousand from 2 to 12

	if(num1000 > 2 && num1000 < 13){
		str = ones[num1000 - 1]+" ألاف "
		//re call function with out thousands
		return str + numToArabic(number - (num1000 * 1000))
	}

	// thousand larger than 12

	if (num1000 > 12 && num1000 < 1000){
		return numToArabic(num1000) + " الف و " + numToArabic(number - (num1000 * 1000))
	}

	//hundreds

	if (num100 > 0) {
		str += hundreds[num100 - 1] + " و "
		return str + numToArabic(number - (num100*100))
	}

	//ones
	//13 14 15 16 17 8 19
	if (number > 12 && number < 20) {
		str +=  ones[num - 1]+ " " +tens[num10 - 1]
		return str

	}
	//21 22 .....
	if (number > 20 && (number % 10 > 0 )) {
		str += ones[num - 1] + " و " + tens[num10 - 1]
		return str
	}
	//20 30 40 50 60 70 80 90
	if  (number > 19 && (number % 10 == 0)){
		str += tens [num10 - 1]
		return str
	}
	// 1 2 .....12
	if (number < 13 && number > 0) {
		str += ones[number - 1]
		return str
	}
	return str
}
func main()  {
	var number int
	fmt.Print("Enter your number : ")
	fmt.Scanf("%d", &number)        //cin

	arabicString := numToArabic(number)
	fmt.Println(arabicString)
}