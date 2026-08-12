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
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to obtain the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IP blacklist to delete. Custom IP addresses or CIDR blocks are supported, including both IPv4 and IPv6. Separate multiple IP addresses with commas (,).
	//
	// For more information, see [Critical event protection](https://help.aliyun.com/document_detail/425591.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.XX.XX,192.0.XX.XX/24
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
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
	// The ID of the IP blacklist rule for critical event protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20013135
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The ID of the critical event protection scenario template.
	//
	// > This parameter requires the ID of a protection template of the critical event protection type. You can create this type of template only after purchasing the critical event protection upgrade.
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
