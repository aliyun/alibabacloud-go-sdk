// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssStorageAndAccByBucketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketList(v string) *GetOssStorageAndAccByBucketsRequest
	GetBucketList() *string
}

type GetOssStorageAndAccByBucketsRequest struct {
	// The information about the bucket.
	//
	// example:
	//
	// my-bucket
	BucketList *string `json:"BucketList,omitempty" xml:"BucketList,omitempty"`
}

func (s GetOssStorageAndAccByBucketsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssStorageAndAccByBucketsRequest) GoString() string {
	return s.String()
}

func (s *GetOssStorageAndAccByBucketsRequest) GetBucketList() *string {
	return s.BucketList
}

func (s *GetOssStorageAndAccByBucketsRequest) SetBucketList(v string) *GetOssStorageAndAccByBucketsRequest {
	s.BucketList = &v
	return s
}

func (s *GetOssStorageAndAccByBucketsRequest) Validate() error {
	return dara.Validate(s)
}
