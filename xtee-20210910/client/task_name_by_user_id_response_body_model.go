// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskNameByUserIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TaskNameByUserIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *TaskNameByUserIdResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *TaskNameByUserIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *TaskNameByUserIdResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *TaskNameByUserIdResponseBody
	GetResultObject() *bool
}

type TaskNameByUserIdResponseBody struct {
	// Response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s TaskNameByUserIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TaskNameByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *TaskNameByUserIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *TaskNameByUserIdResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *TaskNameByUserIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TaskNameByUserIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TaskNameByUserIdResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *TaskNameByUserIdResponseBody) SetCode(v string) *TaskNameByUserIdResponseBody {
	s.Code = &v
	return s
}

func (s *TaskNameByUserIdResponseBody) SetHttpStatusCode(v string) *TaskNameByUserIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TaskNameByUserIdResponseBody) SetMessage(v string) *TaskNameByUserIdResponseBody {
	s.Message = &v
	return s
}

func (s *TaskNameByUserIdResponseBody) SetRequestId(v string) *TaskNameByUserIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskNameByUserIdResponseBody) SetResultObject(v bool) *TaskNameByUserIdResponseBody {
	s.ResultObject = &v
	return s
}

func (s *TaskNameByUserIdResponseBody) Validate() error {
	return dara.Validate(s)
}
