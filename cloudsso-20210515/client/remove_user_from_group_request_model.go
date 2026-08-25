// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *RemoveUserFromGroupRequest
	GetDirectoryId() *string
	SetGroupId(v string) *RemoveUserFromGroupRequest
	GetGroupId() *string
	SetUserId(v string) *RemoveUserFromGroupRequest
	GetUserId() *string
}

type RemoveUserFromGroupRequest struct {
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

func (s RemoveUserFromGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *RemoveUserFromGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *RemoveUserFromGroupRequest) GetUserId() *string {
	return s.UserId
}

func (s *RemoveUserFromGroupRequest) SetDirectoryId(v string) *RemoveUserFromGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *RemoveUserFromGroupRequest) SetGroupId(v string) *RemoveUserFromGroupRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveUserFromGroupRequest) SetUserId(v string) *RemoveUserFromGroupRequest {
	s.UserId = &v
	return s
}

func (s *RemoveUserFromGroupRequest) Validate() error {
	return dara.Validate(s)
}
