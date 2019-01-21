FROM golang:alpine

LABEL maintainer="peace0phmind@gmail.com"

RUN apk add --no-cache git \
        build-base

