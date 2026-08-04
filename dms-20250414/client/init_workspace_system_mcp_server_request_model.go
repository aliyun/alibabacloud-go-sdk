// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitWorkspaceSystemMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWorkspaceId(v string) *InitWorkspaceSystemMcpServerRequest
	GetWorkspaceId() *string
}

type InitWorkspaceSystemMcpServerRequest struct {
	// The ID of the Data Agent workspace for which to initialize system MCP services. The caller must have at least MEMBER permissions on the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s InitWorkspaceSystemMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s InitWorkspaceSystemMcpServerRequest) GoString() string {
	return s.String()
}

func (s *InitWorkspaceSystemMcpServerRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *InitWorkspaceSystemMcpServerRequest) SetWorkspaceId(v string) *InitWorkspaceSystemMcpServerRequest {
	s.WorkspaceId = &v
	return s
}

func (s *InitWorkspaceSystemMcpServerRequest) Validate() error {
	return dara.Validate(s)
}
