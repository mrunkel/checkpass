package main

import (
	"testing"
)

func TestCheckHIBP(t *testing.T) {

	tables := []struct {
		password      string
		expectedCount int
	}{
		{"ilovemymom", 20141},
		{"apostasy-fryer-osaka-impiety-aspirin-hightail", 0},
	}

	for _, table := range tables {
		count := checkHIBP(table.password)

		if count != table.expectedCount {
			t.Errorf("checkHIBP is incorrect, got: %d, want: %d for %s.", count, table.expectedCount, table.password)
		}

	}

}

func TestFormatResult(t *testing.T) {
	tables := []struct {
		count          int
		expectedResult string
	}{
		{0, "Your password has NOT been found!  Congrats....."},
		{1, "Your password has been found.\tIt has been used 1 time"},
		{2, "Your password has been found.\tIt has been used 2 times"},
	}

	for _, table := range tables {
		result := formatResult(table.count)
		if result != table.expectedResult {
			t.Errorf("formatResult is incorrect got: %s, wanted %s for %d.", result, table.expectedResult, table.count)
		}
	}
}
