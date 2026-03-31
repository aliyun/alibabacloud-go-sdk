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
	// The description of the IP address blacklist.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time after which the IP address blacklist becomes invalid. Unit: seconds.
	//
	// >  If you set this parameter to **0**, the blacklist is permanently valid.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1662603328
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
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
	// The IP addresses that you want to add to the IP address blacklist. You can specify multiple CIDR blocks or IP addresses. IPv4 and IPv6 addresses are supported. Separate the CIDR blocks or IP addresses with commas (,). For more information, see [Protection for major events](https://help.aliyun.com/document_detail/425591.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 12.XX.XX.2,3.XX.XX.3/24
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
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
	// The ID of the IP address blacklist rule for major event protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20012033
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the IP address blacklist rule template for major event protection.
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
