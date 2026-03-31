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
	// This parameter is required.
	//
	// example:
	//
	// address_book
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// global
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11sr5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
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
