// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateQueueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateQueueResponseBody
	GetSuccess() *bool
}

type UpdateQueueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateQueueResponseBody) SetRequestId(v string) *UpdateQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQueueResponseBody) SetSuccess(v bool) *UpdateQueueResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
