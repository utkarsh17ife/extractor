# windows build
set GOOS=windows
set GOARCH=amd64
go build -o ./build/extractor.exe main.go