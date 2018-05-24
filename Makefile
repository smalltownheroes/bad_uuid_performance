# linux: main.go
linux: perf.go
	GOOS=linux go build -o $@ .