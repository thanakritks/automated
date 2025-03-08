# Automated DevSecOps with Golang

This project demonstrates a simple Golang web server integrated with a DevSecOps pipeline using GitHub Actions. The pipeline includes static code analysis, container security scanning, and dynamic application security testing (DAST).

## Project Structure

├── Dockerfile  
├── go.mod  
├── main.go  
├── readme.md  
└── .github/  
    └── workflows/  
        └── security_scan.yml

## Getting Started

### Prerequisites

- Golang installed on your local machine
- Docker installed on your local machine
- GitHub account

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/automated-devsecops-golang.git
    cd automated-devsecops-golang
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Web Server

To run the web server locally, use the following command:
```sh
go run main.go
```

The server will start on `http://localhost:8080`.

### Building the Docker Image

To build the Docker image, use the following command:
```sh
docker build -t yourusername/automated-devsecops-golang .
```

### Running the Docker Container

To run the Docker container, use the following command:
```sh
docker run -p 8080:8080 yourusername/automated-devsecops-golang
```

## GitHub Actions Pipeline

This project uses GitHub Actions for CI/CD. The pipeline is defined in `.github/workflows/security_scan.yml` and includes the following steps:

1. **Static Code Analysis**: Uses `golangci-lint` to perform static code analysis.
2. **Container Security Scanning**: Uses `Trivy` to scan the Docker image for vulnerabilities.
3. **Dynamic Application Security Testing (DAST)**: Uses `OWASP ZAP` to perform dynamic security testing on the running application.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License.

