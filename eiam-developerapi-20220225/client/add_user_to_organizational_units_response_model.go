// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToOrganizationalUnitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserToOrganizationalUnitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserToOrganizationalUnitsResponse
	GetStatusCode() *int32
}

type AddUserToOrganizationalUnitsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AddUserToOrganizationalUnitsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserToOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserToOrganizationalUnitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserToOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *AddUserToOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *AddUserToOrganizationalUnitsResponse) SetStatusCode(v int32) *AddUserToOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToOrganizationalUnitsResponse) Validate() error {
	return dara.Validate(s)
}
