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
	DisableSessionIdReuse         *bool  `json:"disableSessionIdReuse,omitempty" xml:"disableSessionIdReuse,omitempty"`
	SessionConcurrencyPerInstance *int64 `json:"sessionConcurrencyPerInstance,omitempty" xml:"sessionConcurrencyPerInstance,omitempty"`
	SessionIdleTimeoutInSeconds   *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	SessionTTLInSeconds           *int64 `json:"sessionTTLInSeconds,omitempty" xml:"sessionTTLInSeconds,omitempty"`
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
