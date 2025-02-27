FROM golang

WORKDIR /go/src/github.com/QuanTran91/ksm

ADD . /go/src/github.com/QuanTran91/ksm

RUN go get -t -v ./...

CMD ["go", "run", "example/basic.go"]

EXPOSE 8080