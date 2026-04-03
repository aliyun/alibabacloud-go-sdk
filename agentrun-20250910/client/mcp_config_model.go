// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMcpConfig interface {
	dara.Model
	String() string
	GoString() string
	SetBoundConfiguration(v *BoundConfiguration) *McpConfig
	GetBoundConfiguration() *BoundConfiguration
	SetMcpProxyConfiguration(v *McpProxyConfiguration) *McpConfig
	GetMcpProxyConfiguration() *McpProxyConfiguration
	SetProxyEnabled(v bool) *McpConfig
	GetProxyEnabled() *bool
	SetSessionAffinity(v string) *McpConfig
	GetSessionAffinity() *string
	SetSessionAffinityConfig(v string) *McpConfig
	GetSessionAffinityConfig() *string
}

type McpConfig struct {
	// 工具的绑定配置，定义工具与 HTTP 路径和方法的映射关系
	BoundConfiguration *BoundConfiguration `json:"boundConfiguration,omitempty" xml:"boundConfiguration,omitempty"`
	// MCP 代理的详细配置，包括钩子函数等，用于在 MCP 请求处理过程中执行自定义逻辑
	McpProxyConfiguration *McpProxyConfiguration `json:"mcpProxyConfiguration,omitempty" xml:"mcpProxyConfiguration,omitempty"`
	// 是否启用 MCP 代理功能，启用后可以通过代理配置对 MCP 请求进行拦截和处理
	//
	// example:
	//
	// true
	ProxyEnabled *bool `json:"proxyEnabled,omitempty" xml:"proxyEnabled,omitempty"`
	// 会话亲和性策略，用于控制请求的路由方式。NONE：无亲和性，MCP_SSE：基于 SSE 的会话亲和性，MCP_STREAMABLE：基于流式 HTTP 的会话亲和性
	//
	// example:
	//
	// MCP_STREAMABLE
	SessionAffinity *string `json:"sessionAffinity,omitempty" xml:"sessionAffinity,omitempty"`
	// 会话亲和性的详细配置信息，JSON 格式字符串，包含会话保持的具体参数
	SessionAffinityConfig *string `json:"sessionAffinityConfig,omitempty" xml:"sessionAffinityConfig,omitempty"`
}

func (s McpConfig) String() string {
	return dara.Prettify(s)
}

func (s McpConfig) GoString() string {
	return s.String()
}

func (s *McpConfig) GetBoundConfiguration() *BoundConfiguration {
	return s.BoundConfiguration
}

func (s *McpConfig) GetMcpProxyConfiguration() *McpProxyConfiguration {
	return s.McpProxyConfiguration
}

func (s *McpConfig) GetProxyEnabled() *bool {
	return s.ProxyEnabled
}

func (s *McpConfig) GetSessionAffinity() *string {
	return s.SessionAffinity
}

func (s *McpConfig) GetSessionAffinityConfig() *string {
	return s.SessionAffinityConfig
}

func (s *McpConfig) SetBoundConfiguration(v *BoundConfiguration) *McpConfig {
	s.BoundConfiguration = v
	return s
}

func (s *McpConfig) SetMcpProxyConfiguration(v *McpProxyConfiguration) *McpConfig {
	s.McpProxyConfiguration = v
	return s
}

func (s *McpConfig) SetProxyEnabled(v bool) *McpConfig {
	s.ProxyEnabled = &v
	return s
}

func (s *McpConfig) SetSessionAffinity(v string) *McpConfig {
	s.SessionAffinity = &v
	return s
}

func (s *McpConfig) SetSessionAffinityConfig(v string) *McpConfig {
	s.SessionAffinityConfig = &v
	return s
}

func (s *McpConfig) Validate() error {
	if s.BoundConfiguration != nil {
		if err := s.BoundConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.McpProxyConfiguration != nil {
		if err := s.McpProxyConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
