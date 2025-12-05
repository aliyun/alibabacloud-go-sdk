// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionInput interface {
	dara.Model
	String() string
	GoString() string
	SetDisableSessionIdReuse(v bool) *CreateSessionInput
	GetDisableSessionIdReuse() *bool
	SetNasConfig(v *NASConfig) *CreateSessionInput
	GetNasConfig() *NASConfig
	SetSessionId(v string) *CreateSessionInput
	GetSessionId() *string
	SetSessionIdleTimeoutInSeconds(v int64) *CreateSessionInput
	GetSessionIdleTimeoutInSeconds() *int64
	SetSessionTTLInSeconds(v int64) *CreateSessionInput
	GetSessionTTLInSeconds() *int64
}

type CreateSessionInput struct {
	DisableSessionIdReuse *bool      `json:"disableSessionIdReuse,omitempty" xml:"disableSessionIdReuse,omitempty"`
	NasConfig             *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// example:
	//
	// custom-test-session-id
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 1800
	SessionIdleTimeoutInSeconds *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	// example:
	//
	// 21600
	SessionTTLInSeconds *int64 `json:"sessionTTLInSeconds,omitempty" xml:"sessionTTLInSeconds,omitempty"`
}

func (s CreateSessionInput) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionInput) GoString() string {
	return s.String()
}

func (s *CreateSessionInput) GetDisableSessionIdReuse() *bool {
	return s.DisableSessionIdReuse
}

func (s *CreateSessionInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *CreateSessionInput) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateSessionInput) GetSessionIdleTimeoutInSeconds() *int64 {
	return s.SessionIdleTimeoutInSeconds
}

func (s *CreateSessionInput) GetSessionTTLInSeconds() *int64 {
	return s.SessionTTLInSeconds
}

func (s *CreateSessionInput) SetDisableSessionIdReuse(v bool) *CreateSessionInput {
	s.DisableSessionIdReuse = &v
	return s
}

func (s *CreateSessionInput) SetNasConfig(v *NASConfig) *CreateSessionInput {
	s.NasConfig = v
	return s
}

func (s *CreateSessionInput) SetSessionId(v string) *CreateSessionInput {
	s.SessionId = &v
	return s
}

func (s *CreateSessionInput) SetSessionIdleTimeoutInSeconds(v int64) *CreateSessionInput {
	s.SessionIdleTimeoutInSeconds = &v
	return s
}

func (s *CreateSessionInput) SetSessionTTLInSeconds(v int64) *CreateSessionInput {
	s.SessionTTLInSeconds = &v
	return s
}

func (s *CreateSessionInput) Validate() error {
	if s.NasConfig != nil {
		if err := s.NasConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
