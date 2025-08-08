// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeProjectResponse
	GetStatusCode() *int32
	SetBody(v *ResumeProjectResponseBody) *ResumeProjectResponse
	GetBody() *ResumeProjectResponseBody
}

type ResumeProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeProjectResponse) GoString() string {
	return s.String()
}

func (s *ResumeProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeProjectResponse) GetBody() *ResumeProjectResponseBody {
	return s.Body
}

func (s *ResumeProjectResponse) SetHeaders(v map[string]*string) *ResumeProjectResponse {
	s.Headers = v
	return s
}

func (s *ResumeProjectResponse) SetStatusCode(v int32) *ResumeProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeProjectResponse) SetBody(v *ResumeProjectResponseBody) *ResumeProjectResponse {
	s.Body = v
	return s
}

func (s *ResumeProjectResponse) Validate() error {
	return dara.Validate(s)
}
