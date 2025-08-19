// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnableInstanceMetrics(v bool) *LogConfig
	GetEnableInstanceMetrics() *bool
	SetEnableRequestMetrics(v bool) *LogConfig
	GetEnableRequestMetrics() *bool
	SetLogBeginRule(v string) *LogConfig
	GetLogBeginRule() *string
	SetLogstore(v string) *LogConfig
	GetLogstore() *string
	SetProject(v string) *LogConfig
	GetProject() *string
}

type LogConfig struct {
	// example:
	//
	// true
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	// example:
	//
	// true
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// example:
	//
	// DefaultRegex
	LogBeginRule *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	// example:
	//
	// test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// example:
	//
	// test-project
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s LogConfig) String() string {
	return dara.Prettify(s)
}

func (s LogConfig) GoString() string {
	return s.String()
}

func (s *LogConfig) GetEnableInstanceMetrics() *bool {
	return s.EnableInstanceMetrics
}

func (s *LogConfig) GetEnableRequestMetrics() *bool {
	return s.EnableRequestMetrics
}

func (s *LogConfig) GetLogBeginRule() *string {
	return s.LogBeginRule
}

func (s *LogConfig) GetLogstore() *string {
	return s.Logstore
}

func (s *LogConfig) GetProject() *string {
	return s.Project
}

func (s *LogConfig) SetEnableInstanceMetrics(v bool) *LogConfig {
	s.EnableInstanceMetrics = &v
	return s
}

func (s *LogConfig) SetEnableRequestMetrics(v bool) *LogConfig {
	s.EnableRequestMetrics = &v
	return s
}

func (s *LogConfig) SetLogBeginRule(v string) *LogConfig {
	s.LogBeginRule = &v
	return s
}

func (s *LogConfig) SetLogstore(v string) *LogConfig {
	s.Logstore = &v
	return s
}

func (s *LogConfig) SetProject(v string) *LogConfig {
	s.Project = &v
	return s
}

func (s *LogConfig) Validate() error {
	return dara.Validate(s)
}
