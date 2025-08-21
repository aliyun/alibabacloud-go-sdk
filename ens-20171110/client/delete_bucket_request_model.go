// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *DeleteBucketRequest
	GetBucketName() *string
}

type DeleteBucketRequest struct {
	// The name of the bucket that you want to delete. You can delete only one bucket at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s DeleteBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBucketRequest) GoString() string {
	return s.String()
}

func (s *DeleteBucketRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *DeleteBucketRequest) SetBucketName(v string) *DeleteBucketRequest {
	s.BucketName = &v
	return s
}

func (s *DeleteBucketRequest) Validate() error {
	return dara.Validate(s)
}
