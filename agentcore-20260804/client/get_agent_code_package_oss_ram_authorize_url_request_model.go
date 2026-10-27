// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentCodePackageOssRamAuthorizeUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOssBucketName(v string) *GetAgentCodePackageOssRamAuthorizeUrlRequest
	GetOssBucketName() *string
}

type GetAgentCodePackageOssRamAuthorizeUrlRequest struct {
	// The name of the source OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent-bucket
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
}

func (s GetAgentCodePackageOssRamAuthorizeUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageOssRamAuthorizeUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlRequest) SetOssBucketName(v string) *GetAgentCodePackageOssRamAuthorizeUrlRequest {
	s.OssBucketName = &v
	return s
}

func (s *GetAgentCodePackageOssRamAuthorizeUrlRequest) Validate() error {
	return dara.Validate(s)
}
