// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteOrganizationalUnitRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *DeleteOrganizationalUnitRequest
	GetOrganizationalUnitId() *string
}

type DeleteOrganizationalUnitRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
}

func (s DeleteOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteOrganizationalUnitRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *DeleteOrganizationalUnitRequest) SetInstanceId(v string) *DeleteOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *DeleteOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *DeleteOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
