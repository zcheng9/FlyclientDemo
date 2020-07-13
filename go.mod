module github.com/zcheng9/FlyclientDemo

go 1.13

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200429183012-4b2356b1ed79

require (
	github.com/aristanetworks/goarista v0.0.0-20200513152637-638451432ae4
	github.com/arnaucube/go-snark v0.0.4
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/dchest/siphash v1.2.1
	github.com/go-stack/stack v1.8.0
	github.com/influxdata/influxdb v1.8.0
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/peterh/liner v1.2.0
	github.com/stretchr/testify v1.5.1
	github.com/syndtr/goleveldb v1.0.0
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/sys v0.0.0-20200219091948-cb0a6d8edb6c
)
