// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualMFADevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListVirtualMFADevicesResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListVirtualMFADevicesResponseBody
	GetMarker() *string
	SetRequestId(v string) *ListVirtualMFADevicesResponseBody
	GetRequestId() *string
	SetVirtualMFADevices(v *ListVirtualMFADevicesResponseBodyVirtualMFADevices) *ListVirtualMFADevicesResponseBody
	GetVirtualMFADevices() *ListVirtualMFADevicesResponseBodyVirtualMFADevices
}

type ListVirtualMFADevicesResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when `IsTruncated` is `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32272612-DF82-485E-8BA9-AFA4E0C3D0BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the MFA device.
	VirtualMFADevices *ListVirtualMFADevicesResponseBodyVirtualMFADevices `json:"VirtualMFADevices,omitempty" xml:"VirtualMFADevices,omitempty" type:"Struct"`
}

func (s ListVirtualMFADevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListVirtualMFADevicesResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListVirtualMFADevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirtualMFADevicesResponseBody) GetVirtualMFADevices() *ListVirtualMFADevicesResponseBodyVirtualMFADevices {
	return s.VirtualMFADevices
}

func (s *ListVirtualMFADevicesResponseBody) SetIsTruncated(v bool) *ListVirtualMFADevicesResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBody) SetMarker(v string) *ListVirtualMFADevicesResponseBody {
	s.Marker = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBody) SetRequestId(v string) *ListVirtualMFADevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBody) SetVirtualMFADevices(v *ListVirtualMFADevicesResponseBodyVirtualMFADevices) *ListVirtualMFADevicesResponseBody {
	s.VirtualMFADevices = v
	return s
}

func (s *ListVirtualMFADevicesResponseBody) Validate() error {
	if s.VirtualMFADevices != nil {
		if err := s.VirtualMFADevices.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVirtualMFADevicesResponseBodyVirtualMFADevices struct {
	VirtualMFADevice []*ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice `json:"VirtualMFADevice,omitempty" xml:"VirtualMFADevice,omitempty" type:"Repeated"`
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevices) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevices) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevices) GetVirtualMFADevice() []*ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice {
	return s.VirtualMFADevice
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevices) SetVirtualMFADevice(v []*ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) *ListVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.VirtualMFADevice = v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevices) Validate() error {
	if s.VirtualMFADevice != nil {
		for _, item := range s.VirtualMFADevice {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice struct {
	// The time when the MFA device was activated.
	//
	// example:
	//
	// 2020-10-16T06:02:09Z
	ActivateDate *string `json:"ActivateDate,omitempty" xml:"ActivateDate,omitempty"`
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::177242285274****:mfa/test
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The information of the RAM user that has an MFA device bound.
	User *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) GetActivateDate() *string {
	return s.ActivateDate
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) GetUser() *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser {
	return s.User
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) SetActivateDate(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice {
	s.ActivateDate = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) SetSerialNumber(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice {
	s.SerialNumber = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) SetUser(v *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice {
	s.User = v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser struct {
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@177242285274****.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) GetUserId() *string {
	return s.UserId
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) SetDisplayName(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser {
	s.DisplayName = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) SetUserId(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser {
	s.UserId = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) SetUserPrincipalName(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser {
	s.UserPrincipalName = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) Validate() error {
	return dara.Validate(s)
}
