// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignJobsAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncTaskId(v string) *AssignJobsAsyncResponseBody
	GetAsyncTaskId() *string
	SetCode(v string) *AssignJobsAsyncResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AssignJobsAsyncResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AssignJobsAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *AssignJobsAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AssignJobsAsyncResponseBody
	GetSuccess() *bool
}

type AssignJobsAsyncResponseBody struct {
	// Asynchronous task ID.
	//
	// example:
	//
	// 12f3dd08-0c55-44ce-9b64-e69d35ed3a76
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// API status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API prompt message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssignJobsAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignJobsAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *AssignJobsAsyncResponseBody) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *AssignJobsAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *AssignJobsAsyncResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AssignJobsAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssignJobsAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignJobsAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AssignJobsAsyncResponseBody) SetAsyncTaskId(v string) *AssignJobsAsyncResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *AssignJobsAsyncResponseBody) SetCode(v string) *AssignJobsAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *AssignJobsAsyncResponseBody) SetHttpStatusCode(v int32) *AssignJobsAsyncResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AssignJobsAsyncResponseBody) SetMessage(v string) *AssignJobsAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *AssignJobsAsyncResponseBody) SetRequestId(v string) *AssignJobsAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignJobsAsyncResponseBody) SetSuccess(v bool) *AssignJobsAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *AssignJobsAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
