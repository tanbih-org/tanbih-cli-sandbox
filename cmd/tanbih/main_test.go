package main

import (
	"bytes"
	"testing"
)

func TestRunDefault(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run(nil, &stdout, &stderr)

	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}

	if stdout.String() == "" {
		t.Fatal("expected application output")
	}
}

func TestRunVersion(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"--version"}, &stdout, &stderr)

	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}

	if !bytes.Contains(stdout.Bytes(), []byte("Tanbih")) {
		t.Fatal("expected version output to contain Tanbih")
	}
}

func TestRunUnknownCommand(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"unknown"}, &stdout, &stderr)

	if code != 2 {
		t.Fatalf("expected exit code 2, got %d", code)
	}

	if stderr.Len() == 0 {
		t.Fatal("expected an error message")
	}
}
