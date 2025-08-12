// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostAvailabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyHostAvailabilityResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyHostAvailabilityResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyHostAvailabilityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyHostAvailabilityResponseBody
	GetSuccess() *bool
}

type ModifyHostAvailabilityResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16E815A3-47E1-4290-87F9-D5C99471FF45
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

func (s ModifyHostAvailabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostAvailabilityResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyHostAvailabilityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyHostAvailabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHostAvailabilityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyHostAvailabilityResponseBody) SetCode(v string) *ModifyHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyHostAvailabilityResponseBody) SetMessage(v string) *ModifyHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyHostAvailabilityResponseBody) SetRequestId(v string) *ModifyHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostAvailabilityResponseBody) SetSuccess(v bool) *ModifyHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyHostAvailabilityResponseBody) Validate() error {
	return dara.Validate(s)
}
