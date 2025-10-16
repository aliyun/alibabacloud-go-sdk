// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTarget interface {
	dara.Model
	String() string
	GoString() string
	SetLlmConfig(v *LLMAPIConfiguration) *Target
	GetLlmConfig() *LLMAPIConfiguration
	SetMcpAPI(v *MCPAPI) *Target
	GetMcpAPI() *MCPAPI
	SetTargetType(v string) *Target
	GetTargetType() *string
}

type Target struct {
	LlmConfig  *LLMAPIConfiguration `json:"llmConfig,omitempty" xml:"llmConfig,omitempty"`
	McpAPI     *MCPAPI              `json:"mcpAPI,omitempty" xml:"mcpAPI,omitempty"`
	TargetType *string              `json:"targetType,omitempty" xml:"targetType,omitempty"`
}

func (s Target) String() string {
	return dara.Prettify(s)
}

func (s Target) GoString() string {
	return s.String()
}

func (s *Target) GetLlmConfig() *LLMAPIConfiguration {
	return s.LlmConfig
}

func (s *Target) GetMcpAPI() *MCPAPI {
	return s.McpAPI
}

func (s *Target) GetTargetType() *string {
	return s.TargetType
}

func (s *Target) SetLlmConfig(v *LLMAPIConfiguration) *Target {
	s.LlmConfig = v
	return s
}

func (s *Target) SetMcpAPI(v *MCPAPI) *Target {
	s.McpAPI = v
	return s
}

func (s *Target) SetTargetType(v string) *Target {
	s.TargetType = &v
	return s
}

func (s *Target) Validate() error {
	if s.LlmConfig != nil {
		if err := s.LlmConfig.Validate(); err != nil {
			return err
		}
	}
	if s.McpAPI != nil {
		if err := s.McpAPI.Validate(); err != nil {
			return err
		}
	}
	return nil
}
