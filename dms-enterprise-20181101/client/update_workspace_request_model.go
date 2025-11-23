// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateWorkspaceRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateWorkspaceRequest
	GetDescription() *string
	SetWorkspaceId(v int64) *UpdateWorkspaceRequest
	GetWorkspaceId() *int64
	SetWorkspaceName(v string) *UpdateWorkspaceRequest
	GetWorkspaceName() *string
}

type UpdateWorkspaceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// token-xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new description of the workspace.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The new name of the workspace.
	//
	// example:
	//
	// workspace_xxx
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s UpdateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateWorkspaceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkspaceRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *UpdateWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *UpdateWorkspaceRequest) SetClientToken(v string) *UpdateWorkspaceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetDescription(v string) *UpdateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetWorkspaceId(v int64) *UpdateWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetWorkspaceName(v string) *UpdateWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *UpdateWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
