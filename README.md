# Read flags from config file

[![GoDoc](https://godoc.org/github.com/hyperjiang/flagfile?status.svg)](https://godoc.org/github.com/hyperjiang/flagfile)
[![Build Status](https://travis-ci.org/hyperjiang/flagfile.svg?branch=master)](https://travis-ci.org/hyperjiang/flagfile)
[![](https://goreportcard.com/badge/github.com/hyperjiang/flagfile)](https://goreportcard.com/report/github.com/hyperjiang/flagfile)
[![codecov](https://codecov.io/gh/hyperjiang/flagfile/branch/master/graph/badge.svg)](https://codecov.io/gh/hyperjiang/flagfile)



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
