package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO: Read config from yaml file
	Config = config{
		AliyunCfg: AliyunCfg{
			Endpoint:        "oss-cn-hangzhou.aliyuncs.com",
			AccessKeyID:     "",
			AccessKeySecret: "",
			BucketName:      "",
		},
	}

	client := AliyunInit(Config.AliyunCfg)
	bucket, err := GetBucket(client, Config.AliyunCfg.BucketName)
	if err != nil {
		panic(err)
	}

	objects, err := ListObjects(bucket)
	if err != nil {
		panic(err)
	}

	//fmt.Println(len(objects))

	err = os.Mkdir("output", 0777)
	if err != nil {
		panic(err)
	}

	for idx, object := range objects {
		fmt.Printf("Downloading %d/%d: %s\n", idx+1, len(objects), object.Key)
		bytes, err := DownloadObject(object, bucket)
		if err != nil {
			panic(err)
		}
		err = Write(bytes, "output/"+object.Key)
		if err != nil {
			panic(err)
		}
	}

}
