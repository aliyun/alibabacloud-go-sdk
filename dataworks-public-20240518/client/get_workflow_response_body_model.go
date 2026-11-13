// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkflowResponseBody
	GetRequestId() *string
	SetWorkflow(v *GetWorkflowResponseBodyWorkflow) *GetWorkflowResponseBody
	GetWorkflow() *GetWorkflowResponseBodyWorkflow
}

type GetWorkflowResponseBody struct {
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the workflow.
	Workflow *GetWorkflowResponseBodyWorkflow `json:"Workflow,omitempty" xml:"Workflow,omitempty" type:"Struct"`
}

func (s GetWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowResponseBody) GetWorkflow() *GetWorkflowResponseBodyWorkflow {
	return s.Workflow
}

func (s *GetWorkflowResponseBody) SetRequestId(v string) *GetWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowResponseBody) SetWorkflow(v *GetWorkflowResponseBodyWorkflow) *GetWorkflowResponseBody {
	s.Workflow = v
	return s
}

func (s *GetWorkflowResponseBody) Validate() error {
	if s.Workflow != nil {
		if err := s.Workflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowResponseBodyWorkflow struct {
	// The client unique code of the workflow, which is used to implement asynchronous and idempotent operations. If you do not specify this parameter when creating the workflow, the system automatically generates one and uniquely binds the code to the resource ID. If you specify this parameter when updating or deleting the resource, it must be the same as the client unique code used during creation.
	//
	// example:
	//
	// Workflow_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The creation time.
	//
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The account ID of the user who created the workflow.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The dependency information.
	Dependencies []*GetWorkflowResponseBodyWorkflowDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// Test workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project environment. Valid values:
	//
	// - Prod: production.
	//
	// - Dev: development.
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The unique identifier of the workflow.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time.
	//
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account ID of the user who modified the workflow.
	//
	// example:
	//
	// 1000
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name.
	//
	// example:
	//
	// Workflow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output information.
	Outputs *GetWorkflowResponseBodyWorkflowOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// The account ID of the owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The list of parameters.
	//
	// example:
	//
	// para1=$bizdate para2=$[yyyymmdd]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the project.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of tags.
	Tags []*GetWorkflowResponseBodyWorkflowTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The list of tasks.
	Tasks []*GetWorkflowResponseBodyWorkflowTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The trigger method.
	Trigger *GetWorkflowResponseBodyWorkflowTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s GetWorkflowResponseBodyWorkflow) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyWorkflow) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyWorkflow) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *GetWorkflowResponseBodyWorkflow) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetWorkflowResponseBodyWorkflow) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetWorkflowResponseBodyWorkflow) GetDependencies() []*GetWorkflowResponseBodyWorkflowDependencies {
	return s.Dependencies
}

func (s *GetWorkflowResponseBodyWorkflow) GetDescription() *string {
	return s.Description
}

func (s *GetWorkflowResponseBodyWorkflow) GetEnvType() *string {
	return s.EnvType
}

func (s *GetWorkflowResponseBodyWorkflow) GetId() *int64 {
	return s.Id
}

func (s *GetWorkflowResponseBodyWorkflow) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetWorkflowResponseBodyWorkflow) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetWorkflowResponseBodyWorkflow) GetName() *string {
	return s.Name
}

func (s *GetWorkflowResponseBodyWorkflow) GetOutputs() *GetWorkflowResponseBodyWorkflowOutputs {
	return s.Outputs
}

func (s *GetWorkflowResponseBodyWorkflow) GetOwner() *string {
	return s.Owner
}

func (s *GetWorkflowResponseBodyWorkflow) GetParameters() *string {
	return s.Parameters
}

func (s *GetWorkflowResponseBodyWorkflow) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetWorkflowResponseBodyWorkflow) GetTags() []*GetWorkflowResponseBodyWorkflowTags {
	return s.Tags
}

