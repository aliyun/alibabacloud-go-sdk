// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssStorageAndAccByBucketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketList(v []*GetOssStorageAndAccByBucketsResponseBodyBucketList) *GetOssStorageAndAccByBucketsResponseBody
	GetBucketList() []*GetOssStorageAndAccByBucketsResponseBodyBucketList
	SetRequestId(v string) *GetOssStorageAndAccByBucketsResponseBody
	GetRequestId() *string
}

type GetOssStorageAndAccByBucketsResponseBody struct {
	// The information about the bucket.
	BucketList []*GetOssStorageAndAccByBucketsResponseBodyBucketList `json:"BucketList,omitempty" xml:"BucketList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 112F4860-F1B2-58DD-8FC0-75F19DA1C4BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssStorageAndAccByBucketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssStorageAndAccByBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssStorageAndAccByBucketsResponseBody) GetBucketList() []*GetOssStorageAndAccByBucketsResponseBodyBucketList {
	return s.BucketList
}

func (s *GetOssStorageAndAccByBucketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssStorageAndAccByBucketsResponseBody) SetBucketList(v []*GetOssStorageAndAccByBucketsResponseBodyBucketList) *GetOssStorageAndAccByBucketsResponseBody {
	s.BucketList = v
	return s
}

func (s *GetOssStorageAndAccByBucketsResponseBody) SetRequestId(v string) *GetOssStorageAndAccByBucketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssStorageAndAccByBucketsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOssStorageAndAccByBucketsResponseBodyBucketList struct {
	// The number of times that the bucket is accessed.
	//
	// example:
	//
	// 1000
	Acc *int64 `json:"Acc,omitempty" xml:"Acc,omitempty"`
	// The name of the bucket.
	//
	// example:
	//
	// my-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The storage usage of the bucket. Unit: bytes.
	//
	// example:
	//
	// 1024
	StorageUsageByte *int64 `json:"StorageUsageByte,omitempty" xml:"StorageUsageByte,omitempty"`
}

func (s GetOssStorageAndAccByBucketsResponseBodyBucketList) String() string {
	return dara.Prettify(s)
}

func (s GetOssStorageAndAccByBucketsResponseBodyBucketList) GoString() string {
	return s.String()
}

func (s *GetOssStorageAndAccByBucketsResponseBodyBucketList) GetAcc() *int64 {
	return s.Acc
}

func (s *GetOssStorageAndAccByBucketsResponseBodyBucketList) GetBucket() *string {
	return s.Bucket
}

func (s *GetOssStorageAndAccByBucketsResponseBodyBucketList) GetStorageUsageByte() *int64 {
	return s.StorageUsageByte
}

func (s *GetOssStorageAndAccByBucketsResponseBodyBucketList) SetAcc(v int64) *GetOssStorageAndAccByBucketsResponseBodyBucketList {
	s.Acc = &v
	return s
}

func (s *GetOssStorageAndAccByBucketsResponseBodyBucketList) SetBucket(v string) *GetOssStorageAndAccByBucketsResponseBodyBucketList {
	s.Bucket = &v
	return s
}

func (s *GetOssStorageAndAccByBucketsResponseBodyBucketList) SetStorageUsageByte(v int64) *GetOssStorageAndAccByBucketsResponseBodyBucketList {
	s.StorageUsageByte = &v
	return s
}

func (s *GetOssStorageAndAccByBucketsResponseBodyBucketList) Validate() error {
	return dara.Validate(s)
}
