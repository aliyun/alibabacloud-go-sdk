// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAssumeSlrRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAssumeSlrRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAssumeSlrRoleResponse
	GetStatusCode() *int32
	SetBody(v *CheckAssumeSlrRoleResponseBody) *CheckAssumeSlrRoleResponse
	GetBody() *CheckAssumeSlrRoleResponseBody
}

type CheckAssumeSlrRoleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAssumeSlrRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAssumeSlrRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAssumeSlrRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckAssumeSlrRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAssumeSlrRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAssumeSlrRoleResponse) GetBody() *CheckAssumeSlrRoleResponseBody {
	return s.Body
}

func (s *CheckAssumeSlrRoleResponse) SetHeaders(v map[string]*string) *CheckAssumeSlrRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckAssumeSlrRoleResponse) SetStatusCode(v int32) *CheckAssumeSlrRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAssumeSlrRoleResponse) SetBody(v *CheckAssumeSlrRoleResponseBody) *CheckAssumeSlrRoleResponse {
	s.Body = v
	return s
}

func (s *CheckAssumeSlrRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
