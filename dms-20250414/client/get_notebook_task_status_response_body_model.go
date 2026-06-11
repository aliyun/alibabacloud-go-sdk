// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotebookTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNotebookTaskStatusResponseBody
	GetCode() *string
	SetData(v *GetNotebookTaskStatusResponseBodyData) *GetNotebookTaskStatusResponseBody
	GetData() *GetNotebookTaskStatusResponseBodyData
	SetErrorCode(v string) *GetNotebookTaskStatusResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *GetNotebookTaskStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetNotebookTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNotebookTaskStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNotebookTaskStatusResponseBody
	GetSuccess() *bool
}

type GetNotebookTaskStatusResponseBody struct {
	// The status code. A value of Success indicates that the request was successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The scheduling result.
	Data *GetNotebookTaskStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNotebookTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNotebookTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetNotebookTaskStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNotebookTaskStatusResponseBody) GetData() *GetNotebookTaskStatusResponseBodyData {
	return s.Data
}

func (s *GetNotebookTaskStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetNotebookTaskStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNotebookTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNotebookTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNotebookTaskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNotebookTaskStatusResponseBody) SetCode(v string) *GetNotebookTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetNotebookTaskStatusResponseBody) SetData(v *GetNotebookTaskStatusResponseBodyData) *GetNotebookTaskStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetNotebookTaskStatusResponseBody) SetErrorCode(v string) *GetNotebookTaskStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNotebookTaskStatusResponseBody) SetHttpStatusCode(v int32) *GetNotebookTaskStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNotebookTaskStatusResponseBody) SetMessage(v string) *GetNotebookTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetNotebookTaskStatusResponseBody) SetRequestId(v string) *GetNotebookTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNotebookTaskStatusResponseBody) SetSuccess(v bool) *GetNotebookTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetNotebookTaskStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNotebookTaskStatusResponseBodyData struct {
	// The URL to preview the scheduling result.
	//
	// example:
	//
	// https://dms.aliyun.com/web-ide?***
	NotebookSchedulePreviewUrl *string `json:"NotebookSchedulePreviewUrl,omitempty" xml:"NotebookSchedulePreviewUrl,omitempty"`
	// The progress of the scheduling task.
	//
	// example:
	//
	// 5/6
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The output of the scheduling task.
	//
	// example:
	//
	// test
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The status of the scheduling result.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetNotebookTaskStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNotebookTaskStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNotebookTaskStatusResponseBodyData) GetNotebookSchedulePreviewUrl() *string {
	return s.NotebookSchedulePreviewUrl
}

func (s *GetNotebookTaskStatusResponseBodyData) GetProgress() *string {
	return s.Progress
}

func (s *GetNotebookTaskStatusResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *GetNotebookTaskStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetNotebookTaskStatusResponseBodyData) SetNotebookSchedulePreviewUrl(v string) *GetNotebookTaskStatusResponseBodyData {
	s.NotebookSchedulePreviewUrl = &v
	return s
}

func (s *GetNotebookTaskStatusResponseBodyData) SetProgress(v string) *GetNotebookTaskStatusResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetNotebookTaskStatusResponseBodyData) SetResult(v string) *GetNotebookTaskStatusResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetNotebookTaskStatusResponseBodyData) SetStatus(v string) *GetNotebookTaskStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetNotebookTaskStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
