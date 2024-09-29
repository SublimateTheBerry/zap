package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Stats struct {
	TotalRequests      int
	Successful         int
	Failed             int
	TotalDuration      time.Duration
	TotalResponseTime  time.Duration
}

func sendRequest(method, url string, headers map[string]string, body []byte, stats *Stats, silence bool) {
	defer wg.Done()
	start := time.Now()

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		stats.Failed++
		return
	}
	defer resp.Body.Close()

	duration := time.Since(start)
	stats.TotalDuration += duration
	stats.TotalRequests++

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		stats.Successful++
		stats.TotalResponseTime += duration
	} else {
		stats.Failed++
	}

	if !silence {
		fmt.Printf("Response status: %s, Time taken: %v\n", resp.Status, duration)
	}
}

func main() {
	var url string
	var connections int
	var duration time.Duration
	var method string
	var headersInput string
	var silence bool

	flag.StringVar(&url, "url", "", "URL to test (required)")
	flag.IntVar(&connections, "c", 1, "Number of parallel connections (optional, default: 1)")
	flag.DurationVar(&duration, "d", 30*time.Second, "Duration of the test (optional, default: 30s)")
	flag.StringVar(&method, "m", "GET", "HTTP method (GET, POST, PUT, DELETE) (optional, default: GET)")
	flag.StringVar(&headersInput, "H", "", "Headers in format 'Key1:Value1,Key2:Value2' (optional)")
	flag.BoolVar(&silence, "silence", false, "Suppress output of individual request results (optional)")
	flag.BoolVar(&silence, "slc", false, "Suppress output of individual request results (optional)")

	flag.Parse()

	if url == "" {
		fmt.Println("Usage: (./)zap -url <URL> -c <connections> -d <duration> -m <method> -H <headers> [--silence|-slc]")
		fmt.Println("  -url: URL to test (required)")
		fmt.Println("  -c: Number of parallel connections (optional, default: 1)")
		fmt.Println("  -d: Duration of the test (optional, default: 30s)")
		fmt.Println("  -m: HTTP method (GET, POST, PUT, DELETE) (optional, default: GET)")
		fmt.Println("  -H: Headers in format 'Key1:Value1,Key2:Value2' (optional)")
		fmt.Println("  --silence, -slc: Suppress output of individual request results (optional)")
		return
	}

	headers := make(map[string]string)
	if headersInput != "" {
		for _, header := range split(headersInput, ",") {
			parts := split(header, ":")
			if len(parts) == 2 {
				headers[parts[0]] = parts[1]
			}
		}
	}

	stats := &Stats{}
	startTime := time.Now()
	for time.Since(startTime) < duration {
		for i := 0; i < connections; i++ {
			wg.Add(1)
			go sendRequest(method, url, headers, nil, stats, silence)
		}
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()

	fmt.Println("Load testing completed.")
	fmt.Printf("Total requests: %d\n", stats.TotalRequests)
	fmt.Printf("Successful responses: %d\n", stats.Successful)
	fmt.Printf("Failed responses: %d\n", stats.Failed)
	fmt.Printf("Total duration: %v\n", stats.TotalDuration)

	if stats.Successful > 0 {
		averageResponseTime := stats.TotalResponseTime / time.Duration(stats.Successful)
		fmt.Printf("Average response time: %v\n", averageResponseTime)
	} else {
		fmt.Println("No successful responses to calculate average response time.")
	}
}

func split(s, sep string) []string {
	var result []string
	for _, part := range strings.Split(s, sep) {
		result = append(result, strings.TrimSpace(part))
	}
	return result
}
