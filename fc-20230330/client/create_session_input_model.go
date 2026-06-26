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
	SetJuiceFsConfig(v *JuiceFsConfig) *CreateSessionInput
	GetJuiceFsConfig() *JuiceFsConfig
	SetNasConfig(v *NASConfig) *CreateSessionInput
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *CreateSessionInput
	GetOssMountConfig() *OSSMountConfig
	SetPolarFsConfig(v *PolarFsConfig) *CreateSessionInput
	GetPolarFsConfig() *PolarFsConfig
	SetSessionId(v string) *CreateSessionInput
	GetSessionId() *string
	SetSessionIdleTimeoutInSeconds(v int64) *CreateSessionInput
	GetSessionIdleTimeoutInSeconds() *int64
	SetSessionTTLInSeconds(v int64) *CreateSessionInput
	GetSessionTTLInSeconds() *int64
}

type CreateSessionInput struct {
	// A value of false (the default) allows an expired session ID to be reused for a new session, which the system then binds to a new instance. If set to true, an expired session ID cannot be reused.
	//
	// example:
	//
	// false
	DisableSessionIdReuse *bool          `json:"disableSessionIdReuse,omitempty" xml:"disableSessionIdReuse,omitempty"`
	JuiceFsConfig         *JuiceFsConfig `json:"juiceFsConfig,omitempty" xml:"juiceFsConfig,omitempty"`
	// Allows instances in the session to access specified NAS resources.
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// Allows instances in the session to access specified OSS resources.
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// Allows instances in the session to access specified PolarFS resources.
	PolarFsConfig *PolarFsConfig `json:"polarFsConfig,omitempty" xml:"polarFsConfig,omitempty"`
	// A customizable session ID. If you do not specify a value, the server generates one. This parameter applies only to the HEADER_FIELD affinity mode. The value must be 0 to 64 characters long. The first character must be a character in **a-zA-Z0-9_**. Subsequent characters can be any character in **a-zA-Z0-9_-**.
	//
	// example:
	//
	// custom-test-session-id
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The session idle timeout in seconds.
	//
	// example:
	//
	// 1800
	SessionIdleTimeoutInSeconds *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	// The session lifetime in seconds.
	//
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

func (s *CreateSessionInput) GetJuiceFsConfig() *JuiceFsConfig {
	return s.JuiceFsConfig
}

func (s *CreateSessionInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *CreateSessionInput) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *CreateSessionInput) GetPolarFsConfig() *PolarFsConfig {
	return s.PolarFsConfig
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

func (s *CreateSessionInput) SetJuiceFsConfig(v *JuiceFsConfig) *CreateSessionInput {
	s.JuiceFsConfig = v
	return s
}

func (s *CreateSessionInput) SetNasConfig(v *NASConfig) *CreateSessionInput {
	s.NasConfig = v
	return s
}

func (s *CreateSessionInput) SetOssMountConfig(v *OSSMountConfig) *CreateSessionInput {
	s.OssMountConfig = v
	return s
}

func (s *CreateSessionInput) SetPolarFsConfig(v *PolarFsConfig) *CreateSessionInput {
	s.PolarFsConfig = v
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
	if s.JuiceFsConfig != nil {
		if err := s.JuiceFsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.NasConfig != nil {
		if err := s.NasConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OssMountConfig != nil {
		if err := s.OssMountConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PolarFsConfig != nil {
		if err := s.PolarFsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
