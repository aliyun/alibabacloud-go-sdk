// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupRequest
	GetDescription() *string
	SetGroupName(v string) *CreateGroupRequest
	GetGroupName() *string
	SetIsRoot(v bool) *CreateGroupRequest
	GetIsRoot() *bool
	SetParentGroupId(v string) *CreateGroupRequest
	GetParentGroupId() *string
}

type CreateGroupRequest struct {
	// The description of the group. The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// test group description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the group. The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test group
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// Specifies whether the group is a root group. A root group cannot be added to any other group. In most cases, a root group is the top-level organization in the organizational structure.
	//
	// example:
	//
	// false
	IsRoot *bool `json:"is_root,omitempty" xml:"is_root,omitempty"`
	// The ID of the parent group to which the group is added. If this parameter is specified, the system automatically adds the group to the specified parent group after the group is created.
	//
	// example:
	//
	// 2e43ec8427dd45f19431b7504649a1b3
	ParentGroupId *string `json:"parent_group_id,omitempty" xml:"parent_group_id,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGroupRequest) GetIsRoot() *bool {
	return s.IsRoot
}

func (s *CreateGroupRequest) GetParentGroupId() *string {
	return s.ParentGroupId
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) SetIsRoot(v bool) *CreateGroupRequest {
	s.IsRoot = &v
	return s
}

func (s *CreateGroupRequest) SetParentGroupId(v string) *CreateGroupRequest {
	s.ParentGroupId = &v
	return s
}

func (s *CreateGroupRequest) Validate() error {
	return dara.Validate(s)
}
