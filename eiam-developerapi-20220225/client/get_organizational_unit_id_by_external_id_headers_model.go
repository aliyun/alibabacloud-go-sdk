// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationalUnitIdByExternalIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOrganizationalUnitIdByExternalIdHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GetOrganizationalUnitIdByExternalIdHeaders
	GetAuthorization() *string
}

type GetOrganizationalUnitIdByExternalIdHeaders struct {
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

func (s GetOrganizationalUnitIdByExternalIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdHeaders) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOrganizationalUnitIdByExternalIdHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetOrganizationalUnitIdByExternalIdHeaders) SetCommonHeaders(v map[string]*string) *GetOrganizationalUnitIdByExternalIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdHeaders) SetAuthorization(v string) *GetOrganizationalUnitIdByExternalIdHeaders {
	s.Authorization = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdHeaders) Validate() error {
	return dara.Validate(s)
}
