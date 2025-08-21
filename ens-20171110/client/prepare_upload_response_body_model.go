// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepareUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *PrepareUploadResponseBody
	GetBucketName() *string
	SetEndpoint(v string) *PrepareUploadResponseBody
	GetEndpoint() *string
	SetRequestId(v string) *PrepareUploadResponseBody
	GetRequestId() *string
}

type PrepareUploadResponseBody struct {
	// The name of the bucket. This parameter is available only when the OSS SDK is used.
	//
	// example:
	//
	// test-xxxxxx
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The endpoint. This parameter is available only when the OSS SDK is used.
	//
	// example:
	//
	// eos.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6666C5A5-75ED-422E-A022-7121FA18C968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PrepareUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrepareUploadResponseBody) GoString() string {
	return s.String()
}

func (s *PrepareUploadResponseBody) GetBucketName() *string {
	return s.BucketName
}

func (s *PrepareUploadResponseBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *PrepareUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PrepareUploadResponseBody) SetBucketName(v string) *PrepareUploadResponseBody {
	s.BucketName = &v
	return s
}

func (s *PrepareUploadResponseBody) SetEndpoint(v string) *PrepareUploadResponseBody {
	s.Endpoint = &v
	return s
}

func (s *PrepareUploadResponseBody) SetRequestId(v string) *PrepareUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrepareUploadResponseBody) Validate() error {
	return dara.Validate(s)
}
