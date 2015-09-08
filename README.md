# httpcluster
Make multiple http api endpoint work like a cluster

[![Build Status](https://travis-ci.org/lowstz/httpcluster.svg)](https://travis-ci.org/lowstz/httpcluster) [![GoDoc](https://godoc.org/github.com/lowstz/httpcluster?status.svg)](https://godoc.org/github.com/lowstz/httpcluster)

## Usage
``` go

package main

import (
	"log"

	"github.com/lowstz/httpcluster"
)

func main() {
	var httpClusterUrl = "http://192.168.1.2:8000,http://192.168.1.3:8000"
	cluster, err := httpcluster.NewHttpCluster(httpClusterUrl, "/ping")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	log.Printf("cluster size: %d", cluster.Size())
}
```
