// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromOrganizationalUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RemoveUserFromOrganizationalUnitsRequest
	GetInstanceId() *string
	SetOrganizationalUnitIds(v []*string) *RemoveUserFromOrganizationalUnitsRequest
	GetOrganizationalUnitIds() []*string
	SetUserId(v string) *RemoveUserFromOrganizationalUnitsRequest
	GetUserId() *string
}

type RemoveUserFromOrganizationalUnitsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization IDs. You can remove an account from a maximum of 100 organizations.
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

func (s RemoveUserFromOrganizationalUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromOrganizationalUnitsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveUserFromOrganizationalUnitsRequest) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *RemoveUserFromOrganizationalUnitsRequest) GetUserId() *string {
	return s.UserId
}

func (s *RemoveUserFromOrganizationalUnitsRequest) SetInstanceId(v string) *RemoveUserFromOrganizationalUnitsRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *RemoveUserFromOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsRequest) SetUserId(v string) *RemoveUserFromOrganizationalUnitsRequest {
	s.UserId = &v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsRequest) Validate() error {
	return dara.Validate(s)
}
