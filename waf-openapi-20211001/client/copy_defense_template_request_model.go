// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDefenseTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CopyDefenseTemplateRequest
	GetInstanceId() *string
	SetRegionId(v string) *CopyDefenseTemplateRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CopyDefenseTemplateRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *CopyDefenseTemplateRequest
	GetTemplateId() *int64
}

type CopyDefenseTemplateRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-lbj****x10g
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
	// The ID of the protection template that you want to copy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CopyDefenseTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *CopyDefenseTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CopyDefenseTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopyDefenseTemplateRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CopyDefenseTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *CopyDefenseTemplateRequest) SetInstanceId(v string) *CopyDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) SetRegionId(v string) *CopyDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *CopyDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) SetTemplateId(v int64) *CopyDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *CopyDefenseTemplateRequest) Validate() error {
	return dara.Validate(s)
}
