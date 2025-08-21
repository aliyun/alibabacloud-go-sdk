// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeviceControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeviceControlResponseBody
	GetCode() *int32
	SetMessage(v string) *DeviceControlResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeviceControlResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeviceControlResponseBody
	GetResult() *bool
}

type DeviceControlResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeviceControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeviceControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeviceControlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeviceControlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeviceControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeviceControlResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeviceControlResponseBody) SetCode(v int32) *DeviceControlResponseBody {
	s.Code = &v
	return s
}

func (s *DeviceControlResponseBody) SetMessage(v string) *DeviceControlResponseBody {
	s.Message = &v
	return s
}

func (s *DeviceControlResponseBody) SetRequestId(v string) *DeviceControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeviceControlResponseBody) SetResult(v bool) *DeviceControlResponseBody {
	s.Result = &v
	return s
}

func (s *DeviceControlResponseBody) Validate() error {
	return dara.Validate(s)
}
