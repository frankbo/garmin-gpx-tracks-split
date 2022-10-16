#!/usr/bin/env bash

set -e
set -u
set -o pipefail

mkdir output

# Windows
env GOOS=windows GOARCH=amd64 go build -o output/app.exe .

#Macos
env GOOS=darwin GOARCH=amd64 go build -o output/app-darwin-amd64 .
env GOOS=darwin GOARCH=arm64 go build -o output/app-darwin-arm64 .

#Linux
env GOOS=linux GOARCH=amd64 go build -o output/app-linux-amd64 .
