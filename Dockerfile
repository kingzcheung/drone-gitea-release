FROM golang:1.14.2-alpine3.11 as builder

WORKDIR /release/linux/amd64/

COPY . .

RUN export GOOS=linux \
    && export GOARCH=amd64 \
    && export CGO_ENABLED=0 \
    && export GO111MODULE=on \
    && go mod tidy \
    && go build -v -a -tags netgo -o /release/linux/amd64/drone-gitea-release

FROM alpine:3.10

RUN apk --no-cache add ca-certificates

COPY --from=builder /release/linux/amd64/drone-gitea-release /bin/

ENTRYPOINT ["/bin/drone-gitea-release"]