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
	// Details about the locked snapshots.
	LockedSnapshotsInfo []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo `json:"LockedSnapshotsInfo,omitempty" xml:"LockedSnapshotsInfo,omitempty" type:"Repeated"`
	// A token to retrieve the next page of results. If this parameter is empty, all results have been returned.
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
	// The cool-off period for compliance mode, in hours.
	//
	// example:
	//
	// 3
	CoolOffPeriod *int32 `json:"CoolOffPeriod,omitempty" xml:"CoolOffPeriod,omitempty"`
	// The time when the cool-off period ends. The time follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard and is displayed in UTC in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2025-10-15T13:00:00Z
	CoolOffPeriodExpiredTime *string `json:"CoolOffPeriodExpiredTime,omitempty" xml:"CoolOffPeriodExpiredTime,omitempty"`
	// The time when the snapshot was locked. The time follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard and is displayed in UTC in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2025-10-15T10:00:00Z
	LockCreationTime *string `json:"LockCreationTime,omitempty" xml:"LockCreationTime,omitempty"`
	// The lock duration in days. The lock automatically expires when this period ends.
	//
	// example:
	//
	// 1
	LockDuration *int32 `json:"LockDuration,omitempty" xml:"LockDuration,omitempty"`
	// The time when the lock duration starts. The time follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard and is displayed in UTC in the yyyy-MM-ddTHH:mm:ssZ format. If a snapshot in the progressing state is locked, its lock duration begins after it enters the accomplished state.
	//
	// example:
	//
	// 2025-10-15T10:00:00Z
	LockDurationStartTime *string `json:"LockDurationStartTime,omitempty" xml:"LockDurationStartTime,omitempty"`
	// The time when the lock expires. The time follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard and is displayed in UTC in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2025-10-16T10:00:00Z
	LockExpiredTime *string `json:"LockExpiredTime,omitempty" xml:"LockExpiredTime,omitempty"`
	// The lock mode. Valid value:
	//
	// - compliance: The snapshot is locked in compliance mode. A snapshot locked in compliance mode cannot be unlocked and can only be deleted after its lock duration expires. You cannot shorten the lock duration, but users with the required RAM permissions can extend the lock duration at any time. When you lock a snapshot in compliance mode, you can optionally specify a cool-off period.
	//
	// example:
	//
	// compliance
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The lock status. Valid values:
	//
	// - compliance-cooloff: The snapshot is locked in compliance mode and is in the cool-off period. The snapshot cannot be deleted. However, users with the required RAM permissions can unlock the snapshot and adjust the cool-off period or lock duration.
	//
	// - compliance: The snapshot is locked in compliance mode, and its cool-off period has ended. The snapshot cannot be unlocked or deleted. However, users with the required RAM permissions can extend the lock duration.
	//
	// - expired: The lock has expired, and the snapshot can be deleted.
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
