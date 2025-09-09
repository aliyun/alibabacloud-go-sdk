// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserStatusResponseBody
	GetRequestId() *string
	SetUserStatus(v *DescribeUserStatusResponseBodyUserStatus) *DescribeUserStatusResponseBody
	GetUserStatus() *DescribeUserStatusResponseBodyUserStatus
}

type DescribeUserStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the current account.
	UserStatus *DescribeUserStatusResponseBodyUserStatus `json:"UserStatus,omitempty" xml:"UserStatus,omitempty" type:"Struct"`
}

func (s DescribeUserStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserStatusResponseBody) GetUserStatus() *DescribeUserStatusResponseBodyUserStatus {
	return s.UserStatus
}

func (s *DescribeUserStatusResponseBody) SetRequestId(v string) *DescribeUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserStatusResponseBody) SetUserStatus(v *DescribeUserStatusResponseBodyUserStatus) *DescribeUserStatusResponseBody {
	s.UserStatus = v
	return s
}

func (s *DescribeUserStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserStatusResponseBodyUserStatus struct {
	// The AccessKey ID of the current account.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// Indicates whether the SQL Explorer feature can be disabled. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	AuditClosable *bool `json:"AuditClosable,omitempty" xml:"AuditClosable,omitempty"`
	// Indicates whether the audit resources can be released.
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	AuditReleasable *bool `json:"AuditReleasable,omitempty" xml:"AuditReleasable,omitempty"`
	// Indicates whether DSC has permission to access user resources within the current account. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Authed *bool `json:"Authed,omitempty" xml:"Authed,omitempty"`
	// The billing method of DCS that is purchased by using the current account. Valid values:
	//
	// 	- **PREPAY**: subscription
	//
	// 	- **POSTPAY**: pay-as-you-go
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The permissions that the current account has. Valid values:
	//
	// 	- **0**: The current account has the administrative permissions or read-only permissions on Data Security Center.
	//
	// 	- **1**: The current account has the permissions to manage data domains.
	//
	// example:
	//
	// 1
	DataManagerRole *int32 `json:"DataManagerRole,omitempty" xml:"DataManagerRole,omitempty"`
	// The ID of the data security center instance purchased by the main account.
	//
	// example:
	//
	// sddp-cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of instances within the current account.
	//
	// example:
	//
	// 32
	InstanceNum *int32 `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 10
	InstanceTotalCount *int64 `json:"InstanceTotalCount,omitempty" xml:"InstanceTotalCount,omitempty"`
	// Indicates whether the data security lab feature is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	LabStatus *int32 `json:"LabStatus,omitempty" xml:"LabStatus,omitempty"`
	// OSS total storage capacity. Unit: Bytes.
	//
	// example:
	//
	// 2048
	OssTotalSize *int64 `json:"OssTotalSize,omitempty" xml:"OssTotalSize,omitempty"`
	// Accumulate the number of days to protect user assets.
	//
	// example:
	//
	// 2
	ProtectionDays *int32 `json:"ProtectionDays,omitempty" xml:"ProtectionDays,omitempty"`
	// Indicates whether DSC is purchased. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Purchased *bool `json:"Purchased,omitempty" xml:"Purchased,omitempty"`
	// The grace period between when DSC is expired and when DSC is released. Unit: days.
	//
	// example:
	//
	// 15
	ReleaseDays *int32 `json:"ReleaseDays,omitempty" xml:"ReleaseDays,omitempty"`
	// The time when the audit resources are released. Unit: milliseconds.
	//
	// example:
	//
	// 15000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The remaining period for which the data assets within the current account can be protected by DSC.
	//
	// example:
	//
	// 131
	RemainDays *int32 `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	// Indicates whether the current account uses a free trial of DSC. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Trail *bool `json:"Trail,omitempty" xml:"Trail,omitempty"`
	// Indicates whether the agent audit feature is used. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	UseAgentAudit *bool `json:"UseAgentAudit,omitempty" xml:"UseAgentAudit,omitempty"`
	// The number of instances that are used.
	//
	// example:
	//
	// 125
	UseInstanceNum *int32 `json:"UseInstanceNum,omitempty" xml:"UseInstanceNum,omitempty"`
	// The occupied space of the Object Storage Service (OSS) bucket. Unit: bytes.
	//
	// example:
	//
	// 234
	UseOssSize *int64 `json:"UseOssSize,omitempty" xml:"UseOssSize,omitempty"`
}

func (s DescribeUserStatusResponseBodyUserStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserStatusResponseBodyUserStatus) GoString() string {
	return s.String()
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetAuditClosable() *bool {
	return s.AuditClosable
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetAuditReleasable() *bool {
	return s.AuditReleasable
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetAuthed() *bool {
	return s.Authed
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetDataManagerRole() *int32 {
	return s.DataManagerRole
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetInstanceNum() *int32 {
	return s.InstanceNum
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetInstanceTotalCount() *int64 {
	return s.InstanceTotalCount
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetLabStatus() *int32 {
	return s.LabStatus
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetOssTotalSize() *int64 {
	return s.OssTotalSize
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetProtectionDays() *int32 {
	return s.ProtectionDays
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetPurchased() *bool {
	return s.Purchased
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetReleaseDays() *int32 {
	return s.ReleaseDays
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetRemainDays() *int32 {
	return s.RemainDays
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetTrail() *bool {
	return s.Trail
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetUseAgentAudit() *bool {
	return s.UseAgentAudit
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetUseInstanceNum() *int32 {
	return s.UseInstanceNum
}

func (s *DescribeUserStatusResponseBodyUserStatus) GetUseOssSize() *int64 {
	return s.UseOssSize
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAccessKeyId(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAuditClosable(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.AuditClosable = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAuditReleasable(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.AuditReleasable = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetAuthed(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.Authed = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetChargeType(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.ChargeType = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetDataManagerRole(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.DataManagerRole = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetInstanceId(v string) *DescribeUserStatusResponseBodyUserStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetInstanceNum(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.InstanceNum = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetInstanceTotalCount(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.InstanceTotalCount = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetLabStatus(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.LabStatus = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetOssTotalSize(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.OssTotalSize = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetProtectionDays(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.ProtectionDays = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetPurchased(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.Purchased = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetReleaseDays(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.ReleaseDays = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetReleaseTime(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetRemainDays(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.RemainDays = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetTrail(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.Trail = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetUseAgentAudit(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.UseAgentAudit = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetUseInstanceNum(v int32) *DescribeUserStatusResponseBodyUserStatus {
	s.UseInstanceNum = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) SetUseOssSize(v int64) *DescribeUserStatusResponseBodyUserStatus {
	s.UseOssSize = &v
	return s
}

func (s *DescribeUserStatusResponseBodyUserStatus) Validate() error {
	return dara.Validate(s)
}
