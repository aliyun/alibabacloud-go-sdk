// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveRecordTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateLiveRecordTemplateShrinkRequest
	GetName() *string
	SetRecordFormatShrink(v string) *CreateLiveRecordTemplateShrinkRequest
	GetRecordFormatShrink() *string
}

type CreateLiveRecordTemplateShrinkRequest struct {
	// The name of the template.
	//
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of recording formats.
	//
	// This parameter is required.
	RecordFormatShrink *string `json:"RecordFormat,omitempty" xml:"RecordFormat,omitempty"`
}

func (s CreateLiveRecordTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveRecordTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRecordTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateLiveRecordTemplateShrinkRequest) GetRecordFormatShrink() *string {
	return s.RecordFormatShrink
}

func (s *CreateLiveRecordTemplateShrinkRequest) SetName(v string) *CreateLiveRecordTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateLiveRecordTemplateShrinkRequest) SetRecordFormatShrink(v string) *CreateLiveRecordTemplateShrinkRequest {
	s.RecordFormatShrink = &v
	return s
}

func (s *CreateLiveRecordTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
