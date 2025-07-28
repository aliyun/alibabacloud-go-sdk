// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserEventTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeFreeUserEventTypesRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeFreeUserEventTypesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeFreeUserEventTypesRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeFreeUserEventTypesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-bl0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
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

func (s DescribeFreeUserEventTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventTypesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFreeUserEventTypesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFreeUserEventTypesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeFreeUserEventTypesRequest) SetInstanceId(v string) *DescribeFreeUserEventTypesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFreeUserEventTypesRequest) SetRegionId(v string) *DescribeFreeUserEventTypesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFreeUserEventTypesRequest) SetResourceManagerResourceGroupId(v string) *DescribeFreeUserEventTypesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeFreeUserEventTypesRequest) Validate() error {
	return dara.Validate(s)
}
