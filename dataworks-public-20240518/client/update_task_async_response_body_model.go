// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *UpdateTaskAsyncResponseBody
	GetOperationId() *string
	SetRequestId(v string) *UpdateTaskAsyncResponseBody
	GetRequestId() *string
}

type UpdateTaskAsyncResponseBody struct {
	// The operation ID, which is used to obtain the result of the asynchronous node update. You can call the UpdateTaskAsync operation to obtain the result.
	//
	// example:
	//
	// e15ad21c-b0e9-4792-8f55-b037xxxxxxxx
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The request ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 10000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTaskAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskAsyncResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *UpdateTaskAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskAsyncResponseBody) SetOperationId(v string) *UpdateTaskAsyncResponseBody {
	s.OperationId = &v
	return s
}

func (s *UpdateTaskAsyncResponseBody) SetRequestId(v string) *UpdateTaskAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
