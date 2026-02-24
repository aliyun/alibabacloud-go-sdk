// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateConfigRulesRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *ListAggregateConfigRulesRequest
	GetCompliancePackId() *string
	SetComplianceType(v string) *ListAggregateConfigRulesRequest
	GetComplianceType() *string
	SetConfigRuleName(v string) *ListAggregateConfigRulesRequest
	GetConfigRuleName() *string
	SetConfigRuleState(v string) *ListAggregateConfigRulesRequest
	GetConfigRuleState() *string
	SetKeyword(v string) *ListAggregateConfigRulesRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListAggregateConfigRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAggregateConfigRulesRequest
	GetPageSize() *int32
	SetResourceTypes(v string) *ListAggregateConfigRulesRequest
	GetResourceTypes() *string
	SetRiskLevel(v int32) *ListAggregateConfigRulesRequest
	GetRiskLevel() *int32
	SetSortBy(v string) *ListAggregateConfigRulesRequest
	GetSortBy() *string
	SetTag(v []*ListAggregateConfigRulesRequestTag) *ListAggregateConfigRulesRequest
	GetTag() []*ListAggregateConfigRulesRequestTag
}

type ListAggregateConfigRulesRequest struct {
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
	Tag []*ListAggregateConfigRulesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAggregateConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateConfigRulesRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListAggregateConfigRulesRequest) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListAggregateConfigRulesRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListAggregateConfigRulesRequest) GetConfigRuleState() *string {
	return s.ConfigRuleState
}

func (s *ListAggregateConfigRulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAggregateConfigRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAggregateConfigRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAggregateConfigRulesRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListAggregateConfigRulesRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListAggregateConfigRulesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListAggregateConfigRulesRequest) GetTag() []*ListAggregateConfigRulesRequestTag {
	return s.Tag
}

func (s *ListAggregateConfigRulesRequest) SetAggregatorId(v string) *ListAggregateConfigRulesRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetCompliancePackId(v string) *ListAggregateConfigRulesRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetComplianceType(v string) *ListAggregateConfigRulesRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetConfigRuleName(v string) *ListAggregateConfigRulesRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetConfigRuleState(v string) *ListAggregateConfigRulesRequest {
	s.ConfigRuleState = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetKeyword(v string) *ListAggregateConfigRulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetPageNumber(v int32) *ListAggregateConfigRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetPageSize(v int32) *ListAggregateConfigRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetResourceTypes(v string) *ListAggregateConfigRulesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetRiskLevel(v int32) *ListAggregateConfigRulesRequest {
	s.RiskLevel = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetSortBy(v string) *ListAggregateConfigRulesRequest {
	s.SortBy = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetTag(v []*ListAggregateConfigRulesRequestTag) *ListAggregateConfigRulesRequest {
	s.Tag = v
	return s
}

func (s *ListAggregateConfigRulesRequest) Validate() error {
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

type ListAggregateConfigRulesRequestTag struct {
	// The key of a resource tag.
	//
	// You can add a maximum of 20 tag keys.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of a resource tag.
	//
	// You can add a maximum of 20 tag values.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAggregateConfigRulesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRulesRequestTag) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAggregateConfigRulesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAggregateConfigRulesRequestTag) SetKey(v string) *ListAggregateConfigRulesRequestTag {
	s.Key = &v
	return s
}

func (s *ListAggregateConfigRulesRequestTag) SetValue(v string) *ListAggregateConfigRulesRequestTag {
	s.Value = &v
	return s
}

func (s *ListAggregateConfigRulesRequestTag) Validate() error {
	return dara.Validate(s)
}
