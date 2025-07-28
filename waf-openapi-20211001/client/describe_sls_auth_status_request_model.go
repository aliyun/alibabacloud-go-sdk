// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsAuthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSlsAuthStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeSlsAuthStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSlsAuthStatusRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeSlsAuthStatusRequest struct {
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-g4t3g****04
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
	// rg-aek2okfav****iq
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeSlsAuthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSlsAuthStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSlsAuthStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSlsAuthStatusRequest) SetInstanceId(v string) *DescribeSlsAuthStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetRegionId(v string) *DescribeSlsAuthStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) SetResourceManagerResourceGroupId(v string) *DescribeSlsAuthStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSlsAuthStatusRequest) Validate() error {
	return dara.Validate(s)
}
