name: ci-golang-workflow

on:
  pull_request:
    branches:
      - develop

jobs:
  check-application:
    runs-on: ubuntu-latest

    strategy:
      matrix:
        go: ['1.14', '1.15']
    
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: ${{ matrix.go }} # will run this code twice to test the code with the both go version
      
      - run: go test
      - run: go run math.go
