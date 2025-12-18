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
	// The compliance package ID.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/606968.html).
	//
	// >  You must configure either the `CompliancePackId` or `ConfigRuleId` parameter.
	//
	// example:
	//
	// cp-fe416457e0d90022****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The compliance evaluation result of the rule. Valid values:
	//
	// 	- COMPLIANT: The resources are evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resources are evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to the resources.
	//
	// 	- INSUFFICIENT_DATA: No resource data is available.
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
	// 	- ACTIVE: The rule is enabled.
	//
	// 	- DELETING: The rule is being deleted.
	//
	// 	- EVALUATING: The rule is being used to evaluate resource configurations.
	//
	// 	- INACTIVE: The rule is disabled.
	//
	// example:
	//
	// ACTIVE
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// The query keyword.
	//
	// You can perform a fuzzy search by rule ID, rule name, rule description, or managed rule ID.
	//
	// example:
	//
	// ecs
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// Page numbers start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. A minimum of 1 entry can be returned per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resources to be evaluated based on the rule.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// The risk level of the resources that are not compliant with the rule. Valid values:
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
