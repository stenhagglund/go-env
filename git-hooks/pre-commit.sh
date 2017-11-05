#!/bin/bash

# Check that all .go files added are correctly formatted
unformatted=$(gofmt -l .)
if ! [ -z "$unformatted" ]; then
    # Some files are not gofmt'd. Print message and fail.
    echo >&2 "Failed: gofmt -l ."
    echo >&2 "$unformatted"
    exit 1
fi

# Check that all .go files are linted correctly
golint -set_exit_status .
if [ $? != 0 ]; then
    # Some files are not linted. Print message and fail.

    echo >&2 "Failed: golint ."
    exit 1
fi

# Check that all .go files are linted correctly
go vet .
if [ $? != 0 ]; then
    # Some files are not vetted. Print message and fail.
    echo >&2 "Failed: go vet ."
    exit 1
fi

# Check that all .go files are linted correctly
testoutput=$(go test .)
if [ $? != 0 ]; then
    # Some files failed tests. Print message and fail.

    echo >&2 "Failed: go test ."
    echo >&2 "$testoutput"
    exit 1
fi

exit 0