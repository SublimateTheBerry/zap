<h1 align="center">⚡ Zap — Stress Test Tool for APIs & Websites ⚡</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22.7-blue.svg?style=for-the-badge" alt="Go Version">
  <img src="https://img.shields.io/github/release-date/SublimateTheBerry/zap?style=for-the-badge" alt="Release date">
  <img src="https://img.shields.io/badge/Platform-Cross--platform-important?style=for-the-badge" alt="Cross-platform">
</p>

<p align="center">
  <b>Zap</b> is a high-performance, customizable tool for load-testing APIs and websites. Simulate heavy traffic and measure your service's resilience under pressure, with real-time insights into performance metrics like success rates, failures, and response times.
</p>

---

## 🌍 Cross-Platform Compatibility

**Zap** is a cross-platform tool, meaning it can run on:

- 🖥 **Windows**
- 🐧 **Linux**
- 🍏 **macOS**

Instructions for running the executable may vary slightly by platform:

- **Windows**: Run with `zap` in the terminal.
- **Linux / macOS**: Run with `./zap` in the terminal.

---

## 🌟 Features

- 🔄 **Multiple HTTP Methods**: Supports `GET`, `POST`, `PUT`, `DELETE` for versatile testing scenarios.
- ⚡ **Concurrent Requests**: Easily scale up with multiple parallel connections.
- ⏱ **Customizable Test Duration**: Control the test duration to suit your needs.
- 📝 **Custom Headers Support**: Add and modify request headers dynamically.
- 📊 **Detailed Stats**: Track total requests, successes, failures, and average response time.
- 🔇 **Silent Mode**: Suppress request-by-request output for clean and focused testing.

---

<h2>🚀 Quick Start</h2>

Run your first stress test with the following command:

```
# For Linux/macOS
./zap -url <URL> [-c <connections>] [-d <duration>] [-m <method>] [-H <headers>] [--silence | -slc]

# For Windows
zap -url <URL> [-c <connections>] [-d <duration>] [-m <method>] [-H <headers>] [--silence | -slc]
```

### Example:

```
# For Linux/macOS
./zap -url https://example.com/api -c 10 -d 1m -m POST -H "Authorization:Bearer token,Content-Type:application/json"

# For Windows
zap -url https://example.com/api -c 10 -d 1m -m POST -H "Authorization:Bearer token,Content-Type:application/json"
```

---

<h2>🔧 How to Use</h2>

You can use Zap by either downloading the pre-built binaries from the releases or building the executable yourself.

### Option 1: Download Pre-built Executable

1. Go to the [Releases](https://github.com/SublimateTheBerry/zap/releases) page of the repository.
2. Download the latest version of the binary.
3. Run the executable:
   - **Windows**: Run it from the terminal using `zap`.
   - **Linux/macOS**: Make sure the file is executable (`chmod +x zap`), then run `./zap`.

### Option 2: Build from Source

If you'd like to build Zap from source, follow these steps:

1. Ensure that [Go](https://golang.org/dl/) (version 1.22.7 or higher) is installed on your system.
2. Clone the repository:

   ```
   git clone https://github.com/SublimateTheBerry/zap.git
   cd zap
   ```

3. Build the executable:

   ```
   go build -o main.go zap
   ```

4. Run the tool using the appropriate command for your operating system:
   - **Windows**: `zap`
   - **Linux/macOS**: `./zap`

---

<h2>📊 Output</h2>

At the end of the test, Zap provides a detailed summary:

- **Total requests**: The total number of requests sent.
- **Successful responses**: The number of responses with status codes `2xx`.
- **Failed responses**: The number of non-successful responses or failures.
- **Total duration**: The total time taken for all requests.
- **Average response time**: Calculated for successful requests.

---

<h2>🤝 Contributing</h2>

Feel free to open issues or submit pull requests for any improvements, new features, or bug fixes. Contributions are welcome!

---

<h2>📜 License</h2>

Zap is licensed under the Custom License. See the [LICENSE](LICENSE) file for more details.
