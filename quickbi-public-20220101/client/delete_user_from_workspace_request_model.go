// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserFromWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *DeleteUserFromWorkspaceRequest
	GetUserId() *string
	SetWorkspaceId(v string) *DeleteUserFromWorkspaceRequest
	GetWorkspaceId() *string
}

type DeleteUserFromWorkspaceRequest struct {
	// The ID of the user to be deleted. Note that this UserID is for Quick BI, not the Alibaba Cloud UID.
	//
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteUserFromWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserFromWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserFromWorkspaceRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteUserFromWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteUserFromWorkspaceRequest) SetUserId(v string) *DeleteUserFromWorkspaceRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserFromWorkspaceRequest) SetWorkspaceId(v string) *DeleteUserFromWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteUserFromWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
