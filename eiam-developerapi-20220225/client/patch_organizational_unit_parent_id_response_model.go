// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchOrganizationalUnitParentIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PatchOrganizationalUnitParentIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PatchOrganizationalUnitParentIdResponse
	GetStatusCode() *int32
}

type PatchOrganizationalUnitParentIdResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s PatchOrganizationalUnitParentIdResponse) String() string {
	return dara.Prettify(s)
}

func (s PatchOrganizationalUnitParentIdResponse) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitParentIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PatchOrganizationalUnitParentIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PatchOrganizationalUnitParentIdResponse) SetHeaders(v map[string]*string) *PatchOrganizationalUnitParentIdResponse {
	s.Headers = v
	return s
}

func (s *PatchOrganizationalUnitParentIdResponse) SetStatusCode(v int32) *PatchOrganizationalUnitParentIdResponse {
	s.StatusCode = &v
	return s
}

func (s *PatchOrganizationalUnitParentIdResponse) Validate() error {
	return dara.Validate(s)
}
