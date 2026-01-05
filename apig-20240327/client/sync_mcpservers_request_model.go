// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMCPServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainIds(v []*string) *SyncMCPServersRequest
	GetDomainIds() []*string
	SetGatewayId(v string) *SyncMCPServersRequest
	GetGatewayId() *string
	SetNacosMcpServers(v []*SyncMCPServersRequestNacosMcpServers) *SyncMCPServersRequest
	GetNacosMcpServers() []*SyncMCPServersRequestNacosMcpServers
	SetNamespace(v string) *SyncMCPServersRequest
	GetNamespace() *string
	SetSourceId(v string) *SyncMCPServersRequest
	GetSourceId() *string
}

type SyncMCPServersRequest struct {
	// The domain ID.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// The ID of the gateway.
	//
	// example:
	//
	// gw-cq7l5s5lhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The synchronized Nacos MCP server list. If the synchronized MCP server is included, add the mcpServerId parameter.
	NacosMcpServers []*SyncMCPServersRequestNacosMcpServers `json:"nacosMcpServers,omitempty" xml:"nacosMcpServers,omitempty" type:"Repeated"`
	// The Nacos namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The source ID.
	//
	// example:
	//
	// src-d40mff6m1hkpt84vep60
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s SyncMCPServersRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncMCPServersRequest) GoString() string {
	return s.String()
}

func (s *SyncMCPServersRequest) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *SyncMCPServersRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *SyncMCPServersRequest) GetNacosMcpServers() []*SyncMCPServersRequestNacosMcpServers {
	return s.NacosMcpServers
}

func (s *SyncMCPServersRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *SyncMCPServersRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *SyncMCPServersRequest) SetDomainIds(v []*string) *SyncMCPServersRequest {
	s.DomainIds = v
	return s
}

func (s *SyncMCPServersRequest) SetGatewayId(v string) *SyncMCPServersRequest {
	s.GatewayId = &v
	return s
}

func (s *SyncMCPServersRequest) SetNacosMcpServers(v []*SyncMCPServersRequestNacosMcpServers) *SyncMCPServersRequest {
	s.NacosMcpServers = v
	return s
}

func (s *SyncMCPServersRequest) SetNamespace(v string) *SyncMCPServersRequest {
	s.Namespace = &v
	return s
}

func (s *SyncMCPServersRequest) SetSourceId(v string) *SyncMCPServersRequest {
	s.SourceId = &v
	return s
}

func (s *SyncMCPServersRequest) Validate() error {
	if s.NacosMcpServers != nil {
		for _, item := range s.NacosMcpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SyncMCPServersRequestNacosMcpServers struct {
	// The exposed URI path. This parameter is required when the protocol parameter is set to SSE or StreamableHTTP and the type parameter is set to RealMCP.
	//
	// example:
	//
	// /sse
	ExposedUriPath *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	// The Nacos instance ID.
	//
	// example:
	//
	// mse-24afmoioxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The MCP server ID.
	//
	// example:
	//
	// mcp-d3s8qo6m1hknegofa3bg
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The name of the MCP server.
	//
	// example:
	//
	// test
	McpServerName *string `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	// The protocol.
	Protocols []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
}

func (s SyncMCPServersRequestNacosMcpServers) String() string {
	return dara.Prettify(s)
}

func (s SyncMCPServersRequestNacosMcpServers) GoString() string {
	return s.String()
}

func (s *SyncMCPServersRequestNacosMcpServers) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *SyncMCPServersRequestNacosMcpServers) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SyncMCPServersRequestNacosMcpServers) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *SyncMCPServersRequestNacosMcpServers) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *SyncMCPServersRequestNacosMcpServers) GetProtocols() []*string {
	return s.Protocols
}

func (s *SyncMCPServersRequestNacosMcpServers) SetExposedUriPath(v string) *SyncMCPServersRequestNacosMcpServers {
	s.ExposedUriPath = &v
	return s
}

func (s *SyncMCPServersRequestNacosMcpServers) SetInstanceId(v string) *SyncMCPServersRequestNacosMcpServers {
	s.InstanceId = &v
	return s
}

func (s *SyncMCPServersRequestNacosMcpServers) SetMcpServerId(v string) *SyncMCPServersRequestNacosMcpServers {
	s.McpServerId = &v
	return s
}

func (s *SyncMCPServersRequestNacosMcpServers) SetMcpServerName(v string) *SyncMCPServersRequestNacosMcpServers {
	s.McpServerName = &v
	return s
}

func (s *SyncMCPServersRequestNacosMcpServers) SetProtocols(v []*string) *SyncMCPServersRequestNacosMcpServers {
	s.Protocols = v
	return s
}

func (s *SyncMCPServersRequestNacosMcpServers) Validate() error {
	return dara.Validate(s)
}
