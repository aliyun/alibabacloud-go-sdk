// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncJob(v *ImportWorkflowDefinitionResponseBodyAsyncJob) *ImportWorkflowDefinitionResponseBody
	GetAsyncJob() *ImportWorkflowDefinitionResponseBodyAsyncJob
	SetRequestId(v string) *ImportWorkflowDefinitionResponseBody
	GetRequestId() *string
}

type ImportWorkflowDefinitionResponseBody struct {
	// The status information of the asynchronous task.
	AsyncJob *ImportWorkflowDefinitionResponseBodyAsyncJob `json:"AsyncJob,omitempty" xml:"AsyncJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF020E7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportWorkflowDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *ImportWorkflowDefinitionResponseBody) GetAsyncJob() *ImportWorkflowDefinitionResponseBodyAsyncJob {
	return s.AsyncJob
}

func (s *ImportWorkflowDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportWorkflowDefinitionResponseBody) SetAsyncJob(v *ImportWorkflowDefinitionResponseBodyAsyncJob) *ImportWorkflowDefinitionResponseBody {
	s.AsyncJob = v
	return s
}

func (s *ImportWorkflowDefinitionResponseBody) SetRequestId(v string) *ImportWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBody) Validate() error {
	return dara.Validate(s)
}

type ImportWorkflowDefinitionResponseBodyAsyncJob struct {
	// Indicates whether the asynchronous task is complete.
	//
	// example:
	//
	// false
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// The time when the asynchronous task was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1706581425000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message returned if the asynchronous task fails.
	//
	// example:
	//
	// target folder already exists: XXXX
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 1234567691239009XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The progress of the asynchronous task. Valid values: 0 to 100.
	//
	// example:
	//
	// 0
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The response.
	//
	// >  The workflow ID is returned.
	//
	// example:
	//
	// 632647691239009XXXX
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// The status of the asynchronous task.
	//
	// Valid values:
	//
	// 	- Running: The asynchronous task is running.
	//
	// 	- Success: The asynchronous task is complete.
	//
	// 	- Fail: The asynchronous task fails.
	//
	// 	- Cancel: The asynchronous task is canceled.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the asynchronous task.
	//
	// Valid values:
	//
	// 	- Create: The asynchronous task is used to create an object.
	//
	// 	- Cancel: The asynchronous task is used to cancel an operation.
	//
	// example:
	//
	// Create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ImportWorkflowDefinitionResponseBodyAsyncJob) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowDefinitionResponseBodyAsyncJob) GoString() string {
	return s.String()
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) GetCompleted() *bool {
	return s.Completed
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) GetError() *string {
	return s.Error
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) GetId() *string {
	return s.Id
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) GetProgress() *int32 {
	return s.Progress
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) GetResponse() *string {
	return s.Response
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) GetStatus() *string {
	return s.Status
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) GetType() *string {
	return s.Type
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetCompleted(v bool) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Completed = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetCreateTime(v int64) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.CreateTime = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetError(v string) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Error = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetId(v string) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Id = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetProgress(v int32) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Progress = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetResponse(v string) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Response = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetStatus(v string) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Status = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetType(v string) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Type = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) Validate() error {
	return dara.Validate(s)
}
