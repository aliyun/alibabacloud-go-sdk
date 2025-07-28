// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudServerRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeHybridCloudServerRegionsRequest
	GetInstanceId() *string
	SetRegionCode(v string) *DescribeHybridCloudServerRegionsRequest
	GetRegionCode() *string
	SetRegionId(v string) *DescribeHybridCloudServerRegionsRequest
	GetRegionId() *string
	SetRegionType(v string) *DescribeHybridCloudServerRegionsRequest
	GetRegionType() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudServerRegionsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudServerRegionsRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-5yd****7009
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The code of the region.
	//
	// >  This parameter is required if you set RegionType to region. The value is the code of the city.
	//
	// example:
	//
	// 410
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the region. Valid values:
	//
	// 	- **operator**: the ISP.
	//
	// 	- **continents**: the continent.
	//
	// 	- **region**: the city.
	//
	// This parameter is required.
	//
	// example:
	//
	// region
	RegionType *string `json:"RegionType,omitempty" xml:"RegionType,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudServerRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudServerRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudServerRegionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudServerRegionsRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *DescribeHybridCloudServerRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudServerRegionsRequest) GetRegionType() *string {
	return s.RegionType
}

func (s *DescribeHybridCloudServerRegionsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudServerRegionsRequest) SetInstanceId(v string) *DescribeHybridCloudServerRegionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudServerRegionsRequest) SetRegionCode(v string) *DescribeHybridCloudServerRegionsRequest {
	s.RegionCode = &v
	return s
}

func (s *DescribeHybridCloudServerRegionsRequest) SetRegionId(v string) *DescribeHybridCloudServerRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudServerRegionsRequest) SetRegionType(v string) *DescribeHybridCloudServerRegionsRequest {
	s.RegionType = &v
	return s
}

func (s *DescribeHybridCloudServerRegionsRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudServerRegionsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudServerRegionsRequest) Validate() error {
	return dara.Validate(s)
}
