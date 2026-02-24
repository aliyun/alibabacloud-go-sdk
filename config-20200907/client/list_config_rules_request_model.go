// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePackId(v string) *ListConfigRulesRequest
	GetCompliancePackId() *string
	SetComplianceType(v string) *ListConfigRulesRequest
	GetComplianceType() *string
	SetConfigRuleName(v string) *ListConfigRulesRequest
	GetConfigRuleName() *string
	SetConfigRuleState(v string) *ListConfigRulesRequest
	GetConfigRuleState() *string
	SetKeyword(v string) *ListConfigRulesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListConfigRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConfigRulesRequest
	GetPageSize() *int32
	SetResourceTypes(v string) *ListConfigRulesRequest
	GetResourceTypes() *string
	SetRiskLevel(v int32) *ListConfigRulesRequest
	GetRiskLevel() *int32
	SetSortBy(v string) *ListConfigRulesRequest
	GetSortBy() *string
	SetTag(v []*ListConfigRulesRequestTag) *ListConfigRulesRequest
	GetTag() []*ListConfigRulesRequestTag
}

type ListConfigRulesRequest struct {
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
	Tag []*ListConfigRulesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ListConfigRulesRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListConfigRulesRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListConfigRulesRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListConfigRulesRequest) GetConfigRuleState() *string {
	return s.ConfigRuleState
}

func (s *ListConfigRulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListConfigRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConfigRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConfigRulesRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListConfigRulesRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListConfigRulesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListConfigRulesRequest) GetTag() []*ListConfigRulesRequestTag {
	return s.Tag
}

func (s *ListConfigRulesRequest) SetCompliancePackId(v string) *ListConfigRulesRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ListConfigRulesRequest) SetComplianceType(v string) *ListConfigRulesRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRulesRequest) SetConfigRuleName(v string) *ListConfigRulesRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *ListConfigRulesRequest) SetConfigRuleState(v string) *ListConfigRulesRequest {
	s.ConfigRuleState = &v
	return s
}

func (s *ListConfigRulesRequest) SetKeyword(v string) *ListConfigRulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListConfigRulesRequest) SetPageNumber(v int32) *ListConfigRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConfigRulesRequest) SetPageSize(v int32) *ListConfigRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListConfigRulesRequest) SetResourceTypes(v string) *ListConfigRulesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListConfigRulesRequest) SetRiskLevel(v int32) *ListConfigRulesRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListConfigRulesRequest) SetSortBy(v string) *ListConfigRulesRequest {
	s.SortBy = &v
	return s
}

func (s *ListConfigRulesRequest) SetTag(v []*ListConfigRulesRequestTag) *ListConfigRulesRequest {
	s.Tag = v
	return s
}

func (s *ListConfigRulesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConfigRulesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConfigRulesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesRequestTag) GoString() string {
	return s.String()
}

func (s *ListConfigRulesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListConfigRulesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListConfigRulesRequestTag) SetKey(v string) *ListConfigRulesRequestTag {
	s.Key = &v
	return s
}

func (s *ListConfigRulesRequestTag) SetValue(v string) *ListConfigRulesRequestTag {
	s.Value = &v
	return s
}

func (s *ListConfigRulesRequestTag) Validate() error {
	return dara.Validate(s)
}
