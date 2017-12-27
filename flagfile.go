/*
Package flagfile is a wrapper of flag for reading flags from config file
*/
package flagfile

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	flagFilePrefix    = "-flagfile="
	tryFlagFilePrefix = "-tryflagfile="
)

// All returns all the flag values
func All() map[string]interface{} {
	r := make(map[string]interface{})

	flag.VisitAll(func(f *flag.Flag) {
		r[f.Name] = f.Value
	})

	return r
}

/*
Parse parses flags, the process will exit if fail to parse. There are 4 possible scenarios:
	1. Print usage infomation and exit: no args; args[1] = "-h"; args[1] = "-help"
	2. Parse the specified flag file: args[1] = "-flagfile=<file>"
	3. Try to parse the specified flag file: args[1] = "-tryflagfile=<file>"
	4. Parse from command line: fallback method
*/
func Parse() {
	if len(os.Args) < 2 || "-h" == os.Args[1] || "-help" == os.Args[1] {
		PrintUsage()
		os.Exit(1)
	} else if strings.HasPrefix(os.Args[1], flagFilePrefix) {
		if err := parseFromFile(os.Args[1][len(flagFilePrefix):]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	} else if strings.HasPrefix(os.Args[1], tryFlagFilePrefix) {
		if err := parseFromFile(os.Args[1][len(tryFlagFilePrefix):]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "Try parsing flag file successfully")
		os.Exit(0)
	}

	// fallback: parse from command line
	flag.Parse()
}

// PrintUsage print usage to standard output
func PrintUsage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.CommandLine.PrintDefaults()
}

// parseFromFile parse flags(C++ gflags styling) from file
func parseFromFile(path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	flags := strings.Split(string(content), "\n")
	if flag.CommandLine.Parse(flags) != nil {
		return err
	}

	return nil
}
