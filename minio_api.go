package main

import (
    "log"
    "fmt"
//    "io/ioutil"
    "context"
    "encoding/json"
    "net/http"
    //"strconv"
    //"time"

    "github.com/gorilla/mux"
    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)
type Name struct {
	Name string `json:"name"`
}
func main() {

    router := mux.NewRouter()
    // Create
    router.HandleFunc("/bucket/create/{Name}", createBucket).Methods("POST")
    // List
    router.HandleFunc("/bucket/list", listBuckets).Methods("GET")
	// Delete
    router.HandleFunc("/bucket/delete/{Name}", deleteBucket).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8081", router))
}

func listBuckets(w http.ResponseWriter, r *http.Request) {
    endpoint := "10.XXX.XX.XX:9000"
    accessKeyID := "minioadmin"
    secretAccessKey := "minioadmin"
    // Initialize minio client object.
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
    })
    if err != nil {
        log.Fatalln(err)
    }
    buckets, err := minioClient.ListBuckets(context.Background())
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, bucket := range buckets {
        fmt.Println(bucket)
        json.NewEncoder(w).Encode(bucket)
        //return
    }
}
func createBucket(w http.ResponseWriter, r *http.Request) {
    endpoint := "10.XXX.XX.XX:9000"
    accessKeyID := "minioadmin"
    secretAccessKey := "minioadmin"
    // Initialize minio client object.
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
    })
    if err != nil {
        log.Fatalln(err)
    }

    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    InputName := params["Name"]
    err = minioClient.MakeBucket(context.Background(), InputName, minio.MakeBucketOptions{Region: ""})
    fmt.Println(params)

    if err != nil {
    fmt.Println(err)
    return
    }
    fmt.Println("Successfully created : ", InputName)
}

func deleteBucket(w http.ResponseWriter, r *http.Request) {
    endpoint := "10.XXX.XX.XX:9000"
    accessKeyID := "minioadmin"
    secretAccessKey := "minioadmin"
    // Initialize minio client object.
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
    })
    if err != nil {
        log.Fatalln(err)
    }

    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    InputName := params["Name"]
    fmt.Println(params)
	err = minioClient.RemoveBucket(context.Background(), InputName)
	if err != nil {
    fmt.Println(err)
    return
	}
    fmt.Println("Successfully Removed : ", InputName)
}