FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /valorant-roulette

EXPOSE 8080

CMD [ "/valorant-roulette" ]