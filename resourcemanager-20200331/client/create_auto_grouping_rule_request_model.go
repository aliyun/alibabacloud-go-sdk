// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoGroupingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeRegionIdsScope(v string) *CreateAutoGroupingRuleRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *CreateAutoGroupingRuleRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *CreateAutoGroupingRuleRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeResourceTypesScope(v string) *CreateAutoGroupingRuleRequest
	GetExcludeResourceTypesScope() *string
	SetRegionIdsScope(v string) *CreateAutoGroupingRuleRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *CreateAutoGroupingRuleRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *CreateAutoGroupingRuleRequest
	GetResourceIdsScope() *string
	SetResourceTypesScope(v string) *CreateAutoGroupingRuleRequest
	GetResourceTypesScope() *string
	SetRuleContents(v []*CreateAutoGroupingRuleRequestRuleContents) *CreateAutoGroupingRuleRequest
	GetRuleContents() []*CreateAutoGroupingRuleRequestRuleContents
	SetRuleDesc(v string) *CreateAutoGroupingRuleRequest
	GetRuleDesc() *string
	SetRuleName(v string) *CreateAutoGroupingRuleRequest
	GetRuleName() *string
	SetRuleType(v string) *CreateAutoGroupingRuleRequest
	GetRuleType() *string
}

type CreateAutoGroupingRuleRequest struct {
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
	RuleContents []*CreateAutoGroupingRuleRequestRuleContents `json:"RuleContents,omitempty" xml:"RuleContents,omitempty" type:"Repeated"`
	// The description of the rule.
	//
	// example:
	//
	// Transfer resources to which the {"env": "online"} and {"project": "A"} tags are added to the resource group rg-aek2********qcy.
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// custom_condition
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s CreateAutoGroupingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoGroupingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoGroupingRuleRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *CreateAutoGroupingRuleRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *CreateAutoGroupingRuleRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *CreateAutoGroupingRuleRequest) GetExcludeResourceTypesScope() *string {
	return s.ExcludeResourceTypesScope
}

func (s *CreateAutoGroupingRuleRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *CreateAutoGroupingRuleRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *CreateAutoGroupingRuleRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *CreateAutoGroupingRuleRequest) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *CreateAutoGroupingRuleRequest) GetRuleContents() []*CreateAutoGroupingRuleRequestRuleContents {
	return s.RuleContents
}

func (s *CreateAutoGroupingRuleRequest) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *CreateAutoGroupingRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateAutoGroupingRuleRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *CreateAutoGroupingRuleRequest) SetExcludeRegionIdsScope(v string) *CreateAutoGroupingRuleRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetExcludeResourceGroupIdsScope(v string) *CreateAutoGroupingRuleRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetExcludeResourceIdsScope(v string) *CreateAutoGroupingRuleRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetExcludeResourceTypesScope(v string) *CreateAutoGroupingRuleRequest {
	s.ExcludeResourceTypesScope = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetRegionIdsScope(v string) *CreateAutoGroupingRuleRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetResourceGroupIdsScope(v string) *CreateAutoGroupingRuleRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetResourceIdsScope(v string) *CreateAutoGroupingRuleRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetResourceTypesScope(v string) *CreateAutoGroupingRuleRequest {
	s.ResourceTypesScope = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetRuleContents(v []*CreateAutoGroupingRuleRequestRuleContents) *CreateAutoGroupingRuleRequest {
	s.RuleContents = v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetRuleDesc(v string) *CreateAutoGroupingRuleRequest {
	s.RuleDesc = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetRuleName(v string) *CreateAutoGroupingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) SetRuleType(v string) *CreateAutoGroupingRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateAutoGroupingRuleRequest) Validate() error {
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

type CreateAutoGroupingRuleRequestRuleContents struct {
	// The condition for the range of resources to be automatically transferred.
	//
	// example:
	//
	// {"children":[{"desired":"{\\"env\\":\\"online\\", \\"project\\":\\"A\\"}","featurePath":"$.tags","featureSource":"RESOURCE","operator":"TagMatchAll"}],"operator":"and"}
	AutoGroupingScopeCondition *string `json:"AutoGroupingScopeCondition,omitempty" xml:"AutoGroupingScopeCondition,omitempty"`
	// The condition for the destination resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"children":[{"desired":"rg-aek2********qcy","featurePath":"$.resourceGroupId","featureSource":"RESOURCE","operator":"StringEquals"}],"operator":"and"}
	TargetResourceGroupCondition *string `json:"TargetResourceGroupCondition,omitempty" xml:"TargetResourceGroupCondition,omitempty"`
}

func (s CreateAutoGroupingRuleRequestRuleContents) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoGroupingRuleRequestRuleContents) GoString() string {
	return s.String()
}

func (s *CreateAutoGroupingRuleRequestRuleContents) GetAutoGroupingScopeCondition() *string {
	return s.AutoGroupingScopeCondition
}

func (s *CreateAutoGroupingRuleRequestRuleContents) GetTargetResourceGroupCondition() *string {
	return s.TargetResourceGroupCondition
}

func (s *CreateAutoGroupingRuleRequestRuleContents) SetAutoGroupingScopeCondition(v string) *CreateAutoGroupingRuleRequestRuleContents {
	s.AutoGroupingScopeCondition = &v
	return s
}

func (s *CreateAutoGroupingRuleRequestRuleContents) SetTargetResourceGroupCondition(v string) *CreateAutoGroupingRuleRequestRuleContents {
	s.TargetResourceGroupCondition = &v
	return s
}

func (s *CreateAutoGroupingRuleRequestRuleContents) Validate() error {
	return dara.Validate(s)
}
