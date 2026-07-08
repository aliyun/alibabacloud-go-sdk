// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMajorProtectionBlackIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteMajorProtectionBlackIpRequest
	GetInstanceId() *string
	SetIpList(v string) *DeleteMajorProtectionBlackIpRequest
	GetIpList() *string
	SetRegionId(v string) *DeleteMajorProtectionBlackIpRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteMajorProtectionBlackIpRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *DeleteMajorProtectionBlackIpRequest
	GetRuleId() *int64
	SetTemplateId(v int64) *DeleteMajorProtectionBlackIpRequest
	GetTemplateId() *int64
}

type DeleteMajorProtectionBlackIpRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to obtain the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP address blacklist to be deleted. Supports custom IP addresses or IP address segments in the blacklist, and supports both IPv4 and IPv6. Multiple IP addresses are separated by commas (,). For more information, see [Critical event protection](https://help.aliyun.com/document_detail/425591.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX,192.0.XX.XX/24
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// - **cn-hangzhou**: The Chinese mainland.
	//
	// - **ap-southeast-1**: Outside the Chinese mainland.
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
	// The ID of the IP address blacklist rule for critical event protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20013135
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the critical event protection template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5332
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteMajorProtectionBlackIpRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMajorProtectionBlackIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteMajorProtectionBlackIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteMajorProtectionBlackIpRequest) GetIpList() *string {
	return s.IpList
}

func (s *DeleteMajorProtectionBlackIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMajorProtectionBlackIpRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteMajorProtectionBlackIpRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteMajorProtectionBlackIpRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DeleteMajorProtectionBlackIpRequest) SetInstanceId(v string) *DeleteMajorProtectionBlackIpRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetIpList(v string) *DeleteMajorProtectionBlackIpRequest {
	s.IpList = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetRegionId(v string) *DeleteMajorProtectionBlackIpRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetResourceManagerResourceGroupId(v string) *DeleteMajorProtectionBlackIpRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetRuleId(v int64) *DeleteMajorProtectionBlackIpRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) SetTemplateId(v int64) *DeleteMajorProtectionBlackIpRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpRequest) Validate() error {
	return dara.Validate(s)
}
