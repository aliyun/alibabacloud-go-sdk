// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIPGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionName(v string) *DescribeSecurityIPGroupRequest
	GetRegionName() *string
}

type DescribeSecurityIPGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeSecurityIPGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPGroupRequest) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeSecurityIPGroupRequest) SetRegionName(v string) *DescribeSecurityIPGroupRequest {
	s.RegionName = &v
	return s
}

func (s *DescribeSecurityIPGroupRequest) Validate() error {
	return dara.Validate(s)
}
