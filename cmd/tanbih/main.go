package main

import (
	"fmt"
	"io"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stdout, "Tanbih CLI is ready.")
		fmt.Fprintln(stdout, "Run 'tanbih --version' to see build information.")
		return 0
	}

	switch args[0] {
	case "version", "--version", "-v":
		fmt.Fprintf(stdout, "Tanbih %s\n", version)
		fmt.Fprintf(stdout, "Commit: %s\n", commit)
		fmt.Fprintf(stdout, "Built: %s\n", date)
		return 0

	case "help", "--help", "-h":
		fmt.Fprintln(stdout, "Tanbih - Prayer times notification CLI")
		fmt.Fprintln(stdout, "Usage: tanbih [--version | --help]")
		return 0

	default:
		fmt.Fprintf(stderr, "Unknown command: %s\n", args[0])
		return 2
	}
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
