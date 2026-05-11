package main

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("expected 5, got %d", got)
	}
}

// codelight-rerun

// token-rerun
