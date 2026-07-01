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
	SetEnableAutoPause(v bool) *CreateSessionInput
	GetEnableAutoPause() *bool
	SetEnableAutoResume(v bool) *CreateSessionInput
	GetEnableAutoResume() *bool
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
	// Specifies whether to disable session ID reuse. Default value: False, which indicates that after a session with a specific SessionID expires, you can send requests with the same SessionID, and the system treats it as a new session bound to a new instance. If this parameter is set to True, the SessionID cannot be reused after the session expires.
	//
	// example:
	//
	// false
	DisableSessionIdReuse *bool          `json:"disableSessionIdReuse,omitempty" xml:"disableSessionIdReuse,omitempty"`
	EnableAutoPause       *bool          `json:"enableAutoPause,omitempty" xml:"enableAutoPause,omitempty"`
	EnableAutoResume      *bool          `json:"enableAutoResume,omitempty" xml:"enableAutoResume,omitempty"`
	JuiceFsConfig         *JuiceFsConfig `json:"juiceFsConfig,omitempty" xml:"juiceFsConfig,omitempty"`
	// The NAS configuration. After this parameter is configured, instances associated with the session can access the specified NAS resources.
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// The OSS configuration. After this parameter is configured, instances associated with the session can access the specified OSS resources.
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// The PolarFs configuration. After this parameter is configured, instances associated with the session can access the specified PolarFs resources.
	PolarFsConfig *PolarFsConfig `json:"polarFsConfig,omitempty" xml:"polarFsConfig,omitempty"`
	// The custom session ID. If this parameter is not specified, the server generates a session ID. If specified, the value is used as the session ID. This parameter applies only to the HEADER_FIELD affinity mode. Format: the length is limited to [0,64]. The first character must be from **a-zA-Z0-9_**, and subsequent characters can be from **a-zA-Z0-9_-**.
	//
	// example:
	//
	// custom-test-session-id
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The session idle timeout period.
	//
	// example:
	//
	// 1800
	SessionIdleTimeoutInSeconds *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	// The session lifetime.
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

func (s *CreateSessionInput) GetEnableAutoPause() *bool {
	return s.EnableAutoPause
}

func (s *CreateSessionInput) GetEnableAutoResume() *bool {
	return s.EnableAutoResume
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

func (s *CreateSessionInput) SetEnableAutoPause(v bool) *CreateSessionInput {
	s.EnableAutoPause = &v
	return s
}

func (s *CreateSessionInput) SetEnableAutoResume(v bool) *CreateSessionInput {
	s.EnableAutoResume = &v
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
