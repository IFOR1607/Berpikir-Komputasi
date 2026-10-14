# STAGE 1: Compile kode Go
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Salin file modul dan install dependensi (jika ada)

# Salin seluruh kode program
COPY . .

# Compile project Go menjadi binary statis (bernama 'main')
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# STAGE 2: Jalankan aplikasi menggunakan image minimal
FROM alpine:latest  

WORKDIR /root/

# Salin hasil compile dari stage builder
COPY --from=builder /app/main .

# Expose port aplikasi Anda (sesuaikan dengan port di main.go)
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
