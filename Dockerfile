FROM golang:1.22.5 AS build
ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY . .

RUN go build


FROM gcr.io/distroless/static-debian12

WORKDIR /app/

COPY --from=build /go/tg-notify .

