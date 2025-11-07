// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v *UpdateTemplateResponseBodyTemplate) *UpdateTemplateResponseBody
	GetTemplate() *UpdateTemplateResponseBodyTemplate
}

type UpdateTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2075899A-585D-4A41-A9B2-28DF4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the template.
	Template *UpdateTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s UpdateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTemplateResponseBody) GetTemplate() *UpdateTemplateResponseBodyTemplate {
	return s.Template
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplate(v *UpdateTemplateResponseBodyTemplate) *UpdateTemplateResponseBody {
	s.Template = v
	return s
}

func (s *UpdateTemplateResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTemplateResponseBodyTemplate struct {
	// The user who created the template.
	//
	// example:
	//
	// root(130920000)
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
	// Indicates whether the template is configured with a trigger.
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
	// The share type of the template. The share type of a user-created template is **Private**.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
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
	// t-94753deed38
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
	// v2
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The user who last modified the information about the template.
	//
	// example:
	//
	// root(1309000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the information about the template was last modified.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateTemplateResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplate) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *UpdateTemplateResponseBodyTemplate) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *UpdateTemplateResponseBodyTemplate) GetDescription() *string {
	return s.Description
}

func (s *UpdateTemplateResponseBodyTemplate) GetHasTrigger() *bool {
	return s.HasTrigger
}

func (s *UpdateTemplateResponseBodyTemplate) GetHash() *string {
	return s.Hash
}

func (s *UpdateTemplateResponseBodyTemplate) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateTemplateResponseBodyTemplate) GetShareType() *string {
	return s.ShareType
}

func (s *UpdateTemplateResponseBodyTemplate) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *UpdateTemplateResponseBodyTemplate) GetTemplateFormat() *string {
	return s.TemplateFormat
}

func (s *UpdateTemplateResponseBodyTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateTemplateResponseBodyTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateTemplateResponseBodyTemplate) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateTemplateResponseBodyTemplate) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *UpdateTemplateResponseBodyTemplate) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *UpdateTemplateResponseBodyTemplate) SetCreatedBy(v string) *UpdateTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetCreatedDate(v string) *UpdateTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetDescription(v string) *UpdateTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetHasTrigger(v bool) *UpdateTemplateResponseBodyTemplate {
	s.HasTrigger = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetHash(v string) *UpdateTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetResourceGroupId(v string) *UpdateTemplateResponseBodyTemplate {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetShareType(v string) *UpdateTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *UpdateTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateFormat(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateId(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateName(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateVersion(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetUpdatedBy(v string) *UpdateTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetUpdatedDate(v string) *UpdateTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) Validate() error {
	return dara.Validate(s)
}
