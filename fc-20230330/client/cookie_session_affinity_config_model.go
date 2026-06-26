// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCookieSessionAffinityConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDisableSessionIdReuse(v bool) *CookieSessionAffinityConfig
	GetDisableSessionIdReuse() *bool
	SetSessionConcurrencyPerInstance(v int64) *CookieSessionAffinityConfig
	GetSessionConcurrencyPerInstance() *int64
	SetSessionIdleTimeoutInSeconds(v int64) *CookieSessionAffinityConfig
	GetSessionIdleTimeoutInSeconds() *int64
	SetSessionTTLInSeconds(v int64) *CookieSessionAffinityConfig
	GetSessionTTLInSeconds() *int64
}

type CookieSessionAffinityConfig struct {
	// The default value is \\`false\\`. When set to \\`false\\`, a request with the same session ID can be sent after the session expires. The system treats this as a new session and attaches it to a new instance. When set to \\`true\\`, the session ID cannot be reused after the session expires.
	//
	// example:
	//
	// false
	DisableSessionIdReuse *bool `json:"disableSessionIdReuse,omitempty" xml:"disableSessionIdReuse,omitempty"`
	// The maximum number of sessions that a single instance can process at the same time. The value must be an integer from 1 to 200.
	//
	// example:
	//
	// 20
	SessionConcurrencyPerInstance *int64 `json:"sessionConcurrencyPerInstance,omitempty" xml:"sessionConcurrencyPerInstance,omitempty"`
	// The duration in seconds that a session can remain idle. If a user is inactive for this period, the session is considered idle. The maximum duration is limited by the session\\"s lifecycle. The value must be between 0 and 21,600.
	//
	// example:
	//
	// 1800
	SessionIdleTimeoutInSeconds *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	// The total lifecycle of a session in seconds, from creation to destruction. After this period, Function Compute automatically destroys the session and no longer guarantees affinity. The value must be an integer from 1 to 21,600.
	//
	// example:
	//
	// 21600
	SessionTTLInSeconds *int64 `json:"sessionTTLInSeconds,omitempty" xml:"sessionTTLInSeconds,omitempty"`
}

func (s CookieSessionAffinityConfig) String() string {
	return dara.Prettify(s)
}

func (s CookieSessionAffinityConfig) GoString() string {
	return s.String()
}

func (s *CookieSessionAffinityConfig) GetDisableSessionIdReuse() *bool {
	return s.DisableSessionIdReuse
}

func (s *CookieSessionAffinityConfig) GetSessionConcurrencyPerInstance() *int64 {
	return s.SessionConcurrencyPerInstance
}

func (s *CookieSessionAffinityConfig) GetSessionIdleTimeoutInSeconds() *int64 {
	return s.SessionIdleTimeoutInSeconds
}

func (s *CookieSessionAffinityConfig) GetSessionTTLInSeconds() *int64 {
	return s.SessionTTLInSeconds
}

func (s *CookieSessionAffinityConfig) SetDisableSessionIdReuse(v bool) *CookieSessionAffinityConfig {
	s.DisableSessionIdReuse = &v
	return s
}

func (s *CookieSessionAffinityConfig) SetSessionConcurrencyPerInstance(v int64) *CookieSessionAffinityConfig {
	s.SessionConcurrencyPerInstance = &v
	return s
}

func (s *CookieSessionAffinityConfig) SetSessionIdleTimeoutInSeconds(v int64) *CookieSessionAffinityConfig {
	s.SessionIdleTimeoutInSeconds = &v
	return s
}

func (s *CookieSessionAffinityConfig) SetSessionTTLInSeconds(v int64) *CookieSessionAffinityConfig {
	s.SessionTTLInSeconds = &v
	return s
}

func (s *CookieSessionAffinityConfig) Validate() error {
	return dara.Validate(s)
}
