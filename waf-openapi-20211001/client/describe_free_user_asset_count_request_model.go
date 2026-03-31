// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserAssetCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeFreeUserAssetCountRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeFreeUserAssetCountRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeFreeUserAssetCountRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeFreeUserAssetCountRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-cs0*****
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

func (s DescribeFreeUserAssetCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserAssetCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserAssetCountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFreeUserAssetCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFreeUserAssetCountRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeFreeUserAssetCountRequest) SetInstanceId(v string) *DescribeFreeUserAssetCountRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFreeUserAssetCountRequest) SetRegionId(v string) *DescribeFreeUserAssetCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFreeUserAssetCountRequest) SetResourceManagerResourceGroupId(v string) *DescribeFreeUserAssetCountRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeFreeUserAssetCountRequest) Validate() error {
	return dara.Validate(s)
}
