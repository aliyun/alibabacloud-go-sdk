// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteDefenseTemplateRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteDefenseTemplateRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteDefenseTemplateRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *DeleteDefenseTemplateRequest
	GetTemplateId() *int64
}

type DeleteDefenseTemplateRequest struct {
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the protection rule template that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3155
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteDefenseTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDefenseTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDefenseTemplateRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteDefenseTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DeleteDefenseTemplateRequest) SetInstanceId(v string) *DeleteDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDefenseTemplateRequest) SetRegionId(v string) *DeleteDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *DeleteDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteDefenseTemplateRequest) SetTemplateId(v int64) *DeleteDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteDefenseTemplateRequest) Validate() error {
	return dara.Validate(s)
}
