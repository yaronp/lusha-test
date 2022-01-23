#!/usr/bin/env bash
# Stops the process if something fails
set -xe

go get "github.com/gin-gonic/gin"
go get "github.com/mailjet/mailjet-apiv3-go"

go build -o bin/application
