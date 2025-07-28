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
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The IDs of the protection rules that you want to delete. Separate the IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2223455,23354,465565
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	// The ID of the protection rule template to which the protection rule that you want to delete belongs.
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
