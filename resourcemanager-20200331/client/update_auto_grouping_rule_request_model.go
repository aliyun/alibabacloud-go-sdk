// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoGroupingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeRegionIdsScope(v string) *UpdateAutoGroupingRuleRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *UpdateAutoGroupingRuleRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *UpdateAutoGroupingRuleRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeResourceTypesScope(v string) *UpdateAutoGroupingRuleRequest
	GetExcludeResourceTypesScope() *string
	SetRegionIdsScope(v string) *UpdateAutoGroupingRuleRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *UpdateAutoGroupingRuleRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *UpdateAutoGroupingRuleRequest
	GetResourceIdsScope() *string
	SetResourceTypesScope(v string) *UpdateAutoGroupingRuleRequest
	GetResourceTypesScope() *string
	SetRuleContents(v []*UpdateAutoGroupingRuleRequestRuleContents) *UpdateAutoGroupingRuleRequest
	GetRuleContents() []*UpdateAutoGroupingRuleRequestRuleContents
	SetRuleDesc(v string) *UpdateAutoGroupingRuleRequest
	GetRuleDesc() *string
	SetRuleId(v string) *UpdateAutoGroupingRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *UpdateAutoGroupingRuleRequest
	GetRuleName() *string
}

type UpdateAutoGroupingRuleRequest struct {
	// The IDs of regions to be excluded. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// cn-beijing,cn-guangzhou
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The IDs of resource groups to be excluded. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// rg-aekz******zj2oi,rg-aekz******r62ua
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The IDs of resources to be excluded. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// pc-uf6p******4h784y,rmq-cn-******ny0y
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The resource types to be excluded. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// mse.cluster,slb.loadbalancer
	ExcludeResourceTypesScope *string `json:"ExcludeResourceTypesScope,omitempty" xml:"ExcludeResourceTypesScope,omitempty"`
	// The IDs of regions. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou,cn-shanghai
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The IDs of resource groups. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// rg-aekz******4b5ea,rg-aek2******fxykq
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The IDs of resources. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// i-2zee******ym49kfmwis,vpc-5ts6******fnw493g849a
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The resource types. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// rds.dbinstance,oss.bucket
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// The content records of the rule.
	//
	// This parameter is required.
	RuleContents []*UpdateAutoGroupingRuleRequestRuleContents `json:"RuleContents,omitempty" xml:"RuleContents,omitempty" type:"Repeated"`
	// The description of the rule.
	//
	// example:
	//
	// Transfer resources to which the {"env": "online"} and {"project": "A"} tags are added to the resource group rg-aek2********qcy.
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The ID of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// gr-acfo******hy6a
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom Transfer Rule for Online Resources of Project A
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s UpdateAutoGroupingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoGroupingRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoGroupingRuleRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *UpdateAutoGroupingRuleRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *UpdateAutoGroupingRuleRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *UpdateAutoGroupingRuleRequest) GetExcludeResourceTypesScope() *string {
	return s.ExcludeResourceTypesScope
}

func (s *UpdateAutoGroupingRuleRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *UpdateAutoGroupingRuleRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *UpdateAutoGroupingRuleRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *UpdateAutoGroupingRuleRequest) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *UpdateAutoGroupingRuleRequest) GetRuleContents() []*UpdateAutoGroupingRuleRequestRuleContents {
	return s.RuleContents
}

func (s *UpdateAutoGroupingRuleRequest) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *UpdateAutoGroupingRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateAutoGroupingRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateAutoGroupingRuleRequest) SetExcludeRegionIdsScope(v string) *UpdateAutoGroupingRuleRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetExcludeResourceGroupIdsScope(v string) *UpdateAutoGroupingRuleRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetExcludeResourceIdsScope(v string) *UpdateAutoGroupingRuleRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetExcludeResourceTypesScope(v string) *UpdateAutoGroupingRuleRequest {
	s.ExcludeResourceTypesScope = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetRegionIdsScope(v string) *UpdateAutoGroupingRuleRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetResourceGroupIdsScope(v string) *UpdateAutoGroupingRuleRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetResourceIdsScope(v string) *UpdateAutoGroupingRuleRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetResourceTypesScope(v string) *UpdateAutoGroupingRuleRequest {
	s.ResourceTypesScope = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetRuleContents(v []*UpdateAutoGroupingRuleRequestRuleContents) *UpdateAutoGroupingRuleRequest {
	s.RuleContents = v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetRuleDesc(v string) *UpdateAutoGroupingRuleRequest {
	s.RuleDesc = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetRuleId(v string) *UpdateAutoGroupingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) SetRuleName(v string) *UpdateAutoGroupingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequest) Validate() error {
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

type UpdateAutoGroupingRuleRequestRuleContents struct {
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
	// This parameter is required.
	//
	// example:
	//
	// {"children":[{"desired":"rg-aek2********qcy","featurePath":"$.resourceGroupId","featureSource":"RESOURCE","operator":"StringEquals"}],"operator":"and"}
	TargetResourceGroupCondition *string `json:"TargetResourceGroupCondition,omitempty" xml:"TargetResourceGroupCondition,omitempty"`
}

func (s UpdateAutoGroupingRuleRequestRuleContents) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoGroupingRuleRequestRuleContents) GoString() string {
	return s.String()
}

func (s *UpdateAutoGroupingRuleRequestRuleContents) GetAutoGroupingScopeCondition() *string {
	return s.AutoGroupingScopeCondition
}

func (s *UpdateAutoGroupingRuleRequestRuleContents) GetRuleContentId() *string {
	return s.RuleContentId
}

func (s *UpdateAutoGroupingRuleRequestRuleContents) GetTargetResourceGroupCondition() *string {
	return s.TargetResourceGroupCondition
}

func (s *UpdateAutoGroupingRuleRequestRuleContents) SetAutoGroupingScopeCondition(v string) *UpdateAutoGroupingRuleRequestRuleContents {
	s.AutoGroupingScopeCondition = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequestRuleContents) SetRuleContentId(v string) *UpdateAutoGroupingRuleRequestRuleContents {
	s.RuleContentId = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequestRuleContents) SetTargetResourceGroupCondition(v string) *UpdateAutoGroupingRuleRequestRuleContents {
	s.TargetResourceGroupCondition = &v
	return s
}

func (s *UpdateAutoGroupingRuleRequestRuleContents) Validate() error {
	return dara.Validate(s)
}
