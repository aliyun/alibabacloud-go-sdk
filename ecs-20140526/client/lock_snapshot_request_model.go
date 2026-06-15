// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *LockSnapshotRequest
	GetClientToken() *string
	SetCoolOffPeriod(v int32) *LockSnapshotRequest
	GetCoolOffPeriod() *int32
	SetDryRun(v bool) *LockSnapshotRequest
	GetDryRun() *bool
	SetLockDuration(v int32) *LockSnapshotRequest
	GetLockDuration() *int32
	SetLockMode(v string) *LockSnapshotRequest
	GetLockMode() *string
	SetOwnerAccount(v string) *LockSnapshotRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *LockSnapshotRequest
	GetOwnerId() *int64
	SetRegionId(v string) *LockSnapshotRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *LockSnapshotRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *LockSnapshotRequest
	GetResourceOwnerId() *int64
	SetSnapshotId(v string) *LockSnapshotRequest
	GetSnapshotId() *string
}

type LockSnapshotRequest struct {
	// A unique, case-sensitive token that you provide to ensure the idempotence of the request. The token can contain only ASCII characters and must not exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/zh/ecs/developer-reference/how-to-ensure-idempotence?spm=a2c4g.11186623.0.0.2a29d467Bh2sO5).
	//
	// example:
	//
	// 5EC38E7D-389F-1925-ABE2-D7925A8F****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cool-off period. In compliance mode, you can specify a cool-off period or set this parameter to 0 to lock the snapshot immediately.
	//
	// During the cool-off period, users with the required RAM permissions can unlock the snapshot, extend or shorten the cool-off period, and extend or shorten the lock duration. The snapshot cannot be deleted during the cool-off period.
	//
	// After the cool-off period ends, you can only extend the lock duration.
	//
	// Unit: hours.
	//
	// Valid values: 0 to 72. A value of 0 indicates that the cool-off period is skipped and the snapshot is locked immediately.
	//
	// If a snapshot is already locked in compliance mode, you must set this parameter to 0 to extend its lock duration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	CoolOffPeriod *int32 `json:"CoolOffPeriod,omitempty" xml:"CoolOffPeriod,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - `true`: Performs a dry run to check the request without executing it. The system checks for required parameters, request format, business constraints, and permissions. If the check passes, the `DryRunOperation` error code is returned. If the check fails, a corresponding error is returned.
	//
	// - `false` (default): Checks the request and, if the checks pass, executes the operation.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The number of days to lock the snapshot. The lock automatically expires at the end of this period.
	//
	// Unit: days.
	//
	// Valid values: 1 to 36500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	LockDuration *int32 `json:"LockDuration,omitempty" xml:"LockDuration,omitempty"`
	// The lock mode. Valid value:
	//
	// - `compliance`: Locks the snapshot in compliance mode. A snapshot locked in compliance mode cannot be unlocked by any user and can be deleted only after its lock duration expires. You cannot shorten the lock duration. However, users with the required RAM permissions can extend the lock duration at any time. When you lock a snapshot in compliance mode, you can optionally specify a cool-off period.
	//
	// This parameter is required.
	//
	// example:
	//
	// compliance
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// 158704318252****
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 158704318252****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/zh/ecs/developer-reference/api-ecs-2014-05-26-describeregions?spm=a2c4g.11186623.0.i2) to get the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 158704318252****
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// example:
	//
	// 158704318252****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-9dp2qojdpdfmgfmf****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s LockSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s LockSnapshotRequest) GoString() string {
	return s.String()
}

func (s *LockSnapshotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *LockSnapshotRequest) GetCoolOffPeriod() *int32 {
	return s.CoolOffPeriod
}

func (s *LockSnapshotRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *LockSnapshotRequest) GetLockDuration() *int32 {
	return s.LockDuration
}

func (s *LockSnapshotRequest) GetLockMode() *string {
	return s.LockMode
}

func (s *LockSnapshotRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *LockSnapshotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *LockSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *LockSnapshotRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *LockSnapshotRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *LockSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *LockSnapshotRequest) SetClientToken(v string) *LockSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *LockSnapshotRequest) SetCoolOffPeriod(v int32) *LockSnapshotRequest {
	s.CoolOffPeriod = &v
	return s
}

func (s *LockSnapshotRequest) SetDryRun(v bool) *LockSnapshotRequest {
	s.DryRun = &v
	return s
}

func (s *LockSnapshotRequest) SetLockDuration(v int32) *LockSnapshotRequest {
	s.LockDuration = &v
	return s
}

func (s *LockSnapshotRequest) SetLockMode(v string) *LockSnapshotRequest {
	s.LockMode = &v
	return s
}

func (s *LockSnapshotRequest) SetOwnerAccount(v string) *LockSnapshotRequest {
	s.OwnerAccount = &v
	return s
}

func (s *LockSnapshotRequest) SetOwnerId(v int64) *LockSnapshotRequest {
	s.OwnerId = &v
	return s
}

func (s *LockSnapshotRequest) SetRegionId(v string) *LockSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *LockSnapshotRequest) SetResourceOwnerAccount(v string) *LockSnapshotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LockSnapshotRequest) SetResourceOwnerId(v int64) *LockSnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LockSnapshotRequest) SetSnapshotId(v string) *LockSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *LockSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
