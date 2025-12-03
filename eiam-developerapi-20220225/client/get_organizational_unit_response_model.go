// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrganizationalUnitResponse
	GetStatusCode() *int32
	SetBody(v *GetOrganizationalUnitResponseBody) *GetOrganizationalUnitResponse
	GetBody() *GetOrganizationalUnitResponseBody
}

type GetOrganizationalUnitResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrganizationalUnitResponse) GetBody() *GetOrganizationalUnitResponseBody {
	return s.Body
}

func (s *GetOrganizationalUnitResponse) SetHeaders(v map[string]*string) *GetOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationalUnitResponse) SetStatusCode(v int32) *GetOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationalUnitResponse) SetBody(v *GetOrganizationalUnitResponseBody) *GetOrganizationalUnitResponse {
	s.Body = v
	return s
}

func (s *GetOrganizationalUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
