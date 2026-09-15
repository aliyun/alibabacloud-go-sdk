// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyWorkspaceOssMountRamAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *VerifyWorkspaceOssMountRamAuthorizationRequest
	GetBucketName() *string
}

type VerifyWorkspaceOssMountRamAuthorizationRequest struct {
	// The name of the OSS bucket.
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
}

func (s VerifyWorkspaceOssMountRamAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyWorkspaceOssMountRamAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *VerifyWorkspaceOssMountRamAuthorizationRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *VerifyWorkspaceOssMountRamAuthorizationRequest) SetBucketName(v string) *VerifyWorkspaceOssMountRamAuthorizationRequest {
	s.BucketName = &v
	return s
}

func (s *VerifyWorkspaceOssMountRamAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
