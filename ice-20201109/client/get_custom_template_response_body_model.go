// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTemplate(v *GetCustomTemplateResponseBodyCustomTemplate) *GetCustomTemplateResponseBody
	GetCustomTemplate() *GetCustomTemplateResponseBodyCustomTemplate
	SetRequestId(v string) *GetCustomTemplateResponseBody
	GetRequestId() *string
}

type GetCustomTemplateResponseBody struct {
	// The template information.
	CustomTemplate *GetCustomTemplateResponseBodyCustomTemplate `json:"CustomTemplate,omitempty" xml:"CustomTemplate,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCustomTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomTemplateResponseBody) GetCustomTemplate() *GetCustomTemplateResponseBodyCustomTemplate {
	return s.CustomTemplate
}

func (s *GetCustomTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomTemplateResponseBody) SetCustomTemplate(v *GetCustomTemplateResponseBodyCustomTemplate) *GetCustomTemplateResponseBody {
	s.CustomTemplate = v
	return s
}

func (s *GetCustomTemplateResponseBody) SetRequestId(v string) *GetCustomTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCustomTemplateResponseBodyCustomTemplate struct {
	// The time when the template was created.
	//
	// example:
	//
	// 2022-01-01T10:00:00Z
	CreateTime   *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FrontendHint *GetCustomTemplateResponseBodyCustomTemplateFrontendHint `json:"FrontendHint,omitempty" xml:"FrontendHint,omitempty" type:"Struct"`
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
	// 2022-01-01T11:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The template state.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subtype ID of the template.
	//
	// example:
	//
	// 1
	Subtype *int32 `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	// The subtype name of the template.
	//
	// example:
	//
	// Normal
	SubtypeName *string `json:"SubtypeName,omitempty" xml:"SubtypeName,omitempty"`
	// The template parameters.
	//
	// example:
	//
	// {"Type":"Normal","FrameType":"normal","Time":0,"Count":10}
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
	// 测试转码模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type ID of the template.
	//
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The type name of the template.
	//
	// example:
	//
	// SnapshotTemplate
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s GetCustomTemplateResponseBodyCustomTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTemplateResponseBodyCustomTemplate) GoString() string {
	return s.String()
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetFrontendHint() *GetCustomTemplateResponseBodyCustomTemplateFrontendHint {
	return s.FrontendHint
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetStatus() *string {
	return s.Status
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetSubtype() *int32 {
	return s.Subtype
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetSubtypeName() *string {
	return s.SubtypeName
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetType() *int32 {
	return s.Type
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) GetTypeName() *string {
	return s.TypeName
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetCreateTime(v string) *GetCustomTemplateResponseBodyCustomTemplate {
	s.CreateTime = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetFrontendHint(v *GetCustomTemplateResponseBodyCustomTemplateFrontendHint) *GetCustomTemplateResponseBodyCustomTemplate {
	s.FrontendHint = v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetIsDefault(v bool) *GetCustomTemplateResponseBodyCustomTemplate {
	s.IsDefault = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetModifiedTime(v string) *GetCustomTemplateResponseBodyCustomTemplate {
	s.ModifiedTime = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetStatus(v string) *GetCustomTemplateResponseBodyCustomTemplate {
	s.Status = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetSubtype(v int32) *GetCustomTemplateResponseBodyCustomTemplate {
	s.Subtype = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetSubtypeName(v string) *GetCustomTemplateResponseBodyCustomTemplate {
	s.SubtypeName = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetTemplateConfig(v string) *GetCustomTemplateResponseBodyCustomTemplate {
	s.TemplateConfig = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetTemplateId(v string) *GetCustomTemplateResponseBodyCustomTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetTemplateName(v string) *GetCustomTemplateResponseBodyCustomTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetType(v int32) *GetCustomTemplateResponseBodyCustomTemplate {
	s.Type = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) SetTypeName(v string) *GetCustomTemplateResponseBodyCustomTemplate {
	s.TypeName = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplate) Validate() error {
	return dara.Validate(s)
}

type GetCustomTemplateResponseBodyCustomTemplateFrontendHint struct {
	TranscodeTemplateHint *GetCustomTemplateResponseBodyCustomTemplateFrontendHintTranscodeTemplateHint `json:"TranscodeTemplateHint,omitempty" xml:"TranscodeTemplateHint,omitempty" type:"Struct"`
}

func (s GetCustomTemplateResponseBodyCustomTemplateFrontendHint) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTemplateResponseBodyCustomTemplateFrontendHint) GoString() string {
	return s.String()
}

func (s *GetCustomTemplateResponseBodyCustomTemplateFrontendHint) GetTranscodeTemplateHint() *GetCustomTemplateResponseBodyCustomTemplateFrontendHintTranscodeTemplateHint {
	return s.TranscodeTemplateHint
}

func (s *GetCustomTemplateResponseBodyCustomTemplateFrontendHint) SetTranscodeTemplateHint(v *GetCustomTemplateResponseBodyCustomTemplateFrontendHintTranscodeTemplateHint) *GetCustomTemplateResponseBodyCustomTemplateFrontendHint {
	s.TranscodeTemplateHint = v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplateFrontendHint) Validate() error {
	return dara.Validate(s)
}

type GetCustomTemplateResponseBodyCustomTemplateFrontendHintTranscodeTemplateHint struct {
	BitrateControlType *string `json:"BitrateControlType,omitempty" xml:"BitrateControlType,omitempty"`
}

func (s GetCustomTemplateResponseBodyCustomTemplateFrontendHintTranscodeTemplateHint) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTemplateResponseBodyCustomTemplateFrontendHintTranscodeTemplateHint) GoString() string {
	return s.String()
}

func (s *GetCustomTemplateResponseBodyCustomTemplateFrontendHintTranscodeTemplateHint) GetBitrateControlType() *string {
	return s.BitrateControlType
}

func (s *GetCustomTemplateResponseBodyCustomTemplateFrontendHintTranscodeTemplateHint) SetBitrateControlType(v string) *GetCustomTemplateResponseBodyCustomTemplateFrontendHintTranscodeTemplateHint {
	s.BitrateControlType = &v
	return s
}

func (s *GetCustomTemplateResponseBodyCustomTemplateFrontendHintTranscodeTemplateHint) Validate() error {
	return dara.Validate(s)
}
