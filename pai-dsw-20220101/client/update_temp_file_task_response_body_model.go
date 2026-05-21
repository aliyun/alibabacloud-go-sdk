// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTempFileTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTempFileTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateTempFileTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateTempFileTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTempFileTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTempFileTaskResponseBody
	GetSuccess() *bool
	SetTempFileTaskId(v string) *UpdateTempFileTaskResponseBody
	GetTempFileTaskId() *string
}

type UpdateTempFileTaskResponseBody struct {
	// example:
	//
	// "200"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// task-05cexxxxxxxxx
	TempFileTaskId *string `json:"TempFileTaskId,omitempty" xml:"TempFileTaskId,omitempty"`
}

func (s UpdateTempFileTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTempFileTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTempFileTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTempFileTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateTempFileTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTempFileTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTempFileTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTempFileTaskResponseBody) GetTempFileTaskId() *string {
	return s.TempFileTaskId
}

func (s *UpdateTempFileTaskResponseBody) SetCode(v string) *UpdateTempFileTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTempFileTaskResponseBody) SetHttpStatusCode(v int32) *UpdateTempFileTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTempFileTaskResponseBody) SetMessage(v string) *UpdateTempFileTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTempFileTaskResponseBody) SetRequestId(v string) *UpdateTempFileTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTempFileTaskResponseBody) SetSuccess(v bool) *UpdateTempFileTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTempFileTaskResponseBody) SetTempFileTaskId(v string) *UpdateTempFileTaskResponseBody {
	s.TempFileTaskId = &v
	return s
}

func (s *UpdateTempFileTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
