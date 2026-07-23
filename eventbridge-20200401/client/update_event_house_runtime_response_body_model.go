// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventHouseRuntimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateEventHouseRuntimeResponseBody
	GetCode() *string
	SetData(v *EventHouseRuntime) *UpdateEventHouseRuntimeResponseBody
	GetData() *EventHouseRuntime
	SetMessage(v string) *UpdateEventHouseRuntimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEventHouseRuntimeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateEventHouseRuntimeResponseBody
	GetSuccess() *bool
}

type UpdateEventHouseRuntimeResponseBody struct {
	// The response code. Success indicates that the operation is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The Runtime information.
	Data *EventHouseRuntime `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventHouseRuntimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventHouseRuntimeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventHouseRuntimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateEventHouseRuntimeResponseBody) GetData() *EventHouseRuntime {
	return s.Data
}

func (s *UpdateEventHouseRuntimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEventHouseRuntimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEventHouseRuntimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEventHouseRuntimeResponseBody) SetCode(v string) *UpdateEventHouseRuntimeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventHouseRuntimeResponseBody) SetData(v *EventHouseRuntime) *UpdateEventHouseRuntimeResponseBody {
	s.Data = v
	return s
}

func (s *UpdateEventHouseRuntimeResponseBody) SetMessage(v string) *UpdateEventHouseRuntimeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEventHouseRuntimeResponseBody) SetRequestId(v string) *UpdateEventHouseRuntimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventHouseRuntimeResponseBody) SetSuccess(v bool) *UpdateEventHouseRuntimeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEventHouseRuntimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
