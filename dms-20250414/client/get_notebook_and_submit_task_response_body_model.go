// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotebookAndSubmitTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetNotebookAndSubmitTaskResponseBody
	GetCode() *string
	SetErrMsg(v string) *GetNotebookAndSubmitTaskResponseBody
	GetErrMsg() *string
	SetHttpStatusCode(v int32) *GetNotebookAndSubmitTaskResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetNotebookAndSubmitTaskResponseBody
	GetRequestId() *string
	SetSessionId(v string) *GetNotebookAndSubmitTaskResponseBody
	GetSessionId() *string
	SetSuccess(v bool) *GetNotebookAndSubmitTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetNotebookAndSubmitTaskResponseBody
	GetTaskId() *string
}

type GetNotebookAndSubmitTaskResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Request Invalid
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5zrs5szpiezlb9m3qxi6zp32h
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// IcICC2nbMpYp9KygS43n010100
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetNotebookAndSubmitTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNotebookAndSubmitTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetNotebookAndSubmitTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetNotebookAndSubmitTaskResponseBody) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *GetNotebookAndSubmitTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetNotebookAndSubmitTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNotebookAndSubmitTaskResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *GetNotebookAndSubmitTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNotebookAndSubmitTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetNotebookAndSubmitTaskResponseBody) SetCode(v string) *GetNotebookAndSubmitTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetNotebookAndSubmitTaskResponseBody) SetErrMsg(v string) *GetNotebookAndSubmitTaskResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetNotebookAndSubmitTaskResponseBody) SetHttpStatusCode(v int32) *GetNotebookAndSubmitTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNotebookAndSubmitTaskResponseBody) SetRequestId(v string) *GetNotebookAndSubmitTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNotebookAndSubmitTaskResponseBody) SetSessionId(v string) *GetNotebookAndSubmitTaskResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetNotebookAndSubmitTaskResponseBody) SetSuccess(v bool) *GetNotebookAndSubmitTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetNotebookAndSubmitTaskResponseBody) SetTaskId(v string) *GetNotebookAndSubmitTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetNotebookAndSubmitTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
