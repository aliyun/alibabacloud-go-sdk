// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMcpServerId(v string) *GetDataAgentMcpRequest
	GetMcpServerId() *string
	SetWorkspaceId(v string) *GetDataAgentMcpRequest
	GetWorkspaceId() *string
}

type GetDataAgentMcpRequest struct {
	// The unique identifier of the MCP Server to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 44lg***z65
	McpServerId *string `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
	// The Data Agent workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// atvx***xmz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDataAgentMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentMcpRequest) GoString() string {
	return s.String()
}

func (s *GetDataAgentMcpRequest) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *GetDataAgentMcpRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDataAgentMcpRequest) SetMcpServerId(v string) *GetDataAgentMcpRequest {
	s.McpServerId = &v
	return s
}

func (s *GetDataAgentMcpRequest) SetWorkspaceId(v string) *GetDataAgentMcpRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDataAgentMcpRequest) Validate() error {
	return dara.Validate(s)
}
