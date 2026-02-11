# CI/CD Demo - Cloud Native Rabat

> **From Laptop to Production: The Cloud Native Way**

An end-to-end demo project showcasing the journey of an application from a developer's laptop to a live, monitored service using containers, CI/CD pipelines, and GitOps.

## Architecture

```
Developer Push → GitHub Actions (CI) → GHCR (Container Registry) → ArgoCD (CD) → Kubernetes
```

### The Pipeline Flow

1. **Developer** pushes code to `main` branch
2. **GitHub Actions** runs tests, builds the container image, and pushes it to GitHub Container Registry (GHCR)
3. **GitHub Actions** updates the image tag in `k8s/deployment.yaml` (GitOps pattern)
4. **ArgoCD** detects the manifest change and syncs the new deployment to Kubernetes

## Project Structure

```
ci-cd-demo/
├── main.go                    # Go web application
├── main_test.go               # Unit tests
├── go.mod                     # Go module
├── Dockerfile                 # Multi-stage container build
├── .dockerignore
├── .github/
│   └── workflows/
│       └── ci.yaml            # GitHub Actions CI pipeline
├── k8s/                       # Kubernetes manifests
│   ├── namespace.yaml
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
└── argocd/
    └── application.yaml       # ArgoCD Application resource
```

## The Application

A lightweight Go web server with the following endpoints:

| Endpoint   | Description          |
|------------|----------------------|
| `GET /`    | Landing page         |
| `GET /health` | Liveness probe    |
| `GET /ready`  | Readiness probe   |
| `GET /info`   | App metadata (JSON) |

## Quick Start

### Run Locally

```bash
go run main.go
# Visit http://localhost:8080
```

### Run Tests

```bash
go test -v ./...
```

### Build Container Image

```bash
docker build --build-arg VERSION=v1.0.0 -t ci-cd-demo:v1.0.0 .
docker run -p 8080:8080 ci-cd-demo:v1.0.0
```

## CI/CD Pipeline (GitHub Actions)

The pipeline (`.github/workflows/ci.yaml`) runs on every push to `main`:

1. **Test** - Runs `go test` with race detection and coverage
2. **Build & Push** - Builds the Docker image and pushes to `ghcr.io/pmady/ci-cd-demo`
3. **Update Manifest** - Updates the image tag in `k8s/deployment.yaml` and commits the change

## GitOps with ArgoCD

### Prerequisites

- A Kubernetes cluster (kind, minikube, or cloud)
- ArgoCD installed ([install guide](https://argo-cd.readthedocs.io/en/stable/getting_started/))

### Deploy with ArgoCD

```bash
# Apply the ArgoCD Application
kubectl apply -f argocd/application.yaml

# ArgoCD will automatically sync and deploy the app
# Check status
argocd app get ci-cd-demo
```

### Manual Kubernetes Deploy (without ArgoCD)

```bash
kubectl apply -f k8s/namespace.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/ingress.yaml
```

## Event

**Cloud Native Rabat** - Demonstrating the complete journey of an application from a developer's laptop to a live, monitored Kubernetes service.

### Talk: The CI/CD Guru

This demo covers:
- Automating container builds with **GitHub Actions**
- Pushing images to **GitHub Container Registry (GHCR)**
- Deploying to **Kubernetes** using the **GitOps** pattern
- Continuous delivery with **ArgoCD** (auto-sync, self-heal, prune)

## License

MIT
