// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRecordTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateLiveRecordTemplateShrinkRequest
	GetName() *string
	SetRecordFormatShrink(v string) *UpdateLiveRecordTemplateShrinkRequest
	GetRecordFormatShrink() *string
	SetTemplateId(v string) *UpdateLiveRecordTemplateShrinkRequest
	GetTemplateId() *string
}

type UpdateLiveRecordTemplateShrinkRequest struct {
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of recording formats.
	//
	// This parameter is required.
	RecordFormatShrink *string `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 69e1f9fe-1e97-11ed-ba64-0c42a1b73d66
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateLiveRecordTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLiveRecordTemplateShrinkRequest) GetRecordFormatShrink() *string {
	return s.RecordFormatShrink
}

func (s *UpdateLiveRecordTemplateShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateLiveRecordTemplateShrinkRequest) SetName(v string) *UpdateLiveRecordTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateLiveRecordTemplateShrinkRequest) SetRecordFormatShrink(v string) *UpdateLiveRecordTemplateShrinkRequest {
	s.RecordFormatShrink = &v
	return s
}

func (s *UpdateLiveRecordTemplateShrinkRequest) SetTemplateId(v string) *UpdateLiveRecordTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateLiveRecordTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
