// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationalUnitIdByExternalIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnitId(v string) *GetOrganizationalUnitIdByExternalIdResponseBody
	GetOrganizationalUnitId() *string
}

type GetOrganizationalUnitIdByExternalIdResponseBody struct {
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
}

func (s GetOrganizationalUnitIdByExternalIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationalUnitIdByExternalIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitIdByExternalIdResponseBody) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *GetOrganizationalUnitIdByExternalIdResponseBody) SetOrganizationalUnitId(v string) *GetOrganizationalUnitIdByExternalIdResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetOrganizationalUnitIdByExternalIdResponseBody) Validate() error {
	return dara.Validate(s)
}
