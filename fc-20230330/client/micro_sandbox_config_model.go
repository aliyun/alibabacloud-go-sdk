// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMicroSandboxConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *MicroSandboxConfig
	GetAcrInstanceId() *string
	SetImage(v string) *MicroSandboxConfig
	GetImage() *string
	SetOsType(v string) *MicroSandboxConfig
	GetOsType() *string
	SetReadyCommand(v string) *MicroSandboxConfig
	GetReadyCommand() *string
	SetRegistryConfig(v *RegistryConfig) *MicroSandboxConfig
	GetRegistryConfig() *RegistryConfig
	SetStartCommand(v string) *MicroSandboxConfig
	GetStartCommand() *string
}

type MicroSandboxConfig struct {
	// The ID of the ACR Enterprise Edition image repository instance. Used in pair with MicroSandbox images. This parameter is optional. If not provided, the server resolves it as needed.
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The image address.
	Image        *string `json:"image,omitempty" xml:"image,omitempty"`
	OsType       *string `json:"osType,omitempty" xml:"osType,omitempty"`
	ReadyCommand *string `json:"readyCommand,omitempty" xml:"readyCommand,omitempty"`
	// The image repository configuration.
	RegistryConfig *RegistryConfig `json:"registryConfig,omitempty" xml:"registryConfig,omitempty"`
	StartCommand   *string         `json:"startCommand,omitempty" xml:"startCommand,omitempty"`
}

func (s MicroSandboxConfig) String() string {
	return dara.Prettify(s)
}

func (s MicroSandboxConfig) GoString() string {
	return s.String()
}

func (s *MicroSandboxConfig) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *MicroSandboxConfig) GetImage() *string {
	return s.Image
}

func (s *MicroSandboxConfig) GetOsType() *string {
	return s.OsType
}

func (s *MicroSandboxConfig) GetReadyCommand() *string {
	return s.ReadyCommand
}

func (s *MicroSandboxConfig) GetRegistryConfig() *RegistryConfig {
	return s.RegistryConfig
}

func (s *MicroSandboxConfig) GetStartCommand() *string {
	return s.StartCommand
}

func (s *MicroSandboxConfig) SetAcrInstanceId(v string) *MicroSandboxConfig {
	s.AcrInstanceId = &v
	return s
}

func (s *MicroSandboxConfig) SetImage(v string) *MicroSandboxConfig {
	s.Image = &v
	return s
}

func (s *MicroSandboxConfig) SetOsType(v string) *MicroSandboxConfig {
	s.OsType = &v
	return s
}

func (s *MicroSandboxConfig) SetReadyCommand(v string) *MicroSandboxConfig {
	s.ReadyCommand = &v
	return s
}

func (s *MicroSandboxConfig) SetRegistryConfig(v *RegistryConfig) *MicroSandboxConfig {
	s.RegistryConfig = v
	return s
}

func (s *MicroSandboxConfig) SetStartCommand(v string) *MicroSandboxConfig {
	s.StartCommand = &v
	return s
}

func (s *MicroSandboxConfig) Validate() error {
	if s.RegistryConfig != nil {
		if err := s.RegistryConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
