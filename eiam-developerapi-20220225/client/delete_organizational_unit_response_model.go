// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOrganizationalUnitResponse
	GetStatusCode() *int32
}

type DeleteOrganizationalUnitResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOrganizationalUnitResponse) SetHeaders(v map[string]*string) *DeleteOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *DeleteOrganizationalUnitResponse) SetStatusCode(v int32) *DeleteOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOrganizationalUnitResponse) Validate() error {
	return dara.Validate(s)
}
