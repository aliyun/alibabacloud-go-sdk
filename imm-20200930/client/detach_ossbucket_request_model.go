// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachOSSBucketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOSSBucket(v string) *DetachOSSBucketRequest
	GetOSSBucket() *string
}

type DetachOSSBucketRequest struct {
	// The OSS bucket that you want to unbind.
	//
	// This parameter is required.
	//
	// example:
	//
	// examplebucket
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
}

func (s DetachOSSBucketRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachOSSBucketRequest) GoString() string {
	return s.String()
}

func (s *DetachOSSBucketRequest) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *DetachOSSBucketRequest) SetOSSBucket(v string) *DetachOSSBucketRequest {
	s.OSSBucket = &v
	return s
}

func (s *DetachOSSBucketRequest) Validate() error {
	return dara.Validate(s)
}
