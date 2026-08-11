// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveTranscodeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v string) *DeleteLiveTranscodeTemplateRequest
	GetTemplateId() *string
}

type DeleteLiveTranscodeTemplateRequest struct {
	// The template ID. You can obtain the ID from the [Intelligent Media Services console](https://ice.console.aliyun.com/summary) > Template Management > Real-time Transcoding Template, or from the response parameters of [CreateLiveTranscodeTemplate](https://help.aliyun.com/document_detail/449217.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteLiveTranscodeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveTranscodeTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveTranscodeTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteLiveTranscodeTemplateRequest) SetTemplateId(v string) *DeleteLiveTranscodeTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteLiveTranscodeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
