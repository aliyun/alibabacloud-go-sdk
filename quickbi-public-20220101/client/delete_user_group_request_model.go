// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserGroupId(v string) *DeleteUserGroupRequest
	GetUserGroupId() *string
}

type DeleteUserGroupRequest struct {
	// The ID of the user group.
	//
	// This parameter is required.
	//
	// example:
	//
	// f5eeb52e-d9c2-4a8b-80e3-47ab55c2****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DeleteUserGroupRequest) SetUserGroupId(v string) *DeleteUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *DeleteUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
