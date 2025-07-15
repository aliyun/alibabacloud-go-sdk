// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutions(v []*ListExecutionsResponseBodyExecutions) *ListExecutionsResponseBody
	GetExecutions() []*ListExecutionsResponseBodyExecutions
	SetMaxResults(v int32) *ListExecutionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListExecutionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExecutionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListExecutionsResponseBody
	GetTotalCount() *int32
}

type ListExecutionsResponseBody struct {
	// The details of the task executions.
	Executions []*ListExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 14A074-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the executions.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBody) GetExecutions() []*ListExecutionsResponseBodyExecutions {
	return s.Executions
}

func (s *ListExecutionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExecutionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExecutionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListExecutionsResponseBody) SetExecutions(v []*ListExecutionsResponseBodyExecutions) *ListExecutionsResponseBody {
	s.Executions = v
	return s
}

func (s *ListExecutionsResponseBody) SetMaxResults(v int32) *ListExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionsResponseBody) SetNextToken(v string) *ListExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsResponseBody) SetRequestId(v string) *ListExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutionsResponseBody) SetTotalCount(v int32) *ListExecutionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExecutionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExecutionsResponseBodyExecutions struct {
	// The type of the execution template. Valid values: Other, TimerTrigger, EventTrigger, and AlarmTrigger.
	//
	// example:
	//
	// Other
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The number of tasks that are counted by execution status.
	//
	// example:
	//
	// {"Failed": 0,"Success": 1,"Total": 2}
	Counters map[string]interface{} `json:"Counters,omitempty" xml:"Counters,omitempty"`
	// The time when the execution was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the tasks that are running.
	CurrentTasks []*ListExecutionsResponseBodyExecutionsCurrentTasks `json:"CurrentTasks,omitempty" xml:"CurrentTasks,omitempty" type:"Repeated"`
	// The description of the execution.
	//
	// example:
	//
	// test execution.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the execution stops running.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The account ID of the user who started the execution of the template.
	//
	// example:
	//
	// 1309252800
	ExecutedBy *string `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	// The unique ID of the execution.
	//
	// example:
	//
	// exec-44d32b45d2a449e
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// Indicates whether the execution contains child executions.
	//
	// example:
	//
	// false
	IsParent *bool `json:"IsParent,omitempty" xml:"IsParent,omitempty"`
	// The time when the template was last successfully triggered.
	//
	// example:
	//
	// 2019-05-27T09:29:18Z
	LastSuccessfulTriggerTime *string `json:"LastSuccessfulTriggerTime,omitempty" xml:"LastSuccessfulTriggerTime,omitempty"`
	// The outputs of last trigger.
	//
	// example:
	//
	// {
	//
	//       "InstanceId": "i-xxx"
	//
	// }
	LastTriggerOutputs *string `json:"LastTriggerOutputs,omitempty" xml:"LastTriggerOutputs,omitempty"`
	// The status of the execution after the template was last triggered.
	//
	// example:
	//
	// Success
	LastTriggerStatus *string `json:"LastTriggerStatus,omitempty" xml:"LastTriggerStatus,omitempty"`
	// The status message of last trigger.
	//
	// example:
	//
	// ""
	LastTriggerStatusMessage *string `json:"LastTriggerStatusMessage,omitempty" xml:"LastTriggerStatusMessage,omitempty"`
	// The time when the template was last successfully triggered.
	//
	// example:
	//
	// 2019-05-27T09:29:18Z
	LastTriggerTime *string `json:"LastTriggerTime,omitempty" xml:"LastTriggerTime,omitempty"`
	// The execution mode.
	//
	// example:
	//
	// Automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The next schedule time for timer trigger execution.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	NextScheduleTime *string `json:"NextScheduleTime,omitempty" xml:"NextScheduleTime,omitempty"`
	// The output of the execution.
	//
	// example:
	//
	// { "InstanceId":"i-xxx" }
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The input parameters of the execution.
	//
	// example:
	//
	// { "Status":"Running" }
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the parent execution.
	//
	// example:
	//
	// exec-xxx
	ParentExecutionId *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	// The role that started the execution of the template.
	//
	// example:
	//
	// OOSServiceRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the resource.
	//
	// example:
	//
	// { 			"Success": 1 		}
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The security check mode. Valid values: Skip, and ConfirmEveryHighRiskAction.
	//
	// example:
	//
	// Skip
	SafetyCheck *string `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	// The time when the execution was started.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The status of the execution. Valid values: Started, Queued, Running, Waiting, Success, Failed, and Cancelled.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of the task execution.
	//
	// example:
	//
	// “”
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The reason for which the status occurs.
	//
	// example:
	//
	// ""
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The tags of the execution.
	//
	// example:
	//
	// {}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The target resource.
	//
	// example:
	//
	// "{"ResourceType": "ALIYUN::ECS::Instance", "Filters": [{"ResourceIds": ["i-bp14z07dg3464980x72o"], "RegionId": "cn-hangzhou", "Type": "ResourceIds"}]}"
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 123
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number of the template.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the execution was updated.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The Waiting state.
	//
	// example:
	//
	// ""
	WaitingStatus *string `json:"WaitingStatus,omitempty" xml:"WaitingStatus,omitempty"`
}

