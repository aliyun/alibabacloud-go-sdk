// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPServerInstallationConfig interface {
	dara.Model
	String() string
	GoString() string
	SetArgs(v string) *MCPServerInstallationConfig
	GetArgs() *string
	SetCommand(v string) *MCPServerInstallationConfig
	GetCommand() *string
	SetEnv(v map[string]*string) *MCPServerInstallationConfig
	GetEnv() map[string]*string
	SetTransportType(v string) *MCPServerInstallationConfig
	GetTransportType() *string
	SetUrl(v string) *MCPServerInstallationConfig
	GetUrl() *string
}

type MCPServerInstallationConfig struct {
	Args          *string            `json:"args,omitempty" xml:"args,omitempty"`
	Command       *string            `json:"command,omitempty" xml:"command,omitempty"`
	Env           map[string]*string `json:"env,omitempty" xml:"env,omitempty"`
	TransportType *string            `json:"transportType,omitempty" xml:"transportType,omitempty"`
	Url           *string            `json:"url,omitempty" xml:"url,omitempty"`
}

func (s MCPServerInstallationConfig) String() string {
	return dara.Prettify(s)
}

func (s MCPServerInstallationConfig) GoString() string {
	return s.String()
}

func (s *MCPServerInstallationConfig) GetArgs() *string {
	return s.Args
}

func (s *MCPServerInstallationConfig) GetCommand() *string {
	return s.Command
}

func (s *MCPServerInstallationConfig) GetEnv() map[string]*string {
	return s.Env
}

func (s *MCPServerInstallationConfig) GetTransportType() *string {
	return s.TransportType
}

func (s *MCPServerInstallationConfig) GetUrl() *string {
	return s.Url
}

func (s *MCPServerInstallationConfig) SetArgs(v string) *MCPServerInstallationConfig {
	s.Args = &v
	return s
}

func (s *MCPServerInstallationConfig) SetCommand(v string) *MCPServerInstallationConfig {
	s.Command = &v
	return s
}

func (s *MCPServerInstallationConfig) SetEnv(v map[string]*string) *MCPServerInstallationConfig {
	s.Env = v
	return s
}

func (s *MCPServerInstallationConfig) SetTransportType(v string) *MCPServerInstallationConfig {
	s.TransportType = &v
	return s
}

func (s *MCPServerInstallationConfig) SetUrl(v string) *MCPServerInstallationConfig {
	s.Url = &v
	return s
}

func (s *MCPServerInstallationConfig) Validate() error {
	return dara.Validate(s)
}
