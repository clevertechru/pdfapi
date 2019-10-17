This is xhtmltopdf wrap api

# Install
Mac OS
```
brew install dep
dep ensure
go run server.go

go build server.go && ./server
```
Debian
```
sudo apt-get update
sudo apt-get -y upgrade

wget https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz
tar -xvf go1.13.1.linux-amd64.tar.gz
sudo mv go /usr/local

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

mkdir -p $HOME/go/src/github.com/clevertechru/
cd $HOME/go/src/github.com/clevertechru/
git clone https://github.com/clevertechru/pdfapi.git && cd ./pdfapi

curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
chmod +x $HOME/go/src/github.com/clevertechru/bin/dep
sudo mv $HOME/go/src/github.com/clevertechru/bin/dep /usr/local/go/bin/
dep ensure

go build server.go && ./server

```
# Use
```
curl -X POST "http://localhost:1323/pdf?url=yandex.ru" -d '{}' --output report.pdf
curl -X POST "http://localhost:1323/pdf" -d '{"url":"yandex.ru"}' -H 'Content-Type: application/json' --output report.pdf
```
