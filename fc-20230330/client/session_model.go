// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSession interface {
	dara.Model
	String() string
	GoString() string
	SetContainerId(v string) *Session
	GetContainerId() *string
	SetCreatedTime(v string) *Session
	GetCreatedTime() *string
	SetDisableSessionIdReuse(v bool) *Session
	GetDisableSessionIdReuse() *bool
	SetFunctionName(v string) *Session
	GetFunctionName() *string
	SetJuiceFsConfig(v *JuiceFsConfig) *Session
	GetJuiceFsConfig() *JuiceFsConfig
	SetLastModifiedTime(v string) *Session
	GetLastModifiedTime() *string
	SetNasConfig(v *NASConfig) *Session
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *Session
	GetOssMountConfig() *OSSMountConfig
	SetPolarFsConfig(v *PolarFsConfig) *Session
	GetPolarFsConfig() *PolarFsConfig
	SetQualifier(v string) *Session
	GetQualifier() *string
	SetSessionAffinityType(v string) *Session
	GetSessionAffinityType() *string
	SetSessionId(v string) *Session
	GetSessionId() *string
	SetSessionIdleTimeoutInSeconds(v int64) *Session
	GetSessionIdleTimeoutInSeconds() *int64
	SetSessionStatus(v string) *Session
	GetSessionStatus() *string
	SetSessionTTLInSeconds(v int64) *Session
	GetSessionTTLInSeconds() *int64
}

type Session struct {
	// The ID of the function instance associated with the session.
	//
	// example:
	//
	// c-68999e02-16a1955c-d2a03d1ccs
	ContainerId *string `json:"containerId,omitempty" xml:"containerId,omitempty"`
	// The time when the session was created.
	//
	// example:
	//
	// 2025-04-01T08:15:27Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// Specifies whether an expired session ID can be reused. If `true`, an expired session ID cannot be reused. If `false` (the default), sending a request with an expired session ID creates a new session bound to a new instance.
	//
	// example:
	//
	// false
	DisableSessionIdReuse *bool `json:"disableSessionIdReuse,omitempty" xml:"disableSessionIdReuse,omitempty"`
	// The name of the function associated with the session.
	//
	// example:
	//
	// functionName1
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The JuiceFS mount configuration, enabling the associated function instance to access specified JuiceFS resources.
	JuiceFsConfig *JuiceFsConfig `json:"juiceFsConfig,omitempty" xml:"juiceFsConfig,omitempty"`
	// The time when the session was last updated.
	//
	// example:
	//
	// 2025-04-01T18:15:27Z
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// The NAS configuration, enabling the associated function instance to access specified NAS resources.
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// The OSS mount configuration, enabling the associated function instance to access specified OSS resources.
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// The PolarFS mount configuration, enabling the associated function instance to access specified PolarFS resources.
	PolarFsConfig *PolarFsConfig `json:"polarFsConfig,omitempty" xml:"polarFsConfig,omitempty"`
	// The qualifier, which specifies a function version or alias. Defaults to `LATEST` if unspecified.
	//
	// example:
	//
	// AliasName1
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// The type of session affinity.
	//
	// example:
	//
	// HEADER_FIELD
	SessionAffinityType *string `json:"sessionAffinityType,omitempty" xml:"sessionAffinityType,omitempty"`
	// The unique identifier for the function session.
	//
	// example:
	//
	// 81f70ae156904eb9b7d43e12f511fe58
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The maximum duration, in seconds, that the session can be idle before it expires.
	//
	// example:
	//
	// 1800
	SessionIdleTimeoutInSeconds *int64 `json:"sessionIdleTimeoutInSeconds,omitempty" xml:"sessionIdleTimeoutInSeconds,omitempty"`
	// The status of the session. `Active` indicates the session is valid, and `Expired` indicates it is no longer valid.
	//
	// example:
	//
	// Active
	SessionStatus *string `json:"sessionStatus,omitempty" xml:"sessionStatus,omitempty"`
	// The maximum lifespan of the session, in seconds.
	//
	// example:
	//
	// 21600
	SessionTTLInSeconds *int64 `json:"sessionTTLInSeconds,omitempty" xml:"sessionTTLInSeconds,omitempty"`
}

func (s Session) String() string {
	return dara.Prettify(s)
}

func (s Session) GoString() string {
	return s.String()
}

func (s *Session) GetContainerId() *string {
	return s.ContainerId
}

func (s *Session) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Session) GetDisableSessionIdReuse() *bool {
	return s.DisableSessionIdReuse
}

func (s *Session) GetFunctionName() *string {
	return s.FunctionName
}

func (s *Session) GetJuiceFsConfig() *JuiceFsConfig {
	return s.JuiceFsConfig
}

func (s *Session) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *Session) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *Session) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *Session) GetPolarFsConfig() *PolarFsConfig {
	return s.PolarFsConfig
}

func (s *Session) GetQualifier() *string {
	return s.Qualifier
}

func (s *Session) GetSessionAffinityType() *string {
	return s.SessionAffinityType
}

func (s *Session) GetSessionId() *string {
	return s.SessionId
}

func (s *Session) GetSessionIdleTimeoutInSeconds() *int64 {
	return s.SessionIdleTimeoutInSeconds
}

func (s *Session) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *Session) GetSessionTTLInSeconds() *int64 {
	return s.SessionTTLInSeconds
}

func (s *Session) SetContainerId(v string) *Session {
	s.ContainerId = &v
	return s
}

func (s *Session) SetCreatedTime(v string) *Session {
	s.CreatedTime = &v
	return s
}

func (s *Session) SetDisableSessionIdReuse(v bool) *Session {
	s.DisableSessionIdReuse = &v
	return s
}

func (s *Session) SetFunctionName(v string) *Session {
	s.FunctionName = &v
	return s
}

func (s *Session) SetJuiceFsConfig(v *JuiceFsConfig) *Session {
	s.JuiceFsConfig = v
	return s
}

func (s *Session) SetLastModifiedTime(v string) *Session {
	s.LastModifiedTime = &v
	return s
}

func (s *Session) SetNasConfig(v *NASConfig) *Session {
	s.NasConfig = v
	return s
}

func (s *Session) SetOssMountConfig(v *OSSMountConfig) *Session {
	s.OssMountConfig = v
	return s
}

func (s *Session) SetPolarFsConfig(v *PolarFsConfig) *Session {
	s.PolarFsConfig = v
	return s
}

func (s *Session) SetQualifier(v string) *Session {
	s.Qualifier = &v
	return s
}

func (s *Session) SetSessionAffinityType(v string) *Session {
	s.SessionAffinityType = &v
	return s
}

func (s *Session) SetSessionId(v string) *Session {
	s.SessionId = &v
	return s
}

func (s *Session) SetSessionIdleTimeoutInSeconds(v int64) *Session {
	s.SessionIdleTimeoutInSeconds = &v
	return s
}

func (s *Session) SetSessionStatus(v string) *Session {
	s.SessionStatus = &v
	return s
}

func (s *Session) SetSessionTTLInSeconds(v int64) *Session {
	s.SessionTTLInSeconds = &v
	return s
}

func (s *Session) Validate() error {
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
