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
	SetId(v string) *GetOssScanConfigRequest
	GetId() *string
}

type GetOssScanConfigRequest struct {
	// The name of the bucket.
	//
	// example:
	//
	// iboxpublic****
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *GetOssScanConfigRequest) GetId() *string {
	return s.Id
}

func (s *GetOssScanConfigRequest) SetBucketName(v string) *GetOssScanConfigRequest {
	s.BucketName = &v
	return s
}

func (s *GetOssScanConfigRequest) SetId(v string) *GetOssScanConfigRequest {
	s.Id = &v
	return s
}

func (s *GetOssScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
