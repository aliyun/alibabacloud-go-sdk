// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUsersToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *AddUsersToGroupRequest
	GetGroupId() *string
	SetInstanceId(v string) *AddUsersToGroupRequest
	GetInstanceId() *string
	SetUserIds(v []*string) *AddUsersToGroupRequest
	GetUserIds() []*string
}

type AddUsersToGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ou_001]
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s AddUsersToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUsersToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddUsersToGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddUsersToGroupRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *AddUsersToGroupRequest) SetGroupId(v string) *AddUsersToGroupRequest {
	s.GroupId = &v
	return s
}

func (s *AddUsersToGroupRequest) SetInstanceId(v string) *AddUsersToGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUsersToGroupRequest) SetUserIds(v []*string) *AddUsersToGroupRequest {
	s.UserIds = v
	return s
}

func (s *AddUsersToGroupRequest) Validate() error {
	return dara.Validate(s)
}
