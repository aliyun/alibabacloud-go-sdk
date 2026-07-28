// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStackResponseBody
	GetRequestId() *string
	SetStack(v *GetStackResponseBodyStack) *GetStackResponseBody
	GetStack() *GetStackResponseBodyStack
}

type GetStackResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C7070EC3-DF66-58BA-A1DD-A8574FF53143
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The stack information.
	Stack *GetStackResponseBodyStack `json:"stack,omitempty" xml:"stack,omitempty" type:"Struct"`
}

func (s GetStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackResponseBody) GetStack() *GetStackResponseBodyStack {
	return s.Stack
}

func (s *GetStackResponseBody) SetRequestId(v string) *GetStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResponseBody) SetStack(v *GetStackResponseBodyStack) *GetStackResponseBody {
	s.Stack = v
	return s
}

func (s *GetStackResponseBody) Validate() error {
	if s.Stack != nil {
		if err := s.Stack.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStackResponseBodyStack struct {
	// The stack configuration.
	Config *GetStackResponseBodyStackConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 2025-07-24T02:58:53Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The current configuration version number, such as v1. The initial value is v1. The version number increments each time the stack is updated or refreshed and the configuration changes.
	//
	// example:
	//
	// v1
	CurrentConfigVersion *string `json:"currentConfigVersion,omitempty" xml:"currentConfigVersion,omitempty"`
	// The description of the stack.
	//
	// example:
	//
	// the description of stack
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The stack name.
	//
	// example:
	//
	// stack-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The RAM role assumed by the system to perform resource change operations during stack deployment.
	//
	// example:
	//
	// TestIacRole
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// The configuration source of the stack. Valid values:
	//
	// - OSS: a template stored in Object Storage Service (OSS).
	//
	// - IAC_SERVICE_MODULE: a template created in the automation service console.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The path value of the configuration source. The value cannot exceed 1000 characters.
	//
	// - If the source is OSS, the value is in the format of oss::<file link>. The file must be a ZIP file. Example: oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip.
	//
	// - If the source is IAC_SERVICE_MODULE, the value is a template ID. Example: mod-xxxxx.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The unique identifier of the stack, which is generated after the stack is created.
	//
	// example:
	//
	// stack-as1d4vld898ppnqxxxxxx
	StackId *string `json:"stackId,omitempty" xml:"stackId,omitempty"`
	// The stack status.
	//
	// | Name | Description |
	//
	// |------|------|
	//
	// | Creating | The stack is being created. |
	//
	// | Created | The stack is created. |
	//
	// | Waiting | The stack is waiting for deployment. |
	//
	// | Deploying | The stack is being deployed. |
	//
	// | Deployed | The stack is deployed. |
	//
	// | Errored | The deployment failed. |
	//
	// | Deleting | The stack is being deleted. |
	//
	// | Deleted | The stack is deleted. |
	//
	// | DeleteFailed | The deletion failed. |
	//
	// | DetectTriggered | Drift detection is triggered. |.
	//
	// example:
	//
	// Deployed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The deployment trigger method of the stack. This field is not publicly available.
	//
	// - SetUpdated: triggered by file changes.
	//
	// example:
	//
	// SetUpdated
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// The directory where the deployment and component configuration files of the stack are located. Set this parameter to / for the root directory.
	//
	// example:
	//
	// /
	WorkingDirectory *string `json:"workingDirectory,omitempty" xml:"workingDirectory,omitempty"`
}

func (s GetStackResponseBodyStack) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyStack) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStack) GetConfig() *GetStackResponseBodyStackConfig {
	return s.Config
}

func (s *GetStackResponseBodyStack) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStackResponseBodyStack) GetCurrentConfigVersion() *string {
	return s.CurrentConfigVersion
}

func (s *GetStackResponseBodyStack) GetDescription() *string {
	return s.Description
}

func (s *GetStackResponseBodyStack) GetName() *string {
	return s.Name
}