func (s ListExecutionsResponseBodyExecutions) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutions) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutions) GetCategory() *string {
	return s.Category
}

func (s *ListExecutionsResponseBodyExecutions) GetCounters() map[string]interface{} {
	return s.Counters
}

func (s *ListExecutionsResponseBodyExecutions) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListExecutionsResponseBodyExecutions) GetCurrentTasks() []*ListExecutionsResponseBodyExecutionsCurrentTasks {
	return s.CurrentTasks
}

func (s *ListExecutionsResponseBodyExecutions) GetDescription() *string {
	return s.Description
}

func (s *ListExecutionsResponseBodyExecutions) GetEndDate() *string {
	return s.EndDate
}

func (s *ListExecutionsResponseBodyExecutions) GetExecutedBy() *string {
	return s.ExecutedBy
}

func (s *ListExecutionsResponseBodyExecutions) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *ListExecutionsResponseBodyExecutions) GetIsParent() *bool {
	return s.IsParent
}

func (s *ListExecutionsResponseBodyExecutions) GetLastSuccessfulTriggerTime() *string {
	return s.LastSuccessfulTriggerTime
}

func (s *ListExecutionsResponseBodyExecutions) GetLastTriggerOutputs() *string {
	return s.LastTriggerOutputs
}

func (s *ListExecutionsResponseBodyExecutions) GetLastTriggerStatus() *string {
	return s.LastTriggerStatus
}

func (s *ListExecutionsResponseBodyExecutions) GetLastTriggerStatusMessage() *string {
	return s.LastTriggerStatusMessage
}

func (s *ListExecutionsResponseBodyExecutions) GetLastTriggerTime() *string {
	return s.LastTriggerTime
}

func (s *ListExecutionsResponseBodyExecutions) GetMode() *string {
	return s.Mode
}

func (s *ListExecutionsResponseBodyExecutions) GetNextScheduleTime() *string {
	return s.NextScheduleTime
}

func (s *ListExecutionsResponseBodyExecutions) GetOutputs() *string {
	return s.Outputs
}

func (s *ListExecutionsResponseBodyExecutions) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *ListExecutionsResponseBodyExecutions) GetParentExecutionId() *string {
	return s.ParentExecutionId
}

func (s *ListExecutionsResponseBodyExecutions) GetRamRole() *string {
	return s.RamRole
}

func (s *ListExecutionsResponseBodyExecutions) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListExecutionsResponseBodyExecutions) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *ListExecutionsResponseBodyExecutions) GetSafetyCheck() *string {
	return s.SafetyCheck
}

func (s *ListExecutionsResponseBodyExecutions) GetStartDate() *string {
	return s.StartDate
}

func (s *ListExecutionsResponseBodyExecutions) GetStatus() *string {
	return s.Status
}

func (s *ListExecutionsResponseBodyExecutions) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListExecutionsResponseBodyExecutions) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListExecutionsResponseBodyExecutions) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListExecutionsResponseBodyExecutions) GetTargets() *string {
	return s.Targets
}

func (s *ListExecutionsResponseBodyExecutions) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListExecutionsResponseBodyExecutions) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListExecutionsResponseBodyExecutions) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ListExecutionsResponseBodyExecutions) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListExecutionsResponseBodyExecutions) GetWaitingStatus() *string {
	return s.WaitingStatus
}

