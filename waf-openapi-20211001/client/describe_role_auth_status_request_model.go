// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleAuthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeRoleAuthStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeRoleAuthStatusRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeRoleAuthStatusRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzhalsanv***
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeRoleAuthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRoleAuthStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRoleAuthStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeRoleAuthStatusRequest) SetRegionId(v string) *DescribeRoleAuthStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRoleAuthStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeRoleAuthStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRoleAuthStatusRequest) Validate() error {
	return dara.Validate(s)
}
