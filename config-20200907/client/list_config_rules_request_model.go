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
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag value must start with a letter but cannot start with `aliyun` or `acs:`. The tag value cannot contain `http://` or `https://`.
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
