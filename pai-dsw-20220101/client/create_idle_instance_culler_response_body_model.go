// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdleInstanceCullerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateIdleInstanceCullerResponseBody
	GetCode() *string
	SetInstanceId(v string) *CreateIdleInstanceCullerResponseBody
	GetInstanceId() *string
	SetMessage(v string) *CreateIdleInstanceCullerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIdleInstanceCullerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateIdleInstanceCullerResponseBody
	GetSuccess() *bool
}

type CreateIdleInstanceCullerResponseBody struct {
	// The status code. Valid values:
	//
	// 	- InternalError: an internal error. All errors, except for parameter validation errors, are classified as internal errors.
	//
	// 	- ValidationError: a parameter validation error.
	//
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The error message.
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

func (s CreateIdleInstanceCullerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdleInstanceCullerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdleInstanceCullerResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateIdleInstanceCullerResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateIdleInstanceCullerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIdleInstanceCullerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIdleInstanceCullerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateIdleInstanceCullerResponseBody) SetCode(v string) *CreateIdleInstanceCullerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIdleInstanceCullerResponseBody) SetInstanceId(v string) *CreateIdleInstanceCullerResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateIdleInstanceCullerResponseBody) SetMessage(v string) *CreateIdleInstanceCullerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIdleInstanceCullerResponseBody) SetRequestId(v string) *CreateIdleInstanceCullerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdleInstanceCullerResponseBody) SetSuccess(v bool) *CreateIdleInstanceCullerResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIdleInstanceCullerResponseBody) Validate() error {
	return dara.Validate(s)
}
