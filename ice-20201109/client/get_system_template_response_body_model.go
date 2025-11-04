// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSystemTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSystemTemplateResponseBody
	GetRequestId() *string
	SetSystemTemplate(v *GetSystemTemplateResponseBodySystemTemplate) *GetSystemTemplateResponseBody
	GetSystemTemplate() *GetSystemTemplateResponseBodySystemTemplate
}

type GetSystemTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template information.
	SystemTemplate *GetSystemTemplateResponseBodySystemTemplate `json:"SystemTemplate,omitempty" xml:"SystemTemplate,omitempty" type:"Struct"`
}

func (s GetSystemTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSystemTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSystemTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSystemTemplateResponseBody) GetSystemTemplate() *GetSystemTemplateResponseBodySystemTemplate {
	return s.SystemTemplate
}

func (s *GetSystemTemplateResponseBody) SetRequestId(v string) *GetSystemTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSystemTemplateResponseBody) SetSystemTemplate(v *GetSystemTemplateResponseBodySystemTemplate) *GetSystemTemplateResponseBody {
	s.SystemTemplate = v
	return s
}

func (s *GetSystemTemplateResponseBody) Validate() error {
	if s.SystemTemplate != nil {
		if err := s.SystemTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSystemTemplateResponseBodySystemTemplate struct {
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
	// {"Container":{"Format":"m3u8"},"TransConfig":{"TransMode":"onepass"},"Video":{"Codec":"H.264","Maxrate":8000,"Preset":"medium","PixFmt":"yuv420p","Width":2048,"Bitrate":3500},"Audio":{"Codec":"aac","Bitrate":160,"Samplerate":44100,"Channels":2}}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The template ID.
	//
	// example:
	//
	// S00000001-100060
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// M3U8-2K
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

func (s GetSystemTemplateResponseBodySystemTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetSystemTemplateResponseBodySystemTemplate) GoString() string {
	return s.String()
}

func (s *GetSystemTemplateResponseBodySystemTemplate) GetStatus() *string {
	return s.Status
}

func (s *GetSystemTemplateResponseBodySystemTemplate) GetSubtype() *int32 {
	return s.Subtype
}

func (s *GetSystemTemplateResponseBodySystemTemplate) GetSubtypeName() *string {
	return s.SubtypeName
}

func (s *GetSystemTemplateResponseBodySystemTemplate) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *GetSystemTemplateResponseBodySystemTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetSystemTemplateResponseBodySystemTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetSystemTemplateResponseBodySystemTemplate) GetType() *int32 {
	return s.Type
}

func (s *GetSystemTemplateResponseBodySystemTemplate) GetTypeName() *string {
	return s.TypeName
}

func (s *GetSystemTemplateResponseBodySystemTemplate) SetStatus(v string) *GetSystemTemplateResponseBodySystemTemplate {
	s.Status = &v
	return s
}

func (s *GetSystemTemplateResponseBodySystemTemplate) SetSubtype(v int32) *GetSystemTemplateResponseBodySystemTemplate {
	s.Subtype = &v
	return s
}

func (s *GetSystemTemplateResponseBodySystemTemplate) SetSubtypeName(v string) *GetSystemTemplateResponseBodySystemTemplate {
	s.SubtypeName = &v
	return s
}

func (s *GetSystemTemplateResponseBodySystemTemplate) SetTemplateConfig(v string) *GetSystemTemplateResponseBodySystemTemplate {
	s.TemplateConfig = &v
	return s
}

func (s *GetSystemTemplateResponseBodySystemTemplate) SetTemplateId(v string) *GetSystemTemplateResponseBodySystemTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetSystemTemplateResponseBodySystemTemplate) SetTemplateName(v string) *GetSystemTemplateResponseBodySystemTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetSystemTemplateResponseBodySystemTemplate) SetType(v int32) *GetSystemTemplateResponseBodySystemTemplate {
	s.Type = &v
	return s
}

func (s *GetSystemTemplateResponseBodySystemTemplate) SetTypeName(v string) *GetSystemTemplateResponseBodySystemTemplate {
	s.TypeName = &v
	return s
}

func (s *GetSystemTemplateResponseBodySystemTemplate) Validate() error {
	return dara.Validate(s)
}
