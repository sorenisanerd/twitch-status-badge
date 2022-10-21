FROM golang:1.19 AS build

WORKDIR /twitch-status

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -mod=readonly -v -o /app

# Now create separate deployment image
FROM gcr.io/distroless/base
ENV GOTRACEBACK=single

WORKDIR /twitch-status
COPY --from=build /app ./app
COPY *.svg ./

ENTRYPOINT ["./app"]