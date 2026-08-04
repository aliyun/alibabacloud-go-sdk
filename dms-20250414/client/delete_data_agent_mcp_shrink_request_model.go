// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentMcpShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMcpServerIdsShrink(v string) *DeleteDataAgentMcpShrinkRequest
	GetMcpServerIdsShrink() *string
	SetWorkspaceId(v string) *DeleteDataAgentMcpShrinkRequest
	GetWorkspaceId() *string
}

type DeleteDataAgentMcpShrinkRequest struct {
	// The list of MCP Server IDs to delete.
	McpServerIdsShrink *string `json:"McpServerIds,omitempty" xml:"McpServerIds,omitempty"`
	// The ID of the Data Agent workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// atvx***xmz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataAgentMcpShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentMcpShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentMcpShrinkRequest) GetMcpServerIdsShrink() *string {
	return s.McpServerIdsShrink
}

func (s *DeleteDataAgentMcpShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDataAgentMcpShrinkRequest) SetMcpServerIdsShrink(v string) *DeleteDataAgentMcpShrinkRequest {
	s.McpServerIdsShrink = &v
	return s
}

func (s *DeleteDataAgentMcpShrinkRequest) SetWorkspaceId(v string) *DeleteDataAgentMcpShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataAgentMcpShrinkRequest) Validate() error {
	return dara.Validate(s)
}
