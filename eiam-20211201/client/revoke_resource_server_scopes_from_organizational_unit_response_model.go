// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourceServerScopesFromOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeResourceServerScopesFromOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeResourceServerScopesFromOrganizationalUnitResponse
	GetStatusCode() *int32
	SetBody(v *RevokeResourceServerScopesFromOrganizationalUnitResponseBody) *RevokeResourceServerScopesFromOrganizationalUnitResponse
	GetBody() *RevokeResourceServerScopesFromOrganizationalUnitResponseBody
}

type RevokeResourceServerScopesFromOrganizationalUnitResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeResourceServerScopesFromOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeResourceServerScopesFromOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourceServerScopesFromOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitResponse) GetBody() *RevokeResourceServerScopesFromOrganizationalUnitResponseBody {
	return s.Body
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitResponse) SetHeaders(v map[string]*string) *RevokeResourceServerScopesFromOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitResponse) SetStatusCode(v int32) *RevokeResourceServerScopesFromOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitResponse) SetBody(v *RevokeResourceServerScopesFromOrganizationalUnitResponseBody) *RevokeResourceServerScopesFromOrganizationalUnitResponse {
	s.Body = v
	return s
}

func (s *RevokeResourceServerScopesFromOrganizationalUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
