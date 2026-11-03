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
	// The log path.
	//
	// example:
	//
	// /test
	LogPath *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	// The log type. The following types of logs are supported:
	//
	// 	- File collection logs
	//
	// 	- Standard output logs
	//
	// example:
	//
	// file_log
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The name of the Logstore. The name must meet the following requirements:
	//
	// 	- The name must be unique in a project.
	//
	// 	- The name can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must start and end with a lowercase letter or a digit.
	//
	// 	- The name must be 3 to 63 characters in length.
	//
	// example:
	//
	// sag-shanghai
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	// The name for the Logtail configuration.
	//
	// example:
	//
	// ******-access-log-collector
	LogtailName *string `json:"LogtailName,omitempty" xml:"LogtailName,omitempty"`
	// The name of the machine group of Simple Log Service.
	//
	// example:
	//
	// Log Service Group
	MachineGroup *string `json:"MachineGroup,omitempty" xml:"MachineGroup,omitempty"`
	// The name of the SLS project.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
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
