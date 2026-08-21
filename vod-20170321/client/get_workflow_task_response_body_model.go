// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetWorkflowTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetWorkflowTaskResponseBody
	GetErrorMessage() *string
	SetFinishTimeUtc(v string) *GetWorkflowTaskResponseBody
	GetFinishTimeUtc() *string
	SetGmtCreateUtc(v string) *GetWorkflowTaskResponseBody
	GetGmtCreateUtc() *string
	SetNodeResults(v string) *GetWorkflowTaskResponseBody
	GetNodeResults() *string
	SetOutputs(v string) *GetWorkflowTaskResponseBody
	GetOutputs() *string
	SetRequestId(v string) *GetWorkflowTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetWorkflowTaskResponseBody
	GetStatus() *string
	SetTaskId(v string) *GetWorkflowTaskResponseBody
	GetTaskId() *string
	SetUserData(v string) *GetWorkflowTaskResponseBody
	GetUserData() *string
	SetWorkflowId(v string) *GetWorkflowTaskResponseBody
	GetWorkflowId() *string
}

type GetWorkflowTaskResponseBody struct {
	// The error code returned when transcoding fails.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when transcoding fails.
	//
	// example:
	//
	// ErrorMessage
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the task was completed. The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2025-08-19T02:28:22Z
	FinishTimeUtc *string `json:"FinishTimeUtc,omitempty" xml:"FinishTimeUtc,omitempty"`
	// The time when the task was created. The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2025-07-28T02:17:26Z
	GmtCreateUtc *string `json:"GmtCreateUtc,omitempty" xml:"GmtCreateUtc,omitempty"`
	// The node results of the workflow task. The value is in JSON format and varies based on the workflow configuration.
	//
	// example:
	//
	// {}
	NodeResults *string `json:"NodeResults,omitempty" xml:"NodeResults,omitempty"`
	// The output information.
	//
	// example:
	//
	// {}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 63E8B7C7-4812-46*****AD-0FA56029AC86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of this review. This indicates the current manual review result. Valid values:
	//
	// - **running**: Running.
	//
	// - **stopped**: Stopped.
	//
	// - **failed**: Failed.
	//
	// - **partial-succeeded**: Partially succeeded.
	//
	// - **succeeded**: Succeeded.
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID used to query the refresh status.
	//
	// example:
	//
	// 70422****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The custom information.
	//
	// example:
	//
	// {}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The workflow ID. You can log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Processing*	- > **Workflow Management*	- to view the ID.
	//
	// example:
	//
	// 613efff3887ec34af685714cc461****
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetWorkflowTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkflowTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetWorkflowTaskResponseBody) GetFinishTimeUtc() *string {
	return s.FinishTimeUtc
}

func (s *GetWorkflowTaskResponseBody) GetGmtCreateUtc() *string {
	return s.GmtCreateUtc
}

func (s *GetWorkflowTaskResponseBody) GetNodeResults() *string {
	return s.NodeResults
}

func (s *GetWorkflowTaskResponseBody) GetOutputs() *string {
	return s.Outputs
}

func (s *GetWorkflowTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetWorkflowTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWorkflowTaskResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetWorkflowTaskResponseBody) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *GetWorkflowTaskResponseBody) SetErrorCode(v string) *GetWorkflowTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetErrorMessage(v string) *GetWorkflowTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetFinishTimeUtc(v string) *GetWorkflowTaskResponseBody {
	s.FinishTimeUtc = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetGmtCreateUtc(v string) *GetWorkflowTaskResponseBody {
	s.GmtCreateUtc = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetNodeResults(v string) *GetWorkflowTaskResponseBody {
	s.NodeResults = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetOutputs(v string) *GetWorkflowTaskResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetRequestId(v string) *GetWorkflowTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetStatus(v string) *GetWorkflowTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetTaskId(v string) *GetWorkflowTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetUserData(v string) *GetWorkflowTaskResponseBody {
	s.UserData = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) SetWorkflowId(v string) *GetWorkflowTaskResponseBody {
	s.WorkflowId = &v
	return s
}

func (s *GetWorkflowTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
