#!/bin/bash

echo "First creating a rust library C type crate"
echo
echo
rustc --crate-type cdylib -o liboutput.so output.rs
echo
echo "Crate successfully create : link=liboutput.so"
echo "Export current location for linked library"
export LD_LIBRARY_PATH=.
echo
echo
echo "Building a main.go project with link to rust crate"
go build main.go
./main

echo 
echo "Terminating program and cleaning up binaries/linked crates"
rm liboutput.so main 