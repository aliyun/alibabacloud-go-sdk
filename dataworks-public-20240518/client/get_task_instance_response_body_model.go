// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTaskInstanceResponseBody
	GetRequestId() *string
	SetTaskInstance(v *GetTaskInstanceResponseBodyTaskInstance) *GetTaskInstanceResponseBody
	GetTaskInstance() *GetTaskInstanceResponseBodyTaskInstance
}

type GetTaskInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the task instance.
	TaskInstance *GetTaskInstanceResponseBodyTaskInstance `json:"TaskInstance,omitempty" xml:"TaskInstance,omitempty" type:"Struct"`
}

func (s GetTaskInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskInstanceResponseBody) GetTaskInstance() *GetTaskInstanceResponseBodyTaskInstance {
	return s.TaskInstance
}

func (s *GetTaskInstanceResponseBody) SetRequestId(v string) *GetTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskInstanceResponseBody) SetTaskInstance(v *GetTaskInstanceResponseBodyTaskInstance) *GetTaskInstanceResponseBody {
	s.TaskInstance = v
	return s
}

func (s *GetTaskInstanceResponseBody) Validate() error {
	if s.TaskInstance != nil {
		if err := s.TaskInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskInstanceResponseBodyTaskInstance struct {
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The data timestamp.
	//
	// example:
	//
	// 1710239005403
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
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
	DataSource *GetTaskInstanceResponseBodyTaskInstanceDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the instance finished running.
	//
	// example:
	//
	// 1710239005403
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input information.
	Inputs *GetTaskInstanceResponseBodyTaskInstanceInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
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
	// The output information.
	Outputs *GetTaskInstanceResponseBodyTaskInstanceOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// The account ID of the task owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The sequence number of the cycle. This parameter indicates the cycle of the task instance on the current day.
	//
	// example:
	//
	// 1
	PeriodNumber *int32 `json:"PeriodNumber,omitempty" xml:"PeriodNumber,omitempty"`
	// The task priority. Valid values: 1 to 8. A larger value indicates a higher priority. Default value: 1.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The environment of the workspace. Valid values:
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
	// The rerun mode. Valid values:
	//
	// 	- AllDenied: The task cannot be rerun regardless of whether the task is successfully run or fails to run.
	//
	// 	- AllAllowed: The task can be rerun regardless of whether the task is successfully run or fails to run.
	//
	// 	- FailureAllowed: The task can be rerun only after it fails to run.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of times the instance is run. By default, the value starts from 1.
	//
	// example:
	//
	// 1
	RunNumber *int32 `json:"RunNumber,omitempty" xml:"RunNumber,omitempty"`
	// The runtime information about the instance.
	Runtime *GetTaskInstanceResponseBodyTaskInstanceRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
	// The information about the resource group with which the instance is associated.
	RuntimeResource *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The script information.
	Script *GetTaskInstanceResponseBodyTaskInstanceScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The time when the instance started to run.
	//
	// example:
	//
	// 1710239005403
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- NotRun: The instance is not run.
	//
	// 	- Running: The instance is running.
	//
	// 	- WaitTime: The instance is waiting for the scheduling time to arrive.
	//
	// 	- CheckingCondition: Branch conditions are being checked for the instance.
	//
	// 	- WaitResource: The instance is waiting for resources.
	//
	// 	- Failure: The instance fails to be run.
	//
	// 	- Success: The instance is successfully run.
	//
	// 	- Checking: Data quality is being checked for the instance.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the task.
	Tags []*GetTaskInstanceResponseBodyTaskInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the task for which the instance is generated.
	//
	// example:
	//
	// 1234
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task for which the instance is generated.
	//
	// example:
	//
	// SQL node
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task for which the instance is generated.
	//
	// example:
	//
	// ODPS_SQL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The timeout period of task running. Unit: seconds.
	//
	// Note: The value of this parameter is rounded up by hour.
	//
	// example:
	//
	// 1
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The running mode of the instance after it is triggered. This parameter takes effect only if the TriggerType parameter is set to Scheduler. Valid values:
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
	TriggerRecurrence *string `json:"TriggerRecurrence,omitempty" xml:"TriggerRecurrence,omitempty"`
	// The scheduling time.
	//
	// example:
	//
	// 1710239005403
	TriggerTime *int64 `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
	// The method to trigger instance scheduling. The value of the Trigger.Type parameter in the response of the GetTask operation is used. Valid values:
	//
	// 	- Scheduler
	//
	// 	- Manual
	//
	// example:
	//
	// Scheduler
	TriggerType               *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	UnifiedWorkflowInstanceId *int64  `json:"UnifiedWorkflowInstanceId,omitempty" xml:"UnifiedWorkflowInstanceId,omitempty"`
	// example:
	//
	// 1710239005403
	WaitingResourceTime *int64 `json:"WaitingResourceTime,omitempty" xml:"WaitingResourceTime,omitempty"`
	// example:
	//
	// 1710239005403
	WaitingTriggerTime *int64 `json:"WaitingTriggerTime,omitempty" xml:"WaitingTriggerTime,omitempty"`
	// The ID of the workflow to which the instance belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The workflow instance ID.
	//
	// example:
	//
	// 1234
	WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
	// The type of the workflow instance. Valid values:
	//
	// 	- SmokeTest
	//
	// 	- SupplementData
	//
	// 	- Manual
	//
	// 	- ManualWorkflow
	//
	// 	- Normal
	//
	// 	- ManualFlow
	//
	// example:
	//
	// Normal
	WorkflowInstanceType *string `json:"WorkflowInstanceType,omitempty" xml:"WorkflowInstanceType,omitempty"`
	// The name of the workflow to which the instance belongs.
	//
	// example:
	//
	// Test workflow
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
}

func (s GetTaskInstanceResponseBodyTaskInstance) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstance) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetDataSource() *GetTaskInstanceResponseBodyTaskInstanceDataSource {
	return s.DataSource
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetDescription() *string {
	return s.Description
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetId() *int64 {
	return s.Id
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetInputs() *GetTaskInstanceResponseBodyTaskInstanceInputs {
	return s.Inputs
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetOutputs() *GetTaskInstanceResponseBodyTaskInstanceOutputs {
	return s.Outputs
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetOwner() *string {
	return s.Owner
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetPeriodNumber() *int32 {
	return s.PeriodNumber
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetPriority() *int32 {
	return s.Priority
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetRerunMode() *string {
	return s.RerunMode
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetRunNumber() *int32 {
	return s.RunNumber
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetRuntime() *GetTaskInstanceResponseBodyTaskInstanceRuntime {
	return s.Runtime
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetRuntimeResource() *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource {
	return s.RuntimeResource
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetScript() *GetTaskInstanceResponseBodyTaskInstanceScript {
	return s.Script
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetStartedTime() *int64 {
	return s.StartedTime
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetStatus() *string {
	return s.Status
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetTags() []*GetTaskInstanceResponseBodyTaskInstanceTags {
	return s.Tags
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetTaskName() *string {
	return s.TaskName
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetTaskType() *string {
	return s.TaskType
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetTriggerTime() *int64 {
	return s.TriggerTime
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetUnifiedWorkflowInstanceId() *int64 {
	return s.UnifiedWorkflowInstanceId
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetWaitingResourceTime() *int64 {
	return s.WaitingResourceTime
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetWaitingTriggerTime() *int64 {
	return s.WaitingTriggerTime
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetWorkflowInstanceId() *int64 {
	return s.WorkflowInstanceId
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetWorkflowInstanceType() *string {
	return s.WorkflowInstanceType
}

func (s *GetTaskInstanceResponseBodyTaskInstance) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetBaselineId(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.BaselineId = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetBizdate(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.Bizdate = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetCreateTime(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.CreateTime = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetCreateUser(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.CreateUser = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetDataSource(v *GetTaskInstanceResponseBodyTaskInstanceDataSource) *GetTaskInstanceResponseBodyTaskInstance {
	s.DataSource = v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetDescription(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.Description = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetFinishedTime(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.FinishedTime = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetId(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.Id = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetInputs(v *GetTaskInstanceResponseBodyTaskInstanceInputs) *GetTaskInstanceResponseBodyTaskInstance {
	s.Inputs = v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetModifyTime(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.ModifyTime = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetModifyUser(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.ModifyUser = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetOutputs(v *GetTaskInstanceResponseBodyTaskInstanceOutputs) *GetTaskInstanceResponseBodyTaskInstance {
	s.Outputs = v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetOwner(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.Owner = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetPeriodNumber(v int32) *GetTaskInstanceResponseBodyTaskInstance {
	s.PeriodNumber = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetPriority(v int32) *GetTaskInstanceResponseBodyTaskInstance {
	s.Priority = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetProjectEnv(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.ProjectEnv = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetProjectId(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.ProjectId = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetRerunMode(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.RerunMode = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetRunNumber(v int32) *GetTaskInstanceResponseBodyTaskInstance {
	s.RunNumber = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetRuntime(v *GetTaskInstanceResponseBodyTaskInstanceRuntime) *GetTaskInstanceResponseBodyTaskInstance {
	s.Runtime = v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetRuntimeResource(v *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource) *GetTaskInstanceResponseBodyTaskInstance {
	s.RuntimeResource = v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetScript(v *GetTaskInstanceResponseBodyTaskInstanceScript) *GetTaskInstanceResponseBodyTaskInstance {
	s.Script = v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetStartedTime(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.StartedTime = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetStatus(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.Status = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetTags(v []*GetTaskInstanceResponseBodyTaskInstanceTags) *GetTaskInstanceResponseBodyTaskInstance {
	s.Tags = v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetTaskId(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.TaskId = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetTaskName(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.TaskName = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetTaskType(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.TaskType = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetTimeout(v int32) *GetTaskInstanceResponseBodyTaskInstance {
	s.Timeout = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetTriggerRecurrence(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.TriggerRecurrence = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetTriggerTime(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.TriggerTime = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetTriggerType(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.TriggerType = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetUnifiedWorkflowInstanceId(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.UnifiedWorkflowInstanceId = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetWaitingResourceTime(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.WaitingResourceTime = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetWaitingTriggerTime(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.WaitingTriggerTime = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetWorkflowId(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.WorkflowId = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetWorkflowInstanceId(v int64) *GetTaskInstanceResponseBodyTaskInstance {
	s.WorkflowInstanceId = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetWorkflowInstanceType(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.WorkflowInstanceType = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) SetWorkflowName(v string) *GetTaskInstanceResponseBodyTaskInstance {
	s.WorkflowName = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstance) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
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
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskInstanceResponseBodyTaskInstanceDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTaskInstanceResponseBodyTaskInstanceDataSource) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstanceDataSource) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstanceDataSource) GetName() *string {
	return s.Name
}

func (s *GetTaskInstanceResponseBodyTaskInstanceDataSource) SetName(v string) *GetTaskInstanceResponseBodyTaskInstanceDataSource {
	s.Name = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceDataSource) Validate() error {
	return dara.Validate(s)
}

type GetTaskInstanceResponseBodyTaskInstanceInputs struct {
	// The variables.
	Variables []*GetTaskInstanceResponseBodyTaskInstanceInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s GetTaskInstanceResponseBodyTaskInstanceInputs) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstanceInputs) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstanceInputs) GetVariables() []*GetTaskInstanceResponseBodyTaskInstanceInputsVariables {
	return s.Variables
}

func (s *GetTaskInstanceResponseBodyTaskInstanceInputs) SetVariables(v []*GetTaskInstanceResponseBodyTaskInstanceInputsVariables) *GetTaskInstanceResponseBodyTaskInstanceInputs {
	s.Variables = v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceInputs) Validate() error {
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

type GetTaskInstanceResponseBodyTaskInstanceInputsVariables struct {
	// The name of the variable.
	//
	// example:
	//
	// Key1
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

func (s GetTaskInstanceResponseBodyTaskInstanceInputsVariables) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstanceInputsVariables) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstanceInputsVariables) GetName() *string {
	return s.Name
}

func (s *GetTaskInstanceResponseBodyTaskInstanceInputsVariables) GetType() *string {
	return s.Type
}

func (s *GetTaskInstanceResponseBodyTaskInstanceInputsVariables) GetValue() *string {
	return s.Value
}

func (s *GetTaskInstanceResponseBodyTaskInstanceInputsVariables) SetName(v string) *GetTaskInstanceResponseBodyTaskInstanceInputsVariables {
	s.Name = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceInputsVariables) SetType(v string) *GetTaskInstanceResponseBodyTaskInstanceInputsVariables {
	s.Type = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceInputsVariables) SetValue(v string) *GetTaskInstanceResponseBodyTaskInstanceInputsVariables {
	s.Value = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceInputsVariables) Validate() error {
	return dara.Validate(s)
}

type GetTaskInstanceResponseBodyTaskInstanceOutputs struct {
	// The task outputs.
	TaskOutputs []*GetTaskInstanceResponseBodyTaskInstanceOutputsTaskOutputs `json:"TaskOutputs,omitempty" xml:"TaskOutputs,omitempty" type:"Repeated"`
	// The variables.
	Variables []*GetTaskInstanceResponseBodyTaskInstanceOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s GetTaskInstanceResponseBodyTaskInstanceOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstanceOutputs) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputs) GetTaskOutputs() []*GetTaskInstanceResponseBodyTaskInstanceOutputsTaskOutputs {
	return s.TaskOutputs
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputs) GetVariables() []*GetTaskInstanceResponseBodyTaskInstanceOutputsVariables {
	return s.Variables
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputs) SetTaskOutputs(v []*GetTaskInstanceResponseBodyTaskInstanceOutputsTaskOutputs) *GetTaskInstanceResponseBodyTaskInstanceOutputs {
	s.TaskOutputs = v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputs) SetVariables(v []*GetTaskInstanceResponseBodyTaskInstanceOutputsVariables) *GetTaskInstanceResponseBodyTaskInstanceOutputs {
	s.Variables = v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputs) Validate() error {
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

type GetTaskInstanceResponseBodyTaskInstanceOutputsTaskOutputs struct {
	// The output identifier.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s GetTaskInstanceResponseBodyTaskInstanceOutputsTaskOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstanceOutputsTaskOutputs) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputsTaskOutputs) GetOutput() *string {
	return s.Output
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputsTaskOutputs) SetOutput(v string) *GetTaskInstanceResponseBodyTaskInstanceOutputsTaskOutputs {
	s.Output = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputsTaskOutputs) Validate() error {
	return dara.Validate(s)
}

type GetTaskInstanceResponseBodyTaskInstanceOutputsVariables struct {
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

func (s GetTaskInstanceResponseBodyTaskInstanceOutputsVariables) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstanceOutputsVariables) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputsVariables) GetName() *string {
	return s.Name
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputsVariables) GetType() *string {
	return s.Type
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputsVariables) GetValue() *string {
	return s.Value
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputsVariables) SetName(v string) *GetTaskInstanceResponseBodyTaskInstanceOutputsVariables {
	s.Name = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputsVariables) SetType(v string) *GetTaskInstanceResponseBodyTaskInstanceOutputsVariables {
	s.Type = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputsVariables) SetValue(v string) *GetTaskInstanceResponseBodyTaskInstanceOutputsVariables {
	s.Value = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceOutputsVariables) Validate() error {
	return dara.Validate(s)
}

type GetTaskInstanceResponseBodyTaskInstanceRuntime struct {
	// The host for running.
	//
	// example:
	//
	// cn-shanghai.1.2
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The instance run ID.
	//
	// example:
	//
	// T3_123
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
}

func (s GetTaskInstanceResponseBodyTaskInstanceRuntime) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstanceRuntime) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntime) GetGateway() *string {
	return s.Gateway
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntime) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntime) SetGateway(v string) *GetTaskInstanceResponseBodyTaskInstanceRuntime {
	s.Gateway = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntime) SetProcessId(v string) *GetTaskInstanceResponseBodyTaskInstanceRuntime {
	s.ProcessId = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntime) Validate() error {
	return dara.Validate(s)
}

type GetTaskInstanceResponseBodyTaskInstanceRuntimeResource struct {
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

func (s GetTaskInstanceResponseBodyTaskInstanceRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstanceRuntimeResource) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource) SetCu(v string) *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource {
	s.Cu = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource) SetImage(v string) *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource {
	s.Image = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource) SetResourceGroupId(v string) *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type GetTaskInstanceResponseBodyTaskInstanceScript struct {
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

func (s GetTaskInstanceResponseBodyTaskInstanceScript) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstanceScript) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstanceScript) GetContent() *string {
	return s.Content
}

func (s *GetTaskInstanceResponseBodyTaskInstanceScript) GetParameters() *string {
	return s.Parameters
}

func (s *GetTaskInstanceResponseBodyTaskInstanceScript) SetContent(v string) *GetTaskInstanceResponseBodyTaskInstanceScript {
	s.Content = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceScript) SetParameters(v string) *GetTaskInstanceResponseBodyTaskInstanceScript {
	s.Parameters = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceScript) Validate() error {
	return dara.Validate(s)
}

type GetTaskInstanceResponseBodyTaskInstanceTags struct {
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

func (s GetTaskInstanceResponseBodyTaskInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceResponseBodyTaskInstanceTags) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceResponseBodyTaskInstanceTags) GetKey() *string {
	return s.Key
}

func (s *GetTaskInstanceResponseBodyTaskInstanceTags) GetValue() *string {
	return s.Value
}

func (s *GetTaskInstanceResponseBodyTaskInstanceTags) SetKey(v string) *GetTaskInstanceResponseBodyTaskInstanceTags {
	s.Key = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceTags) SetValue(v string) *GetTaskInstanceResponseBodyTaskInstanceTags {
	s.Value = &v
	return s
}

func (s *GetTaskInstanceResponseBodyTaskInstanceTags) Validate() error {
	return dara.Validate(s)
}
