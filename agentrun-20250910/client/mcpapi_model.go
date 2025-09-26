// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPAPI interface {
	dara.Model
	String() string
	GoString() string
	SetBackendConfig(v *MCPBackendConfig) *MCPAPI
	GetBackendConfig() *MCPBackendConfig
	SetDescription(v string) *MCPAPI
	GetDescription() *string
	SetExposedUriPath(v string) *MCPAPI
	GetExposedUriPath() *string
	SetMatch(v *MCPMatch) *MCPAPI
	GetMatch() *MCPMatch
	SetMcpStatisticsEnable(v bool) *MCPAPI
	GetMcpStatisticsEnable() *bool
	SetProtocol(v string) *MCPAPI
	GetProtocol() *string
	SetToolId(v string) *MCPAPI
	GetToolId() *string
	SetToolsConfig(v string) *MCPAPI
	GetToolsConfig() *string
}

type MCPAPI struct {
	BackendConfig       *MCPBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty"`
	Description         *string           `json:"description,omitempty" xml:"description,omitempty"`
	ExposedUriPath      *string           `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	Match               *MCPMatch         `json:"match,omitempty" xml:"match,omitempty"`
	McpStatisticsEnable *bool             `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	Protocol            *string           `json:"protocol,omitempty" xml:"protocol,omitempty"`
	ToolId              *string           `json:"toolId,omitempty" xml:"toolId,omitempty"`
	ToolsConfig         *string           `json:"toolsConfig,omitempty" xml:"toolsConfig,omitempty"`
}

func (s MCPAPI) String() string {
	return dara.Prettify(s)
}

func (s MCPAPI) GoString() string {
	return s.String()
}

func (s *MCPAPI) GetBackendConfig() *MCPBackendConfig {
	return s.BackendConfig
}

func (s *MCPAPI) GetDescription() *string {
	return s.Description
}

func (s *MCPAPI) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *MCPAPI) GetMatch() *MCPMatch {
	return s.Match
}

func (s *MCPAPI) GetMcpStatisticsEnable() *bool {
	return s.McpStatisticsEnable
}

func (s *MCPAPI) GetProtocol() *string {
	return s.Protocol
}

func (s *MCPAPI) GetToolId() *string {
	return s.ToolId
}

func (s *MCPAPI) GetToolsConfig() *string {
	return s.ToolsConfig
}

func (s *MCPAPI) SetBackendConfig(v *MCPBackendConfig) *MCPAPI {
	s.BackendConfig = v
	return s
}

func (s *MCPAPI) SetDescription(v string) *MCPAPI {
	s.Description = &v
	return s
}

func (s *MCPAPI) SetExposedUriPath(v string) *MCPAPI {
	s.ExposedUriPath = &v
	return s
}

func (s *MCPAPI) SetMatch(v *MCPMatch) *MCPAPI {
	s.Match = v
	return s
}

func (s *MCPAPI) SetMcpStatisticsEnable(v bool) *MCPAPI {
	s.McpStatisticsEnable = &v
	return s
}

func (s *MCPAPI) SetProtocol(v string) *MCPAPI {
	s.Protocol = &v
	return s
}

func (s *MCPAPI) SetToolId(v string) *MCPAPI {
	s.ToolId = &v
	return s
}

func (s *MCPAPI) SetToolsConfig(v string) *MCPAPI {
	s.ToolsConfig = &v
	return s
}

func (s *MCPAPI) Validate() error {
	return dara.Validate(s)
}
