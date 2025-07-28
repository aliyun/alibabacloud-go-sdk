// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsLogStoreStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSlsLogStoreStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeSlsLogStoreStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSlsLogStoreStatusRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeSlsLogStoreStatusRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-uqm35****02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfmzedqv****ma
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeSlsLogStoreStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsLogStoreStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsLogStoreStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSlsLogStoreStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSlsLogStoreStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSlsLogStoreStatusRequest) SetInstanceId(v string) *DescribeSlsLogStoreStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlsLogStoreStatusRequest) SetRegionId(v string) *DescribeSlsLogStoreStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlsLogStoreStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeSlsLogStoreStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSlsLogStoreStatusRequest) Validate() error {
	return dara.Validate(s)
}
