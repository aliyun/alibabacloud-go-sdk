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
	SetDisablePublicNetworkAccess(v bool) *UpdateFlowInput
	GetDisablePublicNetworkAccess() *bool
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
	SetTracingConfiguration(v *TracingConfiguration) *UpdateFlowInput
	GetTracingConfiguration() *TracingConfiguration
	SetWorkspaceId(v string) *UpdateFlowInput
	GetWorkspaceId() *string
}

type UpdateFlowInput struct {
	// The definition content of the workflow, in JSON or YAML format
	//
	// example:
	//
	// {}
	Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// The description of the workflow, used to explain its purpose and functionality
	//
	// example:
	//
	// Customer service automation flow
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Whether to disable public network access for the workflow, serving as the default policy at the workflow level. When FlowEndpoint is not specified, this value will be inherited
	DisablePublicNetworkAccess *bool `json:"disablePublicNetworkAccess,omitempty" xml:"disablePublicNetworkAccess,omitempty"`
	// The environment variable configuration of the workflow, containing a list of named variables
	EnvironmentConfiguration *EnvironmentConfiguration `json:"environmentConfiguration,omitempty" xml:"environmentConfiguration,omitempty"`
	// The execution role ARN that grants the workflow access permissions to cloud services
	//
	// example:
	//
	// acs:ram::123456789012:role/FlowExecutionRole
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// The external storage location of the workflow, such as an OSS path
	//
	// example:
	//
	// oss://bucket/path
	ExternalStorageLocation *string `json:"externalStorageLocation,omitempty" xml:"externalStorageLocation,omitempty"`
	// The unique identifier name of the workflow, used to distinguish different workflow instances
	//
	// example:
	//
	// my-flow
	FlowName *string `json:"flowName,omitempty" xml:"flowName,omitempty"`
	// The logging configuration of the workflow
	LoggingConfiguration *LoggingConfiguration `json:"loggingConfiguration,omitempty" xml:"loggingConfiguration,omitempty"`
	// The resource group identifier to which the workflow belongs
	//
	// example:
	//
	// rg-acfmxsn4m4a4b4a
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The distributed tracing configuration of the workflow
	TracingConfiguration *TracingConfiguration `json:"tracingConfiguration,omitempty" xml:"tracingConfiguration,omitempty"`
	// The workspace identifier to which the workflow belongs, used for resource isolation and permission management
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

func (s *UpdateFlowInput) GetDisablePublicNetworkAccess() *bool {
	return s.DisablePublicNetworkAccess
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

func (s *UpdateFlowInput) SetDisablePublicNetworkAccess(v bool) *UpdateFlowInput {
	s.DisablePublicNetworkAccess = &v
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
