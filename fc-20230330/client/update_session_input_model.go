// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionInput interface {
	dara.Model
	String() string
	GoString() string
	SetDisableSessionIdReuse(v bool) *UpdateSessionInput
	GetDisableSessionIdReuse() *bool
	SetNasConfig(v *NASConfig) *UpdateSessionInput
	GetNasConfig() *NASConfig
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
	DisableSessionIdReuse *bool           `json:"disableSessionIdReuse,omitempty" xml:"disableSessionIdReuse,omitempty"`
	NasConfig             *NASConfig      `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	OssMountConfig        *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	PolarFsConfig         *PolarFsConfig  `json:"polarFsConfig,omitempty" xml:"polarFsConfig,omitempty"`
	// example:
	//
	// 1800
	SessionIdleTimeoutInSeconds *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
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

func (s *UpdateSessionInput) GetDisableSessionIdReuse() *bool {
	return s.DisableSessionIdReuse
}

func (s *UpdateSessionInput) GetNasConfig() *NASConfig {
	return s.NasConfig
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

func (s *UpdateSessionInput) SetDisableSessionIdReuse(v bool) *UpdateSessionInput {
	s.DisableSessionIdReuse = &v
	return s
}

func (s *UpdateSessionInput) SetNasConfig(v *NASConfig) *UpdateSessionInput {
	s.NasConfig = v
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
