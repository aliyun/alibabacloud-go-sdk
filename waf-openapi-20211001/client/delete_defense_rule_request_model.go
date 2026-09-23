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
	// Specifies whether to enable the DryRun dry run mode. If you do not specify this parameter, a normal request is sent. Valid values:
	//
	// - **true**: A dry run request is sent. The system checks whether the request meets the execution conditions without performing the specified operation. If the dry run fails, the corresponding error code is returned. If the dry run succeeds, the error code Defense.Control.DryRunOperation is returned.
	//
	// - **false**: A normal request is sent. The specified operation is performed after the request passes the check.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object associated with the rule to delete.
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
	// The IDs of the protection rules to delete. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2223455,23354,465565
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	// The ID of the protection template to delete.
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
