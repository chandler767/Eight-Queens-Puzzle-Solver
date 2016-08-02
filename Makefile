all:
	go build -o queen_solver_mac main.go
	GOOS=linux GOARCH=amd64 go build -o queen_solver main.go
	GOOS=windows GOARCH=386 go build -o queen_solver.exe main.go

