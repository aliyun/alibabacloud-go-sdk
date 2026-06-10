// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualMFADevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeVirtualMFADevicesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeVirtualMFADevicesResponseBody
	GetRequestId() *string
	SetVirtualMFADevices(v []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) *DescribeVirtualMFADevicesResponseBody
	GetVirtualMFADevices() []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices
}

type DescribeVirtualMFADevicesResponseBody struct {
	// The token to retrieve the next page of results. If this parameter is empty, no more results are available.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL23as
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FB550AAB-FB36-4A91-93F6-F4374AF65403
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of virtual MFA devices.
	VirtualMFADevices []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices `json:"VirtualMFADevices,omitempty" xml:"VirtualMFADevices,omitempty" type:"Repeated"`
}

func (s DescribeVirtualMFADevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVirtualMFADevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVirtualMFADevicesResponseBody) GetVirtualMFADevices() []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	return s.VirtualMFADevices
}

func (s *DescribeVirtualMFADevicesResponseBody) SetNextToken(v string) *DescribeVirtualMFADevicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBody) SetRequestId(v string) *DescribeVirtualMFADevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBody) SetVirtualMFADevices(v []*DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) *DescribeVirtualMFADevicesResponseBody {
	s.VirtualMFADevices = v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBody) Validate() error {
	if s.VirtualMFADevices != nil {
		for _, item := range s.VirtualMFADevices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVirtualMFADevicesResponseBodyVirtualMFADevices struct {
	AdUser *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser `json:"AdUser,omitempty" xml:"AdUser,omitempty" type:"Struct"`
	// The number of consecutive failed attempts to bind or authenticate the virtual MFA device.
	//
	// example:
	//
	// 1
	ConsecutiveFails *int32 `json:"ConsecutiveFails,omitempty" xml:"ConsecutiveFails,omitempty"`
	// > This parameter is in private preview.
	//
	// example:
	//
	// cn-hangzhou+dir-gx2x1dhsmu52rd****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The AD username of the bound user.
	//
	// example:
	//
	// usertest
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The time when the virtual MFA device was enabled. The time is in the `YYYY-MM-DDThh:mm:ssZ` format and in UTC, as specified by the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard.
	//
	// example:
	//
	// 2020-12-20T14:52:28Z
	GmtEnabled *string `json:"GmtEnabled,omitempty" xml:"GmtEnabled,omitempty"`
	// The time when the locked virtual MFA device is automatically unlocked. The time is in the `YYYY-MM-DDThh:mm:ssZ` format and in UTC, as specified by the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard.
	//
	// example:
	//
	// 2020-12-21T15:21:28Z
	GmtUnlock *string `json:"GmtUnlock,omitempty" xml:"GmtUnlock,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// cn-hangzhou+dir-269345****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The serial number of the virtual MFA device.
	//
	// example:
	//
	// a25f297f-f2e1-4a44-bbf1-5f48a6e5****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The status of the virtual MFA device.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GetAdUser() *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser {
	return s.AdUser
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GetConsecutiveFails() *int32 {
	return s.ConsecutiveFails
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GetGmtEnabled() *string {
	return s.GmtEnabled
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GetGmtUnlock() *string {
	return s.GmtUnlock
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) GetStatus() *string {
	return s.Status
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetAdUser(v *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.AdUser = v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetConsecutiveFails(v int32) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.ConsecutiveFails = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetDirectoryId(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.DirectoryId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetEndUserId(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.EndUserId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetGmtEnabled(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.GmtEnabled = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetGmtUnlock(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.GmtUnlock = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetOfficeSiteId(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetSerialNumber(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.SerialNumber = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) SetStatus(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.Status = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevices) Validate() error {
	if s.AdUser != nil {
		if err := s.AdUser.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser struct {
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	DisplayNameNew    *string `json:"DisplayNameNew,omitempty" xml:"DisplayNameNew,omitempty"`
	EndUser           *string `json:"EndUser,omitempty" xml:"EndUser,omitempty"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) GoString() string {
	return s.String()
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) GetDisplayNameNew() *string {
	return s.DisplayNameNew
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) GetEndUser() *string {
	return s.EndUser
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) SetDisplayName(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser {
	s.DisplayName = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) SetDisplayNameNew(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser {
	s.DisplayNameNew = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) SetEndUser(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser {
	s.EndUser = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) SetUserPrincipalName(v string) *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser {
	s.UserPrincipalName = &v
	return s
}

func (s *DescribeVirtualMFADevicesResponseBodyVirtualMFADevicesAdUser) Validate() error {
	return dara.Validate(s)
}
