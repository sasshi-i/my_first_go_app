FROM golang:latest

RUN mkdir /go/src/my_first_app
WORKDIR /go/src/my_first_app
RUN GO111MODULE=off go get -u github.com/motemen/gore/cmd/gore
ADD . /go/src/my_first_app
