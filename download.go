package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

func GetBucket(AliyunClient *oss.Client, bucketName string) (*oss.Bucket, error) {
	bucket, err := AliyunClient.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func ListObjects(bucket *oss.Bucket) ([]oss.ObjectProperties, error) {
	var objects []oss.ObjectProperties
	marker := ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker), oss.MaxKeys(100)) // 适当调整每页的对象数目
		if err != nil {
			return nil, err
		}
		objects = append(objects, lsRes.Objects...)
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	return objects, nil
}

func DownloadObject(object oss.ObjectProperties, bucket *oss.Bucket) ([]byte, error) {
	body, err := bucket.GetObject(object.Key)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	var bytes []byte
	buf := make([]byte, 1024) // 适当调整缓冲区大小
	for {
		n, err := body.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		bytes = append(bytes, buf[:n]...)
	}
	return bytes, nil
}
