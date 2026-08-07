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
	// The request ID.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the current logon account.
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
	if s.UserStatus != nil {
		if err := s.UserStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserStatusResponseBodyUserStatus struct {
	// The AccessKey ID authorized by the user.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// Indicates whether the current logon account has authorized the service-linked role for DSC asset synchronization. Valid values:
	//
	// - **true**: authorized.
	//
	// - **false**: not authorized.
	//
	// example:
	//
	// true
	AssetRoleAuthed *bool `json:"AssetRoleAuthed,omitempty" xml:"AssetRoleAuthed,omitempty"`
	// Indicates whether SQL Explorer can be disabled. Valid values:
	//
	// - **true**: can be disabled.
	//
	// - **false**: cannot be disabled.
	//
	// example:
	//
	// true
	AuditClosable *bool `json:"AuditClosable,omitempty" xml:"AuditClosable,omitempty"`
	// Indicates whether SQL Explorer can be released. Valid values:
	//
	// - **true**: can be released.
	//
	// - **false**: cannot be released.
	//
	// example:
	//
	// true
	AuditReleasable *bool `json:"AuditReleasable,omitempty" xml:"AuditReleasable,omitempty"`
	// Indicates whether the current logon account has authorized DSC to access RAM. Valid values:
	//
	// - **true**: authorized.
	//
	// - **false**: not authorized.
	//
	// example:
	//
	// true
	Authed *bool `json:"Authed,omitempty" xml:"Authed,omitempty"`
	// The billing method of the DSC service purchased by the current logon account. Valid values:
	//
	// - **PREPAY**: subscription.
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The permissions of the current account. Valid values:
	//
	// - **0**: has management or read-only permissions for Data Security Center.
	//
	// - **1**: has data domain management permissions.
	//
	// example:
	//
	// 1
	DataManagerRole *int32 `json:"DataManagerRole,omitempty" xml:"DataManagerRole,omitempty"`
	// The instance ID of the Data Security Center product purchased by the Alibaba Cloud account.
	//
	// example:
	//
	// sddp-cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of instances in the current logon account.
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
	// Indicates whether the asset lab feature is enabled. Valid values:
	//
	// - **1**: enabled.
	//
	// - **0**: not enabled.
	//
	// example:
	//
	// 1
	LabStatus *int32 `json:"LabStatus,omitempty" xml:"LabStatus,omitempty"`
	// The total OSS storage capacity. Unit: bytes.
	//
	// example:
	//
	// 2048
	OssTotalSize *int64 `json:"OssTotalSize,omitempty" xml:"OssTotalSize,omitempty"`
	// The total number of days that user assets have been protected.
	//
	// example:
	//
	// 22
	ProtectionDays *int32 `json:"ProtectionDays,omitempty" xml:"ProtectionDays,omitempty"`
	// Indicates whether the DSC service is purchased. Valid values:
	//
	// - **true**: purchased.
	//
	// - **false**: not purchased.
	//
	// example:
	//
	// true
	Purchased *bool `json:"Purchased,omitempty" xml:"Purchased,omitempty"`
	// The number of days from expiration to release. Unit: days.
	//
	// example:
	//
	// 15
	ReleaseDays *int32 `json:"ReleaseDays,omitempty" xml:"ReleaseDays,omitempty"`
	// The release time. Unit: milliseconds.
	//
	// example:
	//
	// 15000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The remaining days of the protection period for assets in the current logon account.
	//
	// example:
	//
	// 131
	RemainDays *int32 `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	// Indicates whether the current logon account is using a trial version of DSC. Valid values:
	//
	// - **true**: trial.
	//
	// - **false**: non-trial.
	//
	// example:
	//
	// true
	Trail *bool `json:"Trail,omitempty" xml:"Trail,omitempty"`
	// Indicates whether the Agent audit feature has been used. Valid values:
	//
	// - **1**: has been used.
	//
	// - **0**: has not been used.
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
	// The used OSS storage capacity. Unit: bytes.
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

func (s *DescribeUserStatusResponseBodyUserStatus) GetAssetRoleAuthed() *bool {
	return s.AssetRoleAuthed
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

func (s *DescribeUserStatusResponseBodyUserStatus) SetAssetRoleAuthed(v bool) *DescribeUserStatusResponseBodyUserStatus {
	s.AssetRoleAuthed = &v
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
