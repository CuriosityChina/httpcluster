/*
Copyright 2014-2015 Rohith and Zhanpeng Chen All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package httpcluster

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	FAKE_HTTP_CLUSTER_URL = "http://127.0.0.1:3000,127.0.0.1:3001"
	FAKE_HEALTH_CHECK_URI = "/"
)

var cluster Cluster
var testServerOne *httptest.Server
var testServerTwo *httptest.Server

func FakePingServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "ok\n")
}

func GetFakeCluster() {
	if cluster == nil {
		cluster, _ = NewHttpCluster(FAKE_HTTP_CLUSTER_URL, FAKE_HEALTH_CHECK_URI)
	}
}

func TestUrl(t *testing.T) {
	GetFakeCluster()
	assert.Equal(t, cluster.Url(), FAKE_HTTP_CLUSTER_URL)
}

func TestSize(t *testing.T) {
	GetFakeCluster()
	assert.Equal(t, cluster.Size(), 2)
}

func TestActive(t *testing.T) {
	GetFakeCluster()
	assert.Equal(t, len(cluster.Active()), 2)
}

func TestNonActive(t *testing.T) {
	GetFakeCluster()
	assert.Equal(t, len(cluster.NonActive()), 0)
}

func TestGetMember(t *testing.T) {
	GetFakeCluster()
	member, err := cluster.GetMember()
	assert.Nil(t, err)
	assert.Equal(t, member, "http://127.0.0.1:3000")
}

func TestMarkdown(t *testing.T) {
	GetFakeCluster()
	assert.Equal(t, len(cluster.Active()), 2)
	cluster.MarkDown()
	time.Sleep(30 * time.Millisecond)
	assert.Equal(t, len(cluster.Active()), 2)
}
