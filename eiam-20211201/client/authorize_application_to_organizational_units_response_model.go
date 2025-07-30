// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeApplicationToOrganizationalUnitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeApplicationToOrganizationalUnitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeApplicationToOrganizationalUnitsResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeApplicationToOrganizationalUnitsResponseBody) *AuthorizeApplicationToOrganizationalUnitsResponse
	GetBody() *AuthorizeApplicationToOrganizationalUnitsResponseBody
}

type AuthorizeApplicationToOrganizationalUnitsResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeApplicationToOrganizationalUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeApplicationToOrganizationalUnitsResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeApplicationToOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponse) GetBody() *AuthorizeApplicationToOrganizationalUnitsResponseBody {
	return s.Body
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *AuthorizeApplicationToOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponse) SetStatusCode(v int32) *AuthorizeApplicationToOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponse) SetBody(v *AuthorizeApplicationToOrganizationalUnitsResponseBody) *AuthorizeApplicationToOrganizationalUnitsResponse {
	s.Body = v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponse) Validate() error {
	return dara.Validate(s)
}
