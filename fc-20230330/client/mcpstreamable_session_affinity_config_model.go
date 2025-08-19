// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPStreamableSessionAffinityConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSessionConcurrencyPerInstance(v int64) *MCPStreamableSessionAffinityConfig
	GetSessionConcurrencyPerInstance() *int64
	SetSessionIdleTimeoutInSeconds(v int64) *MCPStreamableSessionAffinityConfig
	GetSessionIdleTimeoutInSeconds() *int64
	SetSessionTTLInSeconds(v int64) *MCPStreamableSessionAffinityConfig
	GetSessionTTLInSeconds() *int64
}

type MCPStreamableSessionAffinityConfig struct {
	SessionConcurrencyPerInstance *int64 `json:"sessionConcurrencyPerInstance,omitempty" xml:"sessionConcurrencyPerInstance,omitempty"`
	SessionIdleTimeoutInSeconds   *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	SessionTTLInSeconds           *int64 `json:"sessionTTLInSeconds,omitempty" xml:"sessionTTLInSeconds,omitempty"`
}

func (s MCPStreamableSessionAffinityConfig) String() string {
	return dara.Prettify(s)
}

func (s MCPStreamableSessionAffinityConfig) GoString() string {
	return s.String()
}

func (s *MCPStreamableSessionAffinityConfig) GetSessionConcurrencyPerInstance() *int64 {
	return s.SessionConcurrencyPerInstance
}

func (s *MCPStreamableSessionAffinityConfig) GetSessionIdleTimeoutInSeconds() *int64 {
	return s.SessionIdleTimeoutInSeconds
}

func (s *MCPStreamableSessionAffinityConfig) GetSessionTTLInSeconds() *int64 {
	return s.SessionTTLInSeconds
}

func (s *MCPStreamableSessionAffinityConfig) SetSessionConcurrencyPerInstance(v int64) *MCPStreamableSessionAffinityConfig {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *MCPStreamableSessionAffinityConfig) SetSessionIdleTimeoutInSeconds(v int64) *MCPStreamableSessionAffinityConfig {
	s.SessionIdleTimeoutInSeconds = &v
	return s
}

func (s *MCPStreamableSessionAffinityConfig) SetSessionTTLInSeconds(v int64) *MCPStreamableSessionAffinityConfig {
	s.SessionTTLInSeconds = &v
	return s
}

func (s *MCPStreamableSessionAffinityConfig) Validate() error {
	return dara.Validate(s)
}
