#!/bin/bash
protoc --proto_path=${pwd}:. --micro_out=. --go_out=. proto/account/*.proto
protoc --proto_path=${pwd}:. --micro_out=. --go_out=. proto/customer/*.proto
protoc --proto_path=${pwd}:. --micro_out=. --go_out=. proto/roles/*.proto