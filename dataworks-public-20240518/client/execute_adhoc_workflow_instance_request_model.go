// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAdhocWorkflowInstanceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizDate(v int64) *ExecuteAdhocWorkflowInstanceRequest
  GetBizDate() *int64 
  SetEnvType(v string) *ExecuteAdhocWorkflowInstanceRequest
  GetEnvType() *string 
  SetName(v string) *ExecuteAdhocWorkflowInstanceRequest
  GetName() *string 
  SetOwner(v string) *ExecuteAdhocWorkflowInstanceRequest
  GetOwner() *string 
  SetProjectId(v int64) *ExecuteAdhocWorkflowInstanceRequest
  GetProjectId() *int64 
  SetTasks(v []*ExecuteAdhocWorkflowInstanceRequestTasks) *ExecuteAdhocWorkflowInstanceRequest
  GetTasks() []*ExecuteAdhocWorkflowInstanceRequestTasks 
}

type ExecuteAdhocWorkflowInstanceRequest struct {
  // The data timestamp.
  // 
  // example:
  // 
  // 1710239005403
  BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
  // The environment of the workspace. Valid values:
  // 
  // 	- Prod: production environment
  // 
  // 	- Dev: development environment
  // 
  // example:
  // 
  // Prod
  EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
  // The name of the workflow instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // WorkflowInstance1
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The account ID of the owner.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1000
  Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
  // The workspace ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 100
  ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
  // The tasks.
  // 
  // This parameter is required.
  Tasks []*ExecuteAdhocWorkflowInstanceRequestTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ExecuteAdhocWorkflowInstanceRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequest) GetBizDate() *int64  {
  return s.BizDate
}

func (s *ExecuteAdhocWorkflowInstanceRequest) GetEnvType() *string  {
  return s.EnvType
}

func (s *ExecuteAdhocWorkflowInstanceRequest) GetName() *string  {
  return s.Name
}

func (s *ExecuteAdhocWorkflowInstanceRequest) GetOwner() *string  {
  return s.Owner
}

func (s *ExecuteAdhocWorkflowInstanceRequest) GetProjectId() *int64  {
  return s.ProjectId
}

func (s *ExecuteAdhocWorkflowInstanceRequest) GetTasks() []*ExecuteAdhocWorkflowInstanceRequestTasks  {
  return s.Tasks
}

