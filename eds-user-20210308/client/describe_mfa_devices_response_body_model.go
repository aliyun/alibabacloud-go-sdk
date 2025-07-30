// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMfaDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMfaDevices(v []*DescribeMfaDevicesResponseBodyMfaDevices) *DescribeMfaDevicesResponseBody
	GetMfaDevices() []*DescribeMfaDevicesResponseBodyMfaDevices
	SetNextToken(v string) *DescribeMfaDevicesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeMfaDevicesResponseBody
	GetRequestId() *string
}

type DescribeMfaDevicesResponseBody struct {
	// The information about the virtual MFA devices.
	MfaDevices []*DescribeMfaDevicesResponseBodyMfaDevices `json:"MfaDevices,omitempty" xml:"MfaDevices,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6na6YlN9asMM31MsMcdQNpp
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 33DBB8EC-6E68-4726-91C4-E09C59D9A7D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMfaDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMfaDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesResponseBody) GetMfaDevices() []*DescribeMfaDevicesResponseBodyMfaDevices {
	return s.MfaDevices
}

func (s *DescribeMfaDevicesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMfaDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMfaDevicesResponseBody) SetMfaDevices(v []*DescribeMfaDevicesResponseBodyMfaDevices) *DescribeMfaDevicesResponseBody {
	s.MfaDevices = v
	return s
}

func (s *DescribeMfaDevicesResponseBody) SetNextToken(v string) *DescribeMfaDevicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMfaDevicesResponseBody) SetRequestId(v string) *DescribeMfaDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMfaDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMfaDevicesResponseBodyMfaDevices struct {
	// The number of consecutive failures to bind the virtual MFA device, or the number of authentication failures based on the virtual MFA device.
	//
	// example:
	//
	// 0
	ConsecutiveFails *int32 `json:"ConsecutiveFails,omitempty" xml:"ConsecutiveFails,omitempty"`
	// The type of the virtual MFA device. The value can only be TOTP_VIRTUAL. This value indicates that the virtual MFA device follows the Time-based One-time Password (TOTP) algorithm.
	//
	// example:
	//
	// TOTP_VIRTUAL
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username of the convenience user that uses the virtual MFA device.
	//
	// example:
	//
	// test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The time when the virtual MFA device was enabled. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-22T06:20:49Z
	GmtEnabled *string `json:"GmtEnabled,omitempty" xml:"GmtEnabled,omitempty"`
	// The time when the locked virtual MFA device was automatically unlocked. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-23T06:20:49Z
	GmtUnlock *string `json:"GmtUnlock,omitempty" xml:"GmtUnlock,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 36
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The serial number of the virtual MFA device. The serial number is unique for each device.
	//
	// example:
	//
	// dc856334-446b-4035-bfbc-18af261e****
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
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMfaDevicesResponseBodyMfaDevices) String() string {
	return dara.Prettify(s)
}

func (s DescribeMfaDevicesResponseBodyMfaDevices) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) GetConsecutiveFails() *int32 {
	return s.ConsecutiveFails
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) GetEmail() *string {
	return s.Email
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) GetGmtEnabled() *string {
	return s.GmtEnabled
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) GetGmtUnlock() *string {
	return s.GmtUnlock
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) GetId() *int64 {
	return s.Id
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) GetStatus() *string {
	return s.Status
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetConsecutiveFails(v int32) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.ConsecutiveFails = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetDeviceType(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.DeviceType = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetEmail(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.Email = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetEndUserId(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.EndUserId = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetGmtEnabled(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.GmtEnabled = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetGmtUnlock(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.GmtUnlock = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetId(v int64) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.Id = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetSerialNumber(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.SerialNumber = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetStatus(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.Status = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) Validate() error {
	return dara.Validate(s)
}
