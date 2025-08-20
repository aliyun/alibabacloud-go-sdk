// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAITaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAITaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAITaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAITaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetAITaskResponseBody
	GetStatus() *string
	SetStatusReason(v string) *GetAITaskResponseBody
	GetStatusReason() *string
	SetSuccess(v string) *GetAITaskResponseBody
	GetSuccess() *string
	SetTaskId(v string) *GetAITaskResponseBody
	GetTaskId() *string
	SetTaskOutput(v map[string]interface{}) *GetAITaskResponseBody
	GetTaskOutput() map[string]interface{}
	SetTaskType(v string) *GetAITaskResponseBody
	GetTaskType() *string
}

type GetAITaskResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Forbidden
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// FAILURE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the AI task is in the state.
	//
	// example:
	//
	// Handler execution unexpected failure
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
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
	// The ID of the AI task.
	//
	// example:
	//
	// t-asasas*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The outputs of the AI task. The outputs include the template.
	//
	// *
	//
	// example:
	//
	// "Result": {
	//
	//     "ROSTemplateFormatVersion": "2015-09-01",
	//
	// }
	TaskOutput map[string]interface{} `json:"TaskOutput,omitempty" xml:"TaskOutput,omitempty"`
	// The type of the AI task.
	//
	// 	- GenerateTemplate: The AI task is used to generate a template.
	//
	// 	- FixTemplate: The AI task is used to fix a template.
	//
	// This parameter is required.
	//
	// example:
	//
	// GenerateTemplate
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetAITaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAITaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAITaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAITaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAITaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAITaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAITaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAITaskResponseBody) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetAITaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetAITaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAITaskResponseBody) GetTaskOutput() map[string]interface{} {
	return s.TaskOutput
}

func (s *GetAITaskResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetAITaskResponseBody) SetCode(v string) *GetAITaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetAITaskResponseBody) SetHttpStatusCode(v int32) *GetAITaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAITaskResponseBody) SetMessage(v string) *GetAITaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetAITaskResponseBody) SetRequestId(v string) *GetAITaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAITaskResponseBody) SetStatus(v string) *GetAITaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetAITaskResponseBody) SetStatusReason(v string) *GetAITaskResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetAITaskResponseBody) SetSuccess(v string) *GetAITaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetAITaskResponseBody) SetTaskId(v string) *GetAITaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetAITaskResponseBody) SetTaskOutput(v map[string]interface{}) *GetAITaskResponseBody {
	s.TaskOutput = v
	return s
}

func (s *GetAITaskResponseBody) SetTaskType(v string) *GetAITaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetAITaskResponseBody) Validate() error {
	return dara.Validate(s)
}
