build:
	GOOS=linux GOARCH=mipsle go build -o blink -ldflags '-s' main.go
	
copy: 
	rsync -P -a blink root@omega-5d69.local:/root/go