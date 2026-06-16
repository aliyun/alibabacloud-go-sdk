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
	SetDisablePublicNetworkAccess(v bool) *Flow
	GetDisablePublicNetworkAccess() *bool
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
	SetTracingConfiguration(v *TracingConfiguration) *Flow
	GetTracingConfiguration() *TracingConfiguration
	SetWorkspaceId(v string) *Flow
	GetWorkspaceId() *string
}

type Flow struct {
	// The time when the workflow was created, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The definition of the workflow.
	//
	// example:
	//
	// {}
	Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// The description of the workflow, which explains the purpose and functionality of the workflow.
	//
	// example:
	//
	// Customer service automation flow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to disable public network access for the workflow. This setting serves as the default policy at the workflow level.
	DisablePublicNetworkAccess *bool `json:"disablePublicNetworkAccess,omitempty" xml:"disablePublicNetworkAccess,omitempty"`
	// The environment variable configuration of the workflow, which contains a list of named variables.
	EnvironmentConfiguration *EnvironmentConfiguration `json:"environmentConfiguration,omitempty" xml:"environmentConfiguration,omitempty"`
	// The ARN of the execution role that grants the workflow permissions to access cloud services.
	//
	// example:
	//
	// acs:ram::123456789012:role/FlowExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// The external storage location of the workflow.
	//
	// example:
	//
	// oss://bucket/path
	ExternalStorageLocation *string `json:"externalStorageLocation,omitempty" xml:"externalStorageLocation,omitempty"`
	// The globally unique Alibaba Cloud Resource Name (ARN) of the workflow.
	//
	// example:
	//
	// acs:agentrun:cn-hangzhou:123456789012:workspaces/ws-xxx/flows/flow-xxx
	FlowArn *string `json:"flowArn,omitempty" xml:"flowArn,omitempty"`
	// The unique identifier of the workflow.
	//
	// example:
	//
	// flow-1234567890abcdef
	FlowId *string `json:"flowId,omitempty" xml:"flowId,omitempty"`
	// The name of the workflow, which is used to identify and distinguish different workflow instances.
	//
	// example:
	//
	// my-flow
	FlowName *string `json:"flowName,omitempty" xml:"flowName,omitempty"`
	// The time when the workflow was last updated, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// The logging configuration of the workflow.
	LoggingConfiguration *LoggingConfiguration `json:"loggingConfiguration,omitempty" xml:"loggingConfiguration,omitempty"`
	// The ID of the resource group to which the workflow belongs.
	//
	// example:
	//
	// rg-acfmxsn4m4a4b4a
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The Tracing Analysis configuration of the workflow.
	TracingConfiguration *TracingConfiguration `json:"tracingConfiguration,omitempty" xml:"tracingConfiguration,omitempty"`
	// The ID of the workspace to which the workflow belongs.
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

func (s *Flow) GetDisablePublicNetworkAccess() *bool {
	return s.DisablePublicNetworkAccess
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

func (s *Flow) SetDisablePublicNetworkAccess(v bool) *Flow {
	s.DisablePublicNetworkAccess = &v
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
