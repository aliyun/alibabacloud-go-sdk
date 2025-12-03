// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrganizationalUnitHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateOrganizationalUnitHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *CreateOrganizationalUnitHeaders
	GetAuthorization() *string
}

type CreateOrganizationalUnitHeaders struct {
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

func (s CreateOrganizationalUnitHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateOrganizationalUnitHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *CreateOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrganizationalUnitHeaders) SetAuthorization(v string) *CreateOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

func (s *CreateOrganizationalUnitHeaders) Validate() error {
	return dara.Validate(s)
}
