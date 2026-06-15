// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLockedSnapshotInfo(v *LockSnapshotResponseBodyLockedSnapshotInfo) *LockSnapshotResponseBody
	GetLockedSnapshotInfo() *LockSnapshotResponseBodyLockedSnapshotInfo
	SetRequestId(v string) *LockSnapshotResponseBody
	GetRequestId() *string
}

type LockSnapshotResponseBody struct {
	// Information about the locked snapshot.
	LockedSnapshotInfo *LockSnapshotResponseBodyLockedSnapshotInfo `json:"LockedSnapshotInfo,omitempty" xml:"LockedSnapshotInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LockSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *LockSnapshotResponseBody) GetLockedSnapshotInfo() *LockSnapshotResponseBodyLockedSnapshotInfo {
	return s.LockedSnapshotInfo
}

func (s *LockSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LockSnapshotResponseBody) SetLockedSnapshotInfo(v *LockSnapshotResponseBodyLockedSnapshotInfo) *LockSnapshotResponseBody {
	s.LockedSnapshotInfo = v
	return s
}

func (s *LockSnapshotResponseBody) SetRequestId(v string) *LockSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockSnapshotResponseBody) Validate() error {
	if s.LockedSnapshotInfo != nil {
		if err := s.LockedSnapshotInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LockSnapshotResponseBodyLockedSnapshotInfo struct {
	// The cool-off period for compliance mode. Unit: hours.
	//
	// example:
	//
	// 3
	CoolOffPeriod *int32 `json:"CoolOffPeriod,omitempty" xml:"CoolOffPeriod,omitempty"`
	// The time the cool-off period for compliance mode ends. The time is in UTC and follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard in `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2025-10-15T13:00:00Z
	CoolOffPeriodExpiredTime *string `json:"CoolOffPeriodExpiredTime,omitempty" xml:"CoolOffPeriodExpiredTime,omitempty"`
	// The time the lock was created. The time is in UTC and follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard in `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2025-10-15T10:00:00Z
	LockCreationTime *string `json:"LockCreationTime,omitempty" xml:"LockCreationTime,omitempty"`
	// The lock duration, in days. The snapshot lock automatically expires at the end of this period.
	//
	// example:
	//
	// 1
	LockDuration *int32 `json:"LockDuration,omitempty" xml:"LockDuration,omitempty"`
	// The time the lock duration starts. The time is in UTC and follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard in `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// If you lock a snapshot that is in the `progressing` state, the lock duration starts only after the snapshot enters the `accomplished` state.
	//
	// example:
	//
	// 2025-10-15T10:00:00Z
	LockDurationStartTime *string `json:"LockDurationStartTime,omitempty" xml:"LockDurationStartTime,omitempty"`
	// The time the lock expires. The time is in UTC and follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard in `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2025-10-16T10:00:00Z
	LockExpiredTime *string `json:"LockExpiredTime,omitempty" xml:"LockExpiredTime,omitempty"`
	// The lock mode. Possible value:
	//
	// - `compliance`: The snapshot is locked in compliance mode. A snapshot in compliance mode cannot be unlocked and can be deleted only after its lock duration expires. You cannot shorten the lock duration, but users with the required Resource Access Management (RAM) permissions can extend it at any time. When you lock a snapshot in compliance mode, you can optionally specify a cool-off period.
	//
	// example:
	//
	// compliance
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The lock status. Possible values:
	//
	// - `compliance-cooloff`: The snapshot is locked in compliance mode but is still in its cool-off period. The snapshot cannot be deleted. However, users with the required Resource Access Management (RAM) permissions can unlock it, change the cool-off period, and adjust the lock duration.
	//
	// - `compliance`: The snapshot is locked in compliance mode, and the cool-off period has ended. The snapshot cannot be unlocked or deleted, but users with the required Resource Access Management (RAM) permissions can extend the lock duration.
	//
	// - `expired`: The snapshot was previously locked, but the lock duration has ended, and the lock has expired. The snapshot is not currently locked and can be deleted.
	//
	// example:
	//
	// compliance-cooloff
	LockStatus *string `json:"LockStatus,omitempty" xml:"LockStatus,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-9dp2qojdpdfmgfmf****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s LockSnapshotResponseBodyLockedSnapshotInfo) String() string {
	return dara.Prettify(s)
}

func (s LockSnapshotResponseBodyLockedSnapshotInfo) GoString() string {
	return s.String()
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) GetCoolOffPeriod() *int32 {
	return s.CoolOffPeriod
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) GetCoolOffPeriodExpiredTime() *string {
	return s.CoolOffPeriodExpiredTime
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) GetLockCreationTime() *string {
	return s.LockCreationTime
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) GetLockDuration() *int32 {
	return s.LockDuration
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) GetLockDurationStartTime() *string {
	return s.LockDurationStartTime
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) GetLockExpiredTime() *string {
	return s.LockExpiredTime
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) GetLockMode() *string {
	return s.LockMode
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) GetLockStatus() *string {
	return s.LockStatus
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) SetCoolOffPeriod(v int32) *LockSnapshotResponseBodyLockedSnapshotInfo {
	s.CoolOffPeriod = &v
	return s
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) SetCoolOffPeriodExpiredTime(v string) *LockSnapshotResponseBodyLockedSnapshotInfo {
	s.CoolOffPeriodExpiredTime = &v
	return s
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) SetLockCreationTime(v string) *LockSnapshotResponseBodyLockedSnapshotInfo {
	s.LockCreationTime = &v
	return s
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) SetLockDuration(v int32) *LockSnapshotResponseBodyLockedSnapshotInfo {
	s.LockDuration = &v
	return s
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) SetLockDurationStartTime(v string) *LockSnapshotResponseBodyLockedSnapshotInfo {
	s.LockDurationStartTime = &v
	return s
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) SetLockExpiredTime(v string) *LockSnapshotResponseBodyLockedSnapshotInfo {
	s.LockExpiredTime = &v
	return s
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) SetLockMode(v string) *LockSnapshotResponseBodyLockedSnapshotInfo {
	s.LockMode = &v
	return s
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) SetLockStatus(v string) *LockSnapshotResponseBodyLockedSnapshotInfo {
	s.LockStatus = &v
	return s
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) SetSnapshotId(v string) *LockSnapshotResponseBodyLockedSnapshotInfo {
	s.SnapshotId = &v
	return s
}

func (s *LockSnapshotResponseBodyLockedSnapshotInfo) Validate() error {
	return dara.Validate(s)
}
