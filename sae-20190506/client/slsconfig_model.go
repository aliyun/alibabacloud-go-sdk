// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSLSConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCollectConfigs(v []*SLSConfigCollectConfigs) *SLSConfig
	GetCollectConfigs() []*SLSConfigCollectConfigs
}

type SLSConfig struct {
	CollectConfigs []*SLSConfigCollectConfigs `json:"collectConfigs,omitempty" xml:"collectConfigs,omitempty" type:"Repeated"`
}

func (s SLSConfig) String() string {
	return dara.Prettify(s)
}

func (s SLSConfig) GoString() string {
	return s.String()
}

func (s *SLSConfig) GetCollectConfigs() []*SLSConfigCollectConfigs {
	return s.CollectConfigs
}

func (s *SLSConfig) SetCollectConfigs(v []*SLSConfigCollectConfigs) *SLSConfig {
	s.CollectConfigs = v
	return s
}

func (s *SLSConfig) Validate() error {
	return dara.Validate(s)
}

type SLSConfigCollectConfigs struct {
	LogPath      *string `json:"logPath,omitempty" xml:"logPath,omitempty"`
	LogType      *string `json:"logType,omitempty" xml:"logType,omitempty"`
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	LogtailName  *string `json:"logtailName,omitempty" xml:"logtailName,omitempty"`
	ProjectName  *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s SLSConfigCollectConfigs) String() string {
	return dara.Prettify(s)
}

func (s SLSConfigCollectConfigs) GoString() string {
	return s.String()
}

func (s *SLSConfigCollectConfigs) GetLogPath() *string {
	return s.LogPath
}

func (s *SLSConfigCollectConfigs) GetLogType() *string {
	return s.LogType
}

func (s *SLSConfigCollectConfigs) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *SLSConfigCollectConfigs) GetLogtailName() *string {
	return s.LogtailName
}

func (s *SLSConfigCollectConfigs) GetProjectName() *string {
	return s.ProjectName
}

func (s *SLSConfigCollectConfigs) SetLogPath(v string) *SLSConfigCollectConfigs {
	s.LogPath = &v
	return s
}

func (s *SLSConfigCollectConfigs) SetLogType(v string) *SLSConfigCollectConfigs {
	s.LogType = &v
	return s
}

func (s *SLSConfigCollectConfigs) SetLogstoreName(v string) *SLSConfigCollectConfigs {
	s.LogstoreName = &v
	return s
}

func (s *SLSConfigCollectConfigs) SetLogtailName(v string) *SLSConfigCollectConfigs {
	s.LogtailName = &v
	return s
}

func (s *SLSConfigCollectConfigs) SetProjectName(v string) *SLSConfigCollectConfigs {
	s.ProjectName = &v
	return s
}

func (s *SLSConfigCollectConfigs) Validate() error {
	return dara.Validate(s)
}
