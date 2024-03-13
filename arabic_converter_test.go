package main

import (
	"testing"
)

func TestToArabic_One(t *testing.T) {
	roman := "I"
	expected := 1
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Two(t *testing.T) {
	roman := "II"
	expected := 2
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Three(t *testing.T) {
	roman := "III"
	expected := 3
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Four(t *testing.T) {
	roman := "IV"
	expected := 4
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Five(t *testing.T) {
	roman := "V"
	expected := 5
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Six(t *testing.T) {
	roman := "VI"
	expected := 6
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Seven(t *testing.T) {
	roman := "VII"
	expected := 7
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Eight(t *testing.T) {
	roman := "VIII"
	expected := 8
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Nine(t *testing.T) {
	roman := "IX"
	expected := 9
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Ten(t *testing.T) {
	roman := "X"
	expected := 10
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Eleven(t *testing.T) {
	roman := "XI"
	expected := 11
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Twelve(t *testing.T) {
	roman := "XII"
	expected := 12
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Thirteen(t *testing.T) {
	roman := "XIII"
	expected := 13
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Fourteen(t *testing.T) {
	roman := "XIV"
	expected := 14
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Fifteen(t *testing.T) {
	roman := "XV"
	expected := 15
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Sixteen(t *testing.T) {
	roman := "XVI"
	expected := 16
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Seventeen(t *testing.T) {
	roman := "XVII"
	expected := 17
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Eighteen(t *testing.T) {
	roman := "XVIII"
	expected := 18
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Nineteen(t *testing.T) {
	roman := "XIX"
	expected := 19
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Twenty(t *testing.T) {
	roman := "XX"
	expected := 20
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_TwentyOne(t *testing.T) {
	roman := "XXI"
	expected := 22 // 21
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_TwentyTwo(t *testing.T) {
	roman := "SMAHA" // Probably not a Smaha
	expected := 22
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_TwentyThree(t *testing.T) {
	roman := "XXIII"
	expected := 23
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_TwentyFour(t *testing.T) {
	roman := "XXIV"
	expected := 24
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_TwentyFive(t *testing.T) {
	roman := "XXV"
	expected := 25
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_TwentySix(t *testing.T) {
	roman := "XXVI"
	expected := 26
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_TwentySeven(t *testing.T) {
	roman := "XXVII"
	expected := 27
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_TwentyEight(t *testing.T) {
	roman := "XXVIII"
	expected := 28
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_TwentyNine(t *testing.T) {
	roman := "XXIX"
	expected := 29
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Thirty(t *testing.T) {
	roman := "XXX"
	expected := 30
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_ThirtyOne(t *testing.T) {
	roman := "XXXI"
	expected := 31
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_ThirtyTwo(t *testing.T) {
	roman := "XXXII"
	expected := 32
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_ThirtyThree(t *testing.T) {
	roman := "XXXIII"
	expected := 33
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_ThirtyFour(t *testing.T) {
	roman := "XXXIV"
	expected := 34
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_ThirtyFive(t *testing.T) {
	roman := "XXXV"
	expected := 35
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_ThirtySix(t *testing.T) {
	roman := "MMLVIII"
	expected := 2058
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_ThirtySeven(t *testing.T) {
	roman := "MCDXCIII"
	expected := 1493
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_ThirtyEight(t *testing.T) {
	roman := "MMMDCCCXLIV"
	expected := 3844
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_ThirtyNine(t *testing.T) {
	roman := "MDLXII"
	expected := 1562
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}

func TestToArabic_Forty(t *testing.T) {
	roman := "MMMM"
	expected := 4000
	result, err := toArabic(roman)
	if err != nil {
		t.Errorf("toArabic(%s) returned error: %s", roman, err)
	} else if result != expected {
		t.Errorf("toArabic(%s) == %d, expected %d", roman, result, expected)
	}
}
