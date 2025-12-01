package fib

import (
	"testing"
)

func TestCustomBenchmark(t *testing.T) {
	// Run the benchmark
	for i := 0; i < 10; i++ {
		_ = Fib(20)
	}

	// Create benchmark tool and add results
	tool := NewBenchmarkTool()
	tool.AddResult("Alpha", 1.1, "block_accept_ms/ggas")
	tool.AddResult("Beta", 2.2, "block_parse_ms/ggas")
	tool.AddResult("Charlie", 3.3, "block_verify_ms/ggas")
	tool.AddResult("Delta", 4.0, "mgas/s")

	// Save results to file
	if err := tool.SaveToFile("benchmark_output.json"); err != nil {
		t.Fatalf("Failed to save benchmark results: %v", err)
	}
}
