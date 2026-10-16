// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoGroupingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAutoGroupingRulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAutoGroupingRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAutoGroupingRulesResponseBody
	GetRequestId() *string
	SetRules(v []*ListAutoGroupingRulesResponseBodyRules) *ListAutoGroupingRulesResponseBody
	GetRules() []*ListAutoGroupingRulesResponseBodyRules
}

type ListAutoGroupingRulesResponseBody struct {
	// The maximum number of entries to return in a single request. Valid values: 1 to 50.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If the response is truncated, you can use NextToken to send another request to retrieve the content after the current truncation point.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of rules.
	Rules []*ListAutoGroupingRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s ListAutoGroupingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutoGroupingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoGroupingRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoGroupingRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoGroupingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutoGroupingRulesResponseBody) GetRules() []*ListAutoGroupingRulesResponseBodyRules {
	return s.Rules
}

func (s *ListAutoGroupingRulesResponseBody) SetMaxResults(v int32) *ListAutoGroupingRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBody) SetNextToken(v string) *ListAutoGroupingRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBody) SetRequestId(v string) *ListAutoGroupingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBody) SetRules(v []*ListAutoGroupingRulesResponseBodyRules) *ListAutoGroupingRulesResponseBody {
	s.Rules = v
	return s
}

func (s *ListAutoGroupingRulesResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAutoGroupingRulesResponseBodyRules struct {
	// The time when the rule was created, in RFC 3339 format with a UTC offset of +08:00.
	//
	// example:
	//
	// 2025-01-01T10:00:00+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The excluded region IDs. Multiple region IDs are separated by commas (,).
	//
	// example:
	//
	// cn-hangzhou,cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The excluded resource group IDs. Multiple resource group IDs are separated by commas (,).
	//
	// example:
	//
	// rg-aekz****ql4b5ea,rg-aek2****akfxykq
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The excluded resource IDs. Multiple resource IDs are separated by commas (,).
	//
	// example:
	//
	// i-2zee******ym49kfmwis,vpc-5ts6******fnw493g849a
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The excluded resource types. Multiple resource types are separated by commas (,).
	//
	// example:
	//
	// ecs.instance,vpc.vpc
	ExcludeResourceTypesScope *string `json:"ExcludeResourceTypesScope,omitempty" xml:"ExcludeResourceTypesScope,omitempty"`
	// The time when the rule was last updated, in RFC 3339 format with a UTC offset of +08:00.
	//
	// example:
	//
	// 2025-01-01T10:00:00+08:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The region IDs in scope. Multiple region IDs are separated by commas (,).
	//
	// example:
	//
	// cn-hangzhou,cn-shanghai
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The resource group IDs in scope. Multiple resource group IDs are separated by commas (,).
	//
	// example:
	//
	// rg-aekz****ql4b5ea,rg-aek2****akfxykq
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The resource IDs in scope. Multiple resource IDs are separated by commas (,).
	//
	// example:
	//
	// i-2zee******ym49kfmwis,vpc-5ts6******fnw493g849a
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The resource types in scope. Multiple resource types are separated by commas (,).
	//
	// example:
	//
	// ecs.instance,vpc.vpc
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// The list of rule contents.
	RuleContents []*ListAutoGroupingRulesResponseBodyRulesRuleContents `json:"RuleContents,omitempty" xml:"RuleContents,omitempty" type:"Repeated"`
	// The rule description.
	//
	// example:
	//
	// Transfer resources with tags {"env": "online"} and {"project": "A"} to resource group rg-aek2********qcy
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// gr-acfo******hy6a
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// Auto group transfer rule for Project A online resources
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule type. Valid values:
	//
	// - custom_condition: custom rule-based group transfer.
	//
	// - associated_transfer: attached-resource allocation group transfer.
	//
	// example:
	//
	// custom_condition
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s ListAutoGroupingRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s ListAutoGroupingRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetExcludeResourceTypesScope() *string {
	return s.ExcludeResourceTypesScope
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetRuleContents() []*ListAutoGroupingRulesResponseBodyRulesRuleContents {
	return s.RuleContents
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListAutoGroupingRulesResponseBodyRules) GetRuleType() *string {
	return s.RuleType
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetCreateTime(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetExcludeRegionIdsScope(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetExcludeResourceGroupIdsScope(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetExcludeResourceIdsScope(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetExcludeResourceTypesScope(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.ExcludeResourceTypesScope = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetModifyTime(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.ModifyTime = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetRegionIdsScope(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.RegionIdsScope = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetResourceGroupIdsScope(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetResourceIdsScope(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.ResourceIdsScope = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetResourceTypesScope(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.ResourceTypesScope = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetRuleContents(v []*ListAutoGroupingRulesResponseBodyRulesRuleContents) *ListAutoGroupingRulesResponseBodyRules {
	s.RuleContents = v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetRuleDesc(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.RuleDesc = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetRuleId(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetRuleName(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) SetRuleType(v string) *ListAutoGroupingRulesResponseBodyRules {
	s.RuleType = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRules) Validate() error {
	if s.RuleContents != nil {
		for _, item := range s.RuleContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAutoGroupingRulesResponseBodyRulesRuleContents struct {
	// The automatic group transfer scope condition, which specifies the scope of resources to be automatically transferred to a group.
	//
	// example:
	//
	// {"children":[{"desired":"{\\"env\\":\\"online\\", \\"project\\":\\"A\\"}","featurePath":"$.tags","featureSource":"RESOURCE","operator":"TagMatchAll"}],"operator":"and"}
	AutoGroupingScopeCondition *string `json:"AutoGroupingScopeCondition,omitempty" xml:"AutoGroupingScopeCondition,omitempty"`
	// The rule content ID.
	//
	// example:
	//
	// grc-acfo******fwybpq
	RuleContentId *string `json:"RuleContentId,omitempty" xml:"RuleContentId,omitempty"`
	// The target resource group condition, which specifies the target resource group for automatic grouping.
	//
	// example:
	//
	// {"children":[{"desired":"rg-aek2********qcy","featurePath":"$.resourceGroupId","featureSource":"RESOURCE","operator":"StringEquals"}],"operator":"and"}
	TargetResourceGroupCondition *string `json:"TargetResourceGroupCondition,omitempty" xml:"TargetResourceGroupCondition,omitempty"`
}

func (s ListAutoGroupingRulesResponseBodyRulesRuleContents) String() string {
	return dara.Prettify(s)
}

func (s ListAutoGroupingRulesResponseBodyRulesRuleContents) GoString() string {
	return s.String()
}

func (s *ListAutoGroupingRulesResponseBodyRulesRuleContents) GetAutoGroupingScopeCondition() *string {
	return s.AutoGroupingScopeCondition
}

func (s *ListAutoGroupingRulesResponseBodyRulesRuleContents) GetRuleContentId() *string {
	return s.RuleContentId
}

func (s *ListAutoGroupingRulesResponseBodyRulesRuleContents) GetTargetResourceGroupCondition() *string {
	return s.TargetResourceGroupCondition
}

func (s *ListAutoGroupingRulesResponseBodyRulesRuleContents) SetAutoGroupingScopeCondition(v string) *ListAutoGroupingRulesResponseBodyRulesRuleContents {
	s.AutoGroupingScopeCondition = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRulesRuleContents) SetRuleContentId(v string) *ListAutoGroupingRulesResponseBodyRulesRuleContents {
	s.RuleContentId = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRulesRuleContents) SetTargetResourceGroupCondition(v string) *ListAutoGroupingRulesResponseBodyRulesRuleContents {
	s.TargetResourceGroupCondition = &v
	return s
}

func (s *ListAutoGroupingRulesResponseBodyRulesRuleContents) Validate() error {
	return dara.Validate(s)
}
