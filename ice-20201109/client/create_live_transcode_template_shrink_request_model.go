// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveTranscodeTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateLiveTranscodeTemplateShrinkRequest
	GetName() *string
	SetTemplateConfigShrink(v string) *CreateLiveTranscodeTemplateShrinkRequest
	GetTemplateConfigShrink() *string
	SetType(v string) *CreateLiveTranscodeTemplateShrinkRequest
	GetType() *string
}

type CreateLiveTranscodeTemplateShrinkRequest struct {
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// my template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template configuration.
	//
	// > The pass parameter requirements vary based on the templatetype (Type). When Type is set to normal, at least one of the width and height parameters must be specified, and the frame rate and bitrate parameters are required. For other template types, specify the parameters based on your requirements.
	TemplateConfigShrink *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The template type. Valid values:
	//
	// - normal: standard.
	//
	// - narrow-band: narrowband HD.
	//
	// - audio-only: audio only.
	//
	// - origin: original quality.
	//
	// This parameter is required.
	//
	// example:
	//
	// normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateLiveTranscodeTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveTranscodeTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveTranscodeTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateLiveTranscodeTemplateShrinkRequest) GetTemplateConfigShrink() *string {
	return s.TemplateConfigShrink
}

func (s *CreateLiveTranscodeTemplateShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateLiveTranscodeTemplateShrinkRequest) SetName(v string) *CreateLiveTranscodeTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateLiveTranscodeTemplateShrinkRequest) SetTemplateConfigShrink(v string) *CreateLiveTranscodeTemplateShrinkRequest {
	s.TemplateConfigShrink = &v
	return s
}

func (s *CreateLiveTranscodeTemplateShrinkRequest) SetType(v string) *CreateLiveTranscodeTemplateShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateLiveTranscodeTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
