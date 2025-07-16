// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTodoTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateTodoTaskResponseBody
	GetResult() *bool
}

type UpdateTodoTaskResponseBody struct {
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

func (s UpdateTodoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTodoTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateTodoTaskResponseBody) SetRequestId(v string) *UpdateTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTodoTaskResponseBody) SetResult(v bool) *UpdateTodoTaskResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateTodoTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
