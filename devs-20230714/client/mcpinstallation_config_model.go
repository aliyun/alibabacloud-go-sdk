// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPInstallationConfig interface {
	dara.Model
	String() string
	GoString() string
	SetMcpServers(v *MCPServerInstallationConfig) *MCPInstallationConfig
	GetMcpServers() *MCPServerInstallationConfig
}

type MCPInstallationConfig struct {
	McpServers *MCPServerInstallationConfig `json:"mcpServers,omitempty" xml:"mcpServers,omitempty"`
}

func (s MCPInstallationConfig) String() string {
	return dara.Prettify(s)
}

func (s MCPInstallationConfig) GoString() string {
	return s.String()
}

func (s *MCPInstallationConfig) GetMcpServers() *MCPServerInstallationConfig {
	return s.McpServers
}

func (s *MCPInstallationConfig) SetMcpServers(v *MCPServerInstallationConfig) *MCPInstallationConfig {
	s.McpServers = v
	return s
}

func (s *MCPInstallationConfig) Validate() error {
	return dara.Validate(s)
}
