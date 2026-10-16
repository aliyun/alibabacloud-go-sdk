// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoGroupingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAutoGroupingRuleResponseBody
	GetRequestId() *string
	SetRule(v *GetAutoGroupingRuleResponseBodyRule) *GetAutoGroupingRuleResponseBody
	GetRule() *GetAutoGroupingRuleResponseBodyRule
}

type GetAutoGroupingRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9EA4F962-1A2E-4AFE-BE0C-B14736FC46CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule information.
	Rule *GetAutoGroupingRuleResponseBodyRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
}

func (s GetAutoGroupingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoGroupingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoGroupingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoGroupingRuleResponseBody) GetRule() *GetAutoGroupingRuleResponseBodyRule {
	return s.Rule
}

func (s *GetAutoGroupingRuleResponseBody) SetRequestId(v string) *GetAutoGroupingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBody) SetRule(v *GetAutoGroupingRuleResponseBodyRule) *GetAutoGroupingRuleResponseBody {
	s.Rule = v
	return s
}

func (s *GetAutoGroupingRuleResponseBody) Validate() error {
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoGroupingRuleResponseBodyRule struct {
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
	// rg-aekz******4b5ea,rg-aek2******fxykq
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
	// The time when the rule was last modified, in RFC 3339 format with a UTC offset of +08:00.
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
	// rg-aekz******4b5ea,rg-aek2******fxykq
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
	RuleContents []*GetAutoGroupingRuleResponseBodyRuleRuleContents `json:"RuleContents,omitempty" xml:"RuleContents,omitempty" type:"Repeated"`
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
	// Automatic group transfer rule for Project A online resources
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule type. Valid values:
	//
	// - custom_condition: Transfer based on custom rules.
	//
	// - associated_transfer: Transfer associated resources together.
	//
	// example:
	//
	// associated_transfer
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s GetAutoGroupingRuleResponseBodyRule) String() string {
	return dara.Prettify(s)
}

func (s GetAutoGroupingRuleResponseBodyRule) GoString() string {
	return s.String()
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetExcludeResourceTypesScope() *string {
	return s.ExcludeResourceTypesScope
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetRuleContents() []*GetAutoGroupingRuleResponseBodyRuleRuleContents {
	return s.RuleContents
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetRuleId() *string {
	return s.RuleId
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetRuleName() *string {
	return s.RuleName
}

func (s *GetAutoGroupingRuleResponseBodyRule) GetRuleType() *string {
	return s.RuleType
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetCreateTime(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.CreateTime = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetExcludeRegionIdsScope(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetExcludeResourceGroupIdsScope(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetExcludeResourceIdsScope(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetExcludeResourceTypesScope(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.ExcludeResourceTypesScope = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetModifyTime(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.ModifyTime = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetRegionIdsScope(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.RegionIdsScope = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetResourceGroupIdsScope(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetResourceIdsScope(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.ResourceIdsScope = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetResourceTypesScope(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.ResourceTypesScope = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetRuleContents(v []*GetAutoGroupingRuleResponseBodyRuleRuleContents) *GetAutoGroupingRuleResponseBodyRule {
	s.RuleContents = v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetRuleDesc(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.RuleDesc = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetRuleId(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.RuleId = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetRuleName(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.RuleName = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) SetRuleType(v string) *GetAutoGroupingRuleResponseBodyRule {
	s.RuleType = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRule) Validate() error {
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

type GetAutoGroupingRuleResponseBodyRuleRuleContents struct {
	// The scope condition for automatic group transfer, which specifies the range of resources to be automatically transferred to a resource group.
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
	// The target resource group condition, which specifies the destination resource group for automatic grouping.
	//
	// example:
	//
	// {"children":[{"desired":"rg-aek2********qcy","featurePath":"$.resourceGroupId","featureSource":"RESOURCE","operator":"StringEquals"}],"operator":"and"}
	TargetResourceGroupCondition *string `json:"TargetResourceGroupCondition,omitempty" xml:"TargetResourceGroupCondition,omitempty"`
}

func (s GetAutoGroupingRuleResponseBodyRuleRuleContents) String() string {
	return dara.Prettify(s)
}

func (s GetAutoGroupingRuleResponseBodyRuleRuleContents) GoString() string {
	return s.String()
}

func (s *GetAutoGroupingRuleResponseBodyRuleRuleContents) GetAutoGroupingScopeCondition() *string {
	return s.AutoGroupingScopeCondition
}

func (s *GetAutoGroupingRuleResponseBodyRuleRuleContents) GetRuleContentId() *string {
	return s.RuleContentId
}

func (s *GetAutoGroupingRuleResponseBodyRuleRuleContents) GetTargetResourceGroupCondition() *string {
	return s.TargetResourceGroupCondition
}

func (s *GetAutoGroupingRuleResponseBodyRuleRuleContents) SetAutoGroupingScopeCondition(v string) *GetAutoGroupingRuleResponseBodyRuleRuleContents {
	s.AutoGroupingScopeCondition = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRuleRuleContents) SetRuleContentId(v string) *GetAutoGroupingRuleResponseBodyRuleRuleContents {
	s.RuleContentId = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRuleRuleContents) SetTargetResourceGroupCondition(v string) *GetAutoGroupingRuleResponseBodyRuleRuleContents {
	s.TargetResourceGroupCondition = &v
	return s
}

func (s *GetAutoGroupingRuleResponseBodyRuleRuleContents) Validate() error {
	return dara.Validate(s)
}
