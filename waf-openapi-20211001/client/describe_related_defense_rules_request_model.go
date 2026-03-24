// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRelatedDefenseRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScene(v string) *DescribeRelatedDefenseRulesRequest
	GetDefenseScene() *string
	SetDefenseType(v string) *DescribeRelatedDefenseRulesRequest
	GetDefenseType() *string
	SetInstanceId(v string) *DescribeRelatedDefenseRulesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeRelatedDefenseRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRelatedDefenseRulesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeRelatedDefenseRulesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeRelatedDefenseRulesRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *DescribeRelatedDefenseRulesRequest
	GetRuleId() *int64
}

type DescribeRelatedDefenseRulesRequest struct {
	// The WAF protection scenario. Only the following value is supported:
	//
	// - **address_book**: address book.
	//
	// This parameter is required.
	//
	// example:
	//
	// address_book
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The type of the protection rule. Only the following value is supported:
	//
	// - **global**: a global rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// global
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of your WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11sr5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to query the next page of results. If more results are available, this parameter is returned.
	//
	// > If this parameter is returned, more results are available. Use the returned NextToken value as a request parameter to retrieve the next page of data. Repeat this process until the **NextToken*	- parameter is not returned. This indicates that all data has been retrieved.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
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
	// The ID of the protection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRelatedDefenseRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRelatedDefenseRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRelatedDefenseRulesRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeRelatedDefenseRulesRequest) GetDefenseType() *string {
	return s.DefenseType
}

func (s *DescribeRelatedDefenseRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRelatedDefenseRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRelatedDefenseRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRelatedDefenseRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRelatedDefenseRulesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeRelatedDefenseRulesRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeRelatedDefenseRulesRequest) SetDefenseScene(v string) *DescribeRelatedDefenseRulesRequest {
	s.DefenseScene = &v
	return s
}

func (s *DescribeRelatedDefenseRulesRequest) SetDefenseType(v string) *DescribeRelatedDefenseRulesRequest {
	s.DefenseType = &v
	return s
}

func (s *DescribeRelatedDefenseRulesRequest) SetInstanceId(v string) *DescribeRelatedDefenseRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRelatedDefenseRulesRequest) SetMaxResults(v int32) *DescribeRelatedDefenseRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRelatedDefenseRulesRequest) SetNextToken(v string) *DescribeRelatedDefenseRulesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRelatedDefenseRulesRequest) SetRegionId(v string) *DescribeRelatedDefenseRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRelatedDefenseRulesRequest) SetResourceManagerResourceGroupId(v string) *DescribeRelatedDefenseRulesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeRelatedDefenseRulesRequest) SetRuleId(v int64) *DescribeRelatedDefenseRulesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeRelatedDefenseRulesRequest) Validate() error {
	return dara.Validate(s)
}
