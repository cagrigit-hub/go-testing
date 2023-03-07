t:
	go test .

ct:
	go test -cover .

cto:
	go test -coverprofile=coverage.out .

cth:
	go tool cover -html=coverage.out

cc:
	make cto
	make cth

testcycle:
	make t
	make cto
	make cth