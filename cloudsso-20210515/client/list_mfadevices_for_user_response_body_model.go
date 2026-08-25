// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMFADevicesForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMFADevices(v []*ListMFADevicesForUserResponseBodyMFADevices) *ListMFADevicesForUserResponseBody
	GetMFADevices() []*ListMFADevicesForUserResponseBodyMFADevices
	SetRequestId(v string) *ListMFADevicesForUserResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListMFADevicesForUserResponseBody
	GetTotalCounts() *int32
}

type ListMFADevicesForUserResponseBody struct {
	// The MFA device list.
	MFADevices []*ListMFADevicesForUserResponseBodyMFADevices `json:"MFADevices,omitempty" xml:"MFADevices,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8B9982ED-FD0D-5622-8EA0-7B768685DCE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of MFA devices.
	//
	// example:
	//
	// 1
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListMFADevicesForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMFADevicesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListMFADevicesForUserResponseBody) GetMFADevices() []*ListMFADevicesForUserResponseBodyMFADevices {
	return s.MFADevices
}

func (s *ListMFADevicesForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMFADevicesForUserResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListMFADevicesForUserResponseBody) SetMFADevices(v []*ListMFADevicesForUserResponseBodyMFADevices) *ListMFADevicesForUserResponseBody {
	s.MFADevices = v
	return s
}

func (s *ListMFADevicesForUserResponseBody) SetRequestId(v string) *ListMFADevicesForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMFADevicesForUserResponseBody) SetTotalCounts(v int32) *ListMFADevicesForUserResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListMFADevicesForUserResponseBody) Validate() error {
	if s.MFADevices != nil {
		for _, item := range s.MFADevices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMFADevicesForUserResponseBodyMFADevices struct {
	// The MFA device ID.
	//
	// example:
	//
	// mfa-00ujhet8pycljj7j****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The MFA device name.
	//
	// example:
	//
	// Alice-MFA1
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	// The MFA device type. Valid values:
	//
	// - TOTP: a virtual MFA device based on the Time-based One-Time Password algorithm.
	//
	// - CrossPlatformPasskey: a cross-platform passkey.
	//
	// - PlatformPasskey: a platform built-in passkey.
	//
	// example:
	//
	// TOTP
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// The effective period. The time is in UTC and follows the RFC 3339 format (YYYY-MM-DDTHH:mm:ssZ).
	//
	// example:
	//
	// 2021-10-29T09:14:06Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The last time the MFA device was used.
	//
	// example:
	//
	// 2026-08-12T07:26:12Z
	LastUseTime *string `json:"LastUseTime,omitempty" xml:"LastUseTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMFADevicesForUserResponseBodyMFADevices) String() string {
	return dara.Prettify(s)
}

func (s ListMFADevicesForUserResponseBodyMFADevices) GoString() string {
	return s.String()
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) GetDeviceName() *string {
	return s.DeviceName
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) GetLastUseTime() *string {
	return s.LastUseTime
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) GetUserId() *string {
	return s.UserId
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetDeviceId(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.DeviceId = &v
	return s
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetDeviceName(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.DeviceName = &v
	return s
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetDeviceType(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.DeviceType = &v
	return s
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetEffectiveTime(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.EffectiveTime = &v
	return s
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetLastUseTime(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.LastUseTime = &v
	return s
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) SetUserId(v string) *ListMFADevicesForUserResponseBodyMFADevices {
	s.UserId = &v
	return s
}

func (s *ListMFADevicesForUserResponseBodyMFADevices) Validate() error {
	return dara.Validate(s)
}
