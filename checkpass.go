package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func checkHIBP(password string) int {

	h := sha1.New()

	io.WriteString(h, password)
	ourHashBytes := h.Sum(nil)
	ourHash := strings.ToUpper(hex.EncodeToString(ourHashBytes))

	key := string(ourHash[:5])
	target := string(ourHash[5:])
	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + key)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseString := ""

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		responseString = string(bodyBytes)

		if err2 != nil {
			panic(err)
		}
	}

	responses := strings.Split(responseString, "\n")

	for _, response := range responses {
		response = strings.TrimSuffix(response, "\r")
		rows := strings.Split(response, ":")
		hash := rows[0]

		if hash == target {
			count, _ := strconv.Atoi(rows[1])
			return count
		}
	}

	return 0

}

func formatResult(count int) string {
	verb := "has NOT been found!  Congrats....."
	usage := " "
	plural := "s"
	if count > 0 {
		if count == 1 {
			plural = ""
		}
		verb = "has been compromised and shouldn't be used any longer.\n"
		verb += "You can read about how to select a good password at "
		verb += "https://runkel.org/2017/09/how-to-pick-a-password/\n"
		usage = "\n\nIt has been found " + strconv.Itoa(count) + " time" + plural + " on the dark web.\n\n\n"
	}

	return fmt.Sprintf("Your password %s%s", verb, usage)
}

func getPassword() string {
	password := ""
	for password == "" {
		fmt.Print("Enter the password to test: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}
		password = string(bytePassword)
	}

	return password
}

func main() {

	if len(os.Args) > 2 {

		fmt.Printf("Usage: " + os.Args[0] + " PasswordToCheck\n\n\n")

	} else {
		password := ""
		if len(os.Args) == 1 {
			password = getPassword()
		} else {
			password = os.Args[1]
		}

		fmt.Print("\n\nChecking the password against the haveibeenpwnd service.\n\n")

		count := checkHIBP(password)

		fmt.Print(formatResult(count))

	}

}
