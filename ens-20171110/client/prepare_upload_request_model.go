// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepareUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *PrepareUploadRequest
	GetBucketName() *string
	SetClientIp(v string) *PrepareUploadRequest
	GetClientIp() *string
}

type PrepareUploadRequest struct {
	// The bucket name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The client\\"s IP address. Use this parameter when the client\\"s IP address differs from the IP address used to call the API. For example, a server obtains authorization and then sends it to a client.
	//
	// example:
	//
	// 180.166.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
}

func (s PrepareUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s PrepareUploadRequest) GoString() string {
	return s.String()
}

func (s *PrepareUploadRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *PrepareUploadRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *PrepareUploadRequest) SetBucketName(v string) *PrepareUploadRequest {
	s.BucketName = &v
	return s
}

func (s *PrepareUploadRequest) SetClientIp(v string) *PrepareUploadRequest {
	s.ClientIp = &v
	return s
}

func (s *PrepareUploadRequest) Validate() error {
	return dara.Validate(s)
}
