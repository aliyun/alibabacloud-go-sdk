// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTemplateList(v []*ListCustomTemplatesResponseBodyCustomTemplateList) *ListCustomTemplatesResponseBody
	GetCustomTemplateList() []*ListCustomTemplatesResponseBodyCustomTemplateList
	SetRequestId(v string) *ListCustomTemplatesResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListCustomTemplatesResponseBody
	GetTotal() *int32
}

type ListCustomTemplatesResponseBody struct {
	// The queried templates.
	CustomTemplateList []*ListCustomTemplatesResponseBodyCustomTemplateList `json:"CustomTemplateList,omitempty" xml:"CustomTemplateList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of templates.
	//
	// example:
	//
	// 20
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCustomTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomTemplatesResponseBody) GetCustomTemplateList() []*ListCustomTemplatesResponseBodyCustomTemplateList {
	return s.CustomTemplateList
}

func (s *ListCustomTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomTemplatesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListCustomTemplatesResponseBody) SetCustomTemplateList(v []*ListCustomTemplatesResponseBodyCustomTemplateList) *ListCustomTemplatesResponseBody {
	s.CustomTemplateList = v
	return s
}

func (s *ListCustomTemplatesResponseBody) SetRequestId(v string) *ListCustomTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomTemplatesResponseBody) SetTotal(v int32) *ListCustomTemplatesResponseBody {
	s.Total = &v
	return s
}

func (s *ListCustomTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCustomTemplatesResponseBodyCustomTemplateList struct {
	// The time when the template was created.
	//
	// example:
	//
	// 2022-07-12T16:17:54Z
	CreateTime   *string                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FrontendHint *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHint `json:"FrontendHint,omitempty" xml:"FrontendHint,omitempty" type:"Struct"`
	// Indicates whether the template is the default template.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The time when the template was last modified.
	//
	// example:
	//
	// 2022-07-12T16:17:54Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The template state.
	//
	// Valid values:
	//
	// 	- Normal
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subtype ID of the template.
	//
	// example:
	//
	// 2
	Subtype *int32 `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	// The subtype name of the template.
	//
	// example:
	//
	// AudioTranscode
	SubtypeName *string `json:"SubtypeName,omitempty" xml:"SubtypeName,omitempty"`
	// The template parameters.
	//
	// example:
	//
	// {"Container":{"Format":"mp3"},"Audio":{"Codec":"mp3","Bitrate":"64","Samplerate":"22050","Channels":"2"}}
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

func (s ListCustomTemplatesResponseBodyCustomTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListCustomTemplatesResponseBodyCustomTemplateList) GoString() string {
	return s.String()
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetFrontendHint() *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHint {
	return s.FrontendHint
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetStatus() *string {
	return s.Status
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetSubtype() *int32 {
	return s.Subtype
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetSubtypeName() *string {
	return s.SubtypeName
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetType() *int32 {
	return s.Type
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) GetTypeName() *string {
	return s.TypeName
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetCreateTime(v string) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.CreateTime = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetFrontendHint(v *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHint) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.FrontendHint = v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetIsDefault(v bool) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.IsDefault = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetModifiedTime(v string) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.ModifiedTime = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetStatus(v string) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.Status = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetSubtype(v int32) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.Subtype = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetSubtypeName(v string) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.SubtypeName = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetTemplateConfig(v string) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.TemplateConfig = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetTemplateId(v string) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.TemplateId = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetTemplateName(v string) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.TemplateName = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetType(v int32) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.Type = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) SetTypeName(v string) *ListCustomTemplatesResponseBodyCustomTemplateList {
	s.TypeName = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateList) Validate() error {
	return dara.Validate(s)
}

type ListCustomTemplatesResponseBodyCustomTemplateListFrontendHint struct {
	TranscodeTemplateHint *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHintTranscodeTemplateHint `json:"TranscodeTemplateHint,omitempty" xml:"TranscodeTemplateHint,omitempty" type:"Struct"`
}

func (s ListCustomTemplatesResponseBodyCustomTemplateListFrontendHint) String() string {
	return dara.Prettify(s)
}

func (s ListCustomTemplatesResponseBodyCustomTemplateListFrontendHint) GoString() string {
	return s.String()
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHint) GetTranscodeTemplateHint() *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHintTranscodeTemplateHint {
	return s.TranscodeTemplateHint
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHint) SetTranscodeTemplateHint(v *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHintTranscodeTemplateHint) *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHint {
	s.TranscodeTemplateHint = v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHint) Validate() error {
	return dara.Validate(s)
}

type ListCustomTemplatesResponseBodyCustomTemplateListFrontendHintTranscodeTemplateHint struct {
	BitrateControlType *string `json:"BitrateControlType,omitempty" xml:"BitrateControlType,omitempty"`
}

func (s ListCustomTemplatesResponseBodyCustomTemplateListFrontendHintTranscodeTemplateHint) String() string {
	return dara.Prettify(s)
}

func (s ListCustomTemplatesResponseBodyCustomTemplateListFrontendHintTranscodeTemplateHint) GoString() string {
	return s.String()
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHintTranscodeTemplateHint) GetBitrateControlType() *string {
	return s.BitrateControlType
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHintTranscodeTemplateHint) SetBitrateControlType(v string) *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHintTranscodeTemplateHint {
	s.BitrateControlType = &v
	return s
}

func (s *ListCustomTemplatesResponseBodyCustomTemplateListFrontendHintTranscodeTemplateHint) Validate() error {
	return dara.Validate(s)
}
