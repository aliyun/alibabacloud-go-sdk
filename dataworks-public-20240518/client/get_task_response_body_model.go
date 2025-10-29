// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
	SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody
	GetTask() *GetTaskResponseBodyTask
}

type GetTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the task.
	Task *GetTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) GetTask() *GetTaskResponseBodyTask {
	return s.Task
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody {
	s.Task = v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResponseBodyTask struct {
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The account ID of the creator.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The information about the associated data source.
	DataSource *GetTaskResponseBodyTaskDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The dependency information.
	Dependencies []*GetTaskResponseBodyTaskDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	// The description of the task.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The instance ID.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input information.
	//
	// if can be null:
	// false
	Inputs *GetTaskResponseBodyTaskInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The instance generation mode. Valid values:
	//
	// 	- T+1
	//
	// 	- Immediately
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1710239005403
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account ID of the modifier.
	//
	// example:
	//
	// 1000
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// SQL node
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output information.
	Outputs *GetTaskResponseBodyTaskOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// The account ID of the task owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the task. Valid values: 1 to 8. A larger value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Deprecated
	//
	// The environment of the workspace. This parameter is deprecated and replaced by the EnvType parameter. Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
	//
	// example:
	//
	// Prod
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The rerun interval. Unit: seconds.
	//
	// example:
	//
	// 60
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// The rerun mode. Valid values:
	//
	// 	- AllDenied: The task cannot be rerun regardless of whether the task is successfully run or fails to be run.
	//
	// 	- FailureAllowed: The task can be rerun only after it fails to be run.
	//
	// 	- AllAllowed: The task can be rerun regardless of whether the task is successfully run or fails to be run.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of times that the task is rerun. This parameter takes effect only if the RerunMode parameter is set to AllAllowed or FailureAllowed.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The configurations of the runtime environment, such as the resource group information.
	RuntimeResource *GetTaskResponseBodyTaskRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The script information.
	Script *GetTaskResponseBodyTaskScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The configurations of the subtasks, such as a do-while node.
	SubTasks *GetTaskResponseBodyTaskSubTasks `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Struct"`
	// The tags.
	Tags []*GetTaskResponseBodyTaskTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period of task running. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The method to trigger task scheduling.
	Trigger *GetTaskResponseBodyTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The type of the task.
	//
	// example:
	//
	// ODPS_SQL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTask) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetTaskResponseBodyTask) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetTaskResponseBodyTask) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetTaskResponseBodyTask) GetDataSource() *GetTaskResponseBodyTaskDataSource {
	return s.DataSource
}

func (s *GetTaskResponseBodyTask) GetDependencies() []*GetTaskResponseBodyTaskDependencies {
	return s.Dependencies
}

func (s *GetTaskResponseBodyTask) GetDescription() *string {
	return s.Description
}

func (s *GetTaskResponseBodyTask) GetEnvType() *string {
	return s.EnvType
}

func (s *GetTaskResponseBodyTask) GetId() *int64 {
	return s.Id
}

func (s *GetTaskResponseBodyTask) GetInputs() *GetTaskResponseBodyTaskInputs {
	return s.Inputs
}

func (s *GetTaskResponseBodyTask) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *GetTaskResponseBodyTask) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetTaskResponseBodyTask) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetTaskResponseBodyTask) GetName() *string {
	return s.Name
}

func (s *GetTaskResponseBodyTask) GetOutputs() *GetTaskResponseBodyTaskOutputs {
	return s.Outputs
}

func (s *GetTaskResponseBodyTask) GetOwner() *string {
	return s.Owner
}

func (s *GetTaskResponseBodyTask) GetPriority() *int32 {
	return s.Priority
}

func (s *GetTaskResponseBodyTask) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetTaskResponseBodyTask) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetTaskResponseBodyTask) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *GetTaskResponseBodyTask) GetRerunMode() *string {
	return s.RerunMode
}

func (s *GetTaskResponseBodyTask) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *GetTaskResponseBodyTask) GetRuntimeResource() *GetTaskResponseBodyTaskRuntimeResource {
	return s.RuntimeResource
}

func (s *GetTaskResponseBodyTask) GetScript() *GetTaskResponseBodyTaskScript {
	return s.Script
}

func (s *GetTaskResponseBodyTask) GetSubTasks() *GetTaskResponseBodyTaskSubTasks {
	return s.SubTasks
}

