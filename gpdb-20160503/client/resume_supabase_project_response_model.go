// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSupabaseProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeSupabaseProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeSupabaseProjectResponse
	GetStatusCode() *int32
	SetBody(v *ResumeSupabaseProjectResponseBody) *ResumeSupabaseProjectResponse
	GetBody() *ResumeSupabaseProjectResponseBody
}

type ResumeSupabaseProjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeSupabaseProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeSupabaseProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeSupabaseProjectResponse) GoString() string {
	return s.String()
}

func (s *ResumeSupabaseProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeSupabaseProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeSupabaseProjectResponse) GetBody() *ResumeSupabaseProjectResponseBody {
	return s.Body
}

func (s *ResumeSupabaseProjectResponse) SetHeaders(v map[string]*string) *ResumeSupabaseProjectResponse {
	s.Headers = v
	return s
}

func (s *ResumeSupabaseProjectResponse) SetStatusCode(v int32) *ResumeSupabaseProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeSupabaseProjectResponse) SetBody(v *ResumeSupabaseProjectResponseBody) *ResumeSupabaseProjectResponse {
	s.Body = v
	return s
}

func (s *ResumeSupabaseProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
