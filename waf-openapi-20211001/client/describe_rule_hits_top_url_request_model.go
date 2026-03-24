// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribeRuleHitsTopUrlRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribeRuleHitsTopUrlRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeRuleHitsTopUrlRequest
	GetRegionId() *string
	SetResource(v string) *DescribeRuleHitsTopUrlRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopUrlRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleType(v string) *DescribeRuleHitsTopUrlRequest
	GetRuleType() *string
	SetStartTimestamp(v string) *DescribeRuleHitsTopUrlRequest
	GetStartTimestamp() *string
}

type DescribeRuleHitsTopUrlRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds. If you do not specify this parameter, the current time is used as the end time.
	//
	// example:
	//
	// 1665386280
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region of the WAF instance. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of protection rule that is triggered. If you do not specify this parameter, data of all rule types is returned.
	//
	// - **blacklist**: The IP address blacklist.
	//
	// - **custom**: A custom rule.
	//
	// - **antiscan**: A scan protection rule.
	//
	// - **cc_system**: An HTTP flood protection rule.
	//
	// - **region_block**: A location blacklist.
	//
	// example:
	//
	// blacklist
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The start of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1665331200
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeRuleHitsTopUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeRuleHitsTopUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRuleHitsTopUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRuleHitsTopUrlRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeRuleHitsTopUrlRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeRuleHitsTopUrlRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeRuleHitsTopUrlRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribeRuleHitsTopUrlRequest) SetEndTimestamp(v string) *DescribeRuleHitsTopUrlRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetInstanceId(v string) *DescribeRuleHitsTopUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetRegionId(v string) *DescribeRuleHitsTopUrlRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetResource(v string) *DescribeRuleHitsTopUrlRequest {
	s.Resource = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetResourceManagerResourceGroupId(v string) *DescribeRuleHitsTopUrlRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetRuleType(v string) *DescribeRuleHitsTopUrlRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) SetStartTimestamp(v string) *DescribeRuleHitsTopUrlRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeRuleHitsTopUrlRequest) Validate() error {
	return dara.Validate(s)
}
