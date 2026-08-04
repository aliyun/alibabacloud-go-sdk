// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallDataAgentMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *InstallDataAgentMcpRequest
	GetDescription() *string
	SetEndpoint(v string) *InstallDataAgentMcpRequest
	GetEndpoint() *string
	SetFromJson(v string) *InstallDataAgentMcpRequest
	GetFromJson() *string
	SetHeaders(v string) *InstallDataAgentMcpRequest
	GetHeaders() *string
	SetName(v string) *InstallDataAgentMcpRequest
	GetName() *string
	SetNeedUidInHeader(v bool) *InstallDataAgentMcpRequest
	GetNeedUidInHeader() *bool
	SetNetType(v string) *InstallDataAgentMcpRequest
	GetNetType() *string
	SetTransportType(v string) *InstallDataAgentMcpRequest
	GetTransportType() *string
	SetVpcId(v string) *InstallDataAgentMcpRequest
	GetVpcId() *string
	SetVswId(v string) *InstallDataAgentMcpRequest
	GetVswId() *string
	SetWorkspaceId(v string) *InstallDataAgentMcpRequest
	GetWorkspaceId() *string
}

type InstallDataAgentMcpRequest struct {
	// The brief description of the artifact. This parameter can be empty.
	//
	// example:
	//
	// query user information by user ID
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint of the MCP instance.
	//
	// example:
	//
	// http://***.com/mcp
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The JSON used to create the MCP.
	//
	// example:
	//
	// {
	//
	//   "mcpServers": {
	//
	//     "sse-server-name": {
	//
	//       "description": "describe sse mcp server",
	//
	//       "type": "sse",
	//
	//       "netType": "vpc",
	//
	//       "vpcId": "vpc-xxxx",
	//
	//       "url": "http://sse-in-vpc.com/sse",
	//
	//       "needUidInHeader": true,
	//
	//       "headers": {
	//
	//         "Authorization": "Bearer <token>"
	//
	//       }
	//
	//     }
	//
	//   }
	//
	// }
	FromJson *string `json:"FromJson,omitempty" xml:"FromJson,omitempty"`
	// The request header settings.
	//
	// example:
	//
	// {"Authorization":"Bearer ***"}
	Headers *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The MCP name.
	//
	// example:
	//
	// query_tool
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to include the Alibaba Cloud UID in the request header.
	//
	// example:
	//
	// true
	NeedUidInHeader *bool `json:"NeedUidInHeader,omitempty" xml:"NeedUidInHeader,omitempty"`
	// The network type. Valid values:
	//
	// - `vpc`: virtual private cloud.
	//
	// - `public`: public network.
	//
	// example:
	//
	// public
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The transport channel type. Valid values: streamablehttp and sse.
	//
	// example:
	//
	// sse
	TransportType *string `json:"TransportType,omitempty" xml:"TransportType,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-uf63***o5
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vb5j***6h
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s InstallDataAgentMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallDataAgentMcpRequest) GoString() string {
	return s.String()
}

func (s *InstallDataAgentMcpRequest) GetDescription() *string {
	return s.Description
}

func (s *InstallDataAgentMcpRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *InstallDataAgentMcpRequest) GetFromJson() *string {
	return s.FromJson
}

func (s *InstallDataAgentMcpRequest) GetHeaders() *string {
	return s.Headers
}

func (s *InstallDataAgentMcpRequest) GetName() *string {
	return s.Name
}

func (s *InstallDataAgentMcpRequest) GetNeedUidInHeader() *bool {
	return s.NeedUidInHeader
}

func (s *InstallDataAgentMcpRequest) GetNetType() *string {
	return s.NetType
}

func (s *InstallDataAgentMcpRequest) GetTransportType() *string {
	return s.TransportType
}

func (s *InstallDataAgentMcpRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *InstallDataAgentMcpRequest) GetVswId() *string {
	return s.VswId
}

func (s *InstallDataAgentMcpRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *InstallDataAgentMcpRequest) SetDescription(v string) *InstallDataAgentMcpRequest {
	s.Description = &v
	return s
}

func (s *InstallDataAgentMcpRequest) SetEndpoint(v string) *InstallDataAgentMcpRequest {
	s.Endpoint = &v
	return s
}

func (s *InstallDataAgentMcpRequest) SetFromJson(v string) *InstallDataAgentMcpRequest {
	s.FromJson = &v
	return s
}

func (s *InstallDataAgentMcpRequest) SetHeaders(v string) *InstallDataAgentMcpRequest {
	s.Headers = &v
	return s
}

func (s *InstallDataAgentMcpRequest) SetName(v string) *InstallDataAgentMcpRequest {
	s.Name = &v
	return s
}

func (s *InstallDataAgentMcpRequest) SetNeedUidInHeader(v bool) *InstallDataAgentMcpRequest {
	s.NeedUidInHeader = &v
	return s
}

func (s *InstallDataAgentMcpRequest) SetNetType(v string) *InstallDataAgentMcpRequest {
	s.NetType = &v
	return s
}

func (s *InstallDataAgentMcpRequest) SetTransportType(v string) *InstallDataAgentMcpRequest {
	s.TransportType = &v
	return s
}

func (s *InstallDataAgentMcpRequest) SetVpcId(v string) *InstallDataAgentMcpRequest {
	s.VpcId = &v
	return s
}

func (s *InstallDataAgentMcpRequest) SetVswId(v string) *InstallDataAgentMcpRequest {
	s.VswId = &v
	return s
}

func (s *InstallDataAgentMcpRequest) SetWorkspaceId(v string) *InstallDataAgentMcpRequest {
	s.WorkspaceId = &v
	return s
}

func (s *InstallDataAgentMcpRequest) Validate() error {
	return dara.Validate(s)
}
