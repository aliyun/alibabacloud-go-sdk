// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateGroupRequest
	GetDescription() *string
	SetGroupId(v string) *UpdateGroupRequest
	GetGroupId() *string
	SetGroupName(v string) *UpdateGroupRequest
	GetGroupName() *string
}

type UpdateGroupRequest struct {
	// The description of the group after modification.
	//
	// example:
	//
	// test group description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the group that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e43ec8427dd45f19431b7504649a1b4
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The name of the group after modification.
	//
	// example:
	//
	// test group
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGroupRequest) SetDescription(v string) *UpdateGroupRequest {
	s.Description = &v
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

func (s *UpdateGroupRequest) Validate() error {
	return dara.Validate(s)
}
