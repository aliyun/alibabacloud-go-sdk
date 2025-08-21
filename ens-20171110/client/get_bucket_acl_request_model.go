// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *GetBucketAclRequest
	GetBucketName() *string
}

type GetBucketAclRequest struct {
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s GetBucketAclRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBucketAclRequest) GoString() string {
	return s.String()
}

func (s *GetBucketAclRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *GetBucketAclRequest) SetBucketName(v string) *GetBucketAclRequest {
	s.BucketName = &v
	return s
}

func (s *GetBucketAclRequest) Validate() error {
	return dara.Validate(s)
}
