// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchOrganizationalUnitParentIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PatchOrganizationalUnitParentIdHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *PatchOrganizationalUnitParentIdHeaders
	GetAuthorization() *string
}

type PatchOrganizationalUnitParentIdHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer access_token.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PatchOrganizationalUnitParentIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s PatchOrganizationalUnitParentIdHeaders) GoString() string {
	return s.String()
}

func (s *PatchOrganizationalUnitParentIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PatchOrganizationalUnitParentIdHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PatchOrganizationalUnitParentIdHeaders) SetCommonHeaders(v map[string]*string) *PatchOrganizationalUnitParentIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchOrganizationalUnitParentIdHeaders) SetAuthorization(v string) *PatchOrganizationalUnitParentIdHeaders {
	s.Authorization = &v
	return s
}

func (s *PatchOrganizationalUnitParentIdHeaders) Validate() error {
	return dara.Validate(s)
}
