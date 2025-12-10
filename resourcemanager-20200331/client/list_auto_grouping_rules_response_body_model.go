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
	// The maximum number of entries returned for a single request. Valid values: 1 to 50.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
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
	// The queried rules.
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
	// The time when the rule was created.
	//
	// example:
	//
	// 2025-01-01T10:00:00+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The IDs of excluded regions. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// cn-hangzhou,cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The IDs of excluded resource groups. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// rg-aekz****ql4b5ea,rg-aek2****akfxykq
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The IDs of excluded resources. Multiple IDs are separated by commas (,).
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
	// The time when the rule was updated.
	//
	// example:
	//
	// 2025-01-01T10:00:00+08:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The IDs of regions. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// cn-hangzhou,cn-shanghai
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The IDs of resource groups. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// rg-aekz****ql4b5ea,rg-aek2****akfxykq
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The IDs of resources. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// i-2zee******ym49kfmwis,vpc-5ts6******fnw493g849a
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The resource types. Multiple resource types are separated by commas (,).
	//
	// example:
	//
	// ecs.instance,vpc.vpc
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// The content records of the rule.
	RuleContents []*ListAutoGroupingRulesResponseBodyRulesRuleContents `json:"RuleContents,omitempty" xml:"RuleContents,omitempty" type:"Repeated"`
	// The description of the rule.
	//
	// example:
	//
	// Transfer resources to which the {"env": "online"} and {"project": "A"} tags are added to the resource group rg-aek2********qcy.
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// gr-acfo******hy6a
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// Custom Transfer Rule for Online Resources of Project A
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- custom_condition: custom transfer rule
	//
	// 	- associated_transfer: transfer rule for associated resources
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
	// The condition for the range of resources that are automatically transferred.
	//
	// example:
	//
	// {"children":[{"desired":"{\\"env\\":\\"online\\", \\"project\\":\\"A\\"}","featurePath":"$.tags","featureSource":"RESOURCE","operator":"TagMatchAll"}],"operator":"and"}
	AutoGroupingScopeCondition *string `json:"AutoGroupingScopeCondition,omitempty" xml:"AutoGroupingScopeCondition,omitempty"`
	// The ID of the content record.
	//
	// example:
	//
	// grc-acfo******fwybpq
	RuleContentId *string `json:"RuleContentId,omitempty" xml:"RuleContentId,omitempty"`
	// The condition for the destination resource group.
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
