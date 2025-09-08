// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeepWriteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeepWriteTaskResponseBody
	GetCode() *string
	SetData(v *GetDeepWriteTaskResponseBodyData) *GetDeepWriteTaskResponseBody
	GetData() *GetDeepWriteTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetDeepWriteTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDeepWriteTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeepWriteTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDeepWriteTaskResponseBody
	GetSuccess() *bool
}

type GetDeepWriteTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetDeepWriteTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDeepWriteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeepWriteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeepWriteTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeepWriteTaskResponseBody) GetData() *GetDeepWriteTaskResponseBodyData {
	return s.Data
}

func (s *GetDeepWriteTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDeepWriteTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeepWriteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeepWriteTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDeepWriteTaskResponseBody) SetCode(v string) *GetDeepWriteTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeepWriteTaskResponseBody) SetData(v *GetDeepWriteTaskResponseBodyData) *GetDeepWriteTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetDeepWriteTaskResponseBody) SetHttpStatusCode(v int32) *GetDeepWriteTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDeepWriteTaskResponseBody) SetMessage(v string) *GetDeepWriteTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeepWriteTaskResponseBody) SetRequestId(v string) *GetDeepWriteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeepWriteTaskResponseBody) SetSuccess(v bool) *GetDeepWriteTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetDeepWriteTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDeepWriteTaskResponseBodyData struct {
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// queued
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// a2267372-6042-46a4-aab0-1670dfc38c94
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDeepWriteTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeepWriteTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeepWriteTaskResponseBodyData) GetInput() *string {
	return s.Input
}

func (s *GetDeepWriteTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDeepWriteTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDeepWriteTaskResponseBodyData) SetInput(v string) *GetDeepWriteTaskResponseBodyData {
	s.Input = &v
	return s
}

func (s *GetDeepWriteTaskResponseBodyData) SetStatus(v string) *GetDeepWriteTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDeepWriteTaskResponseBodyData) SetTaskId(v string) *GetDeepWriteTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDeepWriteTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
