// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPrimaryOrganizationalUnitHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SetUserPrimaryOrganizationalUnitHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *SetUserPrimaryOrganizationalUnitHeaders
	GetAuthorization() *string
}

type SetUserPrimaryOrganizationalUnitHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitHeaders) String() string {
	return dara.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitHeaders) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SetUserPrimaryOrganizationalUnitHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *SetUserPrimaryOrganizationalUnitHeaders) SetCommonHeaders(v map[string]*string) *SetUserPrimaryOrganizationalUnitHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitHeaders) SetAuthorization(v string) *SetUserPrimaryOrganizationalUnitHeaders {
	s.Authorization = &v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitHeaders) Validate() error {
	return dara.Validate(s)
}
