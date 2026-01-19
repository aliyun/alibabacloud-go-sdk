// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceServerScopesToOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeResourceServerScopesToOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeResourceServerScopesToOrganizationalUnitResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeResourceServerScopesToOrganizationalUnitResponseBody) *AuthorizeResourceServerScopesToOrganizationalUnitResponse
	GetBody() *AuthorizeResourceServerScopesToOrganizationalUnitResponseBody
}

type AuthorizeResourceServerScopesToOrganizationalUnitResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeResourceServerScopesToOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeResourceServerScopesToOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceServerScopesToOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitResponse) GetBody() *AuthorizeResourceServerScopesToOrganizationalUnitResponseBody {
	return s.Body
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitResponse) SetHeaders(v map[string]*string) *AuthorizeResourceServerScopesToOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitResponse) SetStatusCode(v int32) *AuthorizeResourceServerScopesToOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitResponse) SetBody(v *AuthorizeResourceServerScopesToOrganizationalUnitResponseBody) *AuthorizeResourceServerScopesToOrganizationalUnitResponse {
	s.Body = v
	return s
}

func (s *AuthorizeResourceServerScopesToOrganizationalUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
