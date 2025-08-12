// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableHostAvailabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableHostAvailabilityResponseBody
	GetCode() *string
	SetMessage(v string) *DisableHostAvailabilityResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableHostAvailabilityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableHostAvailabilityResponseBody
	GetSuccess() *bool
}

type DisableHostAvailabilityResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ACBDBB40-DFB6-4F4C-8957-51FFB233969C
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

func (s DisableHostAvailabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DisableHostAvailabilityResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableHostAvailabilityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableHostAvailabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableHostAvailabilityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableHostAvailabilityResponseBody) SetCode(v string) *DisableHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *DisableHostAvailabilityResponseBody) SetMessage(v string) *DisableHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *DisableHostAvailabilityResponseBody) SetRequestId(v string) *DisableHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableHostAvailabilityResponseBody) SetSuccess(v bool) *DisableHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

func (s *DisableHostAvailabilityResponseBody) Validate() error {
	return dara.Validate(s)
}
