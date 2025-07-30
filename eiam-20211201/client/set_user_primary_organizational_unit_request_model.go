// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPrimaryOrganizationalUnitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetUserPrimaryOrganizationalUnitRequest
	GetInstanceId() *string
	SetOrganizationalUnitId(v string) *SetUserPrimaryOrganizationalUnitRequest
	GetOrganizationalUnitId() *string
	SetUserId(v string) *SetUserPrimaryOrganizationalUnitRequest
	GetUserId() *string
}

type SetUserPrimaryOrganizationalUnitRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the new primary organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The ID of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitRequest) String() string {
	return dara.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetUserPrimaryOrganizationalUnitRequest) GetOrganizationalUnitId() *string {
	return s.OrganizationalUnitId
}

func (s *SetUserPrimaryOrganizationalUnitRequest) GetUserId() *string {
	return s.UserId
}

func (s *SetUserPrimaryOrganizationalUnitRequest) SetInstanceId(v string) *SetUserPrimaryOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *SetUserPrimaryOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitRequest) SetUserId(v string) *SetUserPrimaryOrganizationalUnitRequest {
	s.UserId = &v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitRequest) Validate() error {
	return dara.Validate(s)
}
