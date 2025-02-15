module github.com/flyteorg/boilerplate

go 1.17

require (
	github.com/alvaroloes/enumer v1.1.2
	github.com/flyteorg/flytestdlib v0.3.22
	github.com/golangci/golangci-lint v1.38.0
	github.com/pseudomuto/protoc-gen-doc v1.4.1 // indirect
	github.com/vektra/mockery v0.0.0-20181123154057-e78b021dcbb5
)

replace github.com/vektra/mockery => github.com/enghabu/mockery v0.0.0-20191009061720-9d0c8670c2f0

replace github.com/pseudomuto/protoc-gen-doc => github.com/flyteorg/protoc-gen-doc v1.4.2
