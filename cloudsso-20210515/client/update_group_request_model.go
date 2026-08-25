// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *UpdateGroupRequest
	GetDirectoryId() *string
	SetGroupId(v string) *UpdateGroupRequest
	GetGroupId() *string
	SetNewDescription(v string) *UpdateGroupRequest
	GetNewDescription() *string
	SetNewGroupName(v string) *UpdateGroupRequest
	GetNewGroupName() *string
}

type UpdateGroupRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// g-00jqzghi2n3o5hkh****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The new description of the group.
	//
	// example:
	//
	// This is a group.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new name of the group.
	//
	// example:
	//
	// NewTestGroup
	NewGroupName *string `json:"NewGroupName,omitempty" xml:"NewGroupName,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupRequest) GetNewDescription() *string {
	return s.NewDescription
}

func (s *UpdateGroupRequest) GetNewGroupName() *string {
	return s.NewGroupName
}

func (s *UpdateGroupRequest) SetDirectoryId(v string) *UpdateGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupId(v string) *UpdateGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupRequest) SetNewDescription(v string) *UpdateGroupRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateGroupRequest) SetNewGroupName(v string) *UpdateGroupRequest {
	s.NewGroupName = &v
	return s
}

func (s *UpdateGroupRequest) Validate() error {
	return dara.Validate(s)
}
