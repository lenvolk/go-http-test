FROM golang:1.12.0-alpine3.9

WORKDIR /go/src/app
COPY *.go .
RUN go get -d -v ./....
RUN go install -v ./....

EXPOSE 8090

ENV FOO=BAR

CMD ["app"]