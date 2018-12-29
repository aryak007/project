FROM golang as builder

ENV GO111MODULE=on

WORKDIR /proj

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN ls
FROM alpine
COPY --from=builder /proj/project /proj/
RUN chmod a+x  /proj/project
ENTRYPOINT [ "/proj/project" ]
