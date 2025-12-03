// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToOrganizationalUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnitIds(v []*string) *AddUserToOrganizationalUnitsRequest
	GetOrganizationalUnitIds() []*string
}

type AddUserToOrganizationalUnitsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [ou_wovwffm62xifdziem7an7xxxxx]
	OrganizationalUnitIds []*string `json:"organizationalUnitIds,omitempty" xml:"organizationalUnitIds,omitempty" type:"Repeated"`
}

func (s AddUserToOrganizationalUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserToOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsRequest) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *AddUserToOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *AddUserToOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *AddUserToOrganizationalUnitsRequest) Validate() error {
	return dara.Validate(s)
}
