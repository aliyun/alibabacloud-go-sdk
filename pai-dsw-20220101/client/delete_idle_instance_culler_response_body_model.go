// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdleInstanceCullerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteIdleInstanceCullerResponseBody
	GetCode() *string
	SetInstanceId(v string) *DeleteIdleInstanceCullerResponseBody
	GetInstanceId() *string
	SetMessage(v string) *DeleteIdleInstanceCullerResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIdleInstanceCullerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteIdleInstanceCullerResponseBody
	GetSuccess() *bool
}

type DeleteIdleInstanceCullerResponseBody struct {
	// The status code. Valid values:
	//
	// 	- InternalError: an internal error. All errors, except for parameter validation errors, are classified as internal errors.
	//
	// 	- ValidationError: a parameter validation error.
	//
	// example:
	//
	// ValidationError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The response message.
	//
	// 	- If the request is successful, null is returned.
	//
	// 	- If the request fails, the failure cause is returned.
	//
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
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

func (s DeleteIdleInstanceCullerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdleInstanceCullerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIdleInstanceCullerResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteIdleInstanceCullerResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteIdleInstanceCullerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIdleInstanceCullerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIdleInstanceCullerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIdleInstanceCullerResponseBody) SetCode(v string) *DeleteIdleInstanceCullerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponseBody) SetInstanceId(v string) *DeleteIdleInstanceCullerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponseBody) SetMessage(v string) *DeleteIdleInstanceCullerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponseBody) SetRequestId(v string) *DeleteIdleInstanceCullerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponseBody) SetSuccess(v bool) *DeleteIdleInstanceCullerResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIdleInstanceCullerResponseBody) Validate() error {
	return dara.Validate(s)
}
