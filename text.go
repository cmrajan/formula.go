package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//func ASC()                                                                      {}
func BAHTTEXT()                                                             {}
func DBCS()                                                                 {}
func DOLLAR(number interface{}, decimals interface{})                       {}
func FIXED(number interface{}, decimals interface{}, no_commas interface{}) {}

//func HTML2TEXT(value interface{})                                           {}

func NUMBERVALUE(text interface{}, decimal_separator interface{}, group_separator interface{}) {}
func PRONETIC()                                                                                {}

//func REGEXEXTRACT(text interface{}, regular_expression interface{})                            {}
//func REGEXMATCH(text interface{}, regular_expression interface{}, full interface{})            {}
//func REGEXREPLACE(text interface{}, regular_expression interface{}, replacement interface{})   {}
func SUBSTITUTE(text interface{}, old_text interface{}, new_text interface{}, occurrence interface{}) {
}
func T(value interface{})    {}
func VALUE(text interface{}) {}

func CHAR(number interface{}) string {

	i := ToInt(number)
	if i < 256 {
		return string(i)
	} else {
		return Errvalue
	}
}

func CLEAN(text interface{}) string {
	s := ToStr(text)
	return strconv.Quote(s)
}

func CODE(text interface{}) int {
	s := ToStr(text)
	var r []rune
	r = []rune(s)[0:1]

	return int(r[0])

}

func CONCATENATE(str ...interface{}) string {
	s := ToStrList(str...)
	return strings.Join(s, "")
}

func EXACT(text1 interface{}, text2 interface{}) string {

	s1 := ToStr(text1)
	s2 := ToStr(text2)

	if s1 == s2 {
		return "TRUE"
	} else {
		return "FALSE"
	}

}

func FIND(find_text interface{}, within_text interface{}, position interface{}) {

}

func LEFT(text interface{}, number interface{}) string {
	s := ToStr(text)
	count := ToInt(number)
	if count > len(s) {
		count = len(s)
	}
	return s[0:count]

}

func LEN(text interface{}) int {
	s := ToStr(text)
	// len := strings.Count(s, "")
	// if len > 0 {
	// 	len = len - 1
	// }
	return len(s)

}

func LOWER(text interface{}) string {
	s := ToStr(text)

	return strings.ToLower(s)
}

func MID(text interface{}, start interface{}, number interface{}) string {
	s := ToStr(text)
	from_char := ToInt(start)
	num_char := ToInt(number)
	if (from_char + num_char) > len(s) {
		num_char = len(s) - from_char
	}
	return s[from_char-1 : from_char+num_char]

}

func PROPER(text interface{}) string {
	s := ToStr(text)
	return strings.Title(s)

}

// func SPLIT(text interface{}, separator interface{}) {
// 	s := ToStr(text)
// 	sep := ToStr(separator)

// }
// func REPLACE(text interface{}, position interface{}, length interface{}, new_text interface{}) string {

// }

func REPT(text interface{}, number interface{}) string {
	s := ToStr(text)
	count := ToInt(number)
	return strings.Repeat(s, count)

}

func RIGHT(text interface{}, number interface{}) string {
	s := ToStr(text)
	count := ToInt(number)
	if count > len(s) {
		count = len(s)
	}
	return s[len(s)-(count) : len(s)]

}

func SEARCH(find_text interface{}, within_text interface{}, position ...interface{}) int {
	var idx_position, n_position int
	var position_arg []int
	s_find_text := ToStr(find_text)
	s_find_text = strings.ToLower(s_find_text) // SEARCH is case insensitive
	s_within_text := ToStr(within_text)
	s_within_text = strings.ToLower(s_within_text) // SEARCH is case insensitive
	position_arg = ToIntList(position)

	fmt.Println("position", position)

	fmt.Println("position_arg", position_arg)

	if len(position_arg) > 0 {
		n_position = position_arg[0]
	} else {
		n_position = 0
	}

	re := regexp.MustCompile(s_find_text)

	s_within_text = s_within_text[n_position:len(s_within_text)]
	fmt.Println("s_within_text", s_within_text, "position_arg", position_arg, "n_position", n_position)
	//	var i int
	i := re.FindStringSubmatchIndex(s_within_text)
	if len(i) > 0 {
		idx_position = i[0] + n_position + 1
	} else {

		idx_position = -1
	}
	return idx_position
}

// func TEXT(value interface{}, format interface{}) string {

// }

func TRIM(text interface{}) string {
	s := ToStr(text)

	return strings.Join(strings.Fields(s), " ")

}

func UPPER(text interface{}) string {
	s := ToStr(text)

	return strings.ToUpper(s)
}

func UNICHAR(number interface{}) string {
	i := ToInt(number)
	return string(i)

}

func UNICODE(text interface{}) int {
	s := ToStr(text)
	var r []rune
	r = []rune(s)[0:1] // only rune can be converted to it's int

	return int(r[0])

}

func main() {}
