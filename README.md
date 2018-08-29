# DJSON and Cobra example

A simple example for demonstrating how DJSON can be used for parsing program arguments.

## Install

```bash
git clone https://github.com/moikot/djson.git
dep ensure -vendor-only
go build
```

## Run

```bash
./start.sh
```

The result should be:

```bash
foo: true
bar: "true"
```