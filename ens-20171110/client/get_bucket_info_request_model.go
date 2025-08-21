// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *GetBucketInfoRequest
	GetBucketName() *string
}

type GetBucketInfoRequest struct {
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s GetBucketInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketInfoRequest) GoString() string {
	return s.String()
}

func (s *GetBucketInfoRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *GetBucketInfoRequest) SetBucketName(v string) *GetBucketInfoRequest {
	s.BucketName = &v
	return s
}

func (s *GetBucketInfoRequest) Validate() error {
	return dara.Validate(s)
}
