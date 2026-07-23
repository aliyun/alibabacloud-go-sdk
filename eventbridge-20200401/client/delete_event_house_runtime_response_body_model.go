// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventHouseRuntimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEventHouseRuntimeResponseBody
	GetCode() *string
	SetData(v *EventHouseRuntime) *DeleteEventHouseRuntimeResponseBody
	GetData() *EventHouseRuntime
	SetMessage(v string) *DeleteEventHouseRuntimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEventHouseRuntimeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEventHouseRuntimeResponseBody
	GetSuccess() *bool
}

type DeleteEventHouseRuntimeResponseBody struct {
	// The response code. Success indicates that the operation was successful.
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

func (s DeleteEventHouseRuntimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventHouseRuntimeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventHouseRuntimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEventHouseRuntimeResponseBody) GetData() *EventHouseRuntime {
	return s.Data
}

func (s *DeleteEventHouseRuntimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEventHouseRuntimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventHouseRuntimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEventHouseRuntimeResponseBody) SetCode(v string) *DeleteEventHouseRuntimeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventHouseRuntimeResponseBody) SetData(v *EventHouseRuntime) *DeleteEventHouseRuntimeResponseBody {
	s.Data = v
	return s
}

func (s *DeleteEventHouseRuntimeResponseBody) SetMessage(v string) *DeleteEventHouseRuntimeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventHouseRuntimeResponseBody) SetRequestId(v string) *DeleteEventHouseRuntimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventHouseRuntimeResponseBody) SetSuccess(v bool) *DeleteEventHouseRuntimeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEventHouseRuntimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
