// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostAvailabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHostAvailabilityResponseBody
	GetCode() *string
	SetMessage(v string) *CreateHostAvailabilityResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHostAvailabilityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateHostAvailabilityResponseBody
	GetSuccess() *bool
	SetTaskId(v int64) *CreateHostAvailabilityResponseBody
	GetTaskId() *int64
}

type CreateHostAvailabilityResponseBody struct {
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
	// The ID of the availability monitoring task.
	//
	// example:
	//
	// 12345
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateHostAvailabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHostAvailabilityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostAvailabilityResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHostAvailabilityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHostAvailabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHostAvailabilityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHostAvailabilityResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateHostAvailabilityResponseBody) SetCode(v string) *CreateHostAvailabilityResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHostAvailabilityResponseBody) SetMessage(v string) *CreateHostAvailabilityResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHostAvailabilityResponseBody) SetRequestId(v string) *CreateHostAvailabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHostAvailabilityResponseBody) SetSuccess(v bool) *CreateHostAvailabilityResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHostAvailabilityResponseBody) SetTaskId(v int64) *CreateHostAvailabilityResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateHostAvailabilityResponseBody) Validate() error {
	return dara.Validate(s)
}