func (s *ListExecutionsResponseBodyExecutions) SetCategory(v string) *ListExecutionsResponseBodyExecutions {
	s.Category = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetCounters(v map[string]interface{}) *ListExecutionsResponseBodyExecutions {
	s.Counters = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetCreateDate(v string) *ListExecutionsResponseBodyExecutions {
	s.CreateDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetCurrentTasks(v []*ListExecutionsResponseBodyExecutionsCurrentTasks) *ListExecutionsResponseBodyExecutions {
	s.CurrentTasks = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetDescription(v string) *ListExecutionsResponseBodyExecutions {
	s.Description = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetEndDate(v string) *ListExecutionsResponseBodyExecutions {
	s.EndDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetExecutedBy(v string) *ListExecutionsResponseBodyExecutions {
	s.ExecutedBy = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetExecutionId(v string) *ListExecutionsResponseBodyExecutions {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetIsParent(v bool) *ListExecutionsResponseBodyExecutions {
	s.IsParent = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastSuccessfulTriggerTime(v string) *ListExecutionsResponseBodyExecutions {
	s.LastSuccessfulTriggerTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastTriggerOutputs(v string) *ListExecutionsResponseBodyExecutions {
	s.LastTriggerOutputs = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastTriggerStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.LastTriggerStatus = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastTriggerStatusMessage(v string) *ListExecutionsResponseBodyExecutions {
	s.LastTriggerStatusMessage = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastTriggerTime(v string) *ListExecutionsResponseBodyExecutions {
	s.LastTriggerTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetMode(v string) *ListExecutionsResponseBodyExecutions {
	s.Mode = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetNextScheduleTime(v string) *ListExecutionsResponseBodyExecutions {
	s.NextScheduleTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetOutputs(v string) *ListExecutionsResponseBodyExecutions {
	s.Outputs = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetParameters(v map[string]interface{}) *ListExecutionsResponseBodyExecutions {
	s.Parameters = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetParentExecutionId(v string) *ListExecutionsResponseBodyExecutions {
	s.ParentExecutionId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetRamRole(v string) *ListExecutionsResponseBodyExecutions {
	s.RamRole = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetResourceGroupId(v string) *ListExecutionsResponseBodyExecutions {
	s.ResourceGroupId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetResourceStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.ResourceStatus = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetSafetyCheck(v string) *ListExecutionsResponseBodyExecutions {
	s.SafetyCheck = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStartDate(v string) *ListExecutionsResponseBodyExecutions {
	s.StartDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.Status = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStatusMessage(v string) *ListExecutionsResponseBodyExecutions {
	s.StatusMessage = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStatusReason(v string) *ListExecutionsResponseBodyExecutions {
	s.StatusReason = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTags(v map[string]interface{}) *ListExecutionsResponseBodyExecutions {
	s.Tags = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTargets(v string) *ListExecutionsResponseBodyExecutions {
	s.Targets = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTemplateId(v string) *ListExecutionsResponseBodyExecutions {
	s.TemplateId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTemplateName(v string) *ListExecutionsResponseBodyExecutions {
	s.TemplateName = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTemplateVersion(v string) *ListExecutionsResponseBodyExecutions {
	s.TemplateVersion = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetUpdateDate(v string) *ListExecutionsResponseBodyExecutions {
	s.UpdateDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetWaitingStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.WaitingStatus = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) Validate() error {
	return dara.Validate(s)
}

type ListExecutionsResponseBodyExecutionsCurrentTasks struct {
	// The execution template of the task.
	//
	// example:
	//
	// acs::Template
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The ID of the task execution.
	//
	// example:
	//
	// task-exec-44d32b45d2a49899#1
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// installSLSILogtail
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListExecutionsResponseBodyExecutionsCurrentTasks) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutionsCurrentTasks) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) GetTaskAction() *string {
	return s.TaskAction
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) GetTaskExecutionId() *string {
	return s.TaskExecutionId
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) SetTaskAction(v string) *ListExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskAction = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) SetTaskExecutionId(v string) *ListExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskExecutionId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) SetTaskName(v string) *ListExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskName = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) Validate() error {
	return dara.Validate(s)
}
