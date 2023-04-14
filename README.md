# pprof-disasm-duplication
This repo provides a minimal replication of a duplicate counting bug in golang's pprof tool with the following setup:
 - `go version go1.20.2 darwin/amd64`
 - `macOS 13.2.1`
 - `x86_64`

Replication steps:
1. `go build .`
2. `./pprof-diasm-duplication`
3. `go tool pprof -weblist=closeEnough cpu.pprof`: observe a flat/cumulative time of 1.37s and 2.05s, respectively
4. `go tool pprof -weblist=closeEnough pprof-diasm-duplication cpu.pprof`: observe a flat/cumulative time of 34.25s and 51.25s, respectively
