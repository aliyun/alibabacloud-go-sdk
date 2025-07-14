// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitContainerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCommand(v string) *InitContainerConfig
	GetCommand() *string
	SetCommandArgs(v string) *InitContainerConfig
	GetCommandArgs() *string
	SetConfigMapMountDesc(v string) *InitContainerConfig
	GetConfigMapMountDesc() *string
	SetEnvs(v string) *InitContainerConfig
	GetEnvs() *string
	SetImageUrl(v string) *InitContainerConfig
	GetImageUrl() *string
	SetName(v string) *InitContainerConfig
	GetName() *string
}

type InitContainerConfig struct {
	Command            *string `json:"Command,omitempty" xml:"Command,omitempty"`
	CommandArgs        *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	Envs               *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	ImageUrl           *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s InitContainerConfig) String() string {
	return dara.Prettify(s)
}

func (s InitContainerConfig) GoString() string {
	return s.String()
}

func (s *InitContainerConfig) GetCommand() *string {
	return s.Command
}

func (s *InitContainerConfig) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *InitContainerConfig) GetConfigMapMountDesc() *string {
	return s.ConfigMapMountDesc
}

func (s *InitContainerConfig) GetEnvs() *string {
	return s.Envs
}

func (s *InitContainerConfig) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *InitContainerConfig) GetName() *string {
	return s.Name
}

func (s *InitContainerConfig) SetCommand(v string) *InitContainerConfig {
	s.Command = &v
	return s
}

func (s *InitContainerConfig) SetCommandArgs(v string) *InitContainerConfig {
	s.CommandArgs = &v
	return s
}

func (s *InitContainerConfig) SetConfigMapMountDesc(v string) *InitContainerConfig {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *InitContainerConfig) SetEnvs(v string) *InitContainerConfig {
	s.Envs = &v
	return s
}

func (s *InitContainerConfig) SetImageUrl(v string) *InitContainerConfig {
	s.ImageUrl = &v
	return s
}

func (s *InitContainerConfig) SetName(v string) *InitContainerConfig {
	s.Name = &v
	return s
}

func (s *InitContainerConfig) Validate() error {
	return dara.Validate(s)
}
