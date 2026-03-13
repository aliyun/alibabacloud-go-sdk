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
	// Specifies whether to enable the collection of instance-level metrics. If you enable this feature, you can view core metrics, such as CPU utilization, memory usage, network conditions of instances, and the number of requests that an instance concurrently processes. Valid values: false: disables the collection of instance-level metrics. This is the default value. true: enables the collection of instance-level metrics.
	//
	// example:
	//
	// true
	EnableInstanceMetrics *bool `json:"enableInstanceMetrics,omitempty" xml:"enableInstanceMetrics,omitempty"`
	EnableLlmMetrics      *bool `json:"enableLlmMetrics,omitempty" xml:"enableLlmMetrics,omitempty"`
	// Specifies whether to enable request-level metrics. If you enable this feature, you can view the amount of time and memory consumed for a specific invocation of each function in the service. Valid values: false: disables request-level metrics. true: enables request-level metrics. This is the default value.
	//
	// example:
	//
	// true
	EnableRequestMetrics *bool `json:"enableRequestMetrics,omitempty" xml:"enableRequestMetrics,omitempty"`
	// The log segmentation rule.
	//
	// example:
	//
	// DefaultRegex
	LogBeginRule *string `json:"logBeginRule,omitempty" xml:"logBeginRule,omitempty"`
	// The name of the Logstore of Simple Log Service.
	//
	// example:
	//
	// test-logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// The name of the project in Simple Log Service.
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
