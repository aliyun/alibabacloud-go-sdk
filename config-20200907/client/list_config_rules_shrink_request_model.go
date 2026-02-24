// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *ListConfigRulesShrinkRequest
	GetCompliancePackId() *string
	SetComplianceType(v string) *ListConfigRulesShrinkRequest
	GetComplianceType() *string
	SetConfigRuleName(v string) *ListConfigRulesShrinkRequest
	GetConfigRuleName() *string
	SetConfigRuleState(v string) *ListConfigRulesShrinkRequest
	GetConfigRuleState() *string
	SetKeyword(v string) *ListConfigRulesShrinkRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListConfigRulesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConfigRulesShrinkRequest
	GetPageSize() *int32
	SetResourceTypes(v string) *ListConfigRulesShrinkRequest
	GetResourceTypes() *string
	SetRiskLevel(v int32) *ListConfigRulesShrinkRequest
	GetRiskLevel() *int32
	SetSortBy(v string) *ListConfigRulesShrinkRequest
	GetSortBy() *string
	SetTagShrink(v string) *ListConfigRulesShrinkRequest
	GetTagShrink() *string
}

type ListConfigRulesShrinkRequest struct {
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-fe416457e0d90022****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The compliance evaluation result of the rule. Valid values:
	//
	// - COMPLIANT: Compliant.
	//
	// - NON_COMPLIANT: Non-compliant.
	//
	// - NOT_APPLICABLE: Not applicable.
	//
	// - INSUFFICIENT_DATA: Insufficient data.
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
	// The keyword for the fuzzy query.
	//
	// Supports fuzzy queries by rule ID, rule name, rule description, or rule template identifier.
	//
	// example:
	//
	// ecs
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// The default value is 1.
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
	// The type of resource evaluated by the rule.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// The risk level of the rule. Valid values:
	//
	// - 1: High risk.
	//
	// - 2: Medium risk.
	//
	// - 3: Low risk.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The sorting method. This parameter is not required. Set the value to `CreateDate-Desc` to sort the rules by creation time in descending order.
	//
	// example:
	//
	// CreateDate-Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The tags of the resource.
	//
	// A maximum of 20 tags can be attached.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListConfigRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListConfigRulesShrinkRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListConfigRulesShrinkRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListConfigRulesShrinkRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListConfigRulesShrinkRequest) GetConfigRuleState() *string {
	return s.ConfigRuleState
}

func (s *ListConfigRulesShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListConfigRulesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConfigRulesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConfigRulesShrinkRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListConfigRulesShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListConfigRulesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListConfigRulesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListConfigRulesShrinkRequest) SetCompliancePackId(v string) *ListConfigRulesShrinkRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) SetComplianceType(v string) *ListConfigRulesShrinkRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) SetConfigRuleName(v string) *ListConfigRulesShrinkRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) SetConfigRuleState(v string) *ListConfigRulesShrinkRequest {
	s.ConfigRuleState = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) SetKeyword(v string) *ListConfigRulesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) SetPageNumber(v int32) *ListConfigRulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) SetPageSize(v int32) *ListConfigRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) SetResourceTypes(v string) *ListConfigRulesShrinkRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) SetRiskLevel(v int32) *ListConfigRulesShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) SetSortBy(v string) *ListConfigRulesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) SetTagShrink(v string) *ListConfigRulesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListConfigRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
