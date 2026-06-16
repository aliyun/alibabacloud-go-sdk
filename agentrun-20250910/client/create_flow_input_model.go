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
	SetDisablePublicNetworkAccess(v bool) *CreateFlowInput
	GetDisablePublicNetworkAccess() *bool
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
	// The definition of the workflow in JSON or YAML format.
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
	// Specifies whether to disable public network access for the workflow. This value serves as the default policy at the workflow level. If FlowEndpoint is not specified, this value is inherited.
	DisablePublicNetworkAccess *bool `json:"disablePublicNetworkAccess,omitempty" xml:"disablePublicNetworkAccess,omitempty"`
	// The environment variable configuration of the workflow, which contains a list of named variables.
	EnvironmentConfiguration *EnvironmentConfiguration `json:"environmentConfiguration,omitempty" xml:"environmentConfiguration,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the execution role that grants the workflow permissions to access cloud services.
	//
	// example:
	//
	// acs:ram::123456789012:role/FlowExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// The external storage location of the workflow, such as an OSS path.
	//
	// example:
	//
	// oss://bucket/path
	ExternalStorageLocation *string `json:"externalStorageLocation,omitempty" xml:"externalStorageLocation,omitempty"`
	// The unique name of the workflow, which is used to distinguish different workflow instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-flow
	FlowName *string `json:"flowName,omitempty" xml:"flowName,omitempty"`
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
	// The ID of the workspace to which the workflow belongs. This parameter is used for resource isolation and permission management.
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

func (s *CreateFlowInput) GetDisablePublicNetworkAccess() *bool {
	return s.DisablePublicNetworkAccess
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

func (s *CreateFlowInput) SetDisablePublicNetworkAccess(v bool) *CreateFlowInput {
	s.DisablePublicNetworkAccess = &v
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
