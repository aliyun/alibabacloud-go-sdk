// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteQueuesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQueuesResponseBody
	GetSuccess() *bool
}

type DeleteQueuesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQueuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQueuesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQueuesResponseBody) SetRequestId(v string) *DeleteQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQueuesResponseBody) SetSuccess(v bool) *DeleteQueuesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQueuesResponseBody) Validate() error {
	return dara.Validate(s)
}
