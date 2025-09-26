// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTargetConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetLlmAPIConfig(v *LLMAPIConfiguration) *TargetConfiguration
	GetLlmAPIConfig() *LLMAPIConfiguration
	SetMcpAPIConfig(v *MCPAPIConfiguration) *TargetConfiguration
	GetMcpAPIConfig() *MCPAPIConfiguration
	SetTargetType(v string) *TargetConfiguration
	GetTargetType() *string
}

type TargetConfiguration struct {
	LlmAPIConfig *LLMAPIConfiguration `json:"llmAPIConfig,omitempty" xml:"llmAPIConfig,omitempty"`
	McpAPIConfig *MCPAPIConfiguration `json:"mcpAPIConfig,omitempty" xml:"mcpAPIConfig,omitempty"`
	TargetType   *string              `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s TargetConfiguration) String() string {
	return dara.Prettify(s)
}

func (s TargetConfiguration) GoString() string {
	return s.String()
}

func (s *TargetConfiguration) GetLlmAPIConfig() *LLMAPIConfiguration {
	return s.LlmAPIConfig
}

func (s *TargetConfiguration) GetMcpAPIConfig() *MCPAPIConfiguration {
	return s.McpAPIConfig
}

func (s *TargetConfiguration) GetTargetType() *string {
	return s.TargetType
}

func (s *TargetConfiguration) SetLlmAPIConfig(v *LLMAPIConfiguration) *TargetConfiguration {
	s.LlmAPIConfig = v
	return s
}

func (s *TargetConfiguration) SetMcpAPIConfig(v *MCPAPIConfiguration) *TargetConfiguration {
	s.McpAPIConfig = v
	return s
}

func (s *TargetConfiguration) SetTargetType(v string) *TargetConfiguration {
	s.TargetType = &v
	return s
}

func (s *TargetConfiguration) Validate() error {
	return dara.Validate(s)
}
