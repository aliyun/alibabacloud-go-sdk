// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHeaderFieldSessionAffinityConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAffinityHeaderFieldName(v string) *HeaderFieldSessionAffinityConfig
	GetAffinityHeaderFieldName() *string
	SetDisableSessionIdReuse(v bool) *HeaderFieldSessionAffinityConfig
	GetDisableSessionIdReuse() *bool
	SetSessionConcurrencyPerInstance(v int64) *HeaderFieldSessionAffinityConfig
	GetSessionConcurrencyPerInstance() *int64
	SetSessionIdleTimeoutInSeconds(v int64) *HeaderFieldSessionAffinityConfig
	GetSessionIdleTimeoutInSeconds() *int64
	SetSessionTTLInSeconds(v int64) *HeaderFieldSessionAffinityConfig
	GetSessionTTLInSeconds() *int64
}

type HeaderFieldSessionAffinityConfig struct {
	// The name of the HTTP request header that passes the client session identity. The name must be 5 to 40 characters long, start with a letter, and contain only letters, numbers, hyphens (-), and underscores (_). The name cannot start with the x-fc- prefix.
	//
	// example:
	//
	// test-session-header1
	AffinityHeaderFieldName *string `json:"affinityHeaderFieldName,omitempty" xml:"affinityHeaderFieldName,omitempty"`
	// The default value is False. If set to False, a session ID can be reused in a new request after the original session expires. The system treats this as a new session and attaches it to a new instance. If set to True, an expired session ID cannot be reused.
	//
	// example:
	//
	// false
	DisableSessionIdReuse *bool `json:"disableSessionIdReuse,omitempty" xml:"disableSessionIdReuse,omitempty"`
	// The maximum number of sessions that a single instance can process simultaneously. The value must be an integer from 1 to 200.
	//
	// example:
	//
	// 20
	SessionConcurrencyPerInstance *int64 `json:"sessionConcurrencyPerInstance,omitempty" xml:"sessionConcurrencyPerInstance,omitempty"`
	// The idle timeout period for a session in seconds. A session becomes idle if no operations are performed within this period. The maximum value cannot exceed the session\\"s TTL. The value must be an integer from 0 to 21600.
	//
	// example:
	//
	// 1800
	SessionIdleTimeoutInSeconds *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	// The session\\"s Time to Live (TTL) in seconds. This defines the entire lifecycle of a session, from creation to destruction. After this period expires, Function Compute automatically destroys the session and no longer guarantees affinity. The value must be an integer from 1 to 21600.
	//
	// example:
	//
	// 21600
	SessionTTLInSeconds *int64 `json:"sessionTTLInSeconds,omitempty" xml:"sessionTTLInSeconds,omitempty"`
}

func (s HeaderFieldSessionAffinityConfig) String() string {
	return dara.Prettify(s)
}

func (s HeaderFieldSessionAffinityConfig) GoString() string {
	return s.String()
}

func (s *HeaderFieldSessionAffinityConfig) GetAffinityHeaderFieldName() *string {
	return s.AffinityHeaderFieldName
}

func (s *HeaderFieldSessionAffinityConfig) GetDisableSessionIdReuse() *bool {
	return s.DisableSessionIdReuse
}

func (s *HeaderFieldSessionAffinityConfig) GetSessionConcurrencyPerInstance() *int64 {
	return s.SessionConcurrencyPerInstance
}

func (s *HeaderFieldSessionAffinityConfig) GetSessionIdleTimeoutInSeconds() *int64 {
	return s.SessionIdleTimeoutInSeconds
}

func (s *HeaderFieldSessionAffinityConfig) GetSessionTTLInSeconds() *int64 {
	return s.SessionTTLInSeconds
}

func (s *HeaderFieldSessionAffinityConfig) SetAffinityHeaderFieldName(v string) *HeaderFieldSessionAffinityConfig {
	s.AffinityHeaderFieldName = &v
	return s
}

func (s *HeaderFieldSessionAffinityConfig) SetDisableSessionIdReuse(v bool) *HeaderFieldSessionAffinityConfig {
	s.DisableSessionIdReuse = &v
	return s
}

func (s *HeaderFieldSessionAffinityConfig) SetSessionConcurrencyPerInstance(v int64) *HeaderFieldSessionAffinityConfig {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *HeaderFieldSessionAffinityConfig) SetSessionIdleTimeoutInSeconds(v int64) *HeaderFieldSessionAffinityConfig {
	s.SessionIdleTimeoutInSeconds = &v
	return s
}

func (s *HeaderFieldSessionAffinityConfig) SetSessionTTLInSeconds(v int64) *HeaderFieldSessionAffinityConfig {
	s.SessionTTLInSeconds = &v
	return s
}

func (s *HeaderFieldSessionAffinityConfig) Validate() error {
	return dara.Validate(s)
}
