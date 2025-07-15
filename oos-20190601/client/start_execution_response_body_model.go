// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecution(v *StartExecutionResponseBodyExecution) *StartExecutionResponseBody
	GetExecution() *StartExecutionResponseBodyExecution
	SetRequestId(v string) *StartExecutionResponseBody
	GetRequestId() *string
}

type StartExecutionResponseBody struct {
	// The details of the execution.
	Execution *StartExecutionResponseBodyExecution `json:"Execution,omitempty" xml:"Execution,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBody) GetExecution() *StartExecutionResponseBodyExecution {
	return s.Execution
}

func (s *StartExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartExecutionResponseBody) SetExecution(v *StartExecutionResponseBodyExecution) *StartExecutionResponseBody {
	s.Execution = v
	return s
}

func (s *StartExecutionResponseBody) SetRequestId(v string) *StartExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}

type StartExecutionResponseBodyExecution struct {
	// The number of executions.
	//
	// example:
	//
	// 1
	Counters map[string]interface{} `json:"Counters,omitempty" xml:"Counters,omitempty"`
	// The time when the execution was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about in-progress tasks.
	CurrentTasks []*StartExecutionResponseBodyExecutionCurrentTasks `json:"CurrentTasks,omitempty" xml:"CurrentTasks,omitempty" type:"Repeated"`
	// The description of the execution.
	//
	// example:
	//
	// test execution.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the execution stopped.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The account ID of the user who started the execution of the template.
	//
	// example:
	//
	// root(13092080xx12344)
	ExecutedBy *string `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	// The GUID of the execution.
	//
	// example:
	//
	// exec-xxxyyy
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// Indicates whether the execution is a parent execution.
	//
	// example:
	//
	// false
	IsParent *bool `json:"IsParent,omitempty" xml:"IsParent,omitempty"`
	// The loop mode.
	//
	// example:
	//
	// Automatic
	LoopMode *string `json:"LoopMode,omitempty" xml:"LoopMode,omitempty"`
	// The execution mode.
	//
	// example:
	//
	// Automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
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
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the parent execution.
	//
	// example:
	//
	// exec-xxxx
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
	// The security check mode.
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
	// The status of the execution.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status information of the execution.
	//
	// example:
	//
	// ""
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The tags of the execution.
	//
	// example:
	//
	// {"k1":"v2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// t-1bd341007f
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
	// The time when the execution was last updated.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s StartExecutionResponseBodyExecution) String() string {
	return dara.Prettify(s)
}

func (s StartExecutionResponseBodyExecution) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBodyExecution) GetCounters() map[string]interface{} {
	return s.Counters
}

func (s *StartExecutionResponseBodyExecution) GetCreateDate() *string {
	return s.CreateDate
}

func (s *StartExecutionResponseBodyExecution) GetCurrentTasks() []*StartExecutionResponseBodyExecutionCurrentTasks {
	return s.CurrentTasks
}

func (s *StartExecutionResponseBodyExecution) GetDescription() *string {
	return s.Description
}

func (s *StartExecutionResponseBodyExecution) GetEndDate() *string {
	return s.EndDate
}

func (s *StartExecutionResponseBodyExecution) GetExecutedBy() *string {
	return s.ExecutedBy
}

func (s *StartExecutionResponseBodyExecution) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *StartExecutionResponseBodyExecution) GetIsParent() *bool {
	return s.IsParent
}

func (s *StartExecutionResponseBodyExecution) GetLoopMode() *string {
	return s.LoopMode
}

func (s *StartExecutionResponseBodyExecution) GetMode() *string {
	return s.Mode
}

func (s *StartExecutionResponseBodyExecution) GetOutputs() *string {
	return s.Outputs
}

func (s *StartExecutionResponseBodyExecution) GetParameters() *string {
	return s.Parameters
}

func (s *StartExecutionResponseBodyExecution) GetParentExecutionId() *string {
	return s.ParentExecutionId
}

func (s *StartExecutionResponseBodyExecution) GetRamRole() *string {
	return s.RamRole
}

func (s *StartExecutionResponseBodyExecution) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StartExecutionResponseBodyExecution) GetSafetyCheck() *string {
	return s.SafetyCheck
}

