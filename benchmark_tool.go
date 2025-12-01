package fib

import (
	"encoding/json"
	"os"
	"strconv"
)

type BenchmarkResult struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type BenchmarkTool struct {
	benchmarkResults []BenchmarkResult
}

func NewBenchmarkTool() *BenchmarkTool {
	return &BenchmarkTool{
		benchmarkResults: make([]BenchmarkResult, 0),
	}
}

func (b *BenchmarkTool) AddResult(name string, value float64, unit string) {
	result := BenchmarkResult{
		Name:  name,
		Value: strconv.FormatFloat(value, 'f', -1, 64),
		Unit:  unit,
	}
	b.benchmarkResults = append(b.benchmarkResults, result)
}

func (b *BenchmarkTool) SaveToFile(filePath string) error {
	output, err := json.MarshalIndent(b.benchmarkResults, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, output, 0644)
}
