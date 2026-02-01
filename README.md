# 🪓 WIP - Lumberjack

> **⚠️ WORK IN PROGRESS (WIP):** This project is part of an ongoing Go concurrency lab. Features, architecture, and documentation are subject to frequent changes as I optimize the ingestor.

**Lumberjack** is a high-throughput TCP log ingestor built in Go. It is engineered to "fell" massive amounts of incoming log data, processing thousands of concurrent streams and efficiently stacking them into persistent storage without breaking a sweat.

## 🌲 The Objective
The core challenge of this project is managing **10,000 concurrent TCP connections**. In a naive implementation, slow disk I/O would back up the entire system. **Lumberjack** solves this by using a decoupled architecture:

* **The Scalers:** Lightweight Goroutines handling the heavy lifting of network ingestion.
* **The Flume:** Buffered channels acting as a high-speed transit layer to prevent back-pressure.
* **The Sawmill:** A controlled **Worker Pool** of file writers that manage disk I/O without starving the network layer.

## 🛠 Technical Blueprint

### Concurrency Model
Lumberjack implements a producer-consumer pattern to ensure that network I/O and disk I/O operate at their own optimal speeds:
* **Ingestion:** An `Accept()` loop spawning one Goroutine per connection.
* **Decoupling:** A global buffered channel to prevent "head-of-line blocking."
* **Persistence:** A fixed-size Worker Pool to prevent "disk thrashing" and limit file descriptor overhead.

### Performance Profiling
We don't guess; we measure. Lumberjack is instrumented with `pprof` to visualize performance bottlenecks.
* **CPU Flame Graphs:** Used to identify and optimize expensive parsing operations.
* **Heap Profiles:** Used to ensure the 10,000 connections are managed with minimal memory overhead.

---

## 🚀 Getting Started

### Installation
```bash
git clone [https://github.com/tjozsa/Lumberjack.git](https://github.com/tjozsa/Lumberjack.git)
cd Lumberjack
```
