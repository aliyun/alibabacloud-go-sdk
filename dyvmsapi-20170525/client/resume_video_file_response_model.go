// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeVideoFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeVideoFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeVideoFileResponse
	GetStatusCode() *int32
	SetBody(v *ResumeVideoFileResponseBody) *ResumeVideoFileResponse
	GetBody() *ResumeVideoFileResponseBody
}

type ResumeVideoFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeVideoFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeVideoFileResponse) GoString() string {
	return s.String()
}

func (s *ResumeVideoFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeVideoFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeVideoFileResponse) GetBody() *ResumeVideoFileResponseBody {
	return s.Body
}

func (s *ResumeVideoFileResponse) SetHeaders(v map[string]*string) *ResumeVideoFileResponse {
	s.Headers = v
	return s
}

func (s *ResumeVideoFileResponse) SetStatusCode(v int32) *ResumeVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeVideoFileResponse) SetBody(v *ResumeVideoFileResponseBody) *ResumeVideoFileResponse {
	s.Body = v
	return s
}

func (s *ResumeVideoFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
