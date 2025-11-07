// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTaskExecutionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTaskExecutionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTaskExecutionsResponseBody
	GetRequestId() *string
	SetTaskExecutions(v []*ListTaskExecutionsResponseBodyTaskExecutions) *ListTaskExecutionsResponseBody
	GetTaskExecutions() []*ListTaskExecutionsResponseBodyTaskExecutions
}

type ListTaskExecutionsResponseBody struct {
	// The details of the task executions.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// CDABABABAB-FC28-4D9C-8FB5-68DC6F0486FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution ID of the child node.
	TaskExecutions []*ListTaskExecutionsResponseBodyTaskExecutions `json:"TaskExecutions,omitempty" xml:"TaskExecutions,omitempty" type:"Repeated"`
}

func (s ListTaskExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTaskExecutionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTaskExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskExecutionsResponseBody) GetTaskExecutions() []*ListTaskExecutionsResponseBodyTaskExecutions {
	return s.TaskExecutions
}

func (s *ListTaskExecutionsResponseBody) SetMaxResults(v int32) *ListTaskExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTaskExecutionsResponseBody) SetNextToken(v string) *ListTaskExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTaskExecutionsResponseBody) SetRequestId(v string) *ListTaskExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskExecutionsResponseBody) SetTaskExecutions(v []*ListTaskExecutionsResponseBodyTaskExecutions) *ListTaskExecutionsResponseBody {
	s.TaskExecutions = v
	return s
}

func (s *ListTaskExecutionsResponseBody) Validate() error {
	if s.TaskExecutions != nil {
		for _, item := range s.TaskExecutions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskExecutionsResponseBodyTaskExecutions struct {
	// The output of the execution.
	//
	// example:
	//
	// exec-xxx
	ChildExecutionId *string `json:"ChildExecutionId,omitempty" xml:"ChildExecutionId,omitempty"`
	// The ID of the execution.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The execution ID of the parent node.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The action of the task.
	//
	// example:
	//
	// exec-44d32b45d2a449e49899
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The Input parameters of the task execution.
	//
	// example:
	//
	// {                     "NotifyNote":""                 }
	ExtraData map[string]interface{} `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// {}
	Loop map[string]interface{} `json:"Loop,omitempty" xml:"Loop,omitempty"`
	// The status information of the task execution.
	//
	// example:
	//
	// 2
	LoopBatchNumber *int32 `json:"LoopBatchNumber,omitempty" xml:"LoopBatchNumber,omitempty"`
	// The time when the execution was created.
	//
	// example:
	//
	// i-1234566zxcvvb
	LoopItem *string `json:"LoopItem,omitempty" xml:"LoopItem,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// { "InstanceId":"i-xxx" }
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// task-exec-xxx
	ParentTaskExecutionId *string `json:"ParentTaskExecutionId,omitempty" xml:"ParentTaskExecutionId,omitempty"`
	// Queries task executions. Multiple methods are supported to filter task executions.
	//
	// example:
	//
	// { "Status":"Running" }
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The elements in the loop task.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The time when the task execution stopped running.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The additional information.
	//
	// example:
	//
	// ""
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The execution ID of the task.
	//
	// example:
	//
	// ACS::Sleep
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The time when the execution was last updated.
	//
	// example:
	//
	// task-exec-xxx
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The time when the execution started.
	//
	// example:
	//
	// describeInstance
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The number of times for which the loop task is run.
	//
	// example:
	//
	// xxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The configuration and statistics information of the loop task. This parameter is returned only for the parent node of the loop task.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListTaskExecutionsResponseBodyTaskExecutions) String() string {
	return dara.Prettify(s)
}

func (s ListTaskExecutionsResponseBodyTaskExecutions) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetChildExecutionId() *string {
	return s.ChildExecutionId
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetEndDate() *string {
	return s.EndDate
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetExtraData() map[string]interface{} {
	return s.ExtraData
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetLoop() map[string]interface{} {
	return s.Loop
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetLoopBatchNumber() *int32 {
	return s.LoopBatchNumber
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetLoopItem() *string {
	return s.LoopItem
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetOutputs() *string {
	return s.Outputs
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetParentTaskExecutionId() *string {
	return s.ParentTaskExecutionId
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetProperties() *string {
	return s.Properties
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetStartDate() *string {
	return s.StartDate
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetStatus() *string {
	return s.Status
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetTaskAction() *string {
	return s.TaskAction
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetTaskExecutionId() *string {
	return s.TaskExecutionId
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetTaskName() *string {
	return s.TaskName
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetChildExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ChildExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetCreateDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.CreateDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetEndDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.EndDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetExtraData(v map[string]interface{}) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ExtraData = v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetLoop(v map[string]interface{}) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Loop = v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetLoopBatchNumber(v int32) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.LoopBatchNumber = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetLoopItem(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.LoopItem = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetOutputs(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Outputs = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetParentTaskExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ParentTaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetProperties(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Properties = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetStartDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.StartDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetStatus(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Status = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetStatusMessage(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.StatusMessage = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTaskAction(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TaskAction = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTaskExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTaskName(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TaskName = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTemplateId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TemplateId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetUpdateDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.UpdateDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) Validate() error {
	return dara.Validate(s)
}
