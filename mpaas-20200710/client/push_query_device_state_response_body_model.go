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
	// example:
	//
	// 100
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *PushQueryDeviceStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// ad-000f18w8vmqtzhvbopge-854
	DeliveryToken *string `json:"DeliveryToken,omitempty" xml:"DeliveryToken,omitempty"`
	// example:
	//
	// ad-000f18w8vmqtzhvbopge-854
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// miui
	Manufacturer *string `json:"Manufacturer,omitempty" xml:"Manufacturer,omitempty"`
	// example:
	//
	// android
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// ONLINE
	Statue *string `json:"Statue,omitempty" xml:"Statue,omitempty"`
	// example:
	//
	// IQAAAACy0f7tAABYiMwLEENtr0TKYJEsv7wyu4Ubt9XXwTJAlknnCb1LAzB3wJvoZIcT_nJdaMhEoXJaqQrObAGHLGoU1GOexlTcLWzja-0HfGHKBw
	ThirdToken *string `json:"ThirdToken,omitempty" xml:"ThirdToken,omitempty"`
	// example:
	//
	// push_test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
