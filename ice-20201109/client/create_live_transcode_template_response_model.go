// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveTranscodeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveTranscodeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveTranscodeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveTranscodeTemplateResponseBody) *CreateLiveTranscodeTemplateResponse
	GetBody() *CreateLiveTranscodeTemplateResponseBody
}

type CreateLiveTranscodeTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveTranscodeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveTranscodeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveTranscodeTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveTranscodeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveTranscodeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveTranscodeTemplateResponse) GetBody() *CreateLiveTranscodeTemplateResponseBody {
	return s.Body
}

func (s *CreateLiveTranscodeTemplateResponse) SetHeaders(v map[string]*string) *CreateLiveTranscodeTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveTranscodeTemplateResponse) SetStatusCode(v int32) *CreateLiveTranscodeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveTranscodeTemplateResponse) SetBody(v *CreateLiveTranscodeTemplateResponseBody) *CreateLiveTranscodeTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateLiveTranscodeTemplateResponse) Validate() error {
	return dara.Validate(s)
}
