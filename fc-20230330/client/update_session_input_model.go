// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionInput interface {
	dara.Model
	String() string
	GoString() string
	SetAllowInternetAccess(v bool) *UpdateSessionInput
	GetAllowInternetAccess() *bool
	SetDisableSessionIdReuse(v bool) *UpdateSessionInput
	GetDisableSessionIdReuse() *bool
	SetEnableAutoPause(v bool) *UpdateSessionInput
	GetEnableAutoPause() *bool
	SetEnableAutoResume(v bool) *UpdateSessionInput
	GetEnableAutoResume() *bool
	SetJuiceFsConfig(v *JuiceFsConfig) *UpdateSessionInput
	GetJuiceFsConfig() *JuiceFsConfig
	SetNasConfig(v *NASConfig) *UpdateSessionInput
	GetNasConfig() *NASConfig
	SetNetwork(v *UpdateSessionNetworkConfig) *UpdateSessionInput
	GetNetwork() *UpdateSessionNetworkConfig
	SetOssMountConfig(v *OSSMountConfig) *UpdateSessionInput
	GetOssMountConfig() *OSSMountConfig
	SetPolarFsConfig(v *PolarFsConfig) *UpdateSessionInput
	GetPolarFsConfig() *PolarFsConfig
	SetSessionIdleTimeoutInSeconds(v int64) *UpdateSessionInput
	GetSessionIdleTimeoutInSeconds() *int64
	SetSessionTTLInSeconds(v int64) *UpdateSessionInput
	GetSessionTTLInSeconds() *int64
}

type UpdateSessionInput struct {
	AllowInternetAccess *bool `json:"allowInternetAccess,omitempty" xml:"allowInternetAccess,omitempty"`
	// Specifies whether to disable session ID reuse after the session expires. Default value: False, which indicates that after a session expires, you can use the same session ID to initiate requests. The system treats the request as a new session and binds it to a new instance. If you set this parameter to True, the session ID cannot be reused after the session expires.
	//
	// example:
	//
	// false
	DisableSessionIdReuse *bool                       `json:"disableSessionIdReuse,omitempty" xml:"disableSessionIdReuse,omitempty"`
	EnableAutoPause       *bool                       `json:"enableAutoPause,omitempty" xml:"enableAutoPause,omitempty"`
	EnableAutoResume      *bool                       `json:"enableAutoResume,omitempty" xml:"enableAutoResume,omitempty"`
	JuiceFsConfig         *JuiceFsConfig              `json:"juiceFsConfig,omitempty" xml:"juiceFsConfig,omitempty"`
	NasConfig             *NASConfig                  `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	Network               *UpdateSessionNetworkConfig `json:"network,omitempty" xml:"network,omitempty"`
	OssMountConfig        *OSSMountConfig             `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	PolarFsConfig         *PolarFsConfig              `json:"polarFsConfig,omitempty" xml:"polarFsConfig,omitempty"`
	// The session idle timeout.
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

func (s UpdateSessionInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionInput) GoString() string {
	return s.String()
}

func (s *UpdateSessionInput) GetAllowInternetAccess() *bool {
	return s.AllowInternetAccess
}

func (s *UpdateSessionInput) GetDisableSessionIdReuse() *bool {
	return s.DisableSessionIdReuse
}

func (s *UpdateSessionInput) GetEnableAutoPause() *bool {
	return s.EnableAutoPause
}

func (s *UpdateSessionInput) GetEnableAutoResume() *bool {
	return s.EnableAutoResume
}

func (s *UpdateSessionInput) GetJuiceFsConfig() *JuiceFsConfig {
	return s.JuiceFsConfig
}

func (s *UpdateSessionInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *UpdateSessionInput) GetNetwork() *UpdateSessionNetworkConfig {
	return s.Network
}

func (s *UpdateSessionInput) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *UpdateSessionInput) GetPolarFsConfig() *PolarFsConfig {
	return s.PolarFsConfig
}

func (s *UpdateSessionInput) GetSessionIdleTimeoutInSeconds() *int64 {
	return s.SessionIdleTimeoutInSeconds
}

func (s *UpdateSessionInput) GetSessionTTLInSeconds() *int64 {
	return s.SessionTTLInSeconds
}

func (s *UpdateSessionInput) SetAllowInternetAccess(v bool) *UpdateSessionInput {
	s.AllowInternetAccess = &v
	return s
}

func (s *UpdateSessionInput) SetDisableSessionIdReuse(v bool) *UpdateSessionInput {
	s.DisableSessionIdReuse = &v
	return s
}

func (s *UpdateSessionInput) SetEnableAutoPause(v bool) *UpdateSessionInput {
	s.EnableAutoPause = &v
	return s
}

func (s *UpdateSessionInput) SetEnableAutoResume(v bool) *UpdateSessionInput {
	s.EnableAutoResume = &v
	return s
}

func (s *UpdateSessionInput) SetJuiceFsConfig(v *JuiceFsConfig) *UpdateSessionInput {
	s.JuiceFsConfig = v
	return s
}

func (s *UpdateSessionInput) SetNasConfig(v *NASConfig) *UpdateSessionInput {
	s.NasConfig = v
	return s
}

func (s *UpdateSessionInput) SetNetwork(v *UpdateSessionNetworkConfig) *UpdateSessionInput {
	s.Network = v
	return s
}

func (s *UpdateSessionInput) SetOssMountConfig(v *OSSMountConfig) *UpdateSessionInput {
	s.OssMountConfig = v
	return s
}

func (s *UpdateSessionInput) SetPolarFsConfig(v *PolarFsConfig) *UpdateSessionInput {
	s.PolarFsConfig = v
	return s
}

func (s *UpdateSessionInput) SetSessionIdleTimeoutInSeconds(v int64) *UpdateSessionInput {
	s.SessionIdleTimeoutInSeconds = &v
	return s
}

func (s *UpdateSessionInput) SetSessionTTLInSeconds(v int64) *UpdateSessionInput {
	s.SessionTTLInSeconds = &v
	return s
}

func (s *UpdateSessionInput) Validate() error {
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
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
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
