// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseRuleCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyDefenseRuleCacheRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefenseRuleCacheRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefenseRuleCacheRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *ModifyDefenseRuleCacheRequest
	GetRuleId() *int64
	SetTemplateId(v int64) *ModifyDefenseRuleCacheRequest
	GetTemplateId() *int64
}

type ModifyDefenseRuleCacheRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-zvp****xm2r
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
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
	// The ID of the protection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the protection template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyDefenseRuleCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseRuleCacheRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleCacheRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseRuleCacheRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefenseRuleCacheRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefenseRuleCacheRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyDefenseRuleCacheRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifyDefenseRuleCacheRequest) SetInstanceId(v string) *ModifyDefenseRuleCacheRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseRuleCacheRequest) SetRegionId(v string) *ModifyDefenseRuleCacheRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseRuleCacheRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseRuleCacheRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseRuleCacheRequest) SetRuleId(v int64) *ModifyDefenseRuleCacheRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyDefenseRuleCacheRequest) SetTemplateId(v int64) *ModifyDefenseRuleCacheRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyDefenseRuleCacheRequest) Validate() error {
	return dara.Validate(s)
}
