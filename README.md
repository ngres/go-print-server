<p align="center">
    <img src="https://raw.githubusercontent.com/Nico0302/go-print-server/main/art.svg" height="200">
</p>
<p align="center">
    Web Printing API server written in Go
</p>

## Usage

This print server provides an API endpoint (optionally protected by basic auth) which enables users to print a PDF file from a given url.

```bash
curl -X POST http://localhost:8631/v1/print/url \
   -u 'example:secret' \
   -H 'Content-Type: application/json' \
   -d '{"url":"https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf","preset":"example-preset"}'
```

If [Typst](https://typst.app/open-source/#download) is installed, the incoming print files (PDF, SVG, PNG, JPG, GIF, WEBP) can also be wrapped in a Typst document, enabling format conversions and advanced layouts. Checkout the [config.example.yml](./config/config.example.yml) for more information.

## Setup

1. Clone this repo.
2. `$ go mod download`
3. `$ cp ./config/config.example.yml ./config/config.yml`
4. Fill out your _config.yml_.
5. `$ go run ./cmd/print-server/main.go`
