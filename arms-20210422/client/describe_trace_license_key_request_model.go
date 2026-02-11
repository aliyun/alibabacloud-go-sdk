// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceLicenseKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeTraceLicenseKeyRequest
	GetRegionId() *string
}

type DescribeTraceLicenseKeyRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTraceLicenseKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceLicenseKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeTraceLicenseKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTraceLicenseKeyRequest) SetRegionId(v string) *DescribeTraceLicenseKeyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTraceLicenseKeyRequest) Validate() error {
	return dara.Validate(s)
}
