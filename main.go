package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	version   = "dev"
	startTime = time.Now()
)

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
}

type InfoResponse struct {
	App         string `json:"app"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Event       string `json:"event"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/ready", readyHandler)
	mux.HandleFunc("/info", infoHandler)

	log.Printf("Starting ci-cd-demo server v%s on port %s", version, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <title>CI/CD Demo - Cloud Native Rabat</title>
    <style>
        body { font-family: 'Segoe UI', sans-serif; background: linear-gradient(135deg, #0f0c29, #302b63, #24243e); color: #fff; display: flex; justify-content: center; align-items: center; min-height: 100vh; margin: 0; }
        .card { background: rgba(255,255,255,0.1); backdrop-filter: blur(10px); border-radius: 16px; padding: 40px; max-width: 600px; text-align: center; border: 1px solid rgba(255,255,255,0.2); }
        h1 { font-size: 2.2em; margin-bottom: 10px; }
        .version { color: #00d4ff; font-size: 1.1em; margin-bottom: 20px; }
        .event { color: #ff6b6b; font-size: 1.2em; font-weight: bold; }
        .endpoints { text-align: left; margin-top: 30px; background: rgba(0,0,0,0.3); border-radius: 8px; padding: 20px; }
        .endpoints h3 { color: #00d4ff; }
        .endpoints code { color: #ffd93d; }
        a { color: #00d4ff; text-decoration: none; }
    </style>
</head>
<body>
    <div class="card">
        <h1>CI/CD Demo</h1>
        <p class="event">Cloud Native Rabat</p>
        <p class="version">Version: %s</p>
        <p>From Laptop to Production: The Cloud Native Way</p>
        <div class="endpoints">
            <h3>Available Endpoints</h3>
            <p><code>GET /</code> - This page</p>
            <p><code>GET /health</code> - Liveness probe</p>
            <p><code>GET /ready</code> - Readiness probe</p>
            <p><code>GET /info</code> - Application info</p>
        </div>
    </div>
</body>
</html>`, version)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := HealthResponse{
		Status:  "ok",
		Version: version,
		Uptime:  time.Since(startTime).Round(time.Second).String(),
	}
	json.NewEncoder(w).Encode(resp)
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ready"})
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := InfoResponse{
		App:         "ci-cd-demo",
		Version:     version,
		Description: "End-to-end CI/CD demo for Cloud Native Rabat",
		Event:       "Cloud Native Rabat",
	}
	json.NewEncoder(w).Encode(resp)
}
