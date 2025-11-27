// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeRCSecurityGroupPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeRCSecurityGroupPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeRCSecurityGroupPermissionResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeRCSecurityGroupPermissionResponseBody) *AuthorizeRCSecurityGroupPermissionResponse
	GetBody() *AuthorizeRCSecurityGroupPermissionResponseBody
}

type AuthorizeRCSecurityGroupPermissionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeRCSecurityGroupPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeRCSecurityGroupPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeRCSecurityGroupPermissionResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeRCSecurityGroupPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeRCSecurityGroupPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeRCSecurityGroupPermissionResponse) GetBody() *AuthorizeRCSecurityGroupPermissionResponseBody {
	return s.Body
}

func (s *AuthorizeRCSecurityGroupPermissionResponse) SetHeaders(v map[string]*string) *AuthorizeRCSecurityGroupPermissionResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionResponse) SetStatusCode(v int32) *AuthorizeRCSecurityGroupPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionResponse) SetBody(v *AuthorizeRCSecurityGroupPermissionResponseBody) *AuthorizeRCSecurityGroupPermissionResponse {
	s.Body = v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
