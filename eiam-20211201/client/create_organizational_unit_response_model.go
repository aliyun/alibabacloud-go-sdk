// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrganizationalUnitResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrganizationalUnitResponseBody) *CreateOrganizationalUnitResponse
	GetBody() *CreateOrganizationalUnitResponseBody
}

type CreateOrganizationalUnitResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrganizationalUnitResponse) GetBody() *CreateOrganizationalUnitResponseBody {
	return s.Body
}

func (s *CreateOrganizationalUnitResponse) SetHeaders(v map[string]*string) *CreateOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *CreateOrganizationalUnitResponse) SetStatusCode(v int32) *CreateOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrganizationalUnitResponse) SetBody(v *CreateOrganizationalUnitResponseBody) *CreateOrganizationalUnitResponse {
	s.Body = v
	return s
}

func (s *CreateOrganizationalUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
