package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var Precision int

func main() {
	var number string
	var fromBase, toBase int

	fmt.Print("Please input the number: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Please input the precision: ")
	_, err = fmt.Scanln(&Precision)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Please input a base system of the number: ")
	_, err = fmt.Scanln(&fromBase)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Please input a base system you would like to convert it: ")
	_, err = fmt.Scanln(&toBase)
	if err != nil {
		log.Fatal(err)
	}

	res, err := convertNumber(number, fromBase, toBase)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n\nThe result is: %s\n\n", res)
}

func convertNumber(val string, fromBase int, toBase int) (string, error) {
	intPart, err := convertIntPart(strings.Split(val, ".")[0], fromBase, toBase)
	if err != nil {
		return "", err
	}

	fracDecPart, err := convertFracPartToDecimal(strings.Split(val, ".")[1], fromBase)
	if err != nil {
		return "", err
	}

	fracPart, err := convertFracPartFromDecimal(fracDecPart, toBase)
	if err != nil {
		return "", err
	}

	res := intPart + "." + fracPart
	return res, nil
}

func convertIntPart(val string, fromBase int, toBase int) (string, error) {
	intPart, err := strconv.ParseInt(val, fromBase, 64)
	if err != nil {
		return "", err
	}
	res := strconv.FormatInt(intPart, toBase)
	return res, nil
}

func convertFracPartToDecimal(val string, fromBase int) (string, error) {
	k := -1.0
	floatRes := 0.0
	for _, el := range val {
		if int(el-'0') > fromBase {
			return "", errors.New("wrong value (digit can not be > system)")
		}
		floatRes += float64(el-'0') * math.Pow(float64(fromBase), k)
		k--
	}
	res := strings.Split(strconv.FormatFloat(floatRes, 'f', Precision, 64), ".")[1]
	return res, nil
}

func convertFracPartFromDecimal(val string, toBase int) (string, error) {
	val = "0." + val
	flVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return "", err
	}
	mp := map[int]string{
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}
	res := ""
	for i := 0; i < Precision; i++ {
		flVal = flVal * float64(toBase)
		digit := int(math.Floor(flVal))
		runeVal, ok := mp[digit]
		if ok {
			res += runeVal
		} else {
			res += strconv.Itoa(digit)
		}
		flVal = flVal - math.Floor(flVal)
	}

	return res, nil
}
