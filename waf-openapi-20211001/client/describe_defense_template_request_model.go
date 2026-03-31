// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDefenseTemplateRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDefenseTemplateRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplateRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *DescribeDefenseTemplateRequest
	GetTemplateId() *int64
}

type DescribeDefenseTemplateRequest struct {
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
	// The region of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
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
	// The ID of the protection rule template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1333
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDefenseTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseTemplateRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDefenseTemplateRequest) SetInstanceId(v string) *DescribeDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseTemplateRequest) SetRegionId(v string) *DescribeDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseTemplateRequest) SetTemplateId(v int64) *DescribeDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDefenseTemplateRequest) Validate() error {
	return dara.Validate(s)
}
