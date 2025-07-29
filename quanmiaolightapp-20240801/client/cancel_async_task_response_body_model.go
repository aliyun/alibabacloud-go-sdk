// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelAsyncTaskResponseBody
	GetCode() *string
	SetData(v bool) *CancelAsyncTaskResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *CancelAsyncTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CancelAsyncTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelAsyncTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelAsyncTaskResponseBody
	GetSuccess() *bool
}

type CancelAsyncTaskResponseBody struct {
	// example:
	//
	// NoPermission
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAsyncTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelAsyncTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *CancelAsyncTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CancelAsyncTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAsyncTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelAsyncTaskResponseBody) SetCode(v string) *CancelAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetData(v bool) *CancelAsyncTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetHttpStatusCode(v int32) *CancelAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetMessage(v string) *CancelAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetRequestId(v string) *CancelAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetSuccess(v bool) *CancelAsyncTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
