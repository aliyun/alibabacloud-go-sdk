// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectPolarClawDevicePairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RejectPolarClawDevicePairResponseBody
	GetApplicationId() *string
	SetCode(v int32) *RejectPolarClawDevicePairResponseBody
	GetCode() *int32
	SetDeviceId(v string) *RejectPolarClawDevicePairResponseBody
	GetDeviceId() *string
	SetMessage(v string) *RejectPolarClawDevicePairResponseBody
	GetMessage() *string
	SetPairRequestId(v string) *RejectPolarClawDevicePairResponseBody
	GetPairRequestId() *string
	SetRequestId(v string) *RejectPolarClawDevicePairResponseBody
	GetRequestId() *string
}

type RejectPolarClawDevicePairResponseBody struct {
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
	// The unique device ID.
	//
	// example:
	//
	// device-784x37k0vko734fk
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pairing request ID.
	//
	// example:
	//
	// req-abc-123
	PairRequestId *string `json:"PairRequestId,omitempty" xml:"PairRequestId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectPolarClawDevicePairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RejectPolarClawDevicePairResponseBody) GoString() string {
	return s.String()
}

func (s *RejectPolarClawDevicePairResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RejectPolarClawDevicePairResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RejectPolarClawDevicePairResponseBody) GetDeviceId() *string {
	return s.DeviceId
}

func (s *RejectPolarClawDevicePairResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RejectPolarClawDevicePairResponseBody) GetPairRequestId() *string {
	return s.PairRequestId
}

func (s *RejectPolarClawDevicePairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RejectPolarClawDevicePairResponseBody) SetApplicationId(v string) *RejectPolarClawDevicePairResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *RejectPolarClawDevicePairResponseBody) SetCode(v int32) *RejectPolarClawDevicePairResponseBody {
	s.Code = &v
	return s
}

func (s *RejectPolarClawDevicePairResponseBody) SetDeviceId(v string) *RejectPolarClawDevicePairResponseBody {
	s.DeviceId = &v
	return s
}

func (s *RejectPolarClawDevicePairResponseBody) SetMessage(v string) *RejectPolarClawDevicePairResponseBody {
	s.Message = &v
	return s
}

func (s *RejectPolarClawDevicePairResponseBody) SetPairRequestId(v string) *RejectPolarClawDevicePairResponseBody {
	s.PairRequestId = &v
	return s
}

func (s *RejectPolarClawDevicePairResponseBody) SetRequestId(v string) *RejectPolarClawDevicePairResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectPolarClawDevicePairResponseBody) Validate() error {
	return dara.Validate(s)
}
