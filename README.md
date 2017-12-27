# Read flags from config file

[![GoDoc](https://godoc.org/github.com/hyperjiang/go/flagfile?status.svg)](https://godoc.org/github.com/hyperjiang/go/flagfile)

## Install

```
go get github.com/hyperjiang/flagfile
```

## Supported config file format

Currently `flagfile` only supports [gflags](https://github.com/gflags/gflags) styling config file

## Usage

```
import (
    "flag"
    "fmt"
    "github.com/hyperjiang/flagfile"
)

var str string

flag.StringVar(&str, "a", "default", "usage")

flagfile.Parse()

fmt.Println(str)

```
