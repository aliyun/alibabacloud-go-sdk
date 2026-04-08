// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowDraftInput interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *UpdateFlowDraftInput
	GetDefinition() *string
	SetDescription(v string) *UpdateFlowDraftInput
	GetDescription() *string
	SetEnvironmentConfiguration(v *EnvironmentConfiguration) *UpdateFlowDraftInput
	GetEnvironmentConfiguration() *EnvironmentConfiguration
	SetExecutionRoleArn(v string) *UpdateFlowDraftInput
	GetExecutionRoleArn() *string
	SetExternalStorageLocation(v string) *UpdateFlowDraftInput
	GetExternalStorageLocation() *string
}

type UpdateFlowDraftInput struct {
	// 工作流的 FDL 定义
	Definition *string `json:"definition,omitempty" xml:"definition,omitempty"`
	// 工作流的描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 工作流执行期间可以访问的环境变量配置，包含一组命名变量列表
	EnvironmentConfiguration *EnvironmentConfiguration `json:"environmentConfiguration,omitempty" xml:"environmentConfiguration,omitempty"`
	// 工作流执行时使用的 RAM 角色 ARN
	ExecutionRoleArn *string `json:"executionRoleArn,omitempty" xml:"executionRoleArn,omitempty"`
	// 工作流执行历史的外部存储位置
	ExternalStorageLocation *string `json:"externalStorageLocation,omitempty" xml:"externalStorageLocation,omitempty"`
}

func (s UpdateFlowDraftInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowDraftInput) GoString() string {
	return s.String()
}

func (s *UpdateFlowDraftInput) GetDefinition() *string {
	return s.Definition
}

func (s *UpdateFlowDraftInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateFlowDraftInput) GetEnvironmentConfiguration() *EnvironmentConfiguration {
	return s.EnvironmentConfiguration
}

func (s *UpdateFlowDraftInput) GetExecutionRoleArn() *string {
	return s.ExecutionRoleArn
}

func (s *UpdateFlowDraftInput) GetExternalStorageLocation() *string {
	return s.ExternalStorageLocation
}

func (s *UpdateFlowDraftInput) SetDefinition(v string) *UpdateFlowDraftInput {
	s.Definition = &v
	return s
}

func (s *UpdateFlowDraftInput) SetDescription(v string) *UpdateFlowDraftInput {
	s.Description = &v
	return s
}

func (s *UpdateFlowDraftInput) SetEnvironmentConfiguration(v *EnvironmentConfiguration) *UpdateFlowDraftInput {
	s.EnvironmentConfiguration = v
	return s
}

func (s *UpdateFlowDraftInput) SetExecutionRoleArn(v string) *UpdateFlowDraftInput {
	s.ExecutionRoleArn = &v
	return s
}

func (s *UpdateFlowDraftInput) SetExternalStorageLocation(v string) *UpdateFlowDraftInput {
	s.ExternalStorageLocation = &v
	return s
}

func (s *UpdateFlowDraftInput) Validate() error {
	if s.EnvironmentConfiguration != nil {
		if err := s.EnvironmentConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
