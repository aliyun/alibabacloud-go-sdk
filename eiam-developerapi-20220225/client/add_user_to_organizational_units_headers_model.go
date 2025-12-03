// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToOrganizationalUnitsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddUserToOrganizationalUnitsHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *AddUserToOrganizationalUnitsHeaders
	GetAuthorization() *string
}

type AddUserToOrganizationalUnitsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddUserToOrganizationalUnitsHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddUserToOrganizationalUnitsHeaders) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddUserToOrganizationalUnitsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddUserToOrganizationalUnitsHeaders) SetCommonHeaders(v map[string]*string) *AddUserToOrganizationalUnitsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddUserToOrganizationalUnitsHeaders) SetAuthorization(v string) *AddUserToOrganizationalUnitsHeaders {
	s.Authorization = &v
	return s
}

func (s *AddUserToOrganizationalUnitsHeaders) Validate() error {
	return dara.Validate(s)
}
