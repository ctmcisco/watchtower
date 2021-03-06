package mocks

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
)

// NewMockAPIServer returns a mocked docker api server that responds to some fixed requests
// used in the test suite.
func NewMockAPIServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logrus.Debug("Mock server has received a HTTP call on ", r.URL)
			var response = ""

			if isRequestFor("filters=%7B%22status%22%3A%7B%22running%22%3Atrue%7D%7D&limit=0", r) {
				response = getMockJSONFromDisk("./mocks/data/containers.json")
			} else if isRequestFor("filters=%7B%22status%22%3A%7B%22created%22%3Atrue%2C%22exited%22%3Atrue%2C%22running%22%3Atrue%7D%7D&limit=0", r) {
				response = getMockJSONFromDisk("./mocks/data/containers.json")
			} else if isRequestFor("containers/json?limit=0", r) {
				response = getMockJSONFromDisk("./mocks/data/containers.json")
			} else if isRequestFor("ae8964ba86c7cd7522cf84e09781343d88e0e3543281c747d88b27e246578b65", r) {
				response = getMockJSONFromDisk("./mocks/data/container_stopped.json")
			} else if isRequestFor("b978af0b858aa8855cce46b628817d4ed58e58f2c4f66c9b9c5449134ed4c008", r) {
				response = getMockJSONFromDisk("./mocks/data/container_running.json")
			} else if isRequestFor("sha256:19d07168491a3f9e2798a9bed96544e34d57ddc4757a4ac5bb199dea896c87fd", r) {
				response = getMockJSONFromDisk("./mocks/data/image01.json")
			} else if isRequestFor("sha256:4dbc5f9c07028a985e14d1393e849ea07f68804c4293050d5a641b138db72daa", r) {
				response = getMockJSONFromDisk("./mocks/data/image02.json")
			}
			fmt.Fprintln(w, response)
		},
	))
}

func isRequestFor(urlPart string, r *http.Request) bool {
	return strings.Contains(r.URL.String(), urlPart)
}

func getMockJSONFromDisk(relPath string) string {
	absPath, _ := filepath.Abs(relPath)
	logrus.Error(absPath)
	buf, err := ioutil.ReadFile(absPath)
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return string(buf)
}
