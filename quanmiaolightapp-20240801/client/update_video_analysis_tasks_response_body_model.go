// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoAnalysisTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVideoAnalysisTasksResponseBody
	GetCode() *string
	SetData(v []*UpdateVideoAnalysisTasksResponseBodyData) *UpdateVideoAnalysisTasksResponseBody
	GetData() []*UpdateVideoAnalysisTasksResponseBodyData
	SetHttpStatusCode(v int32) *UpdateVideoAnalysisTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateVideoAnalysisTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVideoAnalysisTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateVideoAnalysisTasksResponseBody
	GetSuccess() *bool
}

type UpdateVideoAnalysisTasksResponseBody struct {
	// example:
	//
	// successful
	Code *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Data []*UpdateVideoAnalysisTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 04DA1A52-4E51-56CB-BA64-FDDA0B53BAE8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVideoAnalysisTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisTasksResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVideoAnalysisTasksResponseBody) GetData() []*UpdateVideoAnalysisTasksResponseBodyData {
	return s.Data
}

func (s *UpdateVideoAnalysisTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateVideoAnalysisTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVideoAnalysisTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVideoAnalysisTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateVideoAnalysisTasksResponseBody) SetCode(v string) *UpdateVideoAnalysisTasksResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVideoAnalysisTasksResponseBody) SetData(v []*UpdateVideoAnalysisTasksResponseBodyData) *UpdateVideoAnalysisTasksResponseBody {
	s.Data = v
	return s
}

func (s *UpdateVideoAnalysisTasksResponseBody) SetHttpStatusCode(v int32) *UpdateVideoAnalysisTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVideoAnalysisTasksResponseBody) SetMessage(v string) *UpdateVideoAnalysisTasksResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVideoAnalysisTasksResponseBody) SetRequestId(v string) *UpdateVideoAnalysisTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVideoAnalysisTasksResponseBody) SetSuccess(v bool) *UpdateVideoAnalysisTasksResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVideoAnalysisTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateVideoAnalysisTasksResponseBodyData struct {
	// example:
	//
	// xx
	TaskErrorCode *string `json:"taskErrorCode,omitempty" xml:"taskErrorCode,omitempty"`
	// example:
	//
	// xx
	TaskErrorMessage *string `json:"taskErrorMessage,omitempty" xml:"taskErrorMessage,omitempty"`
	// example:
	//
	// xx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// CANCELED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s UpdateVideoAnalysisTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisTasksResponseBodyData) GetTaskErrorCode() *string {
	return s.TaskErrorCode
}

func (s *UpdateVideoAnalysisTasksResponseBodyData) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *UpdateVideoAnalysisTasksResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateVideoAnalysisTasksResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *UpdateVideoAnalysisTasksResponseBodyData) SetTaskErrorCode(v string) *UpdateVideoAnalysisTasksResponseBodyData {
	s.TaskErrorCode = &v
	return s
}

func (s *UpdateVideoAnalysisTasksResponseBodyData) SetTaskErrorMessage(v string) *UpdateVideoAnalysisTasksResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *UpdateVideoAnalysisTasksResponseBodyData) SetTaskId(v string) *UpdateVideoAnalysisTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateVideoAnalysisTasksResponseBodyData) SetTaskStatus(v string) *UpdateVideoAnalysisTasksResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *UpdateVideoAnalysisTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
