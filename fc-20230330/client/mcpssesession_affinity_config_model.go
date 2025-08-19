// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPSSESessionAffinityConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSessionConcurrencyPerInstance(v int64) *MCPSSESessionAffinityConfig
	GetSessionConcurrencyPerInstance() *int64
	SetSseEndpointPath(v string) *MCPSSESessionAffinityConfig
	GetSseEndpointPath() *string
}

type MCPSSESessionAffinityConfig struct {
	SessionConcurrencyPerInstance *int64  `json:"sessionConcurrencyPerInstance,omitempty" xml:"sessionConcurrencyPerInstance,omitempty"`
	SseEndpointPath               *string `json:"sseEndpointPath,omitempty" xml:"sseEndpointPath,omitempty"`
}

func (s MCPSSESessionAffinityConfig) String() string {
	return dara.Prettify(s)
}

func (s MCPSSESessionAffinityConfig) GoString() string {
	return s.String()
}

func (s *MCPSSESessionAffinityConfig) GetSessionConcurrencyPerInstance() *int64 {
	return s.SessionConcurrencyPerInstance
}

func (s *MCPSSESessionAffinityConfig) GetSseEndpointPath() *string {
	return s.SseEndpointPath
}

func (s *MCPSSESessionAffinityConfig) SetSessionConcurrencyPerInstance(v int64) *MCPSSESessionAffinityConfig {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *MCPSSESessionAffinityConfig) SetSseEndpointPath(v string) *MCPSSESessionAffinityConfig {
	s.SseEndpointPath = &v
	return s
}

func (s *MCPSSESessionAffinityConfig) Validate() error {
	return dara.Validate(s)
}
