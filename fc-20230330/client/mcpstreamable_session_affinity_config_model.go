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
	// The maximum number of sessions for simultaneous processing by a single instance. Valid values: 1 to 200.
	//
	// example:
	//
	// 20
	SessionConcurrencyPerInstance *int64 `json:"sessionConcurrencyPerInstance,omitempty" xml:"sessionConcurrencyPerInstance,omitempty"`
	// The maximum idle time in seconds before a session enters an idle state due to user inactivity. The maximum duration is the upper limit of a single session lifecycle. Valid values: 0 to 21600.
	//
	// example:
	//
	// 1800
	SessionIdleTimeoutInSeconds *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	// The time-to-live of a session in seconds, covering the entire process from creation and usage to final destruction. If the time-to-live is exceeded, Function Compute automatically destroys the session and no longer guarantees affinity. Valid values: 1 to 21600.
	//
	// example:
	//
	// 21600
	SessionTTLInSeconds *int64 `json:"sessionTTLInSeconds,omitempty" xml:"sessionTTLInSeconds,omitempty"`
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
