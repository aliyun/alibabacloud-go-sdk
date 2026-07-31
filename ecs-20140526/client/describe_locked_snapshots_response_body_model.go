// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLockedSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLockedSnapshotsInfo(v []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) *DescribeLockedSnapshotsResponseBody
	GetLockedSnapshotsInfo() []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo
	SetNextToken(v string) *DescribeLockedSnapshotsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeLockedSnapshotsResponseBody
	GetRequestId() *string
}

type DescribeLockedSnapshotsResponseBody struct {
	// The collection of locked snapshot information.
	LockedSnapshotsInfo []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo `json:"LockedSnapshotsInfo,omitempty" xml:"LockedSnapshotsInfo,omitempty" type:"Repeated"`
	// The pagination token returned in this call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLockedSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLockedSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLockedSnapshotsResponseBody) GetLockedSnapshotsInfo() []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	return s.LockedSnapshotsInfo
}

func (s *DescribeLockedSnapshotsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLockedSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLockedSnapshotsResponseBody) SetLockedSnapshotsInfo(v []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) *DescribeLockedSnapshotsResponseBody {
	s.LockedSnapshotsInfo = v
	return s
}

func (s *DescribeLockedSnapshotsResponseBody) SetNextToken(v string) *DescribeLockedSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBody) SetRequestId(v string) *DescribeLockedSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBody) Validate() error {
	if s.LockedSnapshotsInfo != nil {
		for _, item := range s.LockedSnapshotsInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo struct {
	// The cooling-off period for compliance mode. Unit: hours.
	//
	// example:
	//
	// 3
	CoolOffPeriod *int32 `json:"CoolOffPeriod,omitempty" xml:"CoolOffPeriod,omitempty"`
	// The time when the cooling-off period for compliance mode ends. The time follows the [ISO 8601](https://www.alibabacloud.com/help/en/ecs/developer-reference/iso-8601-time-format) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-10-15T13:00:00Z
	CoolOffPeriodExpiredTime *string `json:"CoolOffPeriodExpiredTime,omitempty" xml:"CoolOffPeriodExpiredTime,omitempty"`
	// The time when the snapshot was locked. The time follows the [ISO 8601](https://www.alibabacloud.com/help/en/ecs/developer-reference/iso-8601-time-format) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-10-15T10:00:00Z
	LockCreationTime *string `json:"LockCreationTime,omitempty" xml:"LockCreationTime,omitempty"`
	// The lock duration. The snapshot lock automatically expires after the lock duration ends. Unit: days.
	//
	// example:
	//
	// 1
	LockDuration *int32 `json:"LockDuration,omitempty" xml:"LockDuration,omitempty"`
	// The start time of the lock duration. The time follows the [ISO 8601](https://www.alibabacloud.com/help/en/ecs/developer-reference/iso-8601-time-format) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC. If a snapshot in the progressing state is locked, the lock duration starts only after the snapshot enters the accomplished state.
	//
	// example:
	//
	// 2025-10-15T10:00:00Z
	LockDurationStartTime *string `json:"LockDurationStartTime,omitempty" xml:"LockDurationStartTime,omitempty"`
	// The time when the lock expires. The time follows the [ISO 8601](https://www.alibabacloud.com/help/en/ecs/developer-reference/iso-8601-time-format) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-10-16T10:00:00Z
	LockExpiredTime *string `json:"LockExpiredTime,omitempty" xml:"LockExpiredTime,omitempty"`
	// The lock mode. Valid values:
	//
	// - compliance: The snapshot is locked in compliance mode. A snapshot locked in compliance mode cannot be unlocked by any user and can be deleted only after the lock duration expires. Users cannot shorten the lock duration, but users with the required RAM permissions can extend the lock duration at any time. When locking a snapshot in compliance mode, you can optionally specify a cooling-off period.
	//
	// example:
	//
	// compliance
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The lock status. Valid values:
	//
	// - compliance-cooloff: The snapshot is locked in compliance mode but is still within the cooling-off period. The snapshot cannot be deleted, but users with the required RAM permissions can unlock the snapshot, extend or shorten the cooling-off period, or extend or shorten the lock duration.
	//
	// - compliance: The snapshot is locked in compliance mode and the cooling-off period has ended. The snapshot cannot be unlocked or deleted, but users with the required RAM permissions can extend the lock duration.
	//
	// - expired: The snapshot was previously locked, but the lock duration has ended and the lock has expired. The snapshot is currently unlocked and can be deleted.
	//
	// example:
	//
	// compliance-cooloff
	LockStatus *string `json:"LockStatus,omitempty" xml:"LockStatus,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GoString() string {
	return s.String()
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetCoolOffPeriod() *int32 {
	return s.CoolOffPeriod
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetCoolOffPeriodExpiredTime() *string {
	return s.CoolOffPeriodExpiredTime
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockCreationTime() *string {
	return s.LockCreationTime
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockDuration() *int32 {
	return s.LockDuration
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockDurationStartTime() *string {
	return s.LockDurationStartTime
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockExpiredTime() *string {
	return s.LockExpiredTime
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockStatus() *string {
	return s.LockStatus
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetCoolOffPeriod(v int32) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.CoolOffPeriod = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetCoolOffPeriodExpiredTime(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.CoolOffPeriodExpiredTime = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockCreationTime(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockCreationTime = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockDuration(v int32) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockDuration = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockDurationStartTime(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockDurationStartTime = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockExpiredTime(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockExpiredTime = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockMode(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockMode = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockStatus(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockStatus = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetSnapshotId(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.SnapshotId = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) Validate() error {
	return dara.Validate(s)
}
