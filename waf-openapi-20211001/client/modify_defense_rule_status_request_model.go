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
	DefenseType *string `json:"DefenseType,omitempty" xml:"DefenseType,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule whose status you want to change.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20002615
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The new status of the protection rule. Valid values:
	//
	// 	- **0:*	- disabled.
	//
	// 	- **1:*	- enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleStatus *int32 `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The ID of the protection rule template to which the protection rule whose status you want to change belongs.
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
