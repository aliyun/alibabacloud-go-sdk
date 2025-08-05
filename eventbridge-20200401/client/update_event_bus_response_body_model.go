// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventBusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateEventBusResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateEventBusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEventBusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateEventBusResponseBody
	GetSuccess() *bool
}

type UpdateEventBusResponseBody struct {
	// The response code. Valid values:
	//
	// 	- Success: The request was successful.
	//
	// 	- Other codes: The request failed. For information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The event bus [xxxx] not existed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// f2099962-1628-45f1-9782-2bf6daad823f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. If the operation was successful, the value true is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventBusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventBusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventBusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateEventBusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEventBusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEventBusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEventBusResponseBody) SetCode(v string) *UpdateEventBusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventBusResponseBody) SetMessage(v string) *UpdateEventBusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEventBusResponseBody) SetRequestId(v string) *UpdateEventBusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventBusResponseBody) SetSuccess(v bool) *UpdateEventBusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEventBusResponseBody) Validate() error {
	return dara.Validate(s)
}
