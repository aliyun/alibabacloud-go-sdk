// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseRuleBlockIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteDefenseRuleBlockIpRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteDefenseRuleBlockIpRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteDefenseRuleBlockIpRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *DeleteDefenseRuleBlockIpRequest
	GetRuleId() *int64
	SetTemplateId(v int64) *DeleteDefenseRuleBlockIpRequest
	GetTemplateId() *int64
}

type DeleteDefenseRuleBlockIpRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-9lb37yxse05
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
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the scan protection rule for which you want to unblock an IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36998
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the protection template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4057
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteDefenseRuleBlockIpRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseRuleBlockIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleBlockIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDefenseRuleBlockIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDefenseRuleBlockIpRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteDefenseRuleBlockIpRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteDefenseRuleBlockIpRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DeleteDefenseRuleBlockIpRequest) SetInstanceId(v string) *DeleteDefenseRuleBlockIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDefenseRuleBlockIpRequest) SetRegionId(v string) *DeleteDefenseRuleBlockIpRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDefenseRuleBlockIpRequest) SetResourceManagerResourceGroupId(v string) *DeleteDefenseRuleBlockIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteDefenseRuleBlockIpRequest) SetRuleId(v int64) *DeleteDefenseRuleBlockIpRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteDefenseRuleBlockIpRequest) SetTemplateId(v int64) *DeleteDefenseRuleBlockIpRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteDefenseRuleBlockIpRequest) Validate() error {
	return dara.Validate(s)
}