func (s *GetWorkflowResponseBodyWorkflow) GetTasks() []*GetWorkflowResponseBodyWorkflowTasks {
	return s.Tasks
}

func (s *GetWorkflowResponseBodyWorkflow) GetTrigger() *GetWorkflowResponseBodyWorkflowTrigger {
	return s.Trigger
}

func (s *GetWorkflowResponseBodyWorkflow) SetClientUniqueCode(v string) *GetWorkflowResponseBodyWorkflow {
	s.ClientUniqueCode = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetCreateTime(v int64) *GetWorkflowResponseBodyWorkflow {
	s.CreateTime = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetCreateUser(v string) *GetWorkflowResponseBodyWorkflow {
	s.CreateUser = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetDependencies(v []*GetWorkflowResponseBodyWorkflowDependencies) *GetWorkflowResponseBodyWorkflow {
	s.Dependencies = v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetDescription(v string) *GetWorkflowResponseBodyWorkflow {
	s.Description = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetEnvType(v string) *GetWorkflowResponseBodyWorkflow {
	s.EnvType = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetId(v int64) *GetWorkflowResponseBodyWorkflow {
	s.Id = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetModifyTime(v int64) *GetWorkflowResponseBodyWorkflow {
	s.ModifyTime = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetModifyUser(v string) *GetWorkflowResponseBodyWorkflow {
	s.ModifyUser = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetName(v string) *GetWorkflowResponseBodyWorkflow {
	s.Name = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetOutputs(v *GetWorkflowResponseBodyWorkflowOutputs) *GetWorkflowResponseBodyWorkflow {
	s.Outputs = v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetOwner(v string) *GetWorkflowResponseBodyWorkflow {
	s.Owner = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetParameters(v string) *GetWorkflowResponseBodyWorkflow {
	s.Parameters = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetProjectId(v int64) *GetWorkflowResponseBodyWorkflow {
	s.ProjectId = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetTags(v []*GetWorkflowResponseBodyWorkflowTags) *GetWorkflowResponseBodyWorkflow {
	s.Tags = v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetTasks(v []*GetWorkflowResponseBodyWorkflowTasks) *GetWorkflowResponseBodyWorkflow {
	s.Tasks = v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) SetTrigger(v *GetWorkflowResponseBodyWorkflowTrigger) *GetWorkflowResponseBodyWorkflow {
	s.Trigger = v
	return s
}

func (s *GetWorkflowResponseBodyWorkflow) Validate() error {
	if s.Dependencies != nil {
		for _, item := range s.Dependencies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Outputs != nil {
		if err := s.Outputs.Validate(); err != nil {
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
	if s.Tasks != nil {
		for _, item := range s.Tasks {
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

type GetWorkflowResponseBodyWorkflowDependencies struct {
	// The dependency type. Valid values:
	//
	// - CrossCycleDependsOnChildren: cross-cycle dependency on first-level child nodes.
	//
	// - CrossCycleDependsOnSelf: cross-cycle dependency on the node itself.
	//
	// - CrossCycleDependsOnOtherNode: cross-cycle dependency on other nodes.
	//
	// - Normal: same-cycle dependency.
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The output identifier of the upstream task. This field is returned when the dependency type is same-cycle dependency and the input content is configured.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	UpstreamOutput *string `json:"UpstreamOutput,omitempty" xml:"UpstreamOutput,omitempty"`
	// The ID of the upstream task. This field is returned when the dependency type is cross-cycle dependency on other nodes or same-cycle dependency without input content configured. It is not returned in other cases.
	//
	// example:
	//
	// 1234
	UpstreamTaskId *int64 `json:"UpstreamTaskId,omitempty" xml:"UpstreamTaskId,omitempty"`
}

func (s GetWorkflowResponseBodyWorkflowDependencies) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyWorkflowDependencies) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyWorkflowDependencies) GetType() *string {
	return s.Type
}

func (s *GetWorkflowResponseBodyWorkflowDependencies) GetUpstreamOutput() *string {
	return s.UpstreamOutput
}

func (s *GetWorkflowResponseBodyWorkflowDependencies) GetUpstreamTaskId() *int64 {
	return s.UpstreamTaskId
}

func (s *GetWorkflowResponseBodyWorkflowDependencies) SetType(v string) *GetWorkflowResponseBodyWorkflowDependencies {
	s.Type = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowDependencies) SetUpstreamOutput(v string) *GetWorkflowResponseBodyWorkflowDependencies {
	s.UpstreamOutput = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowDependencies) SetUpstreamTaskId(v int64) *GetWorkflowResponseBodyWorkflowDependencies {
	s.UpstreamTaskId = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowDependencies) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowResponseBodyWorkflowOutputs struct {
	// The list of task output definitions.
	TaskOutputs []*GetWorkflowResponseBodyWorkflowOutputsTaskOutputs `json:"TaskOutputs,omitempty" xml:"TaskOutputs,omitempty" type:"Repeated"`
}

func (s GetWorkflowResponseBodyWorkflowOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyWorkflowOutputs) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyWorkflowOutputs) GetTaskOutputs() []*GetWorkflowResponseBodyWorkflowOutputsTaskOutputs {
	return s.TaskOutputs
}

func (s *GetWorkflowResponseBodyWorkflowOutputs) SetTaskOutputs(v []*GetWorkflowResponseBodyWorkflowOutputsTaskOutputs) *GetWorkflowResponseBodyWorkflowOutputs {
	s.TaskOutputs = v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowOutputs) Validate() error {
	if s.TaskOutputs != nil {
		for _, item := range s.TaskOutputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkflowResponseBodyWorkflowOutputsTaskOutputs struct {
	// The output identifier.
	//
	// example:
	//
	// pre.odps_sql_demo_0
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s GetWorkflowResponseBodyWorkflowOutputsTaskOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyWorkflowOutputsTaskOutputs) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyWorkflowOutputsTaskOutputs) GetOutput() *string {
	return s.Output
}

func (s *GetWorkflowResponseBodyWorkflowOutputsTaskOutputs) SetOutput(v string) *GetWorkflowResponseBodyWorkflowOutputsTaskOutputs {
	s.Output = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowOutputsTaskOutputs) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowResponseBodyWorkflowTags struct {
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

func (s GetWorkflowResponseBodyWorkflowTags) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyWorkflowTags) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyWorkflowTags) GetKey() *string {
	return s.Key
}

func (s *GetWorkflowResponseBodyWorkflowTags) GetValue() *string {
	return s.Value
}

func (s *GetWorkflowResponseBodyWorkflowTags) SetKey(v string) *GetWorkflowResponseBodyWorkflowTags {
	s.Key = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTags) SetValue(v string) *GetWorkflowResponseBodyWorkflowTags {
	s.Value = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTags) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowResponseBodyWorkflowTasks struct {
	// The ID of the baseline.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The client unique code of the task, which is used to implement asynchronous and idempotent operations. If you do not specify this parameter when creating the task, the system automatically generates one and uniquely binds the code to the resource ID. If you specify this parameter when updating or deleting the resource, it must be the same as the client unique code used during creation.
	//
	// example:
	//
	// Task_0bc5213917368545132902xxxxxxxx
	ClientUniqueCode *string `json:"ClientUniqueCode,omitempty" xml:"ClientUniqueCode,omitempty"`
	// The creation time.
	//
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The account ID of the user who created the task.
	//
	// example:
	//
	// 1000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The associated data source information.
	DataSource *GetWorkflowResponseBodyWorkflowTasksDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project environment. Valid values:
	//
	// - Prod: production.
	//
	// - Dev: development.
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time.
	//
	// The value is a 13-digit number, such as `1710239005403`.
	//
	// example:
	//
	// 1710239005403
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The account ID of the user who modified the task.
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
	// The ID of the project.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The retry time interval, in seconds.
	//
	// example:
	//
	// 60
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// Specifies whether the task can be rerun. Valid values:
	//
	// - AllDenied: The task cannot be rerun regardless of whether it succeeds or fails.
	//
	// - FailureAllowed: The task can be rerun only when it fails.
	//
	// - AllAllowed: The task can be rerun after either success or failure.
	//
	// example:
	//
	// AllAllowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of retries. This parameter takes effect only when the task is configured to allow reruns.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The runtime environment configuration, such as the resource group information.
	RuntimeResource *GetWorkflowResponseBodyWorkflowTasksRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The timeout period for task execution, in seconds.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The running mode when the task is triggered. Valid values:
	//
	// - Pause: paused.
	//
	// - Skip: dry run.
	//
	// - Normal: normal running.
	//
	// example:
	//
	// Normal
	TriggerRecurrence *string `json:"TriggerRecurrence,omitempty" xml:"TriggerRecurrence,omitempty"`
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

func (s GetWorkflowResponseBodyWorkflowTasks) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyWorkflowTasks) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetClientUniqueCode() *string {
	return s.ClientUniqueCode
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetDataSource() *GetWorkflowResponseBodyWorkflowTasksDataSource {
	return s.DataSource
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetDescription() *string {
	return s.Description
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetEnvType() *string {
	return s.EnvType
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetId() *int64 {
	return s.Id
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetName() *string {
	return s.Name
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetOwner() *string {
	return s.Owner
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetPriority() *int32 {
	return s.Priority
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetRerunMode() *string {
	return s.RerunMode
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetRuntimeResource() *GetWorkflowResponseBodyWorkflowTasksRuntimeResource {
	return s.RuntimeResource
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetTimeout() *int32 {
	return s.Timeout
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetType() *string {
	return s.Type
}

func (s *GetWorkflowResponseBodyWorkflowTasks) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetBaselineId(v int64) *GetWorkflowResponseBodyWorkflowTasks {
	s.BaselineId = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetClientUniqueCode(v string) *GetWorkflowResponseBodyWorkflowTasks {
	s.ClientUniqueCode = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetCreateTime(v int64) *GetWorkflowResponseBodyWorkflowTasks {
	s.CreateTime = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetCreateUser(v string) *GetWorkflowResponseBodyWorkflowTasks {
	s.CreateUser = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetDataSource(v *GetWorkflowResponseBodyWorkflowTasksDataSource) *GetWorkflowResponseBodyWorkflowTasks {
	s.DataSource = v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetDescription(v string) *GetWorkflowResponseBodyWorkflowTasks {
	s.Description = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetEnvType(v string) *GetWorkflowResponseBodyWorkflowTasks {
	s.EnvType = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetId(v int64) *GetWorkflowResponseBodyWorkflowTasks {
	s.Id = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetModifyTime(v int64) *GetWorkflowResponseBodyWorkflowTasks {
	s.ModifyTime = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetModifyUser(v string) *GetWorkflowResponseBodyWorkflowTasks {
	s.ModifyUser = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetName(v string) *GetWorkflowResponseBodyWorkflowTasks {
	s.Name = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetOwner(v string) *GetWorkflowResponseBodyWorkflowTasks {
	s.Owner = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetPriority(v int32) *GetWorkflowResponseBodyWorkflowTasks {
	s.Priority = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetProjectId(v int64) *GetWorkflowResponseBodyWorkflowTasks {
	s.ProjectId = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetRerunInterval(v int32) *GetWorkflowResponseBodyWorkflowTasks {
	s.RerunInterval = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetRerunMode(v string) *GetWorkflowResponseBodyWorkflowTasks {
	s.RerunMode = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetRerunTimes(v int32) *GetWorkflowResponseBodyWorkflowTasks {
	s.RerunTimes = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetRuntimeResource(v *GetWorkflowResponseBodyWorkflowTasksRuntimeResource) *GetWorkflowResponseBodyWorkflowTasks {
	s.RuntimeResource = v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetTimeout(v int32) *GetWorkflowResponseBodyWorkflowTasks {
	s.Timeout = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetTriggerRecurrence(v string) *GetWorkflowResponseBodyWorkflowTasks {
	s.TriggerRecurrence = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetType(v string) *GetWorkflowResponseBodyWorkflowTasks {
	s.Type = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) SetWorkflowId(v int64) *GetWorkflowResponseBodyWorkflowTasks {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasks) Validate() error {
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
	return nil
}

type GetWorkflowResponseBodyWorkflowTasksDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// mysql_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetWorkflowResponseBodyWorkflowTasksDataSource) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyWorkflowTasksDataSource) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyWorkflowTasksDataSource) GetName() *string {
	return s.Name
}

func (s *GetWorkflowResponseBodyWorkflowTasksDataSource) SetName(v string) *GetWorkflowResponseBodyWorkflowTasksDataSource {
	s.Name = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasksDataSource) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowResponseBodyWorkflowTasksRuntimeResource struct {
	// The number of compute units (CUs) consumed during task execution.
	//
	// example:
	//
	// 0.25
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The ID of the image configured for task execution.
	//
	// example:
	//
	// i-xxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The identifier of the scheduling resource group configured for task execution.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetWorkflowResponseBodyWorkflowTasksRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyWorkflowTasksRuntimeResource) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyWorkflowTasksRuntimeResource) GetCu() *string {
	return s.Cu
}

func (s *GetWorkflowResponseBodyWorkflowTasksRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *GetWorkflowResponseBodyWorkflowTasksRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetWorkflowResponseBodyWorkflowTasksRuntimeResource) SetCu(v string) *GetWorkflowResponseBodyWorkflowTasksRuntimeResource {
	s.Cu = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasksRuntimeResource) SetImage(v string) *GetWorkflowResponseBodyWorkflowTasksRuntimeResource {
	s.Image = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasksRuntimeResource) SetResourceGroupId(v string) *GetWorkflowResponseBodyWorkflowTasksRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTasksRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowResponseBodyWorkflowTrigger struct {
	// The cron expression. This parameter takes effect only when Type is set to Scheduler.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The expiration time of the periodic trigger. This parameter takes effect only when Type is set to Scheduler.
	//
	// The format is `yyyy-MM-dd HH:mm:ss`, such as `1970-01-01 00:00:00`. The example does not include a time zone identifier.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The running mode when the workflow is triggered. This parameter takes effect only when Type is set to Scheduler. Valid values:
	//
	// - Pause: paused.
	//
	// - Skip: dry run.
	//
	// - Normal: normal running.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The start time for periodic triggers. This parameter takes effect only when Type is set to Scheduler.
	//
	// The format is `yyyy-MM-dd HH:mm:ss`, such as `1970-01-01 00:00:00`. The example does not include a time zone identifier.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The trigger type. Valid values:
	//
	// - Scheduler: triggered by a scheduling period.
	//
	// - Manual: manually triggered.
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWorkflowResponseBodyWorkflowTrigger) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowResponseBodyWorkflowTrigger) GoString() string {
	return s.String()
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) GetCron() *string {
	return s.Cron
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) GetRecurrence() *string {
	return s.Recurrence
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) GetType() *string {
	return s.Type
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) SetCron(v string) *GetWorkflowResponseBodyWorkflowTrigger {
	s.Cron = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) SetEndTime(v string) *GetWorkflowResponseBodyWorkflowTrigger {
	s.EndTime = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) SetRecurrence(v string) *GetWorkflowResponseBodyWorkflowTrigger {
	s.Recurrence = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) SetStartTime(v string) *GetWorkflowResponseBodyWorkflowTrigger {
	s.StartTime = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) SetType(v string) *GetWorkflowResponseBodyWorkflowTrigger {
	s.Type = &v
	return s
}

func (s *GetWorkflowResponseBodyWorkflowTrigger) Validate() error {
	return dara.Validate(s)
}
