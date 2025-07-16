// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskExecutorStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTodoTaskExecutorStatusResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateTodoTaskExecutorStatusResponseBody
	GetResult() *bool
}

type UpdateTodoTaskExecutorStatusResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTodoTaskExecutorStatusResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateTodoTaskExecutorStatusResponseBody) SetRequestId(v string) *UpdateTodoTaskExecutorStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponseBody) SetResult(v bool) *UpdateTodoTaskExecutorStatusResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
