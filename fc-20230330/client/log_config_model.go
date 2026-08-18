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
	SetEnableLlmMetrics(v bool) *LogConfig
	GetEnableLlmMetrics() *bool
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
	// Specifies whether to enable instance-level metrics. After you enable this feature, you can view core metrics such as CPU usage, memory usage, network status, and request count at the instance level. Valid values: false: disables instance-level metrics. This is the default value. true: enables instance-level metrics.
	//
	// example:
	//
	// true
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	// Specifies whether to enable LLM metrics. After you enable this feature, you can view LLM metrics. We recommend that you enable this feature only for LLM inference services. Valid values: false: disables LLM metrics. This is the default value. true: enables LLM metrics.
	EnableLlmMetrics *bool `json:"enableLlmMetrics,omitempty" xml:"enableLlmMetrics,omitempty"`
	// Specifies whether to enable request-level metrics. After you enable this feature, you can view the time and memory consumed by each invocation of all functions in the service. Valid values: false: disables request-level metrics. true: enables request-level metrics. This is the default value.
	//
	// example:
	//
	// true
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// The log line beginning matching rule.
	//
	// example:
	//
	// DefaultRegex
	LogBeginRule *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	// The Logstore name in Simple Log Service.
	//
	// example:
	//
	// test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The project name in Simple Log Service.
	//
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

func (s *LogConfig) GetEnableLlmMetrics() *bool {
	return s.EnableLlmMetrics
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

func (s *LogConfig) SetEnableLlmMetrics(v bool) *LogConfig {
	s.EnableLlmMetrics = &v
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
