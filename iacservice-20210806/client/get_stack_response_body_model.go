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
	// example:
	//
	// C7070EC3-DF66-58BA-A1DD-A8574FF53143
	RequestId *string                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Stack     *GetStackResponseBodyStack `json:"stack,omitempty" xml:"stack,omitempty" type:"Struct"`
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
	Config *GetStackResponseBodyStackConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 2025-07-24T02:58:53Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// v1
	CurrentConfigVersion *string `json:"currentConfigVersion,omitempty" xml:"currentConfigVersion,omitempty"`
	// example:
	//
	// the description of stack
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// stack-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// TestIacRole
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// stack-as1d4vld898ppnqxxxxxx
	StackId *string `json:"stackId,omitempty" xml:"stackId,omitempty"`
	// example:
	//
	// Deployed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// SetUpdated
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
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
	// example:
	//
	// format_version: IaCService/2021-08-06\\ndescription: create ALB\\nupstream_input:\\n  - name: stack_network\\n ...
	ComponentContent *string `json:"componentContent,omitempty" xml:"componentContent,omitempty"`
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
