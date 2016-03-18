BIN=linux32 linux64 windows32.exe windows64.exe osx32 osx64
all:$(BIN)

linux32: passgen.go
	env GOOS=linux GOARCH=386 go build
	mv passgen $@

linux64: passgen.go
	env GOOS=linux GOARCH=amd64 go build
	mv passgen $@

windows32.exe: passgen.go
	env GOOS=windows GOARCH=386 go build
	mv passgen.exe $@

windows64.exe: passgen.go
	env GOOS=windows GOARCH=amd64 go build
	mv passgen.exe $@

osx32: passgen.go
	env GOOS=darwin GOARCH=386 go build
	mv passgen $@

osx64: passgen.go
	env GOOS=darwin GOARCH=amd64 go build
	mv passgen $@

clean:
	rm -rf $(BIN)
