// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSLSCollectConfig interface {
	dara.Model
	String() string
	GoString() string
	SetLogPath(v string) *SLSCollectConfig
	GetLogPath() *string
	SetLogType(v string) *SLSCollectConfig
	GetLogType() *string
	SetLogstoreName(v string) *SLSCollectConfig
	GetLogstoreName() *string
	SetLogtailName(v string) *SLSCollectConfig
	GetLogtailName() *string
	SetMachineGroup(v string) *SLSCollectConfig
	GetMachineGroup() *string
	SetProjectName(v string) *SLSCollectConfig
	GetProjectName() *string
}

type SLSCollectConfig struct {
	LogPath      *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogType      *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	LogtailName  *string `json:"LogtailName,omitempty" xml:"LogtailName,omitempty"`
	MachineGroup *string `json:"MachineGroup,omitempty" xml:"MachineGroup,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s SLSCollectConfig) String() string {
	return dara.Prettify(s)
}

func (s SLSCollectConfig) GoString() string {
	return s.String()
}

func (s *SLSCollectConfig) GetLogPath() *string {
	return s.LogPath
}

func (s *SLSCollectConfig) GetLogType() *string {
	return s.LogType
}

func (s *SLSCollectConfig) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *SLSCollectConfig) GetLogtailName() *string {
	return s.LogtailName
}

func (s *SLSCollectConfig) GetMachineGroup() *string {
	return s.MachineGroup
}

func (s *SLSCollectConfig) GetProjectName() *string {
	return s.ProjectName
}

func (s *SLSCollectConfig) SetLogPath(v string) *SLSCollectConfig {
	s.LogPath = &v
	return s
}

func (s *SLSCollectConfig) SetLogType(v string) *SLSCollectConfig {
	s.LogType = &v
	return s
}

func (s *SLSCollectConfig) SetLogstoreName(v string) *SLSCollectConfig {
	s.LogstoreName = &v
	return s
}

func (s *SLSCollectConfig) SetLogtailName(v string) *SLSCollectConfig {
	s.LogtailName = &v
	return s
}

func (s *SLSCollectConfig) SetMachineGroup(v string) *SLSCollectConfig {
	s.MachineGroup = &v
	return s
}

func (s *SLSCollectConfig) SetProjectName(v string) *SLSCollectConfig {
	s.ProjectName = &v
	return s
}

func (s *SLSCollectConfig) Validate() error {
	return dara.Validate(s)
}
