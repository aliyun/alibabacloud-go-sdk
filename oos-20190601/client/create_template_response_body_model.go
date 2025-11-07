// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v *CreateTemplateResponseBodyTemplate) *CreateTemplateResponseBody
	GetTemplate() *CreateTemplateResponseBodyTemplate
	SetTemplateType(v string) *CreateTemplateResponseBody
	GetTemplateType() *string
}

type CreateTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the template.
	Template *CreateTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	// The type of the template.
	//
	// example:
	//
	// Private
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTemplateResponseBody) GetTemplate() *CreateTemplateResponseBodyTemplate {
	return s.Template
}

func (s *CreateTemplateResponseBody) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplate(v *CreateTemplateResponseBodyTemplate) *CreateTemplateResponseBody {
	s.Template = v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateType(v string) *CreateTemplateResponseBody {
	s.TemplateType = &v
	return s
}

func (s *CreateTemplateResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTemplateResponseBodyTemplate struct {
	// The creator of the template.
	//
	// example:
	//
	// root(13090000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Describe instances of given status
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the template was configured with a trigger.
	//
	// example:
	//
	// true
	HasTrigger *bool `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	// The SHA-256 value of the template content.
	//
	// example:
	//
	// 4bc7d7a21b3e003434b9c223f6e6d2578b5ebfeb5be28c1fcf8a8a1b11907bb4
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. The share type of the template that you create is Private.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags of the resources.
	//
	// example:
	//
	// {     "k1":"v1",     "k2":"v2" }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the template. The system automatically determines whether the format is JSON or YAML.
	//
	// example:
	//
	// JSON
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// t-94753d38
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template. The name of the version consists of the letter v and a number. The number starts from 1.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The Alibaba Cloud account that last modified the information about the template.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the template was last updated.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreateTemplateResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBodyTemplate) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CreateTemplateResponseBodyTemplate) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *CreateTemplateResponseBodyTemplate) GetDescription() *string {
	return s.Description
}

func (s *CreateTemplateResponseBodyTemplate) GetHasTrigger() *bool {
	return s.HasTrigger
}

func (s *CreateTemplateResponseBodyTemplate) GetHash() *string {
	return s.Hash
}

func (s *CreateTemplateResponseBodyTemplate) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTemplateResponseBodyTemplate) GetShareType() *string {
	return s.ShareType
}

func (s *CreateTemplateResponseBodyTemplate) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateTemplateResponseBodyTemplate) GetTemplateFormat() *string {
	return s.TemplateFormat
}

func (s *CreateTemplateResponseBodyTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateTemplateResponseBodyTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateTemplateResponseBodyTemplate) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *CreateTemplateResponseBodyTemplate) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *CreateTemplateResponseBodyTemplate) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *CreateTemplateResponseBodyTemplate) SetCreatedBy(v string) *CreateTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetCreatedDate(v string) *CreateTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetDescription(v string) *CreateTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetHasTrigger(v bool) *CreateTemplateResponseBodyTemplate {
	s.HasTrigger = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetHash(v string) *CreateTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetResourceGroupId(v string) *CreateTemplateResponseBodyTemplate {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetShareType(v string) *CreateTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *CreateTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateFormat(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateId(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateName(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateVersion(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetUpdatedBy(v string) *CreateTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetUpdatedDate(v string) *CreateTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) Validate() error {
	return dara.Validate(s)
}
