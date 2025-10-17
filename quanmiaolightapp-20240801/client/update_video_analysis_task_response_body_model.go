// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVideoAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *UpdateVideoAnalysisTaskResponseBodyData) *UpdateVideoAnalysisTaskResponseBody
	GetData() *UpdateVideoAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *UpdateVideoAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateVideoAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVideoAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateVideoAnalysisTaskResponseBody
	GetSuccess() *bool
}

type UpdateVideoAnalysisTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateVideoAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVideoAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVideoAnalysisTaskResponseBody) GetData() *UpdateVideoAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *UpdateVideoAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateVideoAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVideoAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVideoAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateVideoAnalysisTaskResponseBody) SetCode(v string) *UpdateVideoAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVideoAnalysisTaskResponseBody) SetData(v *UpdateVideoAnalysisTaskResponseBodyData) *UpdateVideoAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *UpdateVideoAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *UpdateVideoAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVideoAnalysisTaskResponseBody) SetMessage(v string) *UpdateVideoAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVideoAnalysisTaskResponseBody) SetRequestId(v string) *UpdateVideoAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVideoAnalysisTaskResponseBody) SetSuccess(v bool) *UpdateVideoAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVideoAnalysisTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateVideoAnalysisTaskResponseBodyData struct {
	TaskErrorMessage *string `json:"taskErrorMessage,omitempty" xml:"taskErrorMessage,omitempty"`
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// CANCELED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s UpdateVideoAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisTaskResponseBodyData) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *UpdateVideoAnalysisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateVideoAnalysisTaskResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *UpdateVideoAnalysisTaskResponseBodyData) SetTaskErrorMessage(v string) *UpdateVideoAnalysisTaskResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *UpdateVideoAnalysisTaskResponseBodyData) SetTaskId(v string) *UpdateVideoAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateVideoAnalysisTaskResponseBodyData) SetTaskStatus(v string) *UpdateVideoAnalysisTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *UpdateVideoAnalysisTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
