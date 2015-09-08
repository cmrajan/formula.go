package main

import (
	"fmt"
	"strconv"
	"strings"
)

//define error conts
const (
	Errnil   = "#NULL!"
	Errdiv0  = "#DIV/0!"
	Errvalue = "#VALUE?"
	Errref   = "#REF!"
	Errname  = "#NAME?"
	Errnum   = "#NUM!"
	Errna    = "#N/A"
	Errerror = "#ERROR!"
	Errdata  = "#GETTING_DATA"
)

func ToInt(i interface{}) int {
	var res int
	if val, ok := i.(int); ok {
		res = val
	} else {

		val, _ := i.(string)
		res, _ = strconv.Atoi(strings.Trim(val, " "))
	}
	return res

}

func ToIntList(i ...interface{}) []int {
	var res []int
	for _, v := range i {
		fmt.Println("incoming v", v)
		switch v.(type) {
		case int:
			fmt.Println("int", v.(int))
			res = append(res, v.(int))
		default:
			fmt.Println("stc")

		}
	}
	return res
}

func ToStr(i interface{}) string {
	var res string
	switch i.(type) {
	case int:
		res = strconv.Itoa(i.(int))
	case string:
		res = i.(string)

	}
	return res

}

func ToStrList(i ...interface{}) []string {
	var res []string
	for _, v := range i {
		switch v.(type) {
		case int:
			res = append(res, strconv.Itoa(v.(int)))
		case string:
			res = append(res, v.(string))
		}
	}
	return res
}
