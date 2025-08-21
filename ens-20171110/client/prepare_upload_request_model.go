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
	// The name of the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The specified IP address. This parameter is applicable to scenarios where the user IP address is inconsistent with the operation calling IP address, such as the scenario where the server obtains authorization and sends the authorization to the client.
	//
	// example:
	//
	// 180.166.XX.XXX
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
