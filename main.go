package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const(
	configExtension = ".cfg"
	configFile = "default" + configExtension
)

func MacaddressTofilename(mac string) string{
	if mac != "" {
		return strings.Replace(strings.ToUpper(mac), ":", "", -1) + configExtension
	}
	return ""
}

func GetMacaddress(ip string) string{
	file, err := os.Open("/proc/net/arp")
	if err != nil {
		return ""
	}
	defer file.Close()

	lines := bufio.NewScanner(file)
	for lines.Scan() {
		fields := strings.Fields(lines.Text())
		if fields[0] == ip {
			return fields[3]
		}
	}
	return ""
}

func GenerateConfig(file string) (*template.Template, error) {
	fileName := file
	fileExists := false
	if file !="" {
		_, err := os.Stat(fileName)
		if err == nil {
			fileExists = true
			fmt.Println("something happened")
		}
	}
	if fileExists {
		return template.ParseFiles(configFile, fileName)
	}
	return template.ParseFiles(configFile)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	ip := strings.Split(r.RemoteAddr, ":")[0]
	mac := GetMacaddress(ip)
	file := MacaddressTofilename(mac)


	tmpl, err := template.ParseFiles(configFile)
	if err != nil{
		fmt.Println(err.Error())
	}
	tmpl2, err := template.ParseFiles(file)
	if err != nil {
		tmpl.New("device").Parse("")
	} else {
		tmpl.New("device").Parse(tmpl2.Root.String())
	}
	tmpl.Execute(w, nil)
}

func main() {
	 http.HandleFunc("/test", handleRequest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}