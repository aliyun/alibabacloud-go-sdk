// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDevices(v []*DescribeDevicesResponseBodyDevices) *DescribeDevicesResponseBody
	GetDevices() []*DescribeDevicesResponseBodyDevices
	SetRequestId(v string) *DescribeDevicesResponseBody
	GetRequestId() *string
}

type DescribeDevicesResponseBody struct {
	// The information about devices that you queried.
	Devices []*DescribeDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 5BEFE642-A383-4A18-8939-FB7DE452****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBody) GetDevices() []*DescribeDevicesResponseBodyDevices {
	return s.Devices
}

func (s *DescribeDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDevicesResponseBody) SetDevices(v []*DescribeDevicesResponseBodyDevices) *DescribeDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *DescribeDevicesResponseBody) SetRequestId(v string) *DescribeDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDevicesResponseBody) Validate() error {
	if s.Devices != nil {
		for _, item := range s.Devices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDevicesResponseBodyDevices struct {
	// The ID of the device. The serial number (SN) of the hardware client or the UUID of the software client.
	//
	// example:
	//
	// 5F52817BE267A43C608D245070D2****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The users who are bound to the device.
	EndUserList []*DescribeDevicesResponseBodyDevicesEndUserList `json:"EndUserList,omitempty" xml:"EndUserList,omitempty" type:"Repeated"`
}

func (s DescribeDevicesResponseBodyDevices) String() string {
	return dara.Prettify(s)
}

func (s DescribeDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyDevices) GetDeviceId() *string {
	return s.DeviceId
}

func (s *DescribeDevicesResponseBodyDevices) GetEndUserList() []*DescribeDevicesResponseBodyDevicesEndUserList {
	return s.EndUserList
}

func (s *DescribeDevicesResponseBodyDevices) SetDeviceId(v string) *DescribeDevicesResponseBodyDevices {
	s.DeviceId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) SetEndUserList(v []*DescribeDevicesResponseBodyDevicesEndUserList) *DescribeDevicesResponseBodyDevices {
	s.EndUserList = v
	return s
}

func (s *DescribeDevicesResponseBodyDevices) Validate() error {
	if s.EndUserList != nil {
		for _, item := range s.EndUserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDevicesResponseBodyDevicesEndUserList struct {
	// The address of the AD office network.
	//
	// example:
	//
	// xn--0zw****
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The ID of the convenient office network.
	//
	// example:
	//
	// cn-hangzhou+dir-jedbpr4sl9l37****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// moli
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The account type of the user.
	//
	// Valid values:
	//
	// 	- AD: enterprise AD account.
	//
	// 	- SIMPLE: convenience account
	//
	// example:
	//
	// SIMPLE
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeDevicesResponseBodyDevicesEndUserList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDevicesResponseBodyDevicesEndUserList) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyDevicesEndUserList) GetAdDomain() *string {
	return s.AdDomain
}

func (s *DescribeDevicesResponseBodyDevicesEndUserList) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeDevicesResponseBodyDevicesEndUserList) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDevicesResponseBodyDevicesEndUserList) GetUserType() *string {
	return s.UserType
}

func (s *DescribeDevicesResponseBodyDevicesEndUserList) SetAdDomain(v string) *DescribeDevicesResponseBodyDevicesEndUserList {
	s.AdDomain = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesEndUserList) SetDirectoryId(v string) *DescribeDevicesResponseBodyDevicesEndUserList {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesEndUserList) SetEndUserId(v string) *DescribeDevicesResponseBodyDevicesEndUserList {
	s.EndUserId = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesEndUserList) SetUserType(v string) *DescribeDevicesResponseBodyDevicesEndUserList {
	s.UserType = &v
	return s
}

func (s *DescribeDevicesResponseBodyDevicesEndUserList) Validate() error {
	return dara.Validate(s)
}
