// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSidecarContainerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAcrInstanceId(v string) *SidecarContainerConfig
	GetAcrInstanceId() *string
	SetCommand(v string) *SidecarContainerConfig
	GetCommand() *string
	SetCommandArgs(v string) *SidecarContainerConfig
	GetCommandArgs() *string
	SetConfigMapMountDesc(v string) *SidecarContainerConfig
	GetConfigMapMountDesc() *string
	SetCpu(v int32) *SidecarContainerConfig
	GetCpu() *int32
	SetEmptyDirDesc(v string) *SidecarContainerConfig
	GetEmptyDirDesc() *string
	SetEnvs(v string) *SidecarContainerConfig
	GetEnvs() *string
	SetImageUrl(v string) *SidecarContainerConfig
	GetImageUrl() *string
	SetMemory(v int32) *SidecarContainerConfig
	GetMemory() *int32
	SetName(v string) *SidecarContainerConfig
	GetName() *string
}

type SidecarContainerConfig struct {
	AcrInstanceId      *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	Command            *string `json:"Command,omitempty" xml:"Command,omitempty"`
	CommandArgs        *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	ConfigMapMountDesc *string `json:"ConfigMapMountDesc,omitempty" xml:"ConfigMapMountDesc,omitempty"`
	Cpu                *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EmptyDirDesc       *string `json:"EmptyDirDesc,omitempty" xml:"EmptyDirDesc,omitempty"`
	Envs               *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	ImageUrl           *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Memory             *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SidecarContainerConfig) String() string {
	return dara.Prettify(s)
}

func (s SidecarContainerConfig) GoString() string {
	return s.String()
}

func (s *SidecarContainerConfig) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *SidecarContainerConfig) GetCommand() *string {
	return s.Command
}

func (s *SidecarContainerConfig) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *SidecarContainerConfig) GetConfigMapMountDesc() *string {
	return s.ConfigMapMountDesc
}

func (s *SidecarContainerConfig) GetCpu() *int32 {
	return s.Cpu
}

func (s *SidecarContainerConfig) GetEmptyDirDesc() *string {
	return s.EmptyDirDesc
}

func (s *SidecarContainerConfig) GetEnvs() *string {
	return s.Envs
}

func (s *SidecarContainerConfig) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SidecarContainerConfig) GetMemory() *int32 {
	return s.Memory
}

func (s *SidecarContainerConfig) GetName() *string {
	return s.Name
}

func (s *SidecarContainerConfig) SetAcrInstanceId(v string) *SidecarContainerConfig {
	s.AcrInstanceId = &v
	return s
}

func (s *SidecarContainerConfig) SetCommand(v string) *SidecarContainerConfig {
	s.Command = &v
	return s
}

func (s *SidecarContainerConfig) SetCommandArgs(v string) *SidecarContainerConfig {
	s.CommandArgs = &v
	return s
}

func (s *SidecarContainerConfig) SetConfigMapMountDesc(v string) *SidecarContainerConfig {
	s.ConfigMapMountDesc = &v
	return s
}

func (s *SidecarContainerConfig) SetCpu(v int32) *SidecarContainerConfig {
	s.Cpu = &v
	return s
}

func (s *SidecarContainerConfig) SetEmptyDirDesc(v string) *SidecarContainerConfig {
	s.EmptyDirDesc = &v
	return s
}

func (s *SidecarContainerConfig) SetEnvs(v string) *SidecarContainerConfig {
	s.Envs = &v
	return s
}

func (s *SidecarContainerConfig) SetImageUrl(v string) *SidecarContainerConfig {
	s.ImageUrl = &v
	return s
}

func (s *SidecarContainerConfig) SetMemory(v int32) *SidecarContainerConfig {
	s.Memory = &v
	return s
}

func (s *SidecarContainerConfig) SetName(v string) *SidecarContainerConfig {
	s.Name = &v
	return s
}

func (s *SidecarContainerConfig) Validate() error {
	return dara.Validate(s)
}
