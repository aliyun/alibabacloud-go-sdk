// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationalUnitHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOrganizationalUnitHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GetOrganizationalUnitHeaders
	GetAuthorization() *string
}

type GetOrganizationalUnitHeaders struct {
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

func (s GetOrganizationalUnitHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOrganizationalUnitHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizationalUnitHeaders) SetAuthorization(v string) *GetOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

func (s *GetOrganizationalUnitHeaders) Validate() error {
	return dara.Validate(s)
}
