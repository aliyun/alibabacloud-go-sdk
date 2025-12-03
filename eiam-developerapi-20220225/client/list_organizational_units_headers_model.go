// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListOrganizationalUnitsHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ListOrganizationalUnitsHeaders
	GetAuthorization() *string
}

type ListOrganizationalUnitsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListOrganizationalUnitsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsHeaders) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListOrganizationalUnitsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListOrganizationalUnitsHeaders) SetCommonHeaders(v map[string]*string) *ListOrganizationalUnitsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOrganizationalUnitsHeaders) SetAuthorization(v string) *ListOrganizationalUnitsHeaders {
	s.Authorization = &v
	return s
}

func (s *ListOrganizationalUnitsHeaders) Validate() error {
	return dara.Validate(s)
}
