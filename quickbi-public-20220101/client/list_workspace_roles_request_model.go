// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspaceRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *ListWorkspaceRolesRequest
	GetWorkspaceId() *string
}

type ListWorkspaceRolesRequest struct {
	// Workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWorkspaceRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspaceRolesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceRolesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspaceRolesRequest) SetWorkspaceId(v string) *ListWorkspaceRolesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspaceRolesRequest) Validate() error {
	return dara.Validate(s)
}
