// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseRuleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseType(v string) *ModifyDefenseRuleStatusRequest
	GetDefenseType() *string
	SetInstanceId(v string) *ModifyDefenseRuleStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefenseRuleStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefenseRuleStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *ModifyDefenseRuleStatusRequest
	GetRuleId() *int64
	SetRuleStatus(v int32) *ModifyDefenseRuleStatusRequest
	GetRuleStatus() *int32
	SetTemplateId(v int64) *ModifyDefenseRuleStatusRequest
	GetTemplateId() *int64
}

type ModifyDefenseRuleStatusRequest struct {
	// The type of the protection rule.
	//
	// example:
	//
	// template
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the current WAF instance.
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
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20002615
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The status of the protection rule to set. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleStatus *int32 `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The ID of the protection rule template.
	//
	// example:
	//
	// 7239
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyDefenseRuleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseRuleStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleStatusRequest) GetDefenseType() *string {
	return s.DefenseType
}

func (s *ModifyDefenseRuleStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseRuleStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefenseRuleStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefenseRuleStatusRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyDefenseRuleStatusRequest) GetRuleStatus() *int32 {
	return s.RuleStatus
}

func (s *ModifyDefenseRuleStatusRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifyDefenseRuleStatusRequest) SetDefenseType(v string) *ModifyDefenseRuleStatusRequest {
	s.DefenseType = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetInstanceId(v string) *ModifyDefenseRuleStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetRegionId(v string) *ModifyDefenseRuleStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseRuleStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetRuleId(v int64) *ModifyDefenseRuleStatusRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetRuleStatus(v int32) *ModifyDefenseRuleStatusRequest {
	s.RuleStatus = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) SetTemplateId(v int64) *ModifyDefenseRuleStatusRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyDefenseRuleStatusRequest) Validate() error {
	return dara.Validate(s)
}
