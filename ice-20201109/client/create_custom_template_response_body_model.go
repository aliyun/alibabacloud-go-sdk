// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTemplate(v *CreateCustomTemplateResponseBodyCustomTemplate) *CreateCustomTemplateResponseBody
	GetCustomTemplate() *CreateCustomTemplateResponseBodyCustomTemplate
	SetRequestId(v string) *CreateCustomTemplateResponseBody
	GetRequestId() *string
}

type CreateCustomTemplateResponseBody struct {
	// The template information.
	CustomTemplate *CreateCustomTemplateResponseBodyCustomTemplate `json:"CustomTemplate,omitempty" xml:"CustomTemplate,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomTemplateResponseBody) GetCustomTemplate() *CreateCustomTemplateResponseBodyCustomTemplate {
	return s.CustomTemplate
}

func (s *CreateCustomTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomTemplateResponseBody) SetCustomTemplate(v *CreateCustomTemplateResponseBodyCustomTemplate) *CreateCustomTemplateResponseBody {
	s.CustomTemplate = v
	return s
}

func (s *CreateCustomTemplateResponseBody) SetRequestId(v string) *CreateCustomTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomTemplateResponseBody) Validate() error {
	if s.CustomTemplate != nil {
		if err := s.CustomTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomTemplateResponseBodyCustomTemplate struct {
	// The time when the template was created.
	//
	// example:
	//
	// 2022-04-19T02:04:31Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the template is the default template.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The time when the template was last modified.
	//
	// example:
	//
	// 2022-04-19T02:04:31Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The template state.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subtype name of the template.
	//
	// example:
	//
	// Remux
	Subtype *string `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	// The template configurations.
	//
	// example:
	//
	// {"Container":{"Format":"flv"},"Video":{},"Audio":{}}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// test-template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type ID of the template.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The type name of the template.
	//
	// example:
	//
	// TranscodeTemplate
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s CreateCustomTemplateResponseBodyCustomTemplate) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomTemplateResponseBodyCustomTemplate) GoString() string {
	return s.String()
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) GetStatus() *string {
	return s.Status
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) GetSubtype() *string {
	return s.Subtype
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) GetType() *int32 {
	return s.Type
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) GetTypeName() *string {
	return s.TypeName
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) SetCreateTime(v string) *CreateCustomTemplateResponseBodyCustomTemplate {
	s.CreateTime = &v
	return s
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) SetIsDefault(v bool) *CreateCustomTemplateResponseBodyCustomTemplate {
	s.IsDefault = &v
	return s
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) SetModifiedTime(v string) *CreateCustomTemplateResponseBodyCustomTemplate {
	s.ModifiedTime = &v
	return s
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) SetStatus(v string) *CreateCustomTemplateResponseBodyCustomTemplate {
	s.Status = &v
	return s
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) SetSubtype(v string) *CreateCustomTemplateResponseBodyCustomTemplate {
	s.Subtype = &v
	return s
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) SetTemplateConfig(v string) *CreateCustomTemplateResponseBodyCustomTemplate {
	s.TemplateConfig = &v
	return s
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) SetTemplateId(v string) *CreateCustomTemplateResponseBodyCustomTemplate {
	s.TemplateId = &v
	return s
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) SetTemplateName(v string) *CreateCustomTemplateResponseBodyCustomTemplate {
	s.TemplateName = &v
	return s
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) SetType(v int32) *CreateCustomTemplateResponseBodyCustomTemplate {
	s.Type = &v
	return s
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) SetTypeName(v string) *CreateCustomTemplateResponseBodyCustomTemplate {
	s.TypeName = &v
	return s
}

func (s *CreateCustomTemplateResponseBodyCustomTemplate) Validate() error {
	return dara.Validate(s)
}
