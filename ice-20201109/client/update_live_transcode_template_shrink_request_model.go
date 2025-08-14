// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveTranscodeTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateLiveTranscodeTemplateShrinkRequest
	GetName() *string
	SetTemplateConfigShrink(v string) *UpdateLiveTranscodeTemplateShrinkRequest
	GetTemplateConfigShrink() *string
	SetTemplateId(v string) *UpdateLiveTranscodeTemplateShrinkRequest
	GetTemplateId() *string
}

type UpdateLiveTranscodeTemplateShrinkRequest struct {
	// The template name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration of the template.
	TemplateConfigShrink *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The template ID. To obtain the template ID, log on to the [Intelligent Media Services (IMS) console](https://ims.console.aliyun.com/summary), choose Real-time Media Processing > Template Management, and then click the Transcoding tab. Alternatively, find the ID from the response parameters of the [CreateLiveTranscodeTemplate](https://help.aliyun.com/document_detail/449217.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateLiveTranscodeTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLiveTranscodeTemplateShrinkRequest) GetTemplateConfigShrink() *string {
	return s.TemplateConfigShrink
}

func (s *UpdateLiveTranscodeTemplateShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateLiveTranscodeTemplateShrinkRequest) SetName(v string) *UpdateLiveTranscodeTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateShrinkRequest) SetTemplateConfigShrink(v string) *UpdateLiveTranscodeTemplateShrinkRequest {
	s.TemplateConfigShrink = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateShrinkRequest) SetTemplateId(v string) *UpdateLiveTranscodeTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
