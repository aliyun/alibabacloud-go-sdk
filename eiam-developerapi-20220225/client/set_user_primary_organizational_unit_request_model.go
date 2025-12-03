// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPrimaryOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationalUnitId(v string) *SetUserPrimaryOrganizationalUnitRequest
	GetOrganizationalUnitId() *string
}

type SetUserPrimaryOrganizationalUnitRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"organizationalUnitId,omitempty" xml:"organizationalUnitId,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *SetUserPrimaryOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *SetUserPrimaryOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
