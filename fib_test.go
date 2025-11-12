package fib

import (
	"testing"
)

func BenchmarkFib20WithBlockchainMetrics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ = Fib(20)
	}

	b.ReportMetric(0, "ns/op")
	b.ReportMetric(1.0, "block_accept_ms/ggas")
	b.ReportMetric(2.0, "block_parse_ms/ggas")
	b.ReportMetric(3.0, "block_verify_ms/ggas")
	b.ReportMetric(4.0, "mgas/s")
	b.ReportMetric(5.0, "ms/ggas")
}
