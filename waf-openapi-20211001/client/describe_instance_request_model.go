// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeInstanceRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeInstanceRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeInstanceRequest struct {
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeInstanceRequest) SetRegionId(v string) *DescribeInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceRequest) SetResourceManagerResourceGroupId(v string) *DescribeInstanceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
