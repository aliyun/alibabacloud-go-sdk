// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultHttpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDefaultHttpsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDefaultHttpsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefaultHttpsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeDefaultHttpsRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7pp26f1****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefaultHttpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultHttpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefaultHttpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefaultHttpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefaultHttpsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefaultHttpsRequest) SetInstanceId(v string) *DescribeDefaultHttpsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefaultHttpsRequest) SetRegionId(v string) *DescribeDefaultHttpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefaultHttpsRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefaultHttpsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefaultHttpsRequest) Validate() error {
	return dara.Validate(s)
}
