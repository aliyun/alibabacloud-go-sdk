// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserRoleInfoInWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserId(v string) *QueryUserRoleInfoInWorkspaceRequest
	GetUserId() *string
	SetWorkspaceId(v string) *QueryUserRoleInfoInWorkspaceRequest
	GetWorkspaceId() *string
}

type QueryUserRoleInfoInWorkspaceRequest struct {
	// Quick BI user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f5698bedeb384b1986afccd9e434****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryUserRoleInfoInWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserRoleInfoInWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *QueryUserRoleInfoInWorkspaceRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserRoleInfoInWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryUserRoleInfoInWorkspaceRequest) SetUserId(v string) *QueryUserRoleInfoInWorkspaceRequest {
	s.UserId = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceRequest) SetWorkspaceId(v string) *QueryUserRoleInfoInWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryUserRoleInfoInWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
