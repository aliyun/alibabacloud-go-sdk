// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTodoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTodoTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteTodoTaskResponseBody
	GetResult() *bool
}

type DeleteTodoTaskResponseBody struct {
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

func (s DeleteTodoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTodoTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteTodoTaskResponseBody) SetRequestId(v string) *DeleteTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTodoTaskResponseBody) SetResult(v bool) *DeleteTodoTaskResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteTodoTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
