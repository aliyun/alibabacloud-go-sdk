// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWorkflowTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListWorkflowTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkflowTasksResponseBody
	GetRequestId() *string
	SetTaskList(v []*ListWorkflowTasksResponseBodyTaskList) *ListWorkflowTasksResponseBody
	GetTaskList() []*ListWorkflowTasksResponseBodyTaskList
	SetTotalCount(v int32) *ListWorkflowTasksResponseBody
	GetTotalCount() *int32
}

type ListWorkflowTasksResponseBody struct {
	// The maximum number of entries returned in this response.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****8EqYpQbZ6Eh7+Zz8DxVYoQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C0C02296-113C-5838-8FE9-8F3A32998DDC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The media workflow tasks.
	TaskList []*ListWorkflowTasksResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	// The total number of entries returned. By default, this parameter is not returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWorkflowTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkflowTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowTasksResponseBody) GetTaskList() []*ListWorkflowTasksResponseBodyTaskList {
	return s.TaskList
}

func (s *ListWorkflowTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkflowTasksResponseBody) SetMaxResults(v int32) *ListWorkflowTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowTasksResponseBody) SetNextToken(v string) *ListWorkflowTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowTasksResponseBody) SetRequestId(v string) *ListWorkflowTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowTasksResponseBody) SetTaskList(v []*ListWorkflowTasksResponseBodyTaskList) *ListWorkflowTasksResponseBody {
	s.TaskList = v
	return s
}

func (s *ListWorkflowTasksResponseBody) SetTotalCount(v int32) *ListWorkflowTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkflowTasksResponseBody) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkflowTasksResponseBodyTaskList struct {
	// The time the task was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-07-15T09:45:48Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time the task was completed. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-12-07T10:53:45Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The task state.
	//
	// **Valid values**:
	//
	// 	- Init: Initializing
	//
	// 	- Failed
	//
	// 	- Canceled
	//
	// 	- Processing
	//
	// 	- Succeed
	//
	// example:
	//
	// Succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the workflow task.
	//
	// example:
	//
	// *****4c93d2f404f8345b16a965*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The input data for the workflow task.
	//
	// example:
	//
	// {\\"Type\\":\\"Media\\",\\"Media\\":\\"****8b40884171efb0d9e7f7f458****\\"}
	TaskInput *string `json:"TaskInput,omitempty" xml:"TaskInput,omitempty"`
	// The custom data that was passed when the task was submitted.
	//
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The information about the workflow template.
	Workflow *ListWorkflowTasksResponseBodyTaskListWorkflow `json:"Workflow,omitempty" xml:"Workflow,omitempty" type:"Struct"`
}

func (s ListWorkflowTasksResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTasksResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *ListWorkflowTasksResponseBodyTaskList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkflowTasksResponseBodyTaskList) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListWorkflowTasksResponseBodyTaskList) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowTasksResponseBodyTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListWorkflowTasksResponseBodyTaskList) GetTaskInput() *string {
	return s.TaskInput
}

func (s *ListWorkflowTasksResponseBodyTaskList) GetUserData() *string {
	return s.UserData
}

func (s *ListWorkflowTasksResponseBodyTaskList) GetWorkflow() *ListWorkflowTasksResponseBodyTaskListWorkflow {
	return s.Workflow
}

func (s *ListWorkflowTasksResponseBodyTaskList) SetCreateTime(v string) *ListWorkflowTasksResponseBodyTaskList {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskList) SetFinishTime(v string) *ListWorkflowTasksResponseBodyTaskList {
	s.FinishTime = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskList) SetStatus(v string) *ListWorkflowTasksResponseBodyTaskList {
	s.Status = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskList) SetTaskId(v string) *ListWorkflowTasksResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskList) SetTaskInput(v string) *ListWorkflowTasksResponseBodyTaskList {
	s.TaskInput = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskList) SetUserData(v string) *ListWorkflowTasksResponseBodyTaskList {
	s.UserData = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskList) SetWorkflow(v *ListWorkflowTasksResponseBodyTaskListWorkflow) *ListWorkflowTasksResponseBodyTaskList {
	s.Workflow = v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskList) Validate() error {
	if s.Workflow != nil {
		if err := s.Workflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowTasksResponseBodyTaskListWorkflow struct {
	// The creation time of the workflow template.
	//
	// example:
	//
	// 2025-03-21T01:48:49Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The source of the media file. Valid values:
	//
	// 	- OSS: an OSS object.
	//
	// 	- Media: a media asset.
	//
	// example:
	//
	// OSS
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The last modification time of the workflow template.
	//
	// example:
	//
	// 2025-02-23 10:19:37 +0800
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the workflow template.
	//
	// example:
	//
	// example-workflow-***
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the workflow template.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the workflow template.
	//
	// example:
	//
	// Common
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// ******2491c84dce913da9fe65******
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkflowTasksResponseBodyTaskListWorkflow) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTasksResponseBodyTaskListWorkflow) GoString() string {
	return s.String()
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) GetMediaType() *string {
	return s.MediaType
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) GetName() *string {
	return s.Name
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) GetType() *string {
	return s.Type
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) SetCreateTime(v string) *ListWorkflowTasksResponseBodyTaskListWorkflow {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) SetMediaType(v string) *ListWorkflowTasksResponseBodyTaskListWorkflow {
	s.MediaType = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) SetModifiedTime(v string) *ListWorkflowTasksResponseBodyTaskListWorkflow {
	s.ModifiedTime = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) SetName(v string) *ListWorkflowTasksResponseBodyTaskListWorkflow {
	s.Name = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) SetStatus(v string) *ListWorkflowTasksResponseBodyTaskListWorkflow {
	s.Status = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) SetType(v string) *ListWorkflowTasksResponseBodyTaskListWorkflow {
	s.Type = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) SetWorkflowId(v string) *ListWorkflowTasksResponseBodyTaskListWorkflow {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowTasksResponseBodyTaskListWorkflow) Validate() error {
	return dara.Validate(s)
}
