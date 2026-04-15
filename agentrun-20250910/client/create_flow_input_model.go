// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowInput interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *CreateFlowInput
	GetDefinition() *string
	SetDescription(v string) *CreateFlowInput
	GetDescription() *string
	SetEnvironmentConfiguration(v *EnvironmentConfiguration) *CreateFlowInput
	GetEnvironmentConfiguration() *EnvironmentConfiguration
	SetExecutionRoleArn(v string) *CreateFlowInput
	GetExecutionRoleArn() *string
	SetExternalStorageLocation(v string) *CreateFlowInput
	GetExternalStorageLocation() *string
	SetFlowName(v string) *CreateFlowInput
	GetFlowName() *string
	SetLoggingConfiguration(v *LoggingConfiguration) *CreateFlowInput
	GetLoggingConfiguration() *LoggingConfiguration
	SetResourceGroupId(v string) *CreateFlowInput
	GetResourceGroupId() *string
	SetTracingConfiguration(v *TracingConfiguration) *CreateFlowInput
	GetTracingConfiguration() *TracingConfiguration
	SetWorkspaceId(v string) *CreateFlowInput
	GetWorkspaceId() *string
}

type CreateFlowInput struct {
	// 工作流的定义内容，采用JSON或YAML格式
	//
	// example:
	//
	// {}
	Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// 工作流的描述信息，用于说明该工作流的用途和功能
	//
	// example:
	//
	// Customer service automation flow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 工作流的环境变量配置，包含一组命名变量列表
	EnvironmentConfiguration *EnvironmentConfiguration `json:"environmentConfiguration,omitempty" xml:"environmentConfiguration,omitempty"`
	// 为工作流提供访问云服务权限的执行角色ARN
	//
	// example:
	//
	// acs:ram::1760720386195983:role/FlowExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// 工作流的外部存储位置，如OSS路径
	//
	// example:
	//
	// oss://bucket/path
	ExternalStorageLocation *string `json:"externalStorageLocation,omitempty" xml:"externalStorageLocation,omitempty"`
	// 工作流的唯一标识名称，用于区分不同的工作流实例
	//
	// This parameter is required.
	//
	// example:
	//
	// my-flow
	FlowName *string `json:"flowName,omitempty" xml:"flowName,omitempty"`
	// 工作流的日志配置
	LoggingConfiguration *LoggingConfiguration `json:"loggingConfiguration,omitempty" xml:"loggingConfiguration,omitempty"`
	// 工作流所属的资源组标识符
	//
	// example:
	//
	// rg-acfmxsn4m4a4b4a
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// 工作流的链路追踪配置
	TracingConfiguration *TracingConfiguration `json:"tracingConfiguration,omitempty" xml:"tracingConfiguration,omitempty"`
	// 工作流所属的工作空间标识符，用于资源隔离和权限管理
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateFlowInput) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowInput) GoString() string {
	return s.String()
}

func (s *CreateFlowInput) GetDefinition() *string {
	return s.Definition
}

func (s *CreateFlowInput) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowInput) GetEnvironmentConfiguration() *EnvironmentConfiguration {
	return s.EnvironmentConfiguration
}

func (s *CreateFlowInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *CreateFlowInput) GetExternalStorageLocation() *string {
	return s.ExternalStorageLocation
}

func (s *CreateFlowInput) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateFlowInput) GetLoggingConfiguration() *LoggingConfiguration {
	return s.LoggingConfiguration
}

func (s *CreateFlowInput) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateFlowInput) GetTracingConfiguration() *TracingConfiguration {
	return s.TracingConfiguration
}

func (s *CreateFlowInput) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateFlowInput) SetDefinition(v string) *CreateFlowInput {
	s.Definition = &v
	return s
}

func (s *CreateFlowInput) SetDescription(v string) *CreateFlowInput {
	s.Description = &v
	return s
}

func (s *CreateFlowInput) SetEnvironmentConfiguration(v *EnvironmentConfiguration) *CreateFlowInput {
	s.EnvironmentConfiguration = v
	return s
}

func (s *CreateFlowInput) SetExecutionRoleArn(v string) *CreateFlowInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *CreateFlowInput) SetExternalStorageLocation(v string) *CreateFlowInput {
	s.ExternalStorageLocation = &v
	return s
}

func (s *CreateFlowInput) SetFlowName(v string) *CreateFlowInput {
	s.FlowName = &v
	return s
}

func (s *CreateFlowInput) SetLoggingConfiguration(v *LoggingConfiguration) *CreateFlowInput {
	s.LoggingConfiguration = v
	return s
}

func (s *CreateFlowInput) SetResourceGroupId(v string) *CreateFlowInput {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateFlowInput) SetTracingConfiguration(v *TracingConfiguration) *CreateFlowInput {
	s.TracingConfiguration = v
	return s
}

func (s *CreateFlowInput) SetWorkspaceId(v string) *CreateFlowInput {
	s.WorkspaceId = &v
	return s
}

func (s *CreateFlowInput) Validate() error {
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
