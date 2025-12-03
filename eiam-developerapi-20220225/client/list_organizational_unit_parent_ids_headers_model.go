// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitParentIdsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListOrganizationalUnitParentIdsHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ListOrganizationalUnitParentIdsHeaders
	GetAuthorization() *string
}

type ListOrganizationalUnitParentIdsHeaders struct {
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

func (s ListOrganizationalUnitParentIdsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitParentIdsHeaders) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentIdsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListOrganizationalUnitParentIdsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListOrganizationalUnitParentIdsHeaders) SetCommonHeaders(v map[string]*string) *ListOrganizationalUnitParentIdsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListOrganizationalUnitParentIdsHeaders) SetAuthorization(v string) *ListOrganizationalUnitParentIdsHeaders {
	s.Authorization = &v
	return s
}

func (s *ListOrganizationalUnitParentIdsHeaders) Validate() error {
	return dara.Validate(s)
}
