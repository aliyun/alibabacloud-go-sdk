// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromOrganizationalUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnitIds(v []*string) *RemoveUserFromOrganizationalUnitsRequest
	GetOrganizationalUnitIds() []*string
}

type RemoveUserFromOrganizationalUnitsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [ou_wovwffm62xifdziem7an7xxxxx]
	OrganizationalUnitIds []*string `json:"organizationalUnitIds,omitempty" xml:"organizationalUnitIds,omitempty" type:"Repeated"`
}

func (s RemoveUserFromOrganizationalUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromOrganizationalUnitsRequest) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *RemoveUserFromOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *RemoveUserFromOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsRequest) Validate() error {
	return dara.Validate(s)
}
