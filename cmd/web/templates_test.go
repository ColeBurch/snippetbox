package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC)
	hd := humanDate(tm)
	if hd != "01 Oct 2023 at 12:00" {
		t.Errorf("Expected %q; Got %q", "01 Oct 2023 at 12:00", hd)
	}
}
