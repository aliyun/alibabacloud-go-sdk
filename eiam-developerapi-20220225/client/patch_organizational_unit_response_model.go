// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PatchOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PatchOrganizationalUnitResponse
	GetStatusCode() *int32
}

type PatchOrganizationalUnitResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PatchOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s PatchOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PatchOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PatchOrganizationalUnitResponse) SetHeaders(v map[string]*string) *PatchOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *PatchOrganizationalUnitResponse) SetStatusCode(v int32) *PatchOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *PatchOrganizationalUnitResponse) Validate() error {
	return dara.Validate(s)
}
