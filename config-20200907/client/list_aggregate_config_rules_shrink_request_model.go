// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateConfigRulesShrinkRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *ListAggregateConfigRulesShrinkRequest
	GetCompliancePackId() *string
	SetComplianceType(v string) *ListAggregateConfigRulesShrinkRequest
	GetComplianceType() *string
	SetConfigRuleName(v string) *ListAggregateConfigRulesShrinkRequest
	GetConfigRuleName() *string
	SetConfigRuleState(v string) *ListAggregateConfigRulesShrinkRequest
	GetConfigRuleState() *string
	SetKeyword(v string) *ListAggregateConfigRulesShrinkRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListAggregateConfigRulesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAggregateConfigRulesShrinkRequest
	GetPageSize() *int32
	SetResourceTypes(v string) *ListAggregateConfigRulesShrinkRequest
	GetResourceTypes() *string
	SetRiskLevel(v int32) *ListAggregateConfigRulesShrinkRequest
	GetRiskLevel() *int32
	SetSortBy(v string) *ListAggregateConfigRulesShrinkRequest
	GetSortBy() *string
	SetTagShrink(v string) *ListAggregateConfigRulesShrinkRequest
	GetTagShrink() *string
}

type ListAggregateConfigRulesShrinkRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-fe416457e0d90022****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The compliance evaluation result. Valid values:
	//
	// - COMPLIANT: The resource is compliant.
	//
	// - NON_COMPLIANT: The resource is non-compliant.
	//
	// - NOT_APPLICABLE: The rule does not apply to the resource.
	//
	// - INSUFFICIENT_DATA: No data is available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// The name of the rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The state of the rule. Valid values:
	//
	// - ACTIVE: The rule is enabled.
	//
	// - DELETING: The rule is being deleted.
	//
	// - EVALUATING: The rule is being evaluated.
	//
	// - INACTIVE: The rule is disabled.
	//
	// example:
	//
	// ACTIVE
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// The keyword for a fuzzy query.
	//
	// The keyword can be a rule ID, rule name, rule description, or rule template identifier.
	//
	// example:
	//
	// ecs
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource type to be evaluated by the rule.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// The risk level of the rule. Valid values:
	//
	// - 1: high
	//
	// - 2: medium
	//
	// - 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The method that is used to sort the rules. By default, this parameter is not specified. Set the value to `CreateDate-Desc` to sort the rules in descending order of their creation time.
	//
	// example:
	//
	// CreateDate-Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The tags of the resource.
	//
	// You can add a maximum of 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListAggregateConfigRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateConfigRulesShrinkRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListAggregateConfigRulesShrinkRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListAggregateConfigRulesShrinkRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListAggregateConfigRulesShrinkRequest) GetConfigRuleState() *string {
	return s.ConfigRuleState
}

func (s *ListAggregateConfigRulesShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAggregateConfigRulesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAggregateConfigRulesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAggregateConfigRulesShrinkRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListAggregateConfigRulesShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListAggregateConfigRulesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListAggregateConfigRulesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListAggregateConfigRulesShrinkRequest) SetAggregatorId(v string) *ListAggregateConfigRulesShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetCompliancePackId(v string) *ListAggregateConfigRulesShrinkRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetComplianceType(v string) *ListAggregateConfigRulesShrinkRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetConfigRuleName(v string) *ListAggregateConfigRulesShrinkRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetConfigRuleState(v string) *ListAggregateConfigRulesShrinkRequest {
	s.ConfigRuleState = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetKeyword(v string) *ListAggregateConfigRulesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetPageNumber(v int32) *ListAggregateConfigRulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetPageSize(v int32) *ListAggregateConfigRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetResourceTypes(v string) *ListAggregateConfigRulesShrinkRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetRiskLevel(v int32) *ListAggregateConfigRulesShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetSortBy(v string) *ListAggregateConfigRulesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) SetTagShrink(v string) *ListAggregateConfigRulesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListAggregateConfigRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
