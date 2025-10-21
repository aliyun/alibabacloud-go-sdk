// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelOutputContentAsyncDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModelOutputContentAsyncDetectResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModelOutputContentAsyncDetectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModelOutputContentAsyncDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModelOutputContentAsyncDetectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelOutputContentAsyncDetectResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *ModelOutputContentAsyncDetectResponseBody
	GetTaskId() *string
}

type ModelOutputContentAsyncDetectResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 操作是否成功。true表示成功，false表示失败。
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 5d85cd38-03b2-49fd-86b2-be85c4b13215
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModelOutputContentAsyncDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentAsyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ModelOutputContentAsyncDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModelOutputContentAsyncDetectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelOutputContentAsyncDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModelOutputContentAsyncDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelOutputContentAsyncDetectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelOutputContentAsyncDetectResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetCode(v string) *ModelOutputContentAsyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetHttpStatusCode(v int32) *ModelOutputContentAsyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetMessage(v string) *ModelOutputContentAsyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetRequestId(v string) *ModelOutputContentAsyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetSuccess(v bool) *ModelOutputContentAsyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) SetTaskId(v string) *ModelOutputContentAsyncDetectResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModelOutputContentAsyncDetectResponseBody) Validate() error {
	return dara.Validate(s)
}
