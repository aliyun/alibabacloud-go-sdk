// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeSessionResponse
	GetStatusCode() *int32
	SetBody(v *Session) *ResumeSessionResponse
	GetBody() *Session
}

type ResumeSessionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Session           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeSessionResponse) GoString() string {
	return s.String()
}

func (s *ResumeSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeSessionResponse) GetBody() *Session {
	return s.Body
}

func (s *ResumeSessionResponse) SetHeaders(v map[string]*string) *ResumeSessionResponse {
	s.Headers = v
	return s
}

func (s *ResumeSessionResponse) SetStatusCode(v int32) *ResumeSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeSessionResponse) SetBody(v *Session) *ResumeSessionResponse {
	s.Body = v
	return s
}

func (s *ResumeSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
