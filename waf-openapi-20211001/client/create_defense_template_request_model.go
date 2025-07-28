// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScene(v string) *CreateDefenseTemplateRequest
	GetDefenseScene() *string
	SetDescription(v string) *CreateDefenseTemplateRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateDefenseTemplateRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateDefenseTemplateRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateDefenseTemplateRequest
	GetResourceManagerResourceGroupId() *string
	SetTemplateName(v string) *CreateDefenseTemplateRequest
	GetTemplateName() *string
	SetTemplateOrigin(v string) *CreateDefenseTemplateRequest
	GetTemplateOrigin() *string
	SetTemplateStatus(v int32) *CreateDefenseTemplateRequest
	GetTemplateStatus() *int32
	SetTemplateType(v string) *CreateDefenseTemplateRequest
	GetTemplateType() *string
	SetUnbindResourceGroups(v []*string) *CreateDefenseTemplateRequest
	GetUnbindResourceGroups() []*string
	SetUnbindResources(v []*string) *CreateDefenseTemplateRequest
	GetUnbindResources() []*string
}

type CreateDefenseTemplateRequest struct {
	// The scenario in which you want to use the protection rule template. For more information, see the description of the **DefenseScene*	- parameter in the [CreateDefenseRule](~~CreateDefenseRule~~) topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The description of the protection rule template.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The name of the protection rule template.
	//
	// This parameter is required.
	//
	// example:
	//
	// test221
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The origin of the protection rule template that you want to create. Set the value to **custom**. The value specifies that the protection rule template is a custom template.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	// The status of the protection rule template. Valid values:
	//
	// 	- **0:*	- disabled.
	//
	// 	- **1:*	- enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the protection rule template. Valid values:
	//
	// 	- **user_default:*	- default template.
	//
	// 	- **user_custom:*	- custom template.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_default
	TemplateType         *string   `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	UnbindResourceGroups []*string `json:"UnbindResourceGroups,omitempty" xml:"UnbindResourceGroups,omitempty" type:"Repeated"`
	UnbindResources      []*string `json:"UnbindResources,omitempty" xml:"UnbindResources,omitempty" type:"Repeated"`
}

func (s CreateDefenseTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseTemplateRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *CreateDefenseTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDefenseTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDefenseTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDefenseTemplateRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateDefenseTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateDefenseTemplateRequest) GetTemplateOrigin() *string {
	return s.TemplateOrigin
}

func (s *CreateDefenseTemplateRequest) GetTemplateStatus() *int32 {
	return s.TemplateStatus
}

func (s *CreateDefenseTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateDefenseTemplateRequest) GetUnbindResourceGroups() []*string {
	return s.UnbindResourceGroups
}

func (s *CreateDefenseTemplateRequest) GetUnbindResources() []*string {
	return s.UnbindResources
}

func (s *CreateDefenseTemplateRequest) SetDefenseScene(v string) *CreateDefenseTemplateRequest {
	s.DefenseScene = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetDescription(v string) *CreateDefenseTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetInstanceId(v string) *CreateDefenseTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetRegionId(v string) *CreateDefenseTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetResourceManagerResourceGroupId(v string) *CreateDefenseTemplateRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateName(v string) *CreateDefenseTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateOrigin(v string) *CreateDefenseTemplateRequest {
	s.TemplateOrigin = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateStatus(v int32) *CreateDefenseTemplateRequest {
	s.TemplateStatus = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetTemplateType(v string) *CreateDefenseTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateDefenseTemplateRequest) SetUnbindResourceGroups(v []*string) *CreateDefenseTemplateRequest {
	s.UnbindResourceGroups = v
	return s
}

func (s *CreateDefenseTemplateRequest) SetUnbindResources(v []*string) *CreateDefenseTemplateRequest {
	s.UnbindResources = v
	return s
}

func (s *CreateDefenseTemplateRequest) Validate() error {
	return dara.Validate(s)
}
