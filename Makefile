.PHONY: build run clean

build:
	go mod tidy
	go build -o tunnel main.go

# 启动服务端 (在洛杉矶服务器运行)
server:
	./tunnel -config config/server.json

# 启动客户端 (在本地运行)
client:
	./tunnel -config config/client.json

clean:
	rm -f tunnel
