// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeLocationServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeLocationServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeLocationServiceResponse
	GetStatusCode() *int32
	SetBody(v *ResumeLocationServiceResponseBody) *ResumeLocationServiceResponse
	GetBody() *ResumeLocationServiceResponseBody
}

type ResumeLocationServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeLocationServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeLocationServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeLocationServiceResponse) GoString() string {
	return s.String()
}

func (s *ResumeLocationServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeLocationServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeLocationServiceResponse) GetBody() *ResumeLocationServiceResponseBody {
	return s.Body
}

func (s *ResumeLocationServiceResponse) SetHeaders(v map[string]*string) *ResumeLocationServiceResponse {
	s.Headers = v
	return s
}

func (s *ResumeLocationServiceResponse) SetStatusCode(v int32) *ResumeLocationServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeLocationServiceResponse) SetBody(v *ResumeLocationServiceResponseBody) *ResumeLocationServiceResponse {
	s.Body = v
	return s
}

func (s *ResumeLocationServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
