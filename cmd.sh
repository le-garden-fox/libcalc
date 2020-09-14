#!/usr/bin/env bash

function main() {
    case $1 in
    "gen")
        protoc libcalc/libcalc.proto --go_out=plugins=grpc:.
        ;;
    *)
        echo "Invalid arguments"
        ;;
    esac
}

main "$@"
