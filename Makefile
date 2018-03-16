GOARCH=amd64

all : linux osx windows

linux : 
	GOOS=linux go build -o build/linux-serv serv.go

osx : 
	GOOS=darwin go build -o build/osx-serv serv.go

windows : 
	GOOS=windows go build -o build/windows-serv.exe serv.go

clean : 
	rm -rf build