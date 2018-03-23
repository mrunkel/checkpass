package main

import (
	"github.com/andreyvit/diff"
	"strings"
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

		{0, "Your password has NOT been found!  Congrats..... "},
		{1, "Your password has been compromised and shouldn't be used any longer.\nYou can read about how to select " +
			"a good password at https://runkel.org/2017/09/how-to-pick-a-password/\n" +
			"\n\nIt has been found 1 time on the dark web.\n\n\n\n"},
		{2, "Your password has been compromised and shouldn't be used any longer.\nYou can read about how to select " +
			"a good password at https://runkel.org/2017/09/how-to-pick-a-password/\n" +
			"\n\nIt has been found 2 times on the dark web.\n\n\n\n"},
	}

	for _, table := range tables {
		result := formatResult(table.count)
		if a, e := strings.TrimSpace(result), strings.TrimSpace(table.expectedResult); a != e {
			t.Errorf("Result not as expected\n%v", diff.LineDiff(e, a))
		}
	}
}
