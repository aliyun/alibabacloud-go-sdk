// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomRuntimeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v []*string) *CustomRuntimeConfig
	GetArgs() []*string
	SetCommand(v []*string) *CustomRuntimeConfig
	GetCommand() []*string
	SetHealthCheckConfig(v *CustomHealthCheckConfig) *CustomRuntimeConfig
	GetHealthCheckConfig() *CustomHealthCheckConfig
	SetPort(v int32) *CustomRuntimeConfig
	GetPort() *int32
}

type CustomRuntimeConfig struct {
	// The arguments that are passed to the startup command.
	Args []*string `json:"args" xml:"args" type:"Repeated"`
	// The startup commands.
	Command []*string `json:"command" xml:"command" type:"Repeated"`
	// The custom health check configuration of the function.
	HealthCheckConfig *CustomHealthCheckConfig `json:"healthCheckConfig,omitempty" xml:"healthCheckConfig,omitempty"`
	// The port on which the HTTP server is listening.
	//
	// example:
	//
	// 9000
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s CustomRuntimeConfig) String() string {
	return dara.Prettify(s)
}

func (s CustomRuntimeConfig) GoString() string {
	return s.String()
}

func (s *CustomRuntimeConfig) GetArgs() []*string {
	return s.Args
}

func (s *CustomRuntimeConfig) GetCommand() []*string {
	return s.Command
}

func (s *CustomRuntimeConfig) GetHealthCheckConfig() *CustomHealthCheckConfig {
	return s.HealthCheckConfig
}

func (s *CustomRuntimeConfig) GetPort() *int32 {
	return s.Port
}

func (s *CustomRuntimeConfig) SetArgs(v []*string) *CustomRuntimeConfig {
	s.Args = v
	return s
}

func (s *CustomRuntimeConfig) SetCommand(v []*string) *CustomRuntimeConfig {
	s.Command = v
	return s
}

func (s *CustomRuntimeConfig) SetHealthCheckConfig(v *CustomHealthCheckConfig) *CustomRuntimeConfig {
	s.HealthCheckConfig = v
	return s
}

func (s *CustomRuntimeConfig) SetPort(v int32) *CustomRuntimeConfig {
	s.Port = &v
	return s
}

func (s *CustomRuntimeConfig) Validate() error {
	if s.HealthCheckConfig != nil {
		if err := s.HealthCheckConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
