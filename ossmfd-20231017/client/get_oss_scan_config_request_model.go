// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssScanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucketName(v string) *GetOssScanConfigRequest
	GetBucketName() *string
}

type GetOssScanConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testBucketName****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
}

func (s GetOssScanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssScanConfigRequest) GoString() string {
	return s.String()
}

func (s *GetOssScanConfigRequest) GetBucketName() *string {
	return s.BucketName
}

func (s *GetOssScanConfigRequest) SetBucketName(v string) *GetOssScanConfigRequest {
	s.BucketName = &v
	return s
}

func (s *GetOssScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
