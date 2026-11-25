// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseType(v string) *DeleteDefenseRuleRequest
	GetDefenseType() *string
	SetDryRun(v bool) *DeleteDefenseRuleRequest
	GetDryRun() *bool
	SetInstanceId(v string) *DeleteDefenseRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteDefenseRuleRequest
	GetRegionId() *string
	SetResource(v string) *DeleteDefenseRuleRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DeleteDefenseRuleRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceType(v string) *DeleteDefenseRuleRequest
	GetResourceType() *string
	SetRuleIds(v string) *DeleteDefenseRuleRequest
	GetRuleIds() *string
	SetTemplateId(v int64) *DeleteDefenseRuleRequest
	GetTemplateId() *int64
}

type DeleteDefenseRuleRequest struct {
	// The type of the protection rule.
	//
	// example:
	//
	// template
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// Specifies whether to enable the dry run mode. If you do not specify this parameter, a normal request is initiated. Valid values:
	//
	// - **true**: initiates a dry run request. The system only checks whether the request meets the execution conditions and does not perform the operations specified in the request. If the dry run fails, the system returns the corresponding error code. If the dry run succeeds, the system returns the Defense.Control.DryRunOperation error code.
	//
	// - **false**: initiates a normal request. If the check succeeds, the system performs the operations specified in the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object to which the rules to be deleted apply.
	//
	// > This parameter is required only when **DefenseType*	- is set to **resource**.
	//
	// example:
	//
	// rencs***-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the protection target. This parameter applies to rules whose **DefenseType*	- is **resource**. Valid values:
	//
	// - **single*	- (default): a single protected object.
	//
	// - **group**: a protection group.
	//
	// example:
	//
	// single
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The IDs of the protection rules to be deleted. Separate multiple rule IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2223455,23354,465565
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	// The ID of the protection template to be deleted.
	//
	// > This parameter is required only when **DefenseType*	- is set to **template**.
	//
	// example:
	//
	// 2221
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleRequest) GetDefenseType() *string {
	return s.DefenseType
}

func (s *DeleteDefenseRuleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteDefenseRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDefenseRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDefenseRuleRequest) GetResource() *string {
	return s.Resource
}

func (s *DeleteDefenseRuleRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteDefenseRuleRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DeleteDefenseRuleRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *DeleteDefenseRuleRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DeleteDefenseRuleRequest) SetDefenseType(v string) *DeleteDefenseRuleRequest {
	s.DefenseType = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetDryRun(v bool) *DeleteDefenseRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetInstanceId(v string) *DeleteDefenseRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetRegionId(v string) *DeleteDefenseRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetResource(v string) *DeleteDefenseRuleRequest {
	s.Resource = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetResourceManagerResourceGroupId(v string) *DeleteDefenseRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetResourceType(v string) *DeleteDefenseRuleRequest {
	s.ResourceType = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetRuleIds(v string) *DeleteDefenseRuleRequest {
	s.RuleIds = &v
	return s
}

func (s *DeleteDefenseRuleRequest) SetTemplateId(v int64) *DeleteDefenseRuleRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteDefenseRuleRequest) Validate() error {
	return dara.Validate(s)
}