func (s *GetTaskResponseBodyTask) GetTags() []*GetTaskResponseBodyTaskTags {
	return s.Tags
}

func (s *GetTaskResponseBodyTask) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetTaskResponseBodyTask) GetTrigger() *GetTaskResponseBodyTaskTrigger {
	return s.Trigger
}

func (s *GetTaskResponseBodyTask) GetType() *string {
	return s.Type
}

func (s *GetTaskResponseBodyTask) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetTaskResponseBodyTask) SetBaselineId(v int64) *GetTaskResponseBodyTask {
	s.BaselineId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetCreateTime(v int64) *GetTaskResponseBodyTask {
	s.CreateTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetCreateUser(v string) *GetTaskResponseBodyTask {
	s.CreateUser = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetDataSource(v *GetTaskResponseBodyTaskDataSource) *GetTaskResponseBodyTask {
	s.DataSource = v
	return s
}

func (s *GetTaskResponseBodyTask) SetDependencies(v []*GetTaskResponseBodyTaskDependencies) *GetTaskResponseBodyTask {
	s.Dependencies = v
	return s
}

func (s *GetTaskResponseBodyTask) SetDescription(v string) *GetTaskResponseBodyTask {
	s.Description = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetEnvType(v string) *GetTaskResponseBodyTask {
	s.EnvType = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetId(v int64) *GetTaskResponseBodyTask {
	s.Id = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetInputs(v *GetTaskResponseBodyTaskInputs) *GetTaskResponseBodyTask {
	s.Inputs = v
	return s
}

func (s *GetTaskResponseBodyTask) SetInstanceMode(v string) *GetTaskResponseBodyTask {
	s.InstanceMode = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetModifyTime(v int64) *GetTaskResponseBodyTask {
	s.ModifyTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetModifyUser(v string) *GetTaskResponseBodyTask {
	s.ModifyUser = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetName(v string) *GetTaskResponseBodyTask {
	s.Name = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetOutputs(v *GetTaskResponseBodyTaskOutputs) *GetTaskResponseBodyTask {
	s.Outputs = v
	return s
}

func (s *GetTaskResponseBodyTask) SetOwner(v string) *GetTaskResponseBodyTask {
	s.Owner = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetPriority(v int32) *GetTaskResponseBodyTask {
	s.Priority = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetProjectEnv(v string) *GetTaskResponseBodyTask {
	s.ProjectEnv = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetProjectId(v int64) *GetTaskResponseBodyTask {
	s.ProjectId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetRerunInterval(v int32) *GetTaskResponseBodyTask {
	s.RerunInterval = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetRerunMode(v string) *GetTaskResponseBodyTask {
	s.RerunMode = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetRerunTimes(v int32) *GetTaskResponseBodyTask {
	s.RerunTimes = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetRuntimeResource(v *GetTaskResponseBodyTaskRuntimeResource) *GetTaskResponseBodyTask {
	s.RuntimeResource = v
	return s
}

func (s *GetTaskResponseBodyTask) SetScript(v *GetTaskResponseBodyTaskScript) *GetTaskResponseBodyTask {
	s.Script = v
	return s
}

func (s *GetTaskResponseBodyTask) SetSubTasks(v *GetTaskResponseBodyTaskSubTasks) *GetTaskResponseBodyTask {
	s.SubTasks = v
	return s
}

func (s *GetTaskResponseBodyTask) SetTags(v []*GetTaskResponseBodyTaskTags) *GetTaskResponseBodyTask {
	s.Tags = v
	return s
}

func (s *GetTaskResponseBodyTask) SetTimeout(v int32) *GetTaskResponseBodyTask {
	s.Timeout = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTrigger(v *GetTaskResponseBodyTaskTrigger) *GetTaskResponseBodyTask {
	s.Trigger = v
	return s
}

func (s *GetTaskResponseBodyTask) SetType(v string) *GetTaskResponseBodyTask {
	s.Type = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetWorkflowId(v int64) *GetTaskResponseBodyTask {
	s.WorkflowId = &v
	return s
}

func (s *GetTaskResponseBodyTask) Validate() error {
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
	if s.SubTasks != nil {
		if err := s.SubTasks.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResponseBodyTaskDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTaskResponseBodyTaskDataSource) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskDataSource) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDataSource) GetName() *string {
	return s.Name
}

func (s *GetTaskResponseBodyTaskDataSource) SetName(v string) *GetTaskResponseBodyTaskDataSource {
	s.Name = &v
	return s
}

func (s *GetTaskResponseBodyTaskDataSource) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskDependencies struct {
	// The dependency type. Valid values:
	//
	// 	- CrossCycleDependsOnChildren: cross-cycle dependency on level-1 descendant nodes
	//
	// 	- CrossCycleDependsOnSelf: cross-cycle dependency on the current node
	//
	// 	- CrossCycleDependsOnOtherNode: cross-cycle dependency on other nodes
	//
	// 	- Normal: same-cycle scheduling dependency
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The identifier of the output of the ancestor task. This parameter is returned only if `same-cycle scheduling dependencies` and the node input are configured.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	UpstreamOutput *string `json:"UpstreamOutput,omitempty" xml:"UpstreamOutput,omitempty"`
	// The ancestor task ID. This parameter is returned only if `cross-cycle scheduling dependencies` or `same-cycle scheduling dependencies` and the node input are not configured.
	//
	// example:
	//
	// 1234
	UpstreamTaskId *string `json:"UpstreamTaskId,omitempty" xml:"UpstreamTaskId,omitempty"`
}

func (s GetTaskResponseBodyTaskDependencies) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskDependencies) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDependencies) GetType() *string {
	return s.Type
}

func (s *GetTaskResponseBodyTaskDependencies) GetUpstreamOutput() *string {
	return s.UpstreamOutput
}

func (s *GetTaskResponseBodyTaskDependencies) GetUpstreamTaskId() *string {
	return s.UpstreamTaskId
}

func (s *GetTaskResponseBodyTaskDependencies) SetType(v string) *GetTaskResponseBodyTaskDependencies {
	s.Type = &v
	return s
}

func (s *GetTaskResponseBodyTaskDependencies) SetUpstreamOutput(v string) *GetTaskResponseBodyTaskDependencies {
	s.UpstreamOutput = &v
	return s
}

func (s *GetTaskResponseBodyTaskDependencies) SetUpstreamTaskId(v string) *GetTaskResponseBodyTaskDependencies {
	s.UpstreamTaskId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDependencies) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskInputs struct {
	// The variables.
	Variables []*GetTaskResponseBodyTaskInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s GetTaskResponseBodyTaskInputs) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskInputs) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskInputs) GetVariables() []*GetTaskResponseBodyTaskInputsVariables {
	return s.Variables
}

func (s *GetTaskResponseBodyTaskInputs) SetVariables(v []*GetTaskResponseBodyTaskInputsVariables) *GetTaskResponseBodyTaskInputs {
	s.Variables = v
	return s
}

func (s *GetTaskResponseBodyTaskInputs) Validate() error {
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

type GetTaskResponseBodyTaskInputsVariables struct {
	// The name of the variable.
	//
	// example:
	//
	// key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type. Valid values:
	//
	// 	- Constant: constant
	//
	// 	- PassThrough: node output
	//
	// 	- System: variable
	//
	// 	- NodeOutput: script output
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// Value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTaskResponseBodyTaskInputsVariables) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskInputsVariables) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskInputsVariables) GetName() *string {
	return s.Name
}

func (s *GetTaskResponseBodyTaskInputsVariables) GetType() *string {
	return s.Type
}

func (s *GetTaskResponseBodyTaskInputsVariables) GetValue() *string {
	return s.Value
}

func (s *GetTaskResponseBodyTaskInputsVariables) SetName(v string) *GetTaskResponseBodyTaskInputsVariables {
	s.Name = &v
	return s
}

func (s *GetTaskResponseBodyTaskInputsVariables) SetType(v string) *GetTaskResponseBodyTaskInputsVariables {
	s.Type = &v
	return s
}

func (s *GetTaskResponseBodyTaskInputsVariables) SetValue(v string) *GetTaskResponseBodyTaskInputsVariables {
	s.Value = &v
	return s
}

func (s *GetTaskResponseBodyTaskInputsVariables) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskOutputs struct {
	// The task outputs.
	TaskOutputs []*GetTaskResponseBodyTaskOutputsTaskOutputs `json:"TaskOutputs,omitempty" xml:"TaskOutputs,omitempty" type:"Repeated"`
	// The variables.
	Variables []*GetTaskResponseBodyTaskOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s GetTaskResponseBodyTaskOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskOutputs) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskOutputs) GetTaskOutputs() []*GetTaskResponseBodyTaskOutputsTaskOutputs {
	return s.TaskOutputs
}

func (s *GetTaskResponseBodyTaskOutputs) GetVariables() []*GetTaskResponseBodyTaskOutputsVariables {
	return s.Variables
}

func (s *GetTaskResponseBodyTaskOutputs) SetTaskOutputs(v []*GetTaskResponseBodyTaskOutputsTaskOutputs) *GetTaskResponseBodyTaskOutputs {
	s.TaskOutputs = v
	return s
}

func (s *GetTaskResponseBodyTaskOutputs) SetVariables(v []*GetTaskResponseBodyTaskOutputsVariables) *GetTaskResponseBodyTaskOutputs {
	s.Variables = v
	return s
}

func (s *GetTaskResponseBodyTaskOutputs) Validate() error {
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

type GetTaskResponseBodyTaskOutputsTaskOutputs struct {
	// The identifier of the output.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s GetTaskResponseBodyTaskOutputsTaskOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskOutputsTaskOutputs) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskOutputsTaskOutputs) GetOutput() *string {
	return s.Output
}

func (s *GetTaskResponseBodyTaskOutputsTaskOutputs) SetOutput(v string) *GetTaskResponseBodyTaskOutputsTaskOutputs {
	s.Output = &v
	return s
}

func (s *GetTaskResponseBodyTaskOutputsTaskOutputs) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskOutputsVariables struct {
	// The name of the variable.
	//
	// example:
	//
	// key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type. Valid values:
	//
	// 	- Constant: constant
	//
	// 	- PassThrough: node output
	//
	// 	- System: variable
	//
	// 	- NodeOutput: script output
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

func (s GetTaskResponseBodyTaskOutputsVariables) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskOutputsVariables) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskOutputsVariables) GetName() *string {
	return s.Name
}

func (s *GetTaskResponseBodyTaskOutputsVariables) GetType() *string {
	return s.Type
}

func (s *GetTaskResponseBodyTaskOutputsVariables) GetValue() *string {
	return s.Value
}

func (s *GetTaskResponseBodyTaskOutputsVariables) SetName(v string) *GetTaskResponseBodyTaskOutputsVariables {
	s.Name = &v
	return s
}

func (s *GetTaskResponseBodyTaskOutputsVariables) SetType(v string) *GetTaskResponseBodyTaskOutputsVariables {
	s.Type = &v
	return s
}

func (s *GetTaskResponseBodyTaskOutputsVariables) SetValue(v string) *GetTaskResponseBodyTaskOutputsVariables {
	s.Value = &v
	return s
}

func (s *GetTaskResponseBodyTaskOutputsVariables) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskRuntimeResource struct {
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
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetTaskResponseBodyTaskRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskRuntimeResource) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *GetTaskResponseBodyTaskRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *GetTaskResponseBodyTaskRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTaskResponseBodyTaskRuntimeResource) SetCu(v string) *GetTaskResponseBodyTaskRuntimeResource {
	s.Cu = &v
	return s
}

func (s *GetTaskResponseBodyTaskRuntimeResource) SetImage(v string) *GetTaskResponseBodyTaskRuntimeResource {
	s.Image = &v
	return s
}

func (s *GetTaskResponseBodyTaskRuntimeResource) SetResourceGroupId(v string) *GetTaskResponseBodyTaskRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTaskResponseBodyTaskRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskScript struct {
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

func (s GetTaskResponseBodyTaskScript) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskScript) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskScript) GetContent() *string {
	return s.Content
}

func (s *GetTaskResponseBodyTaskScript) GetParameters() *string {
	return s.Parameters
}

func (s *GetTaskResponseBodyTaskScript) SetContent(v string) *GetTaskResponseBodyTaskScript {
	s.Content = &v
	return s
}

func (s *GetTaskResponseBodyTaskScript) SetParameters(v string) *GetTaskResponseBodyTaskScript {
	s.Parameters = &v
	return s
}

func (s *GetTaskResponseBodyTaskScript) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskSubTasks struct {
	// The subtasks.
	SubTasks []*GetTaskResponseBodyTaskSubTasksSubTasks `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
	// The type of the subtask. Valid values:
	//
	// 	- DoWhile: do-while node
	//
	// 	- Combined: node group
	//
	// 	- ForEach: for-each node
	//
	// example:
	//
	// Combined
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTaskResponseBodyTaskSubTasks) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskSubTasks) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskSubTasks) GetSubTasks() []*GetTaskResponseBodyTaskSubTasksSubTasks {
	return s.SubTasks
}

func (s *GetTaskResponseBodyTaskSubTasks) GetType() *string {
	return s.Type
}

func (s *GetTaskResponseBodyTaskSubTasks) SetSubTasks(v []*GetTaskResponseBodyTaskSubTasksSubTasks) *GetTaskResponseBodyTaskSubTasks {
	s.SubTasks = v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasks) SetType(v string) *GetTaskResponseBodyTaskSubTasks {
	s.Type = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasks) Validate() error {
	if s.SubTasks != nil {
		for _, item := range s.SubTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskResponseBodyTaskSubTasksSubTasks struct {
	// The baseline ID.
	//
	// example:
	//
	// The baseline ID.
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The account ID of the creator.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The information about the associated data source.
	DataSource *GetTaskResponseBodyTaskSubTasksSubTasksDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description of the task.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The task ID.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1710239005403
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account ID of the modifier.
	//
	// example:
	//
	// 1000
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// SQL node
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account ID of the task owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the task. Valid values: 1 to 8. A larger value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// Deprecated
	//
	// The environment of the workspace. This parameter is deprecated and replaced by the EnvType parameter. Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
	//
	// example:
	//
	// Prod
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The rerun interval. Unit: seconds.
	//
	// example:
	//
	// 180
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// The rerun mode. Valid values:
	//
	// 	- AllDenied: The task cannot be rerun regardless of whether the task is successfully run or fails to be run.
	//
	// 	- FailureAllowed: The task can be rerun only after it fails to be run.
	//
	// 	- AllAllowed: The task can be rerun regardless of whether the task is successfully run or fails to be run.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of times that the task is rerun. This parameter takes effect only if the RerunMode parameter is set to AllAllowed or FailureAllowed.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The runtime environment configuration of the task, such as the resource group.
	RuntimeResource *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The timeout period of task running. Unit: seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The method to trigger task scheduling.
	Trigger *GetTaskResponseBodyTaskSubTasksSubTasksTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The type of the task.
	//
	// example:
	//
	// ODPS_SQL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the workflow to which the task belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetTaskResponseBodyTaskSubTasksSubTasks) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskSubTasksSubTasks) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetDataSource() *GetTaskResponseBodyTaskSubTasksSubTasksDataSource {
	return s.DataSource
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetDescription() *string {
	return s.Description
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetEnvType() *string {
	return s.EnvType
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetId() *int64 {
	return s.Id
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetName() *string {
	return s.Name
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetOwner() *string {
	return s.Owner
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetPriority() *int32 {
	return s.Priority
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetRerunMode() *string {
	return s.RerunMode
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetRuntimeResource() *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource {
	return s.RuntimeResource
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetTrigger() *GetTaskResponseBodyTaskSubTasksSubTasksTrigger {
	return s.Trigger
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetType() *string {
	return s.Type
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetBaselineId(v int64) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.BaselineId = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetCreateTime(v int64) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.CreateTime = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetCreateUser(v string) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.CreateUser = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetDataSource(v *GetTaskResponseBodyTaskSubTasksSubTasksDataSource) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.DataSource = v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetDescription(v string) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.Description = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetEnvType(v string) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.EnvType = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetId(v int64) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.Id = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetModifyTime(v int64) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.ModifyTime = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetModifyUser(v string) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.ModifyUser = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetName(v string) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.Name = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetOwner(v string) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.Owner = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetPriority(v int32) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.Priority = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetProjectEnv(v string) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.ProjectEnv = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetProjectId(v int64) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.ProjectId = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetRerunInterval(v int32) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.RerunInterval = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetRerunMode(v string) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.RerunMode = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetRerunTimes(v int32) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.RerunTimes = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetRuntimeResource(v *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.RuntimeResource = v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetTimeout(v int32) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.Timeout = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetTrigger(v *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.Trigger = v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetType(v string) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.Type = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) SetWorkflowId(v int64) *GetTaskResponseBodyTaskSubTasksSubTasks {
	s.WorkflowId = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasks) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResponseBodyTaskSubTasksSubTasksDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTaskResponseBodyTaskSubTasksSubTasksDataSource) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskSubTasksSubTasksDataSource) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksDataSource) GetName() *string {
	return s.Name
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksDataSource) SetName(v string) *GetTaskResponseBodyTaskSubTasksSubTasksDataSource {
	s.Name = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksDataSource) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource struct {
	// The default number of CUs configured for task running.
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
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource) SetCu(v string) *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource {
	s.Cu = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource) SetImage(v string) *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource {
	s.Image = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource) SetResourceGroupId(v string) *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskSubTasksSubTasksTrigger struct {
	// The CRON expression of the task. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The end time of the time range during which the task is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The running mode of the task after it is triggered. This parameter takes effect only if the Type parameter is set to Scheduler. Valid values:
	//
	// 	- Pause
	//
	// 	- Skip
	//
	// 	- Normal
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The start time of the time range during which the task is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trigger type. Valid values:
	//
	// 	- Scheduler: periodic scheduling
	//
	// 	- Manual: manual scheduling
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTaskResponseBodyTaskSubTasksSubTasksTrigger) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskSubTasksSubTasksTrigger) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) GetCron() *string {
	return s.Cron
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) GetType() *string {
	return s.Type
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) SetCron(v string) *GetTaskResponseBodyTaskSubTasksSubTasksTrigger {
	s.Cron = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) SetEndTime(v string) *GetTaskResponseBodyTaskSubTasksSubTasksTrigger {
	s.EndTime = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) SetRecurrence(v string) *GetTaskResponseBodyTaskSubTasksSubTasksTrigger {
	s.Recurrence = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) SetStartTime(v string) *GetTaskResponseBodyTaskSubTasksSubTasksTrigger {
	s.StartTime = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) SetType(v string) *GetTaskResponseBodyTaskSubTasksSubTasksTrigger {
	s.Type = &v
	return s
}

func (s *GetTaskResponseBodyTaskSubTasksSubTasksTrigger) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTaskResponseBodyTaskTags) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskTags) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskTags) GetKey() *string {
	return s.Key
}

func (s *GetTaskResponseBodyTaskTags) GetValue() *string {
	return s.Value
}

func (s *GetTaskResponseBodyTaskTags) SetKey(v string) *GetTaskResponseBodyTaskTags {
	s.Key = &v
	return s
}

func (s *GetTaskResponseBodyTaskTags) SetValue(v string) *GetTaskResponseBodyTaskTags {
	s.Value = &v
	return s
}

func (s *GetTaskResponseBodyTaskTags) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskTrigger struct {
	// The CRON expression of the task. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The end time of the time range during which the task is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The running mode of the task after it is triggered. This parameter takes effect only if the Type parameter is set to Scheduler. Valid values:
	//
	// 	- Pause
	//
	// 	- Skip
	//
	// 	- Normal
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The start time of the time range during which the task is periodically scheduled. This parameter takes effect only if the Type parameter is set to Scheduler.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trigger type. Valid values:
	//
	// 	- Scheduler: periodic scheduling
	//
	// 	- Manual: manual scheduling
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTaskResponseBodyTaskTrigger) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskTrigger) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskTrigger) GetCron() *string {
	return s.Cron
}

func (s *GetTaskResponseBodyTaskTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTaskResponseBodyTaskTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *GetTaskResponseBodyTaskTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTaskResponseBodyTaskTrigger) GetType() *string {
	return s.Type
}

func (s *GetTaskResponseBodyTaskTrigger) SetCron(v string) *GetTaskResponseBodyTaskTrigger {
	s.Cron = &v
	return s
}

func (s *GetTaskResponseBodyTaskTrigger) SetEndTime(v string) *GetTaskResponseBodyTaskTrigger {
	s.EndTime = &v
	return s
}

func (s *GetTaskResponseBodyTaskTrigger) SetRecurrence(v string) *GetTaskResponseBodyTaskTrigger {
	s.Recurrence = &v
	return s
}

func (s *GetTaskResponseBodyTaskTrigger) SetStartTime(v string) *GetTaskResponseBodyTaskTrigger {
	s.StartTime = &v
	return s
}

func (s *GetTaskResponseBodyTaskTrigger) SetType(v string) *GetTaskResponseBodyTaskTrigger {
	s.Type = &v
	return s
}

func (s *GetTaskResponseBodyTaskTrigger) Validate() error {
	return dara.Validate(s)
}
