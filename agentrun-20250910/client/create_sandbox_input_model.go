// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSandboxInput interface {
	dara.Model
	String() string
	GoString() string
	SetNasConfig(v *NASConfig) *CreateSandboxInput
	GetNasConfig() *NASConfig
	SetOssMountConfig(v *OSSMountConfig) *CreateSandboxInput
	GetOssMountConfig() *OSSMountConfig
	SetPolarFsConfig(v *PolarFsConfig) *CreateSandboxInput
	GetPolarFsConfig() *PolarFsConfig
	SetSandboxId(v string) *CreateSandboxInput
	GetSandboxId() *string
	SetSandboxIdleTimeoutInSeconds(v int32) *CreateSandboxInput
	GetSandboxIdleTimeoutInSeconds() *int32
	SetSandboxIdleTimeoutSeconds(v int32) *CreateSandboxInput
	GetSandboxIdleTimeoutSeconds() *int32
	SetTemplateName(v string) *CreateSandboxInput
	GetTemplateName() *string
}

type CreateSandboxInput struct {
	// The NAS configuration.
	NasConfig *NASConfig `json:"nasConfig,omitempty" xml:"nasConfig,omitempty"`
	// The OSS mount configuration.
	OssMountConfig *OSSMountConfig `json:"ossMountConfig,omitempty" xml:"ossMountConfig,omitempty"`
	// The PolarFS configuration.
	PolarFsConfig *PolarFsConfig `json:"polarFsConfig,omitempty" xml:"polarFsConfig,omitempty"`
	// The unique identifier of the sandbox instance.
	//
	// example:
	//
	// 01KE8DAJ35JC8SKP9CNFRZ8CW7
	SandboxId *string `json:"sandboxId,omitempty" xml:"sandboxId,omitempty"`
	// The idle timeout for the sandbox, in seconds.
	SandboxIdleTimeoutInSeconds *int32 `json:"sandboxIdleTimeoutInSeconds,omitempty" xml:"sandboxIdleTimeoutInSeconds,omitempty"`
	// Deprecated
	//
	// The idle timeout for the sandbox, in seconds.
	//
	// example:
	//
	// 1800
	SandboxIdleTimeoutSeconds *int32 `json:"sandboxIdleTimeoutSeconds,omitempty" xml:"sandboxIdleTimeoutSeconds,omitempty"`
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// templateName
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s CreateSandboxInput) String() string {
	return dara.Prettify(s)
}

func (s CreateSandboxInput) GoString() string {
	return s.String()
}

func (s *CreateSandboxInput) GetNasConfig() *NASConfig {
	return s.NasConfig
}

func (s *CreateSandboxInput) GetOssMountConfig() *OSSMountConfig {
	return s.OssMountConfig
}

func (s *CreateSandboxInput) GetPolarFsConfig() *PolarFsConfig {
	return s.PolarFsConfig
}

func (s *CreateSandboxInput) GetSandboxId() *string {
	return s.SandboxId
}

func (s *CreateSandboxInput) GetSandboxIdleTimeoutInSeconds() *int32 {
	return s.SandboxIdleTimeoutInSeconds
}

func (s *CreateSandboxInput) GetSandboxIdleTimeoutSeconds() *int32 {
	return s.SandboxIdleTimeoutSeconds
}

func (s *CreateSandboxInput) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateSandboxInput) SetNasConfig(v *NASConfig) *CreateSandboxInput {
	s.NasConfig = v
	return s
}

func (s *CreateSandboxInput) SetOssMountConfig(v *OSSMountConfig) *CreateSandboxInput {
	s.OssMountConfig = v
	return s
}

func (s *CreateSandboxInput) SetPolarFsConfig(v *PolarFsConfig) *CreateSandboxInput {
	s.PolarFsConfig = v
	return s
}

func (s *CreateSandboxInput) SetSandboxId(v string) *CreateSandboxInput {
	s.SandboxId = &v
	return s
}

func (s *CreateSandboxInput) SetSandboxIdleTimeoutInSeconds(v int32) *CreateSandboxInput {
	s.SandboxIdleTimeoutInSeconds = &v
	return s
}

func (s *CreateSandboxInput) SetSandboxIdleTimeoutSeconds(v int32) *CreateSandboxInput {
	s.SandboxIdleTimeoutSeconds = &v
	return s
}

func (s *CreateSandboxInput) SetTemplateName(v string) *CreateSandboxInput {
	s.TemplateName = &v
	return s
}

func (s *CreateSandboxInput) Validate() error {
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
