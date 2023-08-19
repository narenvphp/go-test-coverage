#!/usr/bin/env bash

set -xe

mkdir -p -- "coverage"

go test ./size/ --cover -coverprofile=./coverage/all_tests.out
