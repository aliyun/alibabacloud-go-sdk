// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSLSTriggerConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *SLSTriggerConfig
	GetEnable() *bool
	SetFunctionParameter(v map[string]*string) *SLSTriggerConfig
	GetFunctionParameter() map[string]*string
	SetJobConfig(v *JobConfig) *SLSTriggerConfig
	GetJobConfig() *JobConfig
	SetLogConfig(v *SLSTriggerLogConfig) *SLSTriggerConfig
	GetLogConfig() *SLSTriggerLogConfig
	SetSourceConfig(v *SourceConfig) *SLSTriggerConfig
	GetSourceConfig() *SourceConfig
}

type SLSTriggerConfig struct {
	// example:
	//
	// true
	Enable            *bool                `json:"enable,omitempty" xml:"enable,omitempty"`
	FunctionParameter map[string]*string   `json:"functionParameter" xml:"functionParameter"`
	JobConfig         *JobConfig           `json:"jobConfig,omitempty" xml:"jobConfig,omitempty"`
	LogConfig         *SLSTriggerLogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty"`
	SourceConfig      *SourceConfig        `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty"`
}

func (s SLSTriggerConfig) String() string {
	return dara.Prettify(s)
}

func (s SLSTriggerConfig) GoString() string {
	return s.String()
}

func (s *SLSTriggerConfig) GetEnable() *bool {
	return s.Enable
}

func (s *SLSTriggerConfig) GetFunctionParameter() map[string]*string {
	return s.FunctionParameter
}

func (s *SLSTriggerConfig) GetJobConfig() *JobConfig {
	return s.JobConfig
}

func (s *SLSTriggerConfig) GetLogConfig() *SLSTriggerLogConfig {
	return s.LogConfig
}

func (s *SLSTriggerConfig) GetSourceConfig() *SourceConfig {
	return s.SourceConfig
}

func (s *SLSTriggerConfig) SetEnable(v bool) *SLSTriggerConfig {
	s.Enable = &v
	return s
}

func (s *SLSTriggerConfig) SetFunctionParameter(v map[string]*string) *SLSTriggerConfig {
	s.FunctionParameter = v
	return s
}

func (s *SLSTriggerConfig) SetJobConfig(v *JobConfig) *SLSTriggerConfig {
	s.JobConfig = v
	return s
}

func (s *SLSTriggerConfig) SetLogConfig(v *SLSTriggerLogConfig) *SLSTriggerConfig {
	s.LogConfig = v
	return s
}

func (s *SLSTriggerConfig) SetSourceConfig(v *SourceConfig) *SLSTriggerConfig {
	s.SourceConfig = v
	return s
}

func (s *SLSTriggerConfig) Validate() error {
	return dara.Validate(s)
}
