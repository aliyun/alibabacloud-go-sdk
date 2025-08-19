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
	SetSessionConcurrencyPerInstance(v int64) *HeaderFieldSessionAffinityConfig
	GetSessionConcurrencyPerInstance() *int64
	SetSessionIdleTimeoutInSeconds(v int64) *HeaderFieldSessionAffinityConfig
	GetSessionIdleTimeoutInSeconds() *int64
	SetSessionTTLInSeconds(v int64) *HeaderFieldSessionAffinityConfig
	GetSessionTTLInSeconds() *int64
}

type HeaderFieldSessionAffinityConfig struct {
	AffinityHeaderFieldName       *string `json:"affinityHeaderFieldName,omitempty" xml:"affinityHeaderFieldName,omitempty"`
	SessionConcurrencyPerInstance *int64  `json:"sessionConcurrencyPerInstance,omitempty" xml:"sessionConcurrencyPerInstance,omitempty"`
	SessionIdleTimeoutInSeconds   *int64  `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	SessionTTLInSeconds           *int64  `json:"sessionTTLInSeconds,omitempty" xml:"sessionTTLInSeconds,omitempty"`
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
