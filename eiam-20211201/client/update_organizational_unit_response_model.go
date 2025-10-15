// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOrganizationalUnitResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOrganizationalUnitResponseBody) *UpdateOrganizationalUnitResponse
	GetBody() *UpdateOrganizationalUnitResponseBody
}

type UpdateOrganizationalUnitResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOrganizationalUnitResponse) GetBody() *UpdateOrganizationalUnitResponseBody {
	return s.Body
}

func (s *UpdateOrganizationalUnitResponse) SetHeaders(v map[string]*string) *UpdateOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationalUnitResponse) SetStatusCode(v int32) *UpdateOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOrganizationalUnitResponse) SetBody(v *UpdateOrganizationalUnitResponseBody) *UpdateOrganizationalUnitResponse {
	s.Body = v
	return s
}

func (s *UpdateOrganizationalUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
