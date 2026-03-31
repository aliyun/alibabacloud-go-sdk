// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMajorProtectionBlackIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateMajorProtectionBlackIpRequest
	GetDescription() *string
	SetExpiredTime(v int64) *CreateMajorProtectionBlackIpRequest
	GetExpiredTime() *int64
	SetInstanceId(v string) *CreateMajorProtectionBlackIpRequest
	GetInstanceId() *string
	SetIpList(v string) *CreateMajorProtectionBlackIpRequest
	GetIpList() *string
	SetRegionId(v string) *CreateMajorProtectionBlackIpRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateMajorProtectionBlackIpRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *CreateMajorProtectionBlackIpRequest
	GetRuleId() *int64
	SetTemplateId(v int64) *CreateMajorProtectionBlackIpRequest
	GetTemplateId() *int64
}

type CreateMajorProtectionBlackIpRequest struct {
	// The description of the IP address blacklist.
	//
	// example:
	//
	// Protection for major events
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp after which the IP address blacklist becomes invalid. Unit: seconds.
	//
	// >  If you set the parameter to **0**, the IP address blacklist is always valid.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1716528465
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-2r42s6y****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP addresses that you want to add to the IP address blacklist. CIDR blocks and IP addresses are supported. IPv4 and IPv6 addresses are supported. Separate the CIDR blocks or IP addresses with commas (,). For more information, see [Protection for major events](https://help.aliyun.com/document_detail/425591.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX,192.0.XX.XX/24
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
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
	// The ID of the IP address blacklist rule for major event protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 232324
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the IP address blacklist rule template for major event protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2221
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateMajorProtectionBlackIpRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMajorProtectionBlackIpRequest) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *CreateMajorProtectionBlackIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMajorProtectionBlackIpRequest) GetIpList() *string {
	return s.IpList
}

func (s *CreateMajorProtectionBlackIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMajorProtectionBlackIpRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateMajorProtectionBlackIpRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *CreateMajorProtectionBlackIpRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CreateMajorProtectionBlackIpRequest) SetDescription(v string) *CreateMajorProtectionBlackIpRequest {
	s.Description = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetExpiredTime(v int64) *CreateMajorProtectionBlackIpRequest {
	s.ExpiredTime = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetInstanceId(v string) *CreateMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetIpList(v string) *CreateMajorProtectionBlackIpRequest {
	s.IpList = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetRegionId(v string) *CreateMajorProtectionBlackIpRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetResourceManagerResourceGroupId(v string) *CreateMajorProtectionBlackIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetRuleId(v int64) *CreateMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) SetTemplateId(v int64) *CreateMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpRequest) Validate() error {
	return dara.Validate(s)
}
