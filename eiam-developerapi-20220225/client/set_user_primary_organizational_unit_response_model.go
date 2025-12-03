// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPrimaryOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetUserPrimaryOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetUserPrimaryOrganizationalUnitResponse
	GetStatusCode() *int32
}

type SetUserPrimaryOrganizationalUnitResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetUserPrimaryOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetUserPrimaryOrganizationalUnitResponse) SetHeaders(v map[string]*string) *SetUserPrimaryOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitResponse) SetStatusCode(v int32) *SetUserPrimaryOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitResponse) Validate() error {
	return dara.Validate(s)
}
