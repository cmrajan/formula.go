package main

import (
	"testing"
)

func TestCHAR(t *testing.T) {
	var v string
	v = CHAR(65)
	if v != "A" {
		t.Error("Expected A, got ", v)
	}
}

func TestCLEAN(t *testing.T) {
	var v string
	v = CLEAN(`this is a 
text`)
	if v != `"this is a \ntext"` {
		t.Error("Expected this is a \ntext, got ", v)
	}
}

func TestCODE(t *testing.T) {
	var v int
	v = CODE("A")
	if v != 65 {
		t.Error("Expected 65, got ", v)
	}

}

func TestCONCATENATE(t *testing.T) {
	var v string
	v = CONCATENATE("hello", "world")
	if v != "helloworld" {
		t.Error("Expected helloworld, got ", v)
	}

	v = CONCATENATE("1 ", "world")
	if v != "1 world" {
		t.Error("Expected 1world, got ", v)
	}

}

func TestEXACT(t *testing.T) {
	var v string
	v = EXACT("hello", "hello")
	if v != "TRUE" {
		t.Error("Expected TRUE, got ", v)
	}

	v = EXACT("1", "world")
	if v != "FALSE" {
		t.Error("Expected FALSE, got ", v)
	}

}

func TestLEFT(t *testing.T) {
	var v string
	v = LEFT("abcdef", 3)
	if v != "abc" {
		t.Error("Expected abc, got ", v)
	}

	v = LEFT("abcdef", 3000)
	if v != "abcdef" {
		t.Error("Expected abcdef, got ", v)
	}

	v = LEFT("abcdef", 6)
	if v != "abcdef" {
		t.Error("Expected abcdef, got ", v)
	}

}

func TestLEN(t *testing.T) {
	var v int
	v = LEN("count me")
	if v != 8 {
		t.Error("Expected 8, got ", v)
	}

	v = LEN("   count me too   ")
	if v != 18 {
		t.Error("Expected 18, got ", v)
	}

}

func TestLOWER(t *testing.T) {
	var v string
	v = LOWER("HELLO")
	if v != "hello" {
		t.Error("Expected hello, got ", v)
	}

	v = UPPER(1)
	if v != "1" {
		t.Error("Expected 1, got ", v)
	}

}

func TestMID(t *testing.T) {
	var v string
	v = MID("abcdef", 4, 3)
	if v != "def" {
		t.Error("Expected def, got ", v)
	}

	v = MID("abcdef", 4, 3000)
	if v != "def" {
		t.Error("Expected def, got ", v)
	}

	v = MID("abcdef", 6, 6)
	if v != "f" {
		t.Error("Expected f, got ", v)
	}

}

func TestPROPER(t *testing.T) {
	var v string
	v = PROPER("change to title case")
	if v != "Change To Title Case" {
		t.Error("Expected Change To Title Case, got ", v)
	}

	v = PROPER("CHANGE TO TITLE CASE")
	if v != "CHANGE TO TITLE CASE" {
		t.Error("Expected Change To Title Case, got ", v)
	}
	v = PROPER(12121212133434)
	if v != "12121212133434" {
		t.Error("Expected 12121212133434, got ", v)
	}

}

func TestREPT(t *testing.T) {
	var v string
	v = REPT("we", 2)
	if v != "wewe" {
		t.Error("Expected wewe, got ", v)
	}

}

func TestRIGHT(t *testing.T) {
	var v string
	v = RIGHT("abcdef", 3)
	if v != "def" {
		t.Error("Expected def, got ", v)
	}

	v = RIGHT("abcdef", 3000)
	if v != "abcdef" {
		t.Error("Expected abcdef, got ", v)
	}

	v = RIGHT("abcdef", 6)
	if v != "abcdef" {
		t.Error("Expected abcdef, got ", v)
	}

}

func TestSEARCH(t *testing.T) {
	var v int
	v = SEARCH("ab", "tablet")
	if v != 2 {
		t.Error("Expected 2, got ", v)
	}

	v = SEARCH("ab", "tabletab", 3, 2)
	if v != 7 {
		t.Error("Expected 7, got ", v)
	}

	v = SEARCH("ab", "tacbletacb", 0)
	if v != -1 {
		t.Error("Expected 0, got ", v)
	}

}

func TestTRIM(t *testing.T) {
	var v string
	v = TRIM(" new     new    old     ")
	if v != "new new old" {
		t.Error("Expected new new old, got ", v)
	}

}

func TestUNICODE(t *testing.T) {
	var v int
	v = UNICODE("A")
	if v != 65 {
		t.Error("Expected 65, got ", v)
	}

}
func TestUNICHAR(t *testing.T) {
	var v string
	v = UNICHAR(65)
	if v != "A" {
		t.Error("Expected A, got ", v)
	}

}

func TestUPPER(t *testing.T) {
	var v string
	v = UPPER("hello")
	if v != "HELLO" {
		t.Error("Expected HELLO, got ", v)
	}

	v = UPPER(1)
	if v != "1" {
		t.Error("Expected 1, got ", v)
	}

}
