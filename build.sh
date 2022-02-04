#!/bin/bash

GOOS=windows GOARCH=amd64 go build -o bin/static-files_windows_x64.exe

GOOS=windows GOARCH=386 go build -o bin/static-files_windows_i386.exe

GOOS=darwin GOARCH=amd64 go build -o bin/static-files_macos_x64

GOOS=linux GOARCH=amd64 go build -o bin/static-files_linux_x64

GOOS=linux GOARCH=386 go build -o bin/static-files_linux_x86