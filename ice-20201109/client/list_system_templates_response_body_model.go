// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSystemTemplatesResponseBody
	GetRequestId() *string
	SetSystemTemplateList(v []*ListSystemTemplatesResponseBodySystemTemplateList) *ListSystemTemplatesResponseBody
	GetSystemTemplateList() []*ListSystemTemplatesResponseBodySystemTemplateList
	SetTotal(v int32) *ListSystemTemplatesResponseBody
	GetTotal() *int32
}

type ListSystemTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried templates.
	SystemTemplateList []*ListSystemTemplatesResponseBodySystemTemplateList `json:"SystemTemplateList,omitempty" xml:"SystemTemplateList,omitempty" type:"Repeated"`
	// The total number of templates.
	//
	// example:
	//
	// 20
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSystemTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSystemTemplatesResponseBody) GetSystemTemplateList() []*ListSystemTemplatesResponseBodySystemTemplateList {
	return s.SystemTemplateList
}

func (s *ListSystemTemplatesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListSystemTemplatesResponseBody) SetRequestId(v string) *ListSystemTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemTemplatesResponseBody) SetSystemTemplateList(v []*ListSystemTemplatesResponseBodySystemTemplateList) *ListSystemTemplatesResponseBody {
	s.SystemTemplateList = v
	return s
}

func (s *ListSystemTemplatesResponseBody) SetTotal(v int32) *ListSystemTemplatesResponseBody {
	s.Total = &v
	return s
}

func (s *ListSystemTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSystemTemplatesResponseBodySystemTemplateList struct {
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
	// Remux
	SubtypeName *string `json:"SubtypeName,omitempty" xml:"SubtypeName,omitempty"`
	// The template parameters.
	//
	// example:
	//
	// {"Container":{"Format":"flv"},"Video":{},"Audio":{}}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The template ID.
	//
	// example:
	//
	// S00000001-000000
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// FLV-COPY
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

func (s ListSystemTemplatesResponseBodySystemTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListSystemTemplatesResponseBodySystemTemplateList) GoString() string {
	return s.String()
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) GetStatus() *string {
	return s.Status
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) GetSubtype() *int32 {
	return s.Subtype
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) GetSubtypeName() *string {
	return s.SubtypeName
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) GetType() *int32 {
	return s.Type
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) GetTypeName() *string {
	return s.TypeName
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) SetStatus(v string) *ListSystemTemplatesResponseBodySystemTemplateList {
	s.Status = &v
	return s
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) SetSubtype(v int32) *ListSystemTemplatesResponseBodySystemTemplateList {
	s.Subtype = &v
	return s
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) SetSubtypeName(v string) *ListSystemTemplatesResponseBodySystemTemplateList {
	s.SubtypeName = &v
	return s
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) SetTemplateConfig(v string) *ListSystemTemplatesResponseBodySystemTemplateList {
	s.TemplateConfig = &v
	return s
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) SetTemplateId(v string) *ListSystemTemplatesResponseBodySystemTemplateList {
	s.TemplateId = &v
	return s
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) SetTemplateName(v string) *ListSystemTemplatesResponseBodySystemTemplateList {
	s.TemplateName = &v
	return s
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) SetType(v int32) *ListSystemTemplatesResponseBodySystemTemplateList {
	s.Type = &v
	return s
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) SetTypeName(v string) *ListSystemTemplatesResponseBodySystemTemplateList {
	s.TypeName = &v
	return s
}

func (s *ListSystemTemplatesResponseBodySystemTemplateList) Validate() error {
	return dara.Validate(s)
}
