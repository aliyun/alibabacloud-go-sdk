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
	SetDefenseSubScene(v string) *CreateDefenseTemplateRequest
	GetDefenseSubScene() *string
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
	// The protection scenario. For more information, see the **DefenseScene*	- parameter of the [CreateDefenseRule](https://help.aliyun.com/document_detail/461421.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_group
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// example:
	//
	// bot_custom_acl
	DefenseSubScene *string `json:"DefenseSubScene,omitempty" xml:"DefenseSubScene,omitempty"`
	// The description of the protection template.
	//
	// - **bot_custom_acl**: Represents the protection template for advanced custom rules in bot management.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to get the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// - **cn-hangzhou**: Represents the Chinese mainland.
	//
	// - **ap-southeast-1**: Represents regions outside the Chinese mainland.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The name of the protection template. The name must be 1 to 255 characters long and can contain letters, digits, Chinese characters, underscores (_), periods (.), and hyphens (-).
	//
	// > The names of templates for the same protection scenario (**DefenseScene**) must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// test221
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The origin of the protection template. The value must be **custom**.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom
	TemplateOrigin *string `json:"TemplateOrigin,omitempty" xml:"TemplateOrigin,omitempty"`
	// Indicates whether the protection template is enabled. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateStatus *int32 `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
	// The type of the protection template. Valid values:
	//
	// - **user_default**: The user\\"s default template.
	//
	// - **user_custom**: A user-defined template.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_default
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The protected objects to unbind when you create a default template. Use the [**"XX1","XX2",...**] format.
	//
	// > This parameter takes effect only when you create a **default template*	- (**TemplateType*	- is set to **user_default**).
	UnbindResourceGroups []*string `json:"UnbindResourceGroups,omitempty" xml:"UnbindResourceGroups,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud resource group.
	UnbindResources []*string `json:"UnbindResources,omitempty" xml:"UnbindResources,omitempty" type:"Repeated"`
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

func (s *CreateDefenseTemplateRequest) GetDefenseSubScene() *string {
	return s.DefenseSubScene
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

func (s *CreateDefenseTemplateRequest) SetDefenseSubScene(v string) *CreateDefenseTemplateRequest {
	s.DefenseSubScene = &v
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
