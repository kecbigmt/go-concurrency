## poolを使わない場合のベンチマーク

```bash
kecy@macbookair benchmark-network-service % go test ./slow -benchtime=10s -bench=.
goos: darwin
goarch: arm64
pkg: github.com/kecbigmt/go-concurrency/gos-concurrency-building-blocks/the-sync-package/pool/benchmark-network-service/slow
BenchmarkSlowNetworkRequest-8   	      10	1005248033 ns/op
PASS
ok  	github.com/kecbigmt/go-concurrency/gos-concurrency-building-blocks/the-sync-package/pool/benchmark-network-service/slow	11.361s
```

## poolを使った場合のベンチマーク

```bash
kecy@macbookair benchmark-network-service % go test ./fast -benchtime=10s -bench=.
goos: darwin
goarch: arm64
pkg: github.com/kecbigmt/go-concurrency/gos-concurrency-building-blocks/the-sync-package/pool/benchmark-network-service/fast
BenchmarkFastNetworkRequest-8   	    8618	   4245597 ns/op
PASS
ok  	github.com/kecbigmt/go-concurrency/gos-concurrency-building-blocks/the-sync-package/pool/benchmark-network-service/fast	56.565s
```

