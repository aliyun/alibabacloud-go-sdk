// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAITasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpStatusCode(v int32) *ListAITasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAITasksResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAITasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAITasksResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListAITasksResponseBody
	GetSuccess() *string
	SetTasks(v []*ListAITasksResponseBodyTasks) *ListAITasksResponseBody
	GetTasks() []*ListAITasksResponseBodyTasks
}

type ListAITasksResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to complete this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// U12WEI6Ro2ol3wA54rBNS3Cltv2VJyA+7hP4GqbIOhmWU5mWU9ZE3cXLgDaH4KSMRfIYcIVrvtHaAzCoyfo7VQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The AI tasks.
	Tasks []*ListAITasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListAITasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAITasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListAITasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAITasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAITasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAITasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAITasksResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListAITasksResponseBody) GetTasks() []*ListAITasksResponseBodyTasks {
	return s.Tasks
}

func (s *ListAITasksResponseBody) SetHttpStatusCode(v int32) *ListAITasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAITasksResponseBody) SetMessage(v string) *ListAITasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListAITasksResponseBody) SetNextToken(v string) *ListAITasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAITasksResponseBody) SetRequestId(v string) *ListAITasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAITasksResponseBody) SetSuccess(v string) *ListAITasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListAITasksResponseBody) SetTasks(v []*ListAITasksResponseBodyTasks) *ListAITasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListAITasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAITasksResponseBodyTasks struct {
	// The time when the AI task was created. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2023-03-15T03:15:53
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the AI task.
	//
	// example:
	//
	// Create an ECS instance and deploy the Nginx service.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The state of the AI task.
	//
	// 	- PENDING
	//
	// 	- WAITING
	//
	// 	- RUNNING
	//
	// 	- SUCCESS
	//
	// 	- FAILURE
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the AI task is in the state.
	//
	// example:
	//
	// Handler execution unexpected failure
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The ID of the AI task.
	//
	// example:
	//
	// t-asasas*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the AI task.
	//
	// 	- GenerateTemplate: The AI task is used to generate a template.
	//
	// 	- FixTemplate: The AI task is used to fix a template.
	//
	// example:
	//
	// GenerateTemplate
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The time when the AI task was updated. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	//
	// example:
	//
	// 2023-11-20T22:00:50
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAITasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListAITasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListAITasksResponseBodyTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAITasksResponseBodyTasks) GetPrompt() *string {
	return s.Prompt
}

func (s *ListAITasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *ListAITasksResponseBodyTasks) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListAITasksResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAITasksResponseBodyTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *ListAITasksResponseBodyTasks) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAITasksResponseBodyTasks) SetCreateTime(v string) *ListAITasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListAITasksResponseBodyTasks) SetPrompt(v string) *ListAITasksResponseBodyTasks {
	s.Prompt = &v
	return s
}

func (s *ListAITasksResponseBodyTasks) SetStatus(v string) *ListAITasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListAITasksResponseBodyTasks) SetStatusReason(v string) *ListAITasksResponseBodyTasks {
	s.StatusReason = &v
	return s
}

func (s *ListAITasksResponseBodyTasks) SetTaskId(v string) *ListAITasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListAITasksResponseBodyTasks) SetTaskType(v string) *ListAITasksResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *ListAITasksResponseBodyTasks) SetUpdateTime(v string) *ListAITasksResponseBodyTasks {
	s.UpdateTime = &v
	return s
}

func (s *ListAITasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
