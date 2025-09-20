// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSBucketAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOSSBucket(v string) *GetOSSBucketAttachmentRequest
	GetOSSBucket() *string
}

type GetOSSBucketAttachmentRequest struct {
	// The name of the OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// examplebucket
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
}

func (s GetOSSBucketAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOSSBucketAttachmentRequest) GoString() string {
	return s.String()
}

func (s *GetOSSBucketAttachmentRequest) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *GetOSSBucketAttachmentRequest) SetOSSBucket(v string) *GetOSSBucketAttachmentRequest {
	s.OSSBucket = &v
	return s
}

func (s *GetOSSBucketAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
