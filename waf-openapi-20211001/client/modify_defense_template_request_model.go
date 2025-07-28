// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyDefenseTemplateRequest
	GetDescription() *string
	SetInstanceId(v string) *ModifyDefenseTemplateRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefenseTemplateRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefenseTemplateRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateId(v int64) *ModifyDefenseTemplateRequest
	GetTemplateId() *int64
	SetTemplateName(v string) *ModifyDefenseTemplateRequest
	GetTemplateName() *string
}

type ModifyDefenseTemplateRequest struct {
	// The description of the protection rule template whose configurations you want to modify.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the WAF instance.
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
	// The ID of the protection rule template whose configurations you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7392
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the protection rule template whose configurations you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ModifyDefenseTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDefenseTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefenseTemplateRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefenseTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifyDefenseTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifyDefenseTemplateRequest) SetDescription(v string) *ModifyDefenseTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetInstanceId(v string) *ModifyDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetRegionId(v string) *ModifyDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetTemplateId(v int64) *ModifyDefenseTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) SetTemplateName(v string) *ModifyDefenseTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyDefenseTemplateRequest) Validate() error {
	return dara.Validate(s)
}
