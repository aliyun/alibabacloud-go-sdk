// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventHouseRuntimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEventHouseRuntimeResponseBody
	GetCode() *string
	SetData(v *EventHouseRuntime) *GetEventHouseRuntimeResponseBody
	GetData() *EventHouseRuntime
	SetMessage(v string) *GetEventHouseRuntimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEventHouseRuntimeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEventHouseRuntimeResponseBody
	GetSuccess() *bool
}

type GetEventHouseRuntimeResponseBody struct {
	// The response code. Success indicates a successful operation.
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
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEventHouseRuntimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventHouseRuntimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventHouseRuntimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEventHouseRuntimeResponseBody) GetData() *EventHouseRuntime {
	return s.Data
}

func (s *GetEventHouseRuntimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEventHouseRuntimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEventHouseRuntimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEventHouseRuntimeResponseBody) SetCode(v string) *GetEventHouseRuntimeResponseBody {
	s.Code = &v
	return s
}

func (s *GetEventHouseRuntimeResponseBody) SetData(v *EventHouseRuntime) *GetEventHouseRuntimeResponseBody {
	s.Data = v
	return s
}

func (s *GetEventHouseRuntimeResponseBody) SetMessage(v string) *GetEventHouseRuntimeResponseBody {
	s.Message = &v
	return s
}

func (s *GetEventHouseRuntimeResponseBody) SetRequestId(v string) *GetEventHouseRuntimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventHouseRuntimeResponseBody) SetSuccess(v bool) *GetEventHouseRuntimeResponseBody {
	s.Success = &v
	return s
}

func (s *GetEventHouseRuntimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
