// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeProcessesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeProcessesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeProcessesResponse
	GetStatusCode() *int32
	SetBody(v *ResumeProcessesResponseBody) *ResumeProcessesResponse
	GetBody() *ResumeProcessesResponseBody
}

type ResumeProcessesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeProcessesResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeProcessesResponse) GoString() string {
	return s.String()
}

func (s *ResumeProcessesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeProcessesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeProcessesResponse) GetBody() *ResumeProcessesResponseBody {
	return s.Body
}

func (s *ResumeProcessesResponse) SetHeaders(v map[string]*string) *ResumeProcessesResponse {
	s.Headers = v
	return s
}

func (s *ResumeProcessesResponse) SetStatusCode(v int32) *ResumeProcessesResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeProcessesResponse) SetBody(v *ResumeProcessesResponseBody) *ResumeProcessesResponse {
	s.Body = v
	return s
}

func (s *ResumeProcessesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
