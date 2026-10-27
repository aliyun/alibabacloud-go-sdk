// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAgentCodePackageOssRamAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOssBucketName(v string) *VerifyAgentCodePackageOssRamAuthorizationRequest
	GetOssBucketName() *string
}

type VerifyAgentCodePackageOssRamAuthorizationRequest struct {
	// The name of the source OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent-bucket
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
}

func (s VerifyAgentCodePackageOssRamAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyAgentCodePackageOssRamAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *VerifyAgentCodePackageOssRamAuthorizationRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *VerifyAgentCodePackageOssRamAuthorizationRequest) SetOssBucketName(v string) *VerifyAgentCodePackageOssRamAuthorizationRequest {
	s.OssBucketName = &v
	return s
}

func (s *VerifyAgentCodePackageOssRamAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
