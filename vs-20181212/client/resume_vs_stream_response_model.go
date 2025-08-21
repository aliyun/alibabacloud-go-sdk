// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeVsStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeVsStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeVsStreamResponse
	GetStatusCode() *int32
	SetBody(v *ResumeVsStreamResponseBody) *ResumeVsStreamResponse
	GetBody() *ResumeVsStreamResponseBody
}

type ResumeVsStreamResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeVsStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeVsStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeVsStreamResponse) GoString() string {
	return s.String()
}

func (s *ResumeVsStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeVsStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeVsStreamResponse) GetBody() *ResumeVsStreamResponseBody {
	return s.Body
}

func (s *ResumeVsStreamResponse) SetHeaders(v map[string]*string) *ResumeVsStreamResponse {
	s.Headers = v
	return s
}

func (s *ResumeVsStreamResponse) SetStatusCode(v int32) *ResumeVsStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeVsStreamResponse) SetBody(v *ResumeVsStreamResponseBody) *ResumeVsStreamResponse {
	s.Body = v
	return s
}

func (s *ResumeVsStreamResponse) Validate() error {
	return dara.Validate(s)
}
