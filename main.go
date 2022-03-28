package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"os/exec"
	"os"
	"strings"
	"flag"
)

type SecReq struct {
		Location string
		Sec string
}

func main() {
		var r SecReq
		var key string
		var path string

		flag.StringVar(&key, "k", "key", "GPG Key")
		flag.StringVar(&path, "p", "path", "Secret path")

		flag.Parse()

		paths := strings.Split(path, "#")

		if len(paths) != 2 {
				fmt.Printf("%s is invalid", path)
				os.Exit(1)
		}

		r.Location = paths[0]
		r.Sec = paths[1]

		a := getFile(r)

		pkey := getPrivateKey(key)

		sec, err := helper.DecryptMessageArmored(pkey, nil, a)
		sec = strings.Replace(sec, "\n", "", -1)
		check(err)

		fmt.Printf("%s", sec)
}

func check(err error) {
		if err != nil {
				log.Fatal(err)
		}
}

func getFile(r SecReq) string {
		req := fmt.Sprintf("https://%s/_secret/%s", r.Location, r.Sec)
		resp, err := http.Get(req)
		check(err)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		check(err)
		return string(body)
}

func getPrivateKey(sig string) string {
		out, err :=  exec.Command("gpg", "--export-secret-subkeys", "-a", sig).Output()
		check(err)
		return string(out)
}

