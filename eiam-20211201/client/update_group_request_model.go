// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupExternalId(v string) *UpdateGroupRequest
	GetGroupExternalId() *string
	SetGroupId(v string) *UpdateGroupRequest
	GetGroupId() *string
	SetGroupName(v string) *UpdateGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *UpdateGroupRequest
	GetInstanceId() *string
}

type UpdateGroupRequest struct {
	// The external ID of the group.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupExternalId *string `json:"GroupExternalId,omitempty" xml:"GroupExternalId,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// name_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) GetGroupExternalId() *string {
	return s.GroupExternalId
}

func (s *UpdateGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateGroupRequest) SetGroupExternalId(v string) *UpdateGroupRequest {
	s.GroupExternalId = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupId(v string) *UpdateGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupName(v string) *UpdateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupRequest) SetInstanceId(v string) *UpdateGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateGroupRequest) Validate() error {
	return dara.Validate(s)
}
