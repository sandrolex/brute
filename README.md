 # brute
 
 ## build && run
 ```
 make build
 
 bin/brute --concurrent 0 --good 209 --bad 219 --url http://localhost:8080/login --username Admin --wordlist /usr/share/wordlists/rockyou.txt
 ```

 ## Usage

 ```
 Usage of ./bin/brute:
  -bad int
        Content-Lenght of a non success request (default -1)
  -concurrent int
        Number of concurrent requests
  -good int
        Content-Lenght of a success request (default -1)
  -url string
        URL
  -username string
        Username var (default "Admin")
  -wordlist string
        Wordlist file (default "/usr/share/wordlists/rockyou.txt")
```
 
 