func (s *GetStackResponseBodyStack) GetRamRole() *string {
	return s.RamRole
}

func (s *GetStackResponseBodyStack) GetSource() *string {
	return s.Source
}

func (s *GetStackResponseBodyStack) GetSourcePath() *string {
	return s.SourcePath
}

func (s *GetStackResponseBodyStack) GetStackId() *string {
	return s.StackId
}

func (s *GetStackResponseBodyStack) GetStatus() *string {
	return s.Status
}

func (s *GetStackResponseBodyStack) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *GetStackResponseBodyStack) GetWorkingDirectory() *string {
	return s.WorkingDirectory
}

func (s *GetStackResponseBodyStack) SetConfig(v *GetStackResponseBodyStackConfig) *GetStackResponseBodyStack {
	s.Config = v
	return s
}

func (s *GetStackResponseBodyStack) SetCreateTime(v string) *GetStackResponseBodyStack {
	s.CreateTime = &v
	return s
}

func (s *GetStackResponseBodyStack) SetCurrentConfigVersion(v string) *GetStackResponseBodyStack {
	s.CurrentConfigVersion = &v
	return s
}

func (s *GetStackResponseBodyStack) SetDescription(v string) *GetStackResponseBodyStack {
	s.Description = &v
	return s
}

func (s *GetStackResponseBodyStack) SetName(v string) *GetStackResponseBodyStack {
	s.Name = &v
	return s
}

func (s *GetStackResponseBodyStack) SetRamRole(v string) *GetStackResponseBodyStack {
	s.RamRole = &v
	return s
}

func (s *GetStackResponseBodyStack) SetSource(v string) *GetStackResponseBodyStack {
	s.Source = &v
	return s
}

func (s *GetStackResponseBodyStack) SetSourcePath(v string) *GetStackResponseBodyStack {
	s.SourcePath = &v
	return s
}

func (s *GetStackResponseBodyStack) SetStackId(v string) *GetStackResponseBodyStack {
	s.StackId = &v
	return s
}

func (s *GetStackResponseBodyStack) SetStatus(v string) *GetStackResponseBodyStack {
	s.Status = &v
	return s
}

func (s *GetStackResponseBodyStack) SetTriggerStrategy(v string) *GetStackResponseBodyStack {
	s.TriggerStrategy = &v
	return s
}

func (s *GetStackResponseBodyStack) SetWorkingDirectory(v string) *GetStackResponseBodyStack {
	s.WorkingDirectory = &v
	return s
}

func (s *GetStackResponseBodyStack) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStackResponseBodyStackConfig struct {
	// The component configuration.
	//
	// example:
	//
	// format_version: IaCService/2021-08-06\\ndescription: create ALB\\nupstream_input:\\n  - name: stack_network\\n ...
	ComponentContent *string `json:"componentContent,omitempty" xml:"componentContent,omitempty"`
	// The deployment configuration.
	//
	// example:
	//
	// format_version: IaCService/2021-08-06\\ndescription: create ALB \\nvariable:\\n  - name: region\\n    type: string\\n ...
	DeploymentContent *string `json:"deploymentContent,omitempty" xml:"deploymentContent,omitempty"`
}

func (s GetStackResponseBodyStackConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyStackConfig) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackConfig) GetComponentContent() *string {
	return s.ComponentContent
}

func (s *GetStackResponseBodyStackConfig) GetDeploymentContent() *string {
	return s.DeploymentContent
}

func (s *GetStackResponseBodyStackConfig) SetComponentContent(v string) *GetStackResponseBodyStackConfig {
	s.ComponentContent = &v
	return s
}

func (s *GetStackResponseBodyStackConfig) SetDeploymentContent(v string) *GetStackResponseBodyStackConfig {
	s.DeploymentContent = &v
	return s
}

func (s *GetStackResponseBodyStackConfig) Validate() error {
	return dara.Validate(s)
}
