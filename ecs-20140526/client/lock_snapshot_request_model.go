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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/en/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// example:
	//
	// 5EC38E7D-389F-1925-ABE2-D7925A8F****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cool-off period. In compliance mode, you can set a cool-off period or skip the cool-off period to directly lock the snapshot.
	//
	// During the cool-off period, users with the required RAM permissions can unlock the snapshot, extend or shorten the cool-off period, and extend or shorten the lock duration. The snapshot cannot be deleted during the cool-off period.
	//
	// After the cool-off period ends, you can only extend the lock duration.
	//
	// Unit: hours.
	//
	// Valid values: 0 to 72. A value of 0 indicates that the cool-off period is skipped and the snapshot is directly locked.
	//
	// If the snapshot has already entered the compliance mode lock period, set this parameter to 0 to extend the lock duration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	CoolOffPeriod *int32 `json:"CoolOffPeriod,omitempty" xml:"CoolOffPeriod,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: performs only a dry run. The system checks whether required parameters are specified, whether the request format is valid, and whether business restrictions are met. If the check fails, the corresponding error is returned. If the check succeeds, the DryRunOperation error code is returned.
	//
	// - false (default): performs a dry run and sends the request. If the check succeeds, the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The lock duration. The snapshot lock automatically expires after the lock duration ends.
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
	// The lock mode. Valid values:
	//
	// - compliance: Locks the snapshot in compliance mode. A snapshot locked in compliance mode cannot be unlocked by any user and can be deleted only after the lock duration expires. Users cannot shorten the lock duration, but users with the required RAM permissions can extend the lock duration at any time. When locking a snapshot in compliance mode, you can optionally specify a cool-off period.
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
	// The region ID. You can call [DescribeRegions](https://www.alibabacloud.com/help/en/ecs/developer-reference/api-ecs-2014-05-26-describeregions) to query the most recent list of Alibaba Cloud regions.
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
