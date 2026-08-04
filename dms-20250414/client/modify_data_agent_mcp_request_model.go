// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataAgentMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyDataAgentMcpRequest
	GetDescription() *string
	SetEnable(v bool) *ModifyDataAgentMcpRequest
	GetEnable() *bool
	SetEndpoint(v string) *ModifyDataAgentMcpRequest
	GetEndpoint() *string
	SetHeaders(v string) *ModifyDataAgentMcpRequest
	GetHeaders() *string
	SetMcpServerId(v string) *ModifyDataAgentMcpRequest
	GetMcpServerId() *string
	SetName(v string) *ModifyDataAgentMcpRequest
	GetName() *string
	SetNeedUidInHeader(v bool) *ModifyDataAgentMcpRequest
	GetNeedUidInHeader() *bool
	SetTransportType(v string) *ModifyDataAgentMcpRequest
	GetTransportType() *string
	SetWorkspaceId(v string) *ModifyDataAgentMcpRequest
	GetWorkspaceId() *string
}

type ModifyDataAgentMcpRequest struct {
	// The brief description of the artifact. This parameter can be empty.
	//
	// example:
	//
	// project name pass the check
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the MCP server is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The endpoint of the MCP instance.
	//
	// example:
	//
	// http://***.com/mcp
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The request header settings.
	//
	// example:
	//
	// {"Authorization":"Bearer ***"}
	Headers *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The ID of the MCP server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6126jk***h2
	McpServerId *string `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
	// The MCP name.
	//
	// example:
	//
	// Efficiency Diagnostics V3
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to include the Alibaba Cloud UID in the request header.
	//
	// example:
	//
	// true
	NeedUidInHeader *bool `json:"NeedUidInHeader,omitempty" xml:"NeedUidInHeader,omitempty"`
	// The transport channel type. Valid values: streamablehttp, sse.
	//
	// example:
	//
	// sse
	TransportType *string `json:"TransportType,omitempty" xml:"TransportType,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e3p***v4
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ModifyDataAgentMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataAgentMcpRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataAgentMcpRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDataAgentMcpRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyDataAgentMcpRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ModifyDataAgentMcpRequest) GetHeaders() *string {
	return s.Headers
}

func (s *ModifyDataAgentMcpRequest) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ModifyDataAgentMcpRequest) GetName() *string {
	return s.Name
}

func (s *ModifyDataAgentMcpRequest) GetNeedUidInHeader() *bool {
	return s.NeedUidInHeader
}

func (s *ModifyDataAgentMcpRequest) GetTransportType() *string {
	return s.TransportType
}

func (s *ModifyDataAgentMcpRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModifyDataAgentMcpRequest) SetDescription(v string) *ModifyDataAgentMcpRequest {
	s.Description = &v
	return s
}

func (s *ModifyDataAgentMcpRequest) SetEnable(v bool) *ModifyDataAgentMcpRequest {
	s.Enable = &v
	return s
}

func (s *ModifyDataAgentMcpRequest) SetEndpoint(v string) *ModifyDataAgentMcpRequest {
	s.Endpoint = &v
	return s
}

func (s *ModifyDataAgentMcpRequest) SetHeaders(v string) *ModifyDataAgentMcpRequest {
	s.Headers = &v
	return s
}

func (s *ModifyDataAgentMcpRequest) SetMcpServerId(v string) *ModifyDataAgentMcpRequest {
	s.McpServerId = &v
	return s
}

func (s *ModifyDataAgentMcpRequest) SetName(v string) *ModifyDataAgentMcpRequest {
	s.Name = &v
	return s
}

func (s *ModifyDataAgentMcpRequest) SetNeedUidInHeader(v bool) *ModifyDataAgentMcpRequest {
	s.NeedUidInHeader = &v
	return s
}

func (s *ModifyDataAgentMcpRequest) SetTransportType(v string) *ModifyDataAgentMcpRequest {
	s.TransportType = &v
	return s
}

func (s *ModifyDataAgentMcpRequest) SetWorkspaceId(v string) *ModifyDataAgentMcpRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ModifyDataAgentMcpRequest) Validate() error {
	return dara.Validate(s)
}
