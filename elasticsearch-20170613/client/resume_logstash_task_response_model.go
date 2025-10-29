// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeLogstashTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeLogstashTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeLogstashTaskResponse
	GetStatusCode() *int32
	SetBody(v *ResumeLogstashTaskResponseBody) *ResumeLogstashTaskResponse
	GetBody() *ResumeLogstashTaskResponseBody
}

type ResumeLogstashTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeLogstashTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeLogstashTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeLogstashTaskResponse) GoString() string {
	return s.String()
}

func (s *ResumeLogstashTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeLogstashTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeLogstashTaskResponse) GetBody() *ResumeLogstashTaskResponseBody {
	return s.Body
}

func (s *ResumeLogstashTaskResponse) SetHeaders(v map[string]*string) *ResumeLogstashTaskResponse {
	s.Headers = v
	return s
}

func (s *ResumeLogstashTaskResponse) SetStatusCode(v int32) *ResumeLogstashTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeLogstashTaskResponse) SetBody(v *ResumeLogstashTaskResponseBody) *ResumeLogstashTaskResponse {
	s.Body = v
	return s
}

func (s *ResumeLogstashTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
