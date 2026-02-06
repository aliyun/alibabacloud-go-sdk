// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeJobsResponse
	GetStatusCode() *int32
	SetBody(v *ResumeJobsResponseBody) *ResumeJobsResponse
	GetBody() *ResumeJobsResponseBody
}

type ResumeJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeJobsResponse) GoString() string {
	return s.String()
}

func (s *ResumeJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeJobsResponse) GetBody() *ResumeJobsResponseBody {
	return s.Body
}

func (s *ResumeJobsResponse) SetHeaders(v map[string]*string) *ResumeJobsResponse {
	s.Headers = v
	return s
}

func (s *ResumeJobsResponse) SetStatusCode(v int32) *ResumeJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeJobsResponse) SetBody(v *ResumeJobsResponseBody) *ResumeJobsResponse {
	s.Body = v
	return s
}

func (s *ResumeJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
