// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHybridCloudUserRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeHybridCloudUserRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudUserRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudUserRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
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

func (s DescribeHybridCloudUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudUserRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudUserRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudUserRequest) SetInstanceId(v string) *DescribeHybridCloudUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudUserRequest) SetRegionId(v string) *DescribeHybridCloudUserRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudUserRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudUserRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudUserRequest) Validate() error {
	return dara.Validate(s)
}
