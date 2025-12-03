// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOrganizationalUnitHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteOrganizationalUnitHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *DeleteOrganizationalUnitHeaders
	GetAuthorization() *string
}

type DeleteOrganizationalUnitHeaders struct {
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

func (s DeleteOrganizationalUnitHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteOrganizationalUnitHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *DeleteOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteOrganizationalUnitHeaders) SetAuthorization(v string) *DeleteOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteOrganizationalUnitHeaders) Validate() error {
	return dara.Validate(s)
}
