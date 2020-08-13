package main

import (
	"encoding/json"
    "fmt"
    "os"
    "io/ioutil"
    "time"
    "strings"
    "net/http"
    "strconv"
)

type Credhub_certificates struct
{
    Credhub_certificates      []Cert
}

type Cert struct {
	Issuer     string
	Name       string
	Not_after  string
    Not_before string
}

func sendResponse(payload string) {
    body := strings.NewReader(payload)
    req, err := http.NewRequest("POST", "https://demo.wavefront.com/report?f=wavefront", body)
    if err != nil {
        // handle err
    }
    req.Header.Set("Content-Type", "application/octet-stream")
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Authorization", "Bearer 17eb3188-bdb9-41eb-a2bb-1309212c08a0")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        // handle err
    }
    defer resp.Body.Close()
}

func main() {
    timestamp := time.Now().Unix()
    
    if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }
    data, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Can't read file:", os.Args[1])
        panic(err)
    }
    //fmt.Println("File content is:")
    //fmt.Println(string(data))

    var certs Credhub_certificates
    json.Unmarshal([]byte(data), &certs)

	for _, cert := range certs.Credhub_certificates {
        metricLine := "tas.credhub_certificates.expiration " + cert.Not_after + " " + strconv.FormatInt(timestamp, 10) + " source=citc issuer=" + "\"" + cert.Issuer + "\" not_before=" + "\"" + cert.Not_before + "\" name=" + "\"" + cert.Name + "\""
        fmt.Println(metricLine)
        sendResponse(metricLine)
    }

}