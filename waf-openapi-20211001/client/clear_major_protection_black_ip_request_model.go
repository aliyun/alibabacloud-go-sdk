// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearMajorProtectionBlackIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ClearMajorProtectionBlackIpRequest
	GetInstanceId() *string
	SetRegionId(v string) *ClearMajorProtectionBlackIpRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ClearMajorProtectionBlackIpRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *ClearMajorProtectionBlackIpRequest
	GetRuleId() *int64
	SetTemplateId(v int64) *ClearMajorProtectionBlackIpRequest
	GetTemplateId() *int64
}

type ClearMajorProtectionBlackIpRequest struct {
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
	// The ID of the Alibaba Cloud resource group.
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

func (s ClearMajorProtectionBlackIpRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *ClearMajorProtectionBlackIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ClearMajorProtectionBlackIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ClearMajorProtectionBlackIpRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ClearMajorProtectionBlackIpRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ClearMajorProtectionBlackIpRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ClearMajorProtectionBlackIpRequest) SetInstanceId(v string) *ClearMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) SetRegionId(v string) *ClearMajorProtectionBlackIpRequest {
	s.RegionId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) SetResourceManagerResourceGroupId(v string) *ClearMajorProtectionBlackIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) SetRuleId(v int64) *ClearMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) SetTemplateId(v int64) *ClearMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpRequest) Validate() error {
	return dara.Validate(s)
}
