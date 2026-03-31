// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopUaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeRuleHitsTopUaRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeRuleHitsTopUaRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeRuleHitsTopUaRequest
	GetRegionId() *string
	SetResource(v string) *DescribeRuleHitsTopUaRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopUaRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v string) *DescribeRuleHitsTopUaRequest
	GetStartTimestamp() *string
}

type DescribeRuleHitsTopUaRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 1665386280
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
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
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1665331200
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopUaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopUaRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeRuleHitsTopUaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRuleHitsTopUaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRuleHitsTopUaRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeRuleHitsTopUaRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeRuleHitsTopUaRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeRuleHitsTopUaRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopUaRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetInstanceId(v string) *DescribeRuleHitsTopUaRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetRegionId(v string) *DescribeRuleHitsTopUaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetResource(v string) *DescribeRuleHitsTopUaRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopUaRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopUaRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopUaRequest) Validate() error {
	return dara.Validate(s)
}
