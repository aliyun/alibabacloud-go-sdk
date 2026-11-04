// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFaceUserFromUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsolationId(v string) *RemoveFaceUserFromUserGroupRequest
	GetIsolationId() *string
	SetUserGroupId(v string) *RemoveFaceUserFromUserGroupRequest
	GetUserGroupId() *string
	SetUserId(v string) *RemoveFaceUserFromUserGroupRequest
	GetUserId() *string
}

type RemoveFaceUserFromUserGroupRequest struct {
	// The isolation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestIsolationId
	IsolationId *string `json:"IsolationId,omitempty" xml:"IsolationId,omitempty"`
	// The user group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestUserGroupId
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveFaceUserFromUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveFaceUserFromUserGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveFaceUserFromUserGroupRequest) GetIsolationId() *string {
	return s.IsolationId
}

func (s *RemoveFaceUserFromUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *RemoveFaceUserFromUserGroupRequest) GetUserId() *string {
	return s.UserId
}

func (s *RemoveFaceUserFromUserGroupRequest) SetIsolationId(v string) *RemoveFaceUserFromUserGroupRequest {
	s.IsolationId = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupRequest) SetUserGroupId(v string) *RemoveFaceUserFromUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupRequest) SetUserId(v string) *RemoveFaceUserFromUserGroupRequest {
	s.UserId = &v
	return s
}

func (s *RemoveFaceUserFromUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
