// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ResumeInstanceResponseBody) *ResumeInstanceResponse
	GetBody() *ResumeInstanceResponseBody
}

type ResumeInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeInstanceResponse) GetBody() *ResumeInstanceResponseBody {
	return s.Body
}

func (s *ResumeInstanceResponse) SetHeaders(v map[string]*string) *ResumeInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResumeInstanceResponse) SetStatusCode(v int32) *ResumeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeInstanceResponse) SetBody(v *ResumeInstanceResponseBody) *ResumeInstanceResponse {
	s.Body = v
	return s
}

func (s *ResumeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
