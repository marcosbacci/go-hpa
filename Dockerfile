FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/go/hpa

COPY src/hpa/*.go .

ENV CGO_ENABLED 0
RUN go get -d -v
RUN go build -ldflags="-s -w" -o /go/bin

FROM scratch

COPY --from=builder /go/bin/hpa /hpa

ENTRYPOINT ["/hpa"]
