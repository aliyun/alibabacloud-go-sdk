// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowVersion interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *FlowVersion
	GetCreatedAt() *string
	SetDefinition(v string) *FlowVersion
	GetDefinition() *string
	SetDescription(v string) *FlowVersion
	GetDescription() *string
	SetEnvironmentConfiguration(v *EnvironmentConfiguration) *FlowVersion
	GetEnvironmentConfiguration() *EnvironmentConfiguration
	SetExternalStorageLocation(v string) *FlowVersion
	GetExternalStorageLocation() *string
	SetFlowId(v string) *FlowVersion
	GetFlowId() *string
	SetFlowName(v string) *FlowVersion
	GetFlowName() *string
	SetFlowVersion(v string) *FlowVersion
	GetFlowVersion() *string
	SetLoggingConfiguration(v *LoggingConfiguration) *FlowVersion
	GetLoggingConfiguration() *LoggingConfiguration
	SetTracingConfiguration(v *TracingConfiguration) *FlowVersion
	GetTracingConfiguration() *TracingConfiguration
}

type FlowVersion struct {
	// 工作流版本的创建时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 工作流版本的定义内容（完整快照，仅在 GetFlowVersion 时返回）
	//
	// example:
	//
	// Type: StateMachine\nName: my-flow\n...
	Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// 工作流版本的描述信息
	//
	// example:
	//
	// Version 1.0 - Initial release
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 工作流版本的环境变量配置（完整快照，仅在 GetFlowVersion 时返回）
	EnvironmentConfiguration *EnvironmentConfiguration `json:"environmentConfiguration,omitempty" xml:"environmentConfiguration,omitempty"`
	// 工作流版本的外部存储位置（完整快照，仅在 GetFlowVersion 时返回）
	//
	// example:
	//
	// oss://bucket/path
	ExternalStorageLocation *string `json:"externalStorageLocation,omitempty" xml:"externalStorageLocation,omitempty"`
	// 工作流的唯一标识符
	//
	// example:
	//
	// flow-1234567890abcdef
	FlowId   *string `json:"flowId,omitempty" xml:"flowId,omitempty"`
	FlowName *string `json:"flowName,omitempty" xml:"flowName,omitempty"`
	// 工作流版本号
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"flowVersion,omitempty" xml:"flowVersion,omitempty"`
	// 工作流版本的日志配置（完整快照，仅在 GetFlowVersion 时返回）
	LoggingConfiguration *LoggingConfiguration `json:"loggingConfiguration,omitempty" xml:"loggingConfiguration,omitempty"`
	// 工作流版本的链路追踪配置（完整快照，仅在 GetFlowVersion 时返回）
	TracingConfiguration *TracingConfiguration `json:"tracingConfiguration,omitempty" xml:"tracingConfiguration,omitempty"`
}

func (s FlowVersion) String() string {
	return dara.Prettify(s)
}

func (s FlowVersion) GoString() string {
	return s.String()
}

func (s *FlowVersion) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *FlowVersion) GetDefinition() *string {
	return s.Definition
}

func (s *FlowVersion) GetDescription() *string {
	return s.Description
}

func (s *FlowVersion) GetEnvironmentConfiguration() *EnvironmentConfiguration {
	return s.EnvironmentConfiguration
}

func (s *FlowVersion) GetExternalStorageLocation() *string {
	return s.ExternalStorageLocation
}

func (s *FlowVersion) GetFlowId() *string {
	return s.FlowId
}

func (s *FlowVersion) GetFlowName() *string {
	return s.FlowName
}

func (s *FlowVersion) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *FlowVersion) GetLoggingConfiguration() *LoggingConfiguration {
	return s.LoggingConfiguration
}

func (s *FlowVersion) GetTracingConfiguration() *TracingConfiguration {
	return s.TracingConfiguration
}

func (s *FlowVersion) SetCreatedAt(v string) *FlowVersion {
	s.CreatedAt = &v
	return s
}

func (s *FlowVersion) SetDefinition(v string) *FlowVersion {
	s.Definition = &v
	return s
}

func (s *FlowVersion) SetDescription(v string) *FlowVersion {
	s.Description = &v
	return s
}

func (s *FlowVersion) SetEnvironmentConfiguration(v *EnvironmentConfiguration) *FlowVersion {
	s.EnvironmentConfiguration = v
	return s
}

func (s *FlowVersion) SetExternalStorageLocation(v string) *FlowVersion {
	s.ExternalStorageLocation = &v
	return s
}

func (s *FlowVersion) SetFlowId(v string) *FlowVersion {
	s.FlowId = &v
	return s
}

func (s *FlowVersion) SetFlowName(v string) *FlowVersion {
	s.FlowName = &v
	return s
}

func (s *FlowVersion) SetFlowVersion(v string) *FlowVersion {
	s.FlowVersion = &v
	return s
}

func (s *FlowVersion) SetLoggingConfiguration(v *LoggingConfiguration) *FlowVersion {
	s.LoggingConfiguration = v
	return s
}

func (s *FlowVersion) SetTracingConfiguration(v *TracingConfiguration) *FlowVersion {
	s.TracingConfiguration = v
	return s
}

func (s *FlowVersion) Validate() error {
	if s.EnvironmentConfiguration != nil {
		if err := s.EnvironmentConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.LoggingConfiguration != nil {
		if err := s.LoggingConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.TracingConfiguration != nil {
		if err := s.TracingConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
