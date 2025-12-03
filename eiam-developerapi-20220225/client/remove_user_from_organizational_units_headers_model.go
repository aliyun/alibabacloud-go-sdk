// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromOrganizationalUnitsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RemoveUserFromOrganizationalUnitsHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *RemoveUserFromOrganizationalUnitsHeaders
	GetAuthorization() *string
}

type RemoveUserFromOrganizationalUnitsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s RemoveUserFromOrganizationalUnitsHeaders) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromOrganizationalUnitsHeaders) GoString() string {
	return s.String()
}

func (s *RemoveUserFromOrganizationalUnitsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RemoveUserFromOrganizationalUnitsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *RemoveUserFromOrganizationalUnitsHeaders) SetCommonHeaders(v map[string]*string) *RemoveUserFromOrganizationalUnitsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsHeaders) SetAuthorization(v string) *RemoveUserFromOrganizationalUnitsHeaders {
	s.Authorization = &v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsHeaders) Validate() error {
	return dara.Validate(s)
}
