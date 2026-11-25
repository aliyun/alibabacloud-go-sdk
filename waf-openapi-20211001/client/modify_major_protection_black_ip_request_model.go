// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMajorProtectionBlackIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyMajorProtectionBlackIpRequest
	GetDescription() *string
	SetExpiredTime(v int64) *ModifyMajorProtectionBlackIpRequest
	GetExpiredTime() *int64
	SetInstanceId(v string) *ModifyMajorProtectionBlackIpRequest
	GetInstanceId() *string
	SetIpList(v string) *ModifyMajorProtectionBlackIpRequest
	GetIpList() *string
	SetRegionId(v string) *ModifyMajorProtectionBlackIpRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyMajorProtectionBlackIpRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *ModifyMajorProtectionBlackIpRequest
	GetRuleId() *int64
	SetTemplateId(v int64) *ModifyMajorProtectionBlackIpRequest
	GetTemplateId() *int64
}

type ModifyMajorProtectionBlackIpRequest struct {
	// The description.
	//
	// example:
	//
	// Protection for major events
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expiration timestamp. Unit: seconds.
	//
	// > A value of **0*	- indicates that the rule permanently takes effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1662603328
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
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
	// The IP blacklist to be added. You can specify custom IP addresses or CIDR blocks for the IP blacklist. Both IPv4 and IPv6 addresses are supported. Separate multiple IP addresses with commas (,).
	//
	// For more information, see [Critical event protection](https://help.aliyun.com/document_detail/425591.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 12.XX.XX.2,3.XX.XX.3/24
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland
	//
	// - **ap-southeast-1**: outside the Chinese mainland
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
	// The ID of the IP blacklist rule for critical event protection. Before you obtain the rule ID, enable the critical event protection feature. Then, call the [DescribeDefenseTemplates](https://help.aliyun.com/document_detail/2712643.html) operation and set DefenseScene to major_protection to obtain the ID of the critical event protection template. After that, call the [DescribeDefenseRules](https://help.aliyun.com/document_detail/461426.html) operation to obtain the ID of the IP blacklist rule under the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20012033
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the critical event protection template. Before you obtain the template ID, enable the critical event protection feature. Then, call the [DescribeDefenseTemplates](https://help.aliyun.com/document_detail/2712643.html) operation and set DefenseScene to major_protection to obtain the ID of the critical event protection template.
	//
	// > You can create this type of template only after you purchase critical event protection through an upgrade.
	//
	// >
	//
	// > For more information about how to enable critical event protection and create a critical event protection rule template, see [Critical event protection](https://help.aliyun.com/document_detail/425591.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5132
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyMajorProtectionBlackIpRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *ModifyMajorProtectionBlackIpRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyMajorProtectionBlackIpRequest) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *ModifyMajorProtectionBlackIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyMajorProtectionBlackIpRequest) GetIpList() *string {
	return s.IpList
}

func (s *ModifyMajorProtectionBlackIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMajorProtectionBlackIpRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyMajorProtectionBlackIpRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyMajorProtectionBlackIpRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifyMajorProtectionBlackIpRequest) SetDescription(v string) *ModifyMajorProtectionBlackIpRequest {
	s.Description = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetExpiredTime(v int64) *ModifyMajorProtectionBlackIpRequest {
	s.ExpiredTime = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetInstanceId(v string) *ModifyMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetIpList(v string) *ModifyMajorProtectionBlackIpRequest {
	s.IpList = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetRegionId(v string) *ModifyMajorProtectionBlackIpRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetResourceManagerResourceGroupId(v string) *ModifyMajorProtectionBlackIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetRuleId(v int64) *ModifyMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) SetTemplateId(v int64) *ModifyMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpRequest) Validate() error {
	return dara.Validate(s)
}
