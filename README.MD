# TCP Line Reader Server

A simple TCP server written in Go that listens on port 42069 and reads incoming data line by line.

## Overview

This server accepts TCP connections and processes incoming data in a streaming fashion, reading and printing each line as it arrives. The implementation uses a custom line reader that reads data in chunks and splits it by newline characters.

## Features

- TCP server listening on port 42069
- Concurrent line reading using Go channels
- Efficient buffered reading with 8-byte chunks
- Handles multiple connections sequentially

## Usage

Run the TCP listener:

```bash
go run ./cmd/tcplistener
```

To capture the output to a file while also displaying it in the terminal:

```bash
go run ./cmd/tcplistener | tee tcplistener.txt
```

### ⚠️ Security Warning

**Do not share `tcp.txt` or `tcplistener.txt` files!** These files may contain sensitive data captured from TCP connections.

Connect to the server using any TCP client (e.g., `nc`, `telnet`):

```bash
echo "Hello, World!" | nc localhost 42069
```

Or use verbose mode to see connection details:

```bash
nc -v localhost 42069
```

## How It Works

The server:
1. Listens for TCP connections on port 42069
2. Accepts incoming connections
3. Reads data line by line using a channel-based reader
4. Prints each line to stdout
5. Closes the connection when the client disconnects
