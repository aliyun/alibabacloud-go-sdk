// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFreeUserEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeFreeUserEventsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeFreeUserEventsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeFreeUserEventsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeFreeUserEventsRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-27a3****
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

func (s DescribeFreeUserEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFreeUserEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFreeUserEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFreeUserEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFreeUserEventsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeFreeUserEventsRequest) SetInstanceId(v string) *DescribeFreeUserEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFreeUserEventsRequest) SetRegionId(v string) *DescribeFreeUserEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFreeUserEventsRequest) SetResourceManagerResourceGroupId(v string) *DescribeFreeUserEventsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeFreeUserEventsRequest) Validate() error {
	return dara.Validate(s)
}
