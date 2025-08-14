// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveTranscodeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveTranscodeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveTranscodeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveTranscodeTemplateResponseBody) *UpdateLiveTranscodeTemplateResponse
	GetBody() *UpdateLiveTranscodeTemplateResponseBody
}

type UpdateLiveTranscodeTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveTranscodeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveTranscodeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveTranscodeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveTranscodeTemplateResponse) GetBody() *UpdateLiveTranscodeTemplateResponseBody {
	return s.Body
}

func (s *UpdateLiveTranscodeTemplateResponse) SetHeaders(v map[string]*string) *UpdateLiveTranscodeTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveTranscodeTemplateResponse) SetStatusCode(v int32) *UpdateLiveTranscodeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateResponse) SetBody(v *UpdateLiveTranscodeTemplateResponseBody) *UpdateLiveTranscodeTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveTranscodeTemplateResponse) Validate() error {
	return dara.Validate(s)
}
