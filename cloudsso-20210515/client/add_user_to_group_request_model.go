// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *AddUserToGroupRequest
	GetDirectoryId() *string
	SetGroupId(v string) *AddUserToGroupRequest
	GetGroupId() *string
	SetUserId(v string) *AddUserToGroupRequest
	GetUserId() *string
}

type AddUserToGroupRequest struct {
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
	// The ID of the user.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUserToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToGroupRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *AddUserToGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddUserToGroupRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddUserToGroupRequest) SetDirectoryId(v string) *AddUserToGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *AddUserToGroupRequest) SetGroupId(v string) *AddUserToGroupRequest {
	s.GroupId = &v
	return s
}

func (s *AddUserToGroupRequest) SetUserId(v string) *AddUserToGroupRequest {
	s.UserId = &v
	return s
}

func (s *AddUserToGroupRequest) Validate() error {
	return dara.Validate(s)
}
