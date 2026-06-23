// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushQueryDeviceStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PushQueryDeviceStateResponseBody
	GetCode() *string
	SetData(v *PushQueryDeviceStateResponseBodyData) *PushQueryDeviceStateResponseBody
	GetData() *PushQueryDeviceStateResponseBodyData
	SetMessage(v string) *PushQueryDeviceStateResponseBody
	GetMessage() *string
	SetRequestId(v string) *PushQueryDeviceStateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PushQueryDeviceStateResponseBody
	GetSuccess() *bool
}

type PushQueryDeviceStateResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *PushQueryDeviceStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushQueryDeviceStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushQueryDeviceStateResponseBody) GoString() string {
	return s.String()
}

func (s *PushQueryDeviceStateResponseBody) GetCode() *string {
	return s.Code
}

func (s *PushQueryDeviceStateResponseBody) GetData() *PushQueryDeviceStateResponseBodyData {
	return s.Data
}

func (s *PushQueryDeviceStateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PushQueryDeviceStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushQueryDeviceStateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PushQueryDeviceStateResponseBody) SetCode(v string) *PushQueryDeviceStateResponseBody {
	s.Code = &v
	return s
}

func (s *PushQueryDeviceStateResponseBody) SetData(v *PushQueryDeviceStateResponseBodyData) *PushQueryDeviceStateResponseBody {
	s.Data = v
	return s
}

func (s *PushQueryDeviceStateResponseBody) SetMessage(v string) *PushQueryDeviceStateResponseBody {
	s.Message = &v
	return s
}

func (s *PushQueryDeviceStateResponseBody) SetRequestId(v string) *PushQueryDeviceStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushQueryDeviceStateResponseBody) SetSuccess(v bool) *PushQueryDeviceStateResponseBody {
	s.Success = &v
	return s
}

func (s *PushQueryDeviceStateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushQueryDeviceStateResponseBodyData struct {
	DeliveryToken *string `json:"DeliveryToken,omitempty" xml:"DeliveryToken,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Manufacturer  *string `json:"Manufacturer,omitempty" xml:"Manufacturer,omitempty"`
	Platform      *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Statue        *string `json:"Statue,omitempty" xml:"Statue,omitempty"`
	ThirdToken    *string `json:"ThirdToken,omitempty" xml:"ThirdToken,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s PushQueryDeviceStateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PushQueryDeviceStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *PushQueryDeviceStateResponseBodyData) GetDeliveryToken() *string {
	return s.DeliveryToken
}

func (s *PushQueryDeviceStateResponseBodyData) GetDeviceId() *string {
	return s.DeviceId
}

func (s *PushQueryDeviceStateResponseBodyData) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *PushQueryDeviceStateResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *PushQueryDeviceStateResponseBodyData) GetStatue() *string {
	return s.Statue
}

func (s *PushQueryDeviceStateResponseBodyData) GetThirdToken() *string {
	return s.ThirdToken
}

func (s *PushQueryDeviceStateResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *PushQueryDeviceStateResponseBodyData) SetDeliveryToken(v string) *PushQueryDeviceStateResponseBodyData {
	s.DeliveryToken = &v
	return s
}

func (s *PushQueryDeviceStateResponseBodyData) SetDeviceId(v string) *PushQueryDeviceStateResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *PushQueryDeviceStateResponseBodyData) SetManufacturer(v string) *PushQueryDeviceStateResponseBodyData {
	s.Manufacturer = &v
	return s
}

func (s *PushQueryDeviceStateResponseBodyData) SetPlatform(v string) *PushQueryDeviceStateResponseBodyData {
	s.Platform = &v
	return s
}

func (s *PushQueryDeviceStateResponseBodyData) SetStatue(v string) *PushQueryDeviceStateResponseBodyData {
	s.Statue = &v
	return s
}

func (s *PushQueryDeviceStateResponseBodyData) SetThirdToken(v string) *PushQueryDeviceStateResponseBodyData {
	s.ThirdToken = &v
	return s
}

func (s *PushQueryDeviceStateResponseBodyData) SetUserId(v string) *PushQueryDeviceStateResponseBodyData {
	s.UserId = &v
	return s
}

func (s *PushQueryDeviceStateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
