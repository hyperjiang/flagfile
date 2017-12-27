package flagfile

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestParseFromCommandLine(t *testing.T) {
	var str string

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "-a", "test"}

	flag.StringVar(&str, "a", "default value", "usage")

	Parse()

	if str != "test" {
		t.Fail()
	}
}

func TestPrintUsage(t *testing.T) {
	if os.Getenv("RUN_TEST") == "1" {
		Parse()
		return
	}

	cmd := exec.Command(os.Args[0])
	cmd.Env = append(os.Environ(), "RUN_TEST=1")

	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return // expected
	}

	t.Fail()
}

func TestParseFromFile(t *testing.T) {
	var (
		addr          string
		serverTimeout time.Duration
	)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "-flagfile=./example.conf"}

	flag.StringVar(&addr, "addr", ":80", "Address to listen and serve")
	flag.DurationVar(&serverTimeout, "server_timeout", time.Second, "Http server timeout")

	Parse()

	if addr != ":8888" || serverTimeout != 5*time.Second {
		t.Fail()
	}
}

func TestTryFile(t *testing.T) {
	if os.Getenv("RUN_TEST2") == "1" {
		oldArgs := os.Args
		defer func() { os.Args = oldArgs }()
		os.Args = []string{"cmd", "-tryflagfile=./example.conf"}
		Parse()
		return
	}

	cmd := exec.Command(os.Args[0])
	cmd.Env = append(os.Environ(), "RUN_TEST2=1")

	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		t.Fatalf("%v", e)
	}
}

func TestAll(t *testing.T) {
	r := All()
	if fmt.Sprintf("%v", r["a"]) != "test" {
		t.Fail()
	}
}
