// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersFromGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *RemoveUsersFromGroupRequest
	GetGroupId() *string
	SetInstanceId(v string) *RemoveUsersFromGroupRequest
	GetInstanceId() *string
	SetUserIds(v []*string) *RemoveUsersFromGroupRequest
	GetUserIds() []*string
}

type RemoveUsersFromGroupRequest struct {
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
	// The account IDs. A maximum of 100 accounts can be removed from a group.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ou_001]
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s RemoveUsersFromGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveUsersFromGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveUsersFromGroupRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *RemoveUsersFromGroupRequest) SetGroupId(v string) *RemoveUsersFromGroupRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetInstanceId(v string) *RemoveUsersFromGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetUserIds(v []*string) *RemoveUsersFromGroupRequest {
	s.UserIds = v
	return s
}

func (s *RemoveUsersFromGroupRequest) Validate() error {
	return dara.Validate(s)
}
