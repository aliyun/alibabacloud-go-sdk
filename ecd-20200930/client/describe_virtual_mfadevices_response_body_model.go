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
	// A pagination token. It can be used in the next request to retrieve a new page of results.If NextToken is empty, no next page exists.
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
	// Details of the virtual MFA devices.
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
	return dara.Validate(s)
}

type DescribeVirtualMFADevicesResponseBodyVirtualMFADevices struct {
	// The number of consecutive failures to bind the virtual MFA device, or the number of failures on the verification of the virtual MFA device.
	//
	// example:
	//
	// 1
	ConsecutiveFails *int32 `json:"ConsecutiveFails,omitempty" xml:"ConsecutiveFails,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// cn-hangzhou+dir-gx2x1dhsmu52rd****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the AD user who uses the virtual MFA device.
	//
	// example:
	//
	// usertest
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The time when the virtual MFA device was started. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-20T14:52:28Z
	GmtEnabled *string `json:"GmtEnabled,omitempty" xml:"GmtEnabled,omitempty"`
	// The time when a locked virtual MFA device was automatically unlocked. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-21T15:21:28Z
	GmtUnlock *string `json:"GmtUnlock,omitempty" xml:"GmtUnlock,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou+dir-269345****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The serial number of the virtual MFA device, which is a unique identifier.
	//
	// example:
	//
	// a25f297f-f2e1-4a44-bbf1-5f48a6e5****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The status of the virtual MFA device.
	//
	// Valid values:
	//
	// 	- LOCKED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- UNBOUND
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NORMAL
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
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
	return dara.Validate(s)
}
