// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *UpdateContextDatabaseWorkspaceRequest
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *UpdateContextDatabaseWorkspaceRequest
	GetWorkspaceName() *string
}

type UpdateContextDatabaseWorkspaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s UpdateContextDatabaseWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *UpdateContextDatabaseWorkspaceRequest) SetWorkspaceId(v string) *UpdateContextDatabaseWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceRequest) SetWorkspaceName(v string) *UpdateContextDatabaseWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
