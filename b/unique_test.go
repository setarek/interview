package b

import (
	"testing"
)

func TestFindUniqueValue(t *testing.T) {
	l := []int{2, 3, 5, 2, 7, 5, 7}
	result := FindUniqueValue(l)
	expectedValue := 3

	if result != expectedValue {
		t.Error("FindUniqueValue Failed")
	} else {
		t.Log("FindUniqueValue Succeded")
	}

}
