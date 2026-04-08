// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowInput interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *UpdateFlowInput
	GetDefinition() *string
	SetDescription(v string) *UpdateFlowInput
	GetDescription() *string
	SetEnvironmentConfiguration(v *EnvironmentConfiguration) *UpdateFlowInput
	GetEnvironmentConfiguration() *EnvironmentConfiguration
	SetExecutionRoleArn(v string) *UpdateFlowInput
	GetExecutionRoleArn() *string
	SetExternalStorageLocation(v string) *UpdateFlowInput
	GetExternalStorageLocation() *string
	SetFlowName(v string) *UpdateFlowInput
	GetFlowName() *string
	SetLoggingConfiguration(v *LoggingConfiguration) *UpdateFlowInput
	GetLoggingConfiguration() *LoggingConfiguration
	SetResourceGroupId(v string) *UpdateFlowInput
	GetResourceGroupId() *string
	SetTags(v []*string) *UpdateFlowInput
	GetTags() []*string
	SetTracingConfiguration(v *TracingConfiguration) *UpdateFlowInput
	GetTracingConfiguration() *TracingConfiguration
	SetWorkspaceId(v string) *UpdateFlowInput
	GetWorkspaceId() *string
}

type UpdateFlowInput struct {
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
	// 工作流的标签信息，用于资源分类和管理
	//
	// example:
	//
	// production,automation
	Tags []*string `json:"tags" xml:"tags" type:"Repeated"`
	// 工作流的链路追踪配置
	TracingConfiguration *TracingConfiguration `json:"tracingConfiguration,omitempty" xml:"tracingConfiguration,omitempty"`
	// 工作流所属的工作空间标识符，用于资源隔离和权限管理
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateFlowInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowInput) GoString() string {
	return s.String()
}

func (s *UpdateFlowInput) GetDefinition() *string {
	return s.Definition
}

func (s *UpdateFlowInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowInput) GetEnvironmentConfiguration() *EnvironmentConfiguration {
	return s.EnvironmentConfiguration
}

func (s *UpdateFlowInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *UpdateFlowInput) GetExternalStorageLocation() *string {
	return s.ExternalStorageLocation
}

func (s *UpdateFlowInput) GetFlowName() *string {
	return s.FlowName
}

func (s *UpdateFlowInput) GetLoggingConfiguration() *LoggingConfiguration {
	return s.LoggingConfiguration
}

func (s *UpdateFlowInput) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateFlowInput) GetTags() []*string {
	return s.Tags
}

func (s *UpdateFlowInput) GetTracingConfiguration() *TracingConfiguration {
	return s.TracingConfiguration
}

func (s *UpdateFlowInput) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateFlowInput) SetDefinition(v string) *UpdateFlowInput {
	s.Definition = &v
	return s
}

func (s *UpdateFlowInput) SetDescription(v string) *UpdateFlowInput {
	s.Description = &v
	return s
}

func (s *UpdateFlowInput) SetEnvironmentConfiguration(v *EnvironmentConfiguration) *UpdateFlowInput {
	s.EnvironmentConfiguration = v
	return s
}

func (s *UpdateFlowInput) SetExecutionRoleArn(v string) *UpdateFlowInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *UpdateFlowInput) SetExternalStorageLocation(v string) *UpdateFlowInput {
	s.ExternalStorageLocation = &v
	return s
}

func (s *UpdateFlowInput) SetFlowName(v string) *UpdateFlowInput {
	s.FlowName = &v
	return s
}

func (s *UpdateFlowInput) SetLoggingConfiguration(v *LoggingConfiguration) *UpdateFlowInput {
	s.LoggingConfiguration = v
	return s
}

func (s *UpdateFlowInput) SetResourceGroupId(v string) *UpdateFlowInput {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateFlowInput) SetTags(v []*string) *UpdateFlowInput {
	s.Tags = v
	return s
}

func (s *UpdateFlowInput) SetTracingConfiguration(v *TracingConfiguration) *UpdateFlowInput {
	s.TracingConfiguration = v
	return s
}

func (s *UpdateFlowInput) SetWorkspaceId(v string) *UpdateFlowInput {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateFlowInput) Validate() error {
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
