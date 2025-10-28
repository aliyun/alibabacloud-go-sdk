// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeRoleResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeRoleResponseBody) *AuthorizeRoleResponse
	GetBody() *AuthorizeRoleResponseBody
}

type AuthorizeRoleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeRoleResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeRoleResponse) GetBody() *AuthorizeRoleResponseBody {
	return s.Body
}

func (s *AuthorizeRoleResponse) SetHeaders(v map[string]*string) *AuthorizeRoleResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeRoleResponse) SetStatusCode(v int32) *AuthorizeRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeRoleResponse) SetBody(v *AuthorizeRoleResponseBody) *AuthorizeRoleResponse {
	s.Body = v
	return s
}

func (s *AuthorizeRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
