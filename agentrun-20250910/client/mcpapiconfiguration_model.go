// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPAPIConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *MCPAPIConfiguration
	GetDescription() *string
	SetExposedUriPath(v string) *MCPAPIConfiguration
	GetExposedUriPath() *string
	SetMcpStatisticsEnable(v bool) *MCPAPIConfiguration
	GetMcpStatisticsEnable() *bool
	SetProtocol(v string) *MCPAPIConfiguration
	GetProtocol() *string
	SetToolId(v string) *MCPAPIConfiguration
	GetToolId() *string
}

type MCPAPIConfiguration struct {
	Description         *string `json:"description,omitempty" xml:"description,omitempty"`
	ExposedUriPath      *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	McpStatisticsEnable *bool   `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	Protocol            *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	ToolId              *string `json:"toolId,omitempty" xml:"toolId,omitempty"`
}

func (s MCPAPIConfiguration) String() string {
	return dara.Prettify(s)
}

func (s MCPAPIConfiguration) GoString() string {
	return s.String()
}

func (s *MCPAPIConfiguration) GetDescription() *string {
	return s.Description
}

func (s *MCPAPIConfiguration) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *MCPAPIConfiguration) GetMcpStatisticsEnable() *bool {
	return s.McpStatisticsEnable
}

func (s *MCPAPIConfiguration) GetProtocol() *string {
	return s.Protocol
}

func (s *MCPAPIConfiguration) GetToolId() *string {
	return s.ToolId
}

func (s *MCPAPIConfiguration) SetDescription(v string) *MCPAPIConfiguration {
	s.Description = &v
	return s
}

func (s *MCPAPIConfiguration) SetExposedUriPath(v string) *MCPAPIConfiguration {
	s.ExposedUriPath = &v
	return s
}

func (s *MCPAPIConfiguration) SetMcpStatisticsEnable(v bool) *MCPAPIConfiguration {
	s.McpStatisticsEnable = &v
	return s
}

func (s *MCPAPIConfiguration) SetProtocol(v string) *MCPAPIConfiguration {
	s.Protocol = &v
	return s
}

func (s *MCPAPIConfiguration) SetToolId(v string) *MCPAPIConfiguration {
	s.ToolId = &v
	return s
}

func (s *MCPAPIConfiguration) Validate() error {
	return dara.Validate(s)
}
