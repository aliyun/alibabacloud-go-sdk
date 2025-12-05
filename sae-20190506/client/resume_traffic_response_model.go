// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeTrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeTrafficResponse
	GetStatusCode() *int32
	SetBody(v *ResumeTrafficResponseBody) *ResumeTrafficResponse
	GetBody() *ResumeTrafficResponseBody
}

type ResumeTrafficResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeTrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeTrafficResponse) GoString() string {
	return s.String()
}

func (s *ResumeTrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeTrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeTrafficResponse) GetBody() *ResumeTrafficResponseBody {
	return s.Body
}

func (s *ResumeTrafficResponse) SetHeaders(v map[string]*string) *ResumeTrafficResponse {
	s.Headers = v
	return s
}

func (s *ResumeTrafficResponse) SetStatusCode(v int32) *ResumeTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeTrafficResponse) SetBody(v *ResumeTrafficResponseBody) *ResumeTrafficResponse {
	s.Body = v
	return s
}

func (s *ResumeTrafficResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
