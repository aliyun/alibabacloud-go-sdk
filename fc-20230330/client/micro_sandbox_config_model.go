// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMicroSandboxConfig interface {
	dara.Model
	String() string
	GoString() string
	SetOsType(v string) *MicroSandboxConfig
	GetOsType() *string
	SetReadyCommand(v string) *MicroSandboxConfig
	GetReadyCommand() *string
	SetStartCommand(v string) *MicroSandboxConfig
	GetStartCommand() *string
}

type MicroSandboxConfig struct {
	OsType       *string `json:"osType,omitempty" xml:"osType,omitempty"`
	ReadyCommand *string `json:"readyCommand,omitempty" xml:"readyCommand,omitempty"`
	StartCommand *string `json:"startCommand,omitempty" xml:"startCommand,omitempty"`
}

func (s MicroSandboxConfig) String() string {
	return dara.Prettify(s)
}

func (s MicroSandboxConfig) GoString() string {
	return s.String()
}

func (s *MicroSandboxConfig) GetOsType() *string {
	return s.OsType
}

func (s *MicroSandboxConfig) GetReadyCommand() *string {
	return s.ReadyCommand
}

func (s *MicroSandboxConfig) GetStartCommand() *string {
	return s.StartCommand
}

func (s *MicroSandboxConfig) SetOsType(v string) *MicroSandboxConfig {
	s.OsType = &v
	return s
}

func (s *MicroSandboxConfig) SetReadyCommand(v string) *MicroSandboxConfig {
	s.ReadyCommand = &v
	return s
}

func (s *MicroSandboxConfig) SetStartCommand(v string) *MicroSandboxConfig {
	s.StartCommand = &v
	return s
}

func (s *MicroSandboxConfig) Validate() error {
	return dara.Validate(s)
}