func (s *StartExecutionResponseBodyExecution) GetStartDate() *string {
	return s.StartDate
}

func (s *StartExecutionResponseBodyExecution) GetStatus() *string {
	return s.Status
}

func (s *StartExecutionResponseBodyExecution) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *StartExecutionResponseBodyExecution) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *StartExecutionResponseBodyExecution) GetTemplateId() *string {
	return s.TemplateId
}

func (s *StartExecutionResponseBodyExecution) GetTemplateName() *string {
	return s.TemplateName
}

func (s *StartExecutionResponseBodyExecution) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *StartExecutionResponseBodyExecution) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *StartExecutionResponseBodyExecution) SetCounters(v map[string]interface{}) *StartExecutionResponseBodyExecution {
	s.Counters = v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetCreateDate(v string) *StartExecutionResponseBodyExecution {
	s.CreateDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetCurrentTasks(v []*StartExecutionResponseBodyExecutionCurrentTasks) *StartExecutionResponseBodyExecution {
	s.CurrentTasks = v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetDescription(v string) *StartExecutionResponseBodyExecution {
	s.Description = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetEndDate(v string) *StartExecutionResponseBodyExecution {
	s.EndDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetExecutedBy(v string) *StartExecutionResponseBodyExecution {
	s.ExecutedBy = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetExecutionId(v string) *StartExecutionResponseBodyExecution {
	s.ExecutionId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetIsParent(v bool) *StartExecutionResponseBodyExecution {
	s.IsParent = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetLoopMode(v string) *StartExecutionResponseBodyExecution {
	s.LoopMode = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetMode(v string) *StartExecutionResponseBodyExecution {
	s.Mode = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetOutputs(v string) *StartExecutionResponseBodyExecution {
	s.Outputs = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetParameters(v string) *StartExecutionResponseBodyExecution {
	s.Parameters = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetParentExecutionId(v string) *StartExecutionResponseBodyExecution {
	s.ParentExecutionId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetRamRole(v string) *StartExecutionResponseBodyExecution {
	s.RamRole = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetResourceGroupId(v string) *StartExecutionResponseBodyExecution {
	s.ResourceGroupId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetSafetyCheck(v string) *StartExecutionResponseBodyExecution {
	s.SafetyCheck = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetStartDate(v string) *StartExecutionResponseBodyExecution {
	s.StartDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetStatus(v string) *StartExecutionResponseBodyExecution {
	s.Status = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetStatusMessage(v string) *StartExecutionResponseBodyExecution {
	s.StatusMessage = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTags(v map[string]interface{}) *StartExecutionResponseBodyExecution {
	s.Tags = v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTemplateId(v string) *StartExecutionResponseBodyExecution {
	s.TemplateId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTemplateName(v string) *StartExecutionResponseBodyExecution {
	s.TemplateName = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTemplateVersion(v string) *StartExecutionResponseBodyExecution {
	s.TemplateVersion = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetUpdateDate(v string) *StartExecutionResponseBodyExecution {
	s.UpdateDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) Validate() error {
	return dara.Validate(s)
}

type StartExecutionResponseBodyExecutionCurrentTasks struct {
	// The action of the task.
	//
	// example:
	//
	// ACS::TimerTrigger
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The execution ID of the task.
	//
	// example:
	//
	// exec-dsadasdawq
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// testTask
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s StartExecutionResponseBodyExecutionCurrentTasks) String() string {
	return dara.Prettify(s)
}

func (s StartExecutionResponseBodyExecutionCurrentTasks) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) GetTaskAction() *string {
	return s.TaskAction
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) GetTaskExecutionId() *string {
	return s.TaskExecutionId
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) SetTaskAction(v string) *StartExecutionResponseBodyExecutionCurrentTasks {
	s.TaskAction = &v
	return s
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) SetTaskExecutionId(v string) *StartExecutionResponseBodyExecutionCurrentTasks {
	s.TaskExecutionId = &v
	return s
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) SetTaskName(v string) *StartExecutionResponseBodyExecutionCurrentTasks {
	s.TaskName = &v
	return s
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) Validate() error {
	return dara.Validate(s)
}
