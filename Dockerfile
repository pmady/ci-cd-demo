# ---- Stage 1: Build ----
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

ARG VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -X main.version=${VERSION}" -o ci-cd-demo .

# ---- Stage 2: Runtime ----
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/ci-cd-demo /ci-cd-demo

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/ci-cd-demo"]
