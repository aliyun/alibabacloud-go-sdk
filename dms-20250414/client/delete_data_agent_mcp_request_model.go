// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMcpServerIds(v []*string) *DeleteDataAgentMcpRequest
	GetMcpServerIds() []*string
	SetWorkspaceId(v string) *DeleteDataAgentMcpRequest
	GetWorkspaceId() *string
}

type DeleteDataAgentMcpRequest struct {
	// The list of MCP Server IDs to delete.
	McpServerIds []*string `json:"McpServerIds,omitempty" xml:"McpServerIds,omitempty" type:"Repeated"`
	// The ID of the Data Agent workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// atvx***xmz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDataAgentMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentMcpRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentMcpRequest) GetMcpServerIds() []*string {
	return s.McpServerIds
}

func (s *DeleteDataAgentMcpRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDataAgentMcpRequest) SetMcpServerIds(v []*string) *DeleteDataAgentMcpRequest {
	s.McpServerIds = v
	return s
}

func (s *DeleteDataAgentMcpRequest) SetWorkspaceId(v string) *DeleteDataAgentMcpRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDataAgentMcpRequest) Validate() error {
	return dara.Validate(s)
}
