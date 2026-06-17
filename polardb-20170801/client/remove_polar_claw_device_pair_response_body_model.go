// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePolarClawDevicePairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RemovePolarClawDevicePairResponseBody
	GetApplicationId() *string
	SetCode(v int32) *RemovePolarClawDevicePairResponseBody
	GetCode() *int32
	SetDeviceId(v string) *RemovePolarClawDevicePairResponseBody
	GetDeviceId() *string
	SetMessage(v string) *RemovePolarClawDevicePairResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemovePolarClawDevicePairResponseBody
	GetRequestId() *string
}

type RemovePolarClawDevicePairResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The device ID.
	//
	// example:
	//
	// device-mac-789
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 580EF224-9647-59E7-9950-D9EBFD6A2921
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePolarClawDevicePairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePolarClawDevicePairResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePolarClawDevicePairResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RemovePolarClawDevicePairResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RemovePolarClawDevicePairResponseBody) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RemovePolarClawDevicePairResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemovePolarClawDevicePairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePolarClawDevicePairResponseBody) SetApplicationId(v string) *RemovePolarClawDevicePairResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *RemovePolarClawDevicePairResponseBody) SetCode(v int32) *RemovePolarClawDevicePairResponseBody {
	s.Code = &v
	return s
}

func (s *RemovePolarClawDevicePairResponseBody) SetDeviceId(v string) *RemovePolarClawDevicePairResponseBody {
	s.DeviceId = &v
	return s
}

func (s *RemovePolarClawDevicePairResponseBody) SetMessage(v string) *RemovePolarClawDevicePairResponseBody {
	s.Message = &v
	return s
}

func (s *RemovePolarClawDevicePairResponseBody) SetRequestId(v string) *RemovePolarClawDevicePairResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePolarClawDevicePairResponseBody) Validate() error {
	return dara.Validate(s)
}
