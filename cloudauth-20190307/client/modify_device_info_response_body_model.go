// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDay(v string) *ModifyDeviceInfoResponseBody
	GetBeginDay() *string
	SetBizType(v string) *ModifyDeviceInfoResponseBody
	GetBizType() *string
	SetDeviceId(v string) *ModifyDeviceInfoResponseBody
	GetDeviceId() *string
	SetExpiredDay(v string) *ModifyDeviceInfoResponseBody
	GetExpiredDay() *string
	SetRequestId(v string) *ModifyDeviceInfoResponseBody
	GetRequestId() *string
	SetUserDeviceId(v string) *ModifyDeviceInfoResponseBody
	GetUserDeviceId() *string
}

type ModifyDeviceInfoResponseBody struct {
	// If the Duration in the request parameters is not empty, this field represents the start time of the authorization after the device validity period has been extended. One year of Duration is calculated as 365 days. Example: 20180101.
	//
	// example:
	//
	// 20190401
	BeginDay *string `json:"BeginDay,omitempty" xml:"BeginDay,omitempty"`
	// Corresponds to the BizType in the request parameters.
	//
	// example:
	//
	// FACE_TEST
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Corresponds to the DeviceId in the request parameters.
	//
	// example:
	//
	// wd.6ziUffspAeW5FVYbaqmexR-1qwNjM
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// If the Duration in the request parameters is not empty, this field represents the expiration time of the authorization after the device validity period has been extended. One year of Duration is calculated as 365 days. Example: 20180101.
	//
	// example:
	//
	// 20200330
	ExpiredDay *string `json:"ExpiredDay,omitempty" xml:"ExpiredDay,omitempty"`
	// The ID of this request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Corresponds to the UserDeviceId in the request parameters.
	//
	// example:
	//
	// 3iJ1AY$oHcu7mC69
	UserDeviceId *string `json:"UserDeviceId,omitempty" xml:"UserDeviceId,omitempty"`
}

func (s ModifyDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceInfoResponseBody) GetBeginDay() *string {
	return s.BeginDay
}

func (s *ModifyDeviceInfoResponseBody) GetBizType() *string {
	return s.BizType
}

func (s *ModifyDeviceInfoResponseBody) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ModifyDeviceInfoResponseBody) GetExpiredDay() *string {
	return s.ExpiredDay
}

func (s *ModifyDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDeviceInfoResponseBody) GetUserDeviceId() *string {
	return s.UserDeviceId
}

func (s *ModifyDeviceInfoResponseBody) SetBeginDay(v string) *ModifyDeviceInfoResponseBody {
	s.BeginDay = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetBizType(v string) *ModifyDeviceInfoResponseBody {
	s.BizType = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetDeviceId(v string) *ModifyDeviceInfoResponseBody {
	s.DeviceId = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetExpiredDay(v string) *ModifyDeviceInfoResponseBody {
	s.ExpiredDay = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetRequestId(v string) *ModifyDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) SetUserDeviceId(v string) *ModifyDeviceInfoResponseBody {
	s.UserDeviceId = &v
	return s
}

func (s *ModifyDeviceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
