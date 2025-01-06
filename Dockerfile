# Используем официальный образ Go как базовый для сборки
FROM golang:1.23.3-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Устанавливаем gcc и другие зависимости для сборки
RUN apk add --no-cache gcc musl-dev

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники приложения в рабочую директорию
COPY . .

# Собираем приложение
RUN go build -o geoService-service ./cmd/main.go

# Начинаем новую стадию на основе минимального образа
FROM alpine:latest

# Копируем исполняемый файл из первой стадии
COPY --from=builder /app/geoService-service /geoService-service

# Копируем .env файл, если приложение использует его
COPY .env /app/.env

# Открываем порт для GRPC
EXPOSE 30303

# Запускаем приложение
CMD ["./geoService-service"]
