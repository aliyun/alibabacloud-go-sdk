// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelInputContentAsyncDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModelInputContentAsyncDetectResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModelInputContentAsyncDetectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModelInputContentAsyncDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModelInputContentAsyncDetectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelInputContentAsyncDetectResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *ModelInputContentAsyncDetectResponseBody
	GetTaskId() *string
}

type ModelInputContentAsyncDetectResponseBody struct {
	// Result code, 00000 indicates success; others indicate failure.
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

func (s ModelInputContentAsyncDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentAsyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *ModelInputContentAsyncDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModelInputContentAsyncDetectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelInputContentAsyncDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModelInputContentAsyncDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelInputContentAsyncDetectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelInputContentAsyncDetectResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModelInputContentAsyncDetectResponseBody) SetCode(v string) *ModelInputContentAsyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) SetHttpStatusCode(v int32) *ModelInputContentAsyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) SetMessage(v string) *ModelInputContentAsyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) SetRequestId(v string) *ModelInputContentAsyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) SetSuccess(v bool) *ModelInputContentAsyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) SetTaskId(v string) *ModelInputContentAsyncDetectResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModelInputContentAsyncDetectResponseBody) Validate() error {
	return dara.Validate(s)
}
