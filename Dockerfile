# syntax=docker/dockerfile:1.2
FROM alpine:3.21

RUN apk add --no-cache --no-progress ca-certificates tzdata

ARG TARGETPLATFORM
COPY .bin/$TARGETPLATFORM/go-hello /go-hello

VOLUME ["/tmp"]

ENTRYPOINT ["/go-hello"]