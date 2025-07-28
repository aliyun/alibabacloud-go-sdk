// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserEventCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeFreeUserEventCountRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeFreeUserEventCountRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeFreeUserEventCountRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeFreeUserEventCountRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepay_public_intl-sg-vf***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeFreeUserEventCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventCountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFreeUserEventCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFreeUserEventCountRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeFreeUserEventCountRequest) SetInstanceId(v string) *DescribeFreeUserEventCountRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFreeUserEventCountRequest) SetRegionId(v string) *DescribeFreeUserEventCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFreeUserEventCountRequest) SetResourceManagerResourceGroupId(v string) *DescribeFreeUserEventCountRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeFreeUserEventCountRequest) Validate() error {
	return dara.Validate(s)
}
