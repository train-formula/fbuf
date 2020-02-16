#!/bin/bash

set -e

rm -rf ./generated
rm -rf ./fbuf-go/*
mkdir -p ./generated

cd ./proto

proto_folders=(sitemanager)

for protos in ${proto_folders[@]}; do \
	protoc -I. --go_out=plugins=grpc:../generated/ ${protos}/*.proto
done

cd ..

mv ./generated/github.com/train-formula/fbuf-go/* ./fbuf-go/

rm -rf ./generated
