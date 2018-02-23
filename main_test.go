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
			t.Errorf("checkHIBP was incorrect, got: %d, want: %d for %s.", count, table.expectedCount, table.password)
		}

	}

}
