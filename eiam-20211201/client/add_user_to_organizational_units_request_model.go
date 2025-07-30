// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToOrganizationalUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddUserToOrganizationalUnitsRequest
	GetInstanceId() *string
	SetOrganizationalUnitIds(v []*string) *AddUserToOrganizationalUnitsRequest
	GetOrganizationalUnitIds() []*string
	SetUserId(v string) *AddUserToOrganizationalUnitsRequest
	GetUserId() *string
}

type AddUserToOrganizationalUnitsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization IDs. You can add an account to a maximum of 100 organizations.
	//
	// This parameter is required.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUserToOrganizationalUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserToOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddUserToOrganizationalUnitsRequest) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *AddUserToOrganizationalUnitsRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddUserToOrganizationalUnitsRequest) SetInstanceId(v string) *AddUserToOrganizationalUnitsRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUserToOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *AddUserToOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *AddUserToOrganizationalUnitsRequest) SetUserId(v string) *AddUserToOrganizationalUnitsRequest {
	s.UserId = &v
	return s
}

func (s *AddUserToOrganizationalUnitsRequest) Validate() error {
	return dara.Validate(s)
}