func (s *ExecuteAdhocWorkflowInstanceRequest) SetBizDate(v int64) *ExecuteAdhocWorkflowInstanceRequest {
  s.BizDate = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequest) SetEnvType(v string) *ExecuteAdhocWorkflowInstanceRequest {
  s.EnvType = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequest) SetName(v string) *ExecuteAdhocWorkflowInstanceRequest {
  s.Name = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequest) SetOwner(v string) *ExecuteAdhocWorkflowInstanceRequest {
  s.Owner = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequest) SetProjectId(v int64) *ExecuteAdhocWorkflowInstanceRequest {
  s.ProjectId = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequest) SetTasks(v []*ExecuteAdhocWorkflowInstanceRequestTasks) *ExecuteAdhocWorkflowInstanceRequest {
  s.Tasks = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequest) Validate() error {
  if s.Tasks != nil {
    for _, item := range s.Tasks {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteAdhocWorkflowInstanceRequestTasks struct {
  // The unique code of the client. This code uniquely identifies a task.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Task_0bc5213917368545132902xxxxxxxx
  ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
  // The information about the associated data source.
  DataSource *ExecuteAdhocWorkflowInstanceRequestTasksDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
  // The dependency information.
  Dependencies []*ExecuteAdhocWorkflowInstanceRequestTasksDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
  // The input information.
  Inputs *ExecuteAdhocWorkflowInstanceRequestTasksInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
  // The name of the task.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // SQL node.
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The output information.
  Outputs *ExecuteAdhocWorkflowInstanceRequestTasksOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
  // The account ID of the owner.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1000
  Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
  // The configurations of the runtime environment, such as the resource group information.
  // 
  // This parameter is required.
  RuntimeResource *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
  // The script information.
  Script *ExecuteAdhocWorkflowInstanceRequestTasksScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
  // The timeout period of task running. Unit: seconds.
  // 
  // example:
  // 
  // 3600
  Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
  // The type of the task.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ODPS_SQL
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceRequestTasks) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequestTasks) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetClientUniqueCode() *string  {
  return s.ClientUniqueCode
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetDataSource() *ExecuteAdhocWorkflowInstanceRequestTasksDataSource  {
  return s.DataSource
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetDependencies() []*ExecuteAdhocWorkflowInstanceRequestTasksDependencies  {
  return s.Dependencies
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetInputs() *ExecuteAdhocWorkflowInstanceRequestTasksInputs  {
  return s.Inputs
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetName() *string  {
  return s.Name
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetOutputs() *ExecuteAdhocWorkflowInstanceRequestTasksOutputs  {
  return s.Outputs
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetOwner() *string  {
  return s.Owner
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetRuntimeResource() *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource  {
  return s.RuntimeResource
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetScript() *ExecuteAdhocWorkflowInstanceRequestTasksScript  {
  return s.Script
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetTimeout() *int32  {
  return s.Timeout
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) GetType() *string  {
  return s.Type
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetClientUniqueCode(v string) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.ClientUniqueCode = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetDataSource(v *ExecuteAdhocWorkflowInstanceRequestTasksDataSource) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.DataSource = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetDependencies(v []*ExecuteAdhocWorkflowInstanceRequestTasksDependencies) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.Dependencies = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetInputs(v *ExecuteAdhocWorkflowInstanceRequestTasksInputs) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.Inputs = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetName(v string) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.Name = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetOutputs(v *ExecuteAdhocWorkflowInstanceRequestTasksOutputs) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.Outputs = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetOwner(v string) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.Owner = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetRuntimeResource(v *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.RuntimeResource = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetScript(v *ExecuteAdhocWorkflowInstanceRequestTasksScript) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.Script = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetTimeout(v int32) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.Timeout = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) SetType(v string) *ExecuteAdhocWorkflowInstanceRequestTasks {
  s.Type = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasks) Validate() error {
  if s.DataSource != nil {
    if err := s.DataSource.Validate(); err != nil {
      return err
    }
  }
  if s.Dependencies != nil {
    for _, item := range s.Dependencies {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.Inputs != nil {
    if err := s.Inputs.Validate(); err != nil {
      return err
    }
  }
  if s.Outputs != nil {
    if err := s.Outputs.Validate(); err != nil {
      return err
    }
  }
  if s.RuntimeResource != nil {
    if err := s.RuntimeResource.Validate(); err != nil {
      return err
    }
  }
  if s.Script != nil {
    if err := s.Script.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAdhocWorkflowInstanceRequestTasksDataSource struct {
  // The name of the data source.
  // 
  // example:
  // 
  // mysql_test
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksDataSource) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksDataSource) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksDataSource) GetName() *string  {
  return s.Name
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksDataSource) SetName(v string) *ExecuteAdhocWorkflowInstanceRequestTasksDataSource {
  s.Name = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksDataSource) Validate() error {
  return dara.Validate(s)
}

type ExecuteAdhocWorkflowInstanceRequestTasksDependencies struct {
  // The identifier of the output of the ancestor task.
  // 
  // example:
  // 
  // pre.odps_sql_demo_0
  UpstreamOutput *string `json:"UpstreamOutput,omitempty" xml:"UpstreamOutput,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksDependencies) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksDependencies) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksDependencies) GetUpstreamOutput() *string  {
  return s.UpstreamOutput
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksDependencies) SetUpstreamOutput(v string) *ExecuteAdhocWorkflowInstanceRequestTasksDependencies {
  s.UpstreamOutput = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksDependencies) Validate() error {
  return dara.Validate(s)
}

type ExecuteAdhocWorkflowInstanceRequestTasksInputs struct {
  // The variables.
  Variables []*ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksInputs) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksInputs) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksInputs) GetVariables() []*ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables  {
  return s.Variables
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksInputs) SetVariables(v []*ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables) *ExecuteAdhocWorkflowInstanceRequestTasksInputs {
  s.Variables = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksInputs) Validate() error {
  if s.Variables != nil {
    for _, item := range s.Variables {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables struct {
  // The name of the variable.
  // 
  // example:
  // 
  // key1
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The value of the variable. You must configure this parameter in the `The ancestor output: The output variable name of the ancestor task` format.
  // 
  // example:
  // 
  // Value1
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables) GetName() *string  {
  return s.Name
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables) GetValue() *string  {
  return s.Value
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables) SetName(v string) *ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables {
  s.Name = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables) SetValue(v string) *ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables {
  s.Value = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksInputsVariables) Validate() error {
  return dara.Validate(s)
}

type ExecuteAdhocWorkflowInstanceRequestTasksOutputs struct {
  // The task outputs.
  TaskOutputs []*ExecuteAdhocWorkflowInstanceRequestTasksOutputsTaskOutputs `json:"TaskOutputs,omitempty" xml:"TaskOutputs,omitempty" type:"Repeated"`
  // The variables.
  Variables []*ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksOutputs) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksOutputs) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputs) GetTaskOutputs() []*ExecuteAdhocWorkflowInstanceRequestTasksOutputsTaskOutputs  {
  return s.TaskOutputs
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputs) GetVariables() []*ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables  {
  return s.Variables
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputs) SetTaskOutputs(v []*ExecuteAdhocWorkflowInstanceRequestTasksOutputsTaskOutputs) *ExecuteAdhocWorkflowInstanceRequestTasksOutputs {
  s.TaskOutputs = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputs) SetVariables(v []*ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables) *ExecuteAdhocWorkflowInstanceRequestTasksOutputs {
  s.Variables = v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputs) Validate() error {
  if s.TaskOutputs != nil {
    for _, item := range s.TaskOutputs {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.Variables != nil {
    for _, item := range s.Variables {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteAdhocWorkflowInstanceRequestTasksOutputsTaskOutputs struct {
  // The identifier of the output.
  // 
  // example:
  // 
  // pre.odps_sql_demo_0
  Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksOutputsTaskOutputs) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksOutputsTaskOutputs) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputsTaskOutputs) GetOutput() *string  {
  return s.Output
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputsTaskOutputs) SetOutput(v string) *ExecuteAdhocWorkflowInstanceRequestTasksOutputsTaskOutputs {
  s.Output = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputsTaskOutputs) Validate() error {
  return dara.Validate(s)
}

type ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables struct {
  // The name of the variable.
  // 
  // example:
  // 
  // key1
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The type. Valid values:
  // 
  // 	- System
  // 
  // 	- Constant
  // 
  // 	- NodeOutput
  // 
  // 	- PassThrough
  // 
  // example:
  // 
  // Constant
  Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
  // The value of the variable.
  // 
  // example:
  // 
  // value1
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables) GetName() *string  {
  return s.Name
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables) GetType() *string  {
  return s.Type
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables) GetValue() *string  {
  return s.Value
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables) SetName(v string) *ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables {
  s.Name = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables) SetType(v string) *ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables {
  s.Type = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables) SetValue(v string) *ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables {
  s.Value = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksOutputsVariables) Validate() error {
  return dara.Validate(s)
}

type ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource struct {
  // The default number of compute units (CUs) configured for task running.
  // 
  // example:
  // 
  // 0.25
  Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
  // The ID of the image configured for task running.
  // 
  // example:
  // 
  // i-xxxxxx
  Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
  // The ID of the resource group for scheduling configured for task running.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // S_res_group_524258031846018_1684XXXXXXXXX
  ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource) GetCu() *string  {
  return s.Cu
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource) GetImage() *string  {
  return s.Image
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource) GetResourceGroupId() *string  {
  return s.ResourceGroupId
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource) SetCu(v string) *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource {
  s.Cu = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource) SetImage(v string) *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource {
  s.Image = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource) SetResourceGroupId(v string) *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource {
  s.ResourceGroupId = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksRuntimeResource) Validate() error {
  return dara.Validate(s)
}

type ExecuteAdhocWorkflowInstanceRequestTasksScript struct {
  // The script content.
  // 
  // example:
  // 
  // echo "helloWorld"
  Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
  // The script parameters.
  // 
  // example:
  // 
  // para1=$bizdate
  Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksScript) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceRequestTasksScript) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksScript) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksScript) GetParameters() *string  {
  return s.Parameters
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksScript) SetContent(v string) *ExecuteAdhocWorkflowInstanceRequestTasksScript {
  s.Content = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksScript) SetParameters(v string) *ExecuteAdhocWorkflowInstanceRequestTasksScript {
  s.Parameters = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceRequestTasksScript) Validate() error {
  return dara.Validate(s)
}

