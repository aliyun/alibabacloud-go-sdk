// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOrganizationalUnitChildrenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteOrganizationalUnitChildrenRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *DeleteOrganizationalUnitChildrenRequest
	GetOrganizationalUnitId() *string
}

type DeleteOrganizationalUnitChildrenRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Organizational Unit ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
}

func (s DeleteOrganizationalUnitChildrenRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOrganizationalUnitChildrenRequest) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitChildrenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteOrganizationalUnitChildrenRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *DeleteOrganizationalUnitChildrenRequest) SetInstanceId(v string) *DeleteOrganizationalUnitChildrenRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteOrganizationalUnitChildrenRequest) SetOrganizationalUnitId(v string) *DeleteOrganizationalUnitChildrenRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *DeleteOrganizationalUnitChildrenRequest) Validate() error {
	return dara.Validate(s)
}
