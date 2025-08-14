// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWorkflowTaskResponseBody
	GetRequestId() *string
	SetWorkflowTask(v *GetWorkflowTaskResponseBodyWorkflowTask) *GetWorkflowTaskResponseBody
	GetWorkflowTask() *GetWorkflowTaskResponseBodyWorkflowTask
}

type GetWorkflowTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******0C-7870-15FE-B96F-8880BB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the workflow task.
	WorkflowTask *GetWorkflowTaskResponseBodyWorkflowTask `json:"WorkflowTask,omitempty" xml:"WorkflowTask,omitempty" type:"Struct"`
}

func (s GetWorkflowTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowTaskResponseBody) GetWorkflowTask() *GetWorkflowTaskResponseBodyWorkflowTask {
	return s.WorkflowTask
}

func (s *GetWorkflowTaskResponseBody) SetRequestId(v string) *GetWorkflowTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetWorkflowTask(v *GetWorkflowTaskResponseBodyWorkflowTask) *GetWorkflowTaskResponseBody {
	s.WorkflowTask = v
	return s
}

func (s *GetWorkflowTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowTaskResponseBodyWorkflowTask struct {
	// The results for all nodes of the workflow task.
	ActivityResults *string `json:"ActivityResults,omitempty" xml:"ActivityResults,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-04T02:05:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the task was complete. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-04T02:06:19Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The task state.
	//
	// Valid values:
	//
	// 	- Init: The task is being initialized.
	//
	// 	- Failed: The task failed.
	//
	// 	- Canceled: The task is canceled.
	//
	// 	- Processing: The task is in progress.
	//
	// 	- Succeed: The task is successful.
	//
	// example:
	//
	// Succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the workflow task.
	//
	// example:
	//
	// ******4215e042b3966ca5441e******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The input of the workflow task.
	//
	// example:
	//
	// {
	//
	//       "Type": "Media",
	//
	//       "Media": "******30706071edbfe290b488******"
	//
	// }
	TaskInput *string `json:"TaskInput,omitempty" xml:"TaskInput,omitempty"`
	// The user-defined field that was specified when the workflow task was submitted.
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The workflow Information.
	Workflow *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow `json:"Workflow,omitempty" xml:"Workflow,omitempty" type:"Struct"`
}

func (s GetWorkflowTaskResponseBodyWorkflowTask) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTaskResponseBodyWorkflowTask) GoString() string {
	return s.String()
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) GetActivityResults() *string {
	return s.ActivityResults
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) GetStatus() *string {
	return s.Status
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) GetTaskInput() *string {
	return s.TaskInput
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) GetUserData() *string {
	return s.UserData
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) GetWorkflow() *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	return s.Workflow
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) SetActivityResults(v string) *GetWorkflowTaskResponseBodyWorkflowTask {
	s.ActivityResults = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) SetCreateTime(v string) *GetWorkflowTaskResponseBodyWorkflowTask {
	s.CreateTime = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) SetFinishTime(v string) *GetWorkflowTaskResponseBodyWorkflowTask {
	s.FinishTime = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) SetStatus(v string) *GetWorkflowTaskResponseBodyWorkflowTask {
	s.Status = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) SetTaskId(v string) *GetWorkflowTaskResponseBodyWorkflowTask {
	s.TaskId = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) SetTaskInput(v string) *GetWorkflowTaskResponseBodyWorkflowTask {
	s.TaskInput = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) SetUserData(v string) *GetWorkflowTaskResponseBodyWorkflowTask {
	s.UserData = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) SetWorkflow(v *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) *GetWorkflowTaskResponseBodyWorkflowTask {
	s.Workflow = v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTask) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowTaskResponseBodyWorkflowTaskWorkflow struct {
	// The time when the workflow was created.
	//
	// example:
	//
	// 2022-11-27T10:02:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the workflow was last modified.
	//
	// example:
	//
	// 2022-11-29T02:06:19Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The workflow name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The workflow state.
	//
	// Valid values:
	//
	// 	- Active
	//
	// 	- Inactive
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workflow type.
	//
	// Valid values:
	//
	// 	- Customize: custom workflow.
	//
	// 	- System: system workflow.
	//
	// 	- Common: user-created workflow.
	//
	// example:
	//
	// Common
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// ******63dca94c609de02ac0d1******
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) GoString() string {
	return s.String()
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetName() *string {
	return s.Name
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetStatus() *string {
	return s.Status
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetType() *string {
	return s.Type
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetCreateTime(v string) *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.CreateTime = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetModifiedTime(v string) *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.ModifiedTime = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetName(v string) *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.Name = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetStatus(v string) *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.Status = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetType(v string) *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.Type = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) SetWorkflowId(v string) *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowTaskResponseBodyWorkflowTaskWorkflow) Validate() error {
	return dara.Validate(s)
}
