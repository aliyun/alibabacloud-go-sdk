// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTaskResponseBody
	GetSuccess() *bool
}

type UpdateTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTaskResponseBody) SetRequestId(v string) *UpdateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskResponseBody) SetSuccess(v bool) *UpdateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
