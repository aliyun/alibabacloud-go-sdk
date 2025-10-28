// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfficeSiteId(v string) *DescribeSecurityGroupAttributeRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DescribeSecurityGroupAttributeRequest
	GetRegionId() *string
}

type DescribeSecurityGroupAttributeRequest struct {
	// This parameter is required.
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSecurityGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeSecurityGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityGroupAttributeRequest) SetOfficeSiteId(v string) *DescribeSecurityGroupAttributeRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) SetRegionId(v string) *DescribeSecurityGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
