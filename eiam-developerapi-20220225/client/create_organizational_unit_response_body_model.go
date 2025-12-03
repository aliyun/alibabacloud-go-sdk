// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrganizationalUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnitId(v string) *CreateOrganizationalUnitResponseBody
	GetOrganizationalUnitId() *string
}

type CreateOrganizationalUnitResponseBody struct {
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
}

func (s CreateOrganizationalUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitResponseBody) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *CreateOrganizationalUnitResponseBody) SetOrganizationalUnitId(v string) *CreateOrganizationalUnitResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

func (s *CreateOrganizationalUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
