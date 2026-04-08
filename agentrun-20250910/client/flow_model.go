// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlow interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *Flow
	GetCreatedAt() *string
	SetDefinition(v string) *Flow
	GetDefinition() *string
	SetDescription(v string) *Flow
	GetDescription() *string
	SetEnvironmentConfiguration(v *EnvironmentConfiguration) *Flow
	GetEnvironmentConfiguration() *EnvironmentConfiguration
	SetExecutionRoleArn(v string) *Flow
	GetExecutionRoleArn() *string
	SetExternalStorageLocation(v string) *Flow
	GetExternalStorageLocation() *string
	SetFlowArn(v string) *Flow
	GetFlowArn() *string
	SetFlowId(v string) *Flow
	GetFlowId() *string
	SetFlowName(v string) *Flow
	GetFlowName() *string
	SetLastUpdatedAt(v string) *Flow
	GetLastUpdatedAt() *string
	SetLoggingConfiguration(v *LoggingConfiguration) *Flow
	GetLoggingConfiguration() *LoggingConfiguration
	SetResourceGroupId(v string) *Flow
	GetResourceGroupId() *string
	SetTags(v []*string) *Flow
	GetTags() []*string
	SetTracingConfiguration(v *TracingConfiguration) *Flow
	GetTracingConfiguration() *TracingConfiguration
	SetWorkspaceId(v string) *Flow
	GetWorkspaceId() *string
}

type Flow struct {
	// 工作流的创建时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 工作流的定义内容
	//
	// example:
	//
	// {}
	Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// 工作流的描述信息，说明该工作流的用途和功能
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
	// 工作流的外部存储位置
	//
	// example:
	//
	// oss://bucket/path
	ExternalStorageLocation *string `json:"externalStorageLocation,omitempty" xml:"externalStorageLocation,omitempty"`
	// 工作流的全局唯一资源名称
	//
	// example:
	//
	// acs:agentrun:cn-hangzhou:1760720386195983:workspaces/ws-xxx/flows/flow-xxx
	FlowArn *string `json:"flowArn,omitempty" xml:"flowArn,omitempty"`
	// 工作流的唯一标识符
	//
	// example:
	//
	// flow-1234567890abcdef
	FlowId *string `json:"flowId,omitempty" xml:"flowId,omitempty"`
	// 工作流的名称，用于标识和区分不同的工作流实例
	//
	// example:
	//
	// my-flow
	FlowName *string `json:"flowName,omitempty" xml:"flowName,omitempty"`
	// 工作流最后一次更新的时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
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
	// 工作流所属的工作空间标识符
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s Flow) String() string {
	return dara.Prettify(s)
}

func (s Flow) GoString() string {
	return s.String()
}

func (s *Flow) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Flow) GetDefinition() *string {
	return s.Definition
}

func (s *Flow) GetDescription() *string {
	return s.Description
}

func (s *Flow) GetEnvironmentConfiguration() *EnvironmentConfiguration {
	return s.EnvironmentConfiguration
}

func (s *Flow) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *Flow) GetExternalStorageLocation() *string {
	return s.ExternalStorageLocation
}

func (s *Flow) GetFlowArn() *string {
	return s.FlowArn
}

func (s *Flow) GetFlowId() *string {
	return s.FlowId
}

func (s *Flow) GetFlowName() *string {
	return s.FlowName
}

func (s *Flow) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *Flow) GetLoggingConfiguration() *LoggingConfiguration {
	return s.LoggingConfiguration
}

func (s *Flow) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *Flow) GetTags() []*string {
	return s.Tags
}

func (s *Flow) GetTracingConfiguration() *TracingConfiguration {
	return s.TracingConfiguration
}

func (s *Flow) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Flow) SetCreatedAt(v string) *Flow {
	s.CreatedAt = &v
	return s
}

func (s *Flow) SetDefinition(v string) *Flow {
	s.Definition = &v
	return s
}

func (s *Flow) SetDescription(v string) *Flow {
	s.Description = &v
	return s
}

func (s *Flow) SetEnvironmentConfiguration(v *EnvironmentConfiguration) *Flow {
	s.EnvironmentConfiguration = v
	return s
}

func (s *Flow) SetExecutionRoleArn(v string) *Flow {
	s.ExecutionRoleArn = &v
	return s
}

func (s *Flow) SetExternalStorageLocation(v string) *Flow {
	s.ExternalStorageLocation = &v
	return s
}

func (s *Flow) SetFlowArn(v string) *Flow {
	s.FlowArn = &v
	return s
}

func (s *Flow) SetFlowId(v string) *Flow {
	s.FlowId = &v
	return s
}

func (s *Flow) SetFlowName(v string) *Flow {
	s.FlowName = &v
	return s
}

func (s *Flow) SetLastUpdatedAt(v string) *Flow {
	s.LastUpdatedAt = &v
	return s
}

func (s *Flow) SetLoggingConfiguration(v *LoggingConfiguration) *Flow {
	s.LoggingConfiguration = v
	return s
}

func (s *Flow) SetResourceGroupId(v string) *Flow {
	s.ResourceGroupId = &v
	return s
}

func (s *Flow) SetTags(v []*string) *Flow {
	s.Tags = v
	return s
}

func (s *Flow) SetTracingConfiguration(v *TracingConfiguration) *Flow {
	s.TracingConfiguration = v
	return s
}

func (s *Flow) SetWorkspaceId(v string) *Flow {
	s.WorkspaceId = &v
	return s
}

func (s *Flow) Validate() error {
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
