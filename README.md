# R1CS proto

- go get -u github.com/golang/dep/cmd/dep
- dep ensure
- go get -u github.com/QED-it/r1cs_proto 
- to generate protos, run *cd $GOPATH/src/github.com/QED-it/r1cs_proto && protoc --go_out=$GOPATH/src \*.proto*
- go run saver.go
