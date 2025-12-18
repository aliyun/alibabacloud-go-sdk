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
	// The compliance package ID.
	//
	// example:
	//
	// cp-fe416457e0d90022****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The compliance evaluation result. Valid values:
	//
	// 	- COMPLIANT: The resources are evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resources are evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to the resources.
	//
	// 	- INSUFFICIENT_DATA: No data is available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test-rule-name
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- ACTIVE: The rule is being used to monitor resource configurations.
	//
	// 	- DELETING: The rule is being deleted.
	//
	// 	- EVALUATING: The rule is triggered and is being used to monitor resource configurations.
	//
	// 	- INACTIVE: The rule is disabled.
	//
	// example:
	//
	// ACTIVE
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// The keyword that is used for queries.
	//
	// You can perform a fuzzy search by rule ID, rule name, rule description, or managed rule ID.
	//
	// example:
	//
	// ecs
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Resource type for the rule to evaluate.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// The risk level of the resources that do not comply with the rule. Valid values:
	//
	// 	- 1: high
	//
	// 	- 2: medium
	//
	// 	- 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	SortBy    *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
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
