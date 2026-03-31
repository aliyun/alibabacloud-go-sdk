// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopTuleTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeRuleHitsTopTuleTypeRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeRuleHitsTopTuleTypeRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeRuleHitsTopTuleTypeRequest
	GetRegionId() *string
	SetResource(v string) *DescribeRuleHitsTopTuleTypeRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopTuleTypeRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v string) *DescribeRuleHitsTopTuleTypeRequest
	GetStartTimestamp() *string
}

type DescribeRuleHitsTopTuleTypeRequest struct {
	// The end point of the time period for which to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 1665386280
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aekzwwkpn****5i
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The start point of the time period for which to query. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1665331200
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopTuleTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeRuleHitsTopTuleTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRuleHitsTopTuleTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRuleHitsTopTuleTypeRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeRuleHitsTopTuleTypeRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeRuleHitsTopTuleTypeRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetInstanceId(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetRegionId(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetResource(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopTuleTypeRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeRequest) Validate() error {
	return dara.Validate(s)
}
