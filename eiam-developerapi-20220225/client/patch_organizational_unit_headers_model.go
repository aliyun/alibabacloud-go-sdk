// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchOrganizationalUnitHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PatchOrganizationalUnitHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *PatchOrganizationalUnitHeaders
	GetAuthorization() *string
}

type PatchOrganizationalUnitHeaders struct {
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

func (s PatchOrganizationalUnitHeaders) String() string {
	return dara.Prettify(s)
}

func (s PatchOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PatchOrganizationalUnitHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PatchOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *PatchOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchOrganizationalUnitHeaders) SetAuthorization(v string) *PatchOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

func (s *PatchOrganizationalUnitHeaders) Validate() error {
	return dara.Validate(s)
}
