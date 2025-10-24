// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMajorProtectionBlackIpV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateMajorProtectionBlackIpV2Request
	GetDescription() *string
	SetExpiredTime(v int64) *CreateMajorProtectionBlackIpV2Request
	GetExpiredTime() *int64
	SetInstanceId(v string) *CreateMajorProtectionBlackIpV2Request
	GetInstanceId() *string
	SetIpList(v string) *CreateMajorProtectionBlackIpV2Request
	GetIpList() *string
	SetRegionId(v string) *CreateMajorProtectionBlackIpV2Request
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateMajorProtectionBlackIpV2Request
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *CreateMajorProtectionBlackIpV2Request
	GetRuleId() *int64
	SetTemplateId(v int64) *CreateMajorProtectionBlackIpV2Request
	GetTemplateId() *int64
}

type CreateMajorProtectionBlackIpV2Request struct {
	// example:
	//
	// Protection for major events
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1716528465
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-2r42s6y****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX,192.0.XX.XX/24
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12399
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2221
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateMajorProtectionBlackIpV2Request) String() string {
	return dara.Prettify(s)
}

func (s CreateMajorProtectionBlackIpV2Request) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpV2Request) GetDescription() *string {
	return s.Description
}

func (s *CreateMajorProtectionBlackIpV2Request) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *CreateMajorProtectionBlackIpV2Request) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMajorProtectionBlackIpV2Request) GetIpList() *string {
	return s.IpList
}

func (s *CreateMajorProtectionBlackIpV2Request) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMajorProtectionBlackIpV2Request) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateMajorProtectionBlackIpV2Request) GetRuleId() *int64 {
	return s.RuleId
}

func (s *CreateMajorProtectionBlackIpV2Request) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CreateMajorProtectionBlackIpV2Request) SetDescription(v string) *CreateMajorProtectionBlackIpV2Request {
	s.Description = &v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Request) SetExpiredTime(v int64) *CreateMajorProtectionBlackIpV2Request {
	s.ExpiredTime = &v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Request) SetInstanceId(v string) *CreateMajorProtectionBlackIpV2Request {
	s.InstanceId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Request) SetIpList(v string) *CreateMajorProtectionBlackIpV2Request {
	s.IpList = &v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Request) SetRegionId(v string) *CreateMajorProtectionBlackIpV2Request {
	s.RegionId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Request) SetResourceManagerResourceGroupId(v string) *CreateMajorProtectionBlackIpV2Request {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Request) SetRuleId(v int64) *CreateMajorProtectionBlackIpV2Request {
	s.RuleId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Request) SetTemplateId(v int64) *CreateMajorProtectionBlackIpV2Request {
	s.TemplateId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpV2Request) Validate() error {
	return dara.Validate(s)
}
