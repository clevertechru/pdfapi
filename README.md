This is xhtmltopdf wrap api

# Install
```
brew install dep
dep ensure
go run server.go

go build server.go && ./server
```
# Use
```
curl -X POST "http://localhost:1323/pdf?url=yandex.ru" -d '{}' --output report.pdf
curl -X POST "http://localhost:1323/pdf" -d '{"url":"yandex.ru"}' -H 'Content-Type: application/json' --output report.pdf
```
