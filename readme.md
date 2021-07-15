# Minio API (Golang)

This repo contains API for Minio bucket operations.

Operations: Create, List and Delete M

#Installation

Clone this repo

#Usage

``` bash
$ curl -X GET http://IP_ADDRESS:8081/bucket/list

$ curl -X POST http://IP_ADDRESS:8081/bucket/create/{bucket_name}

$ curl -X PUT http://IP_ADDRESS:8081/bucket/delete/{bucket_name}

```

Note: Modify the IP ADDRESS with your IP in minio_api.go file