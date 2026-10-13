// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyId(v string) *ModifyAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyId() *string
	SetAutoSnapshotPolicyName(v string) *ModifyAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyName() *string
	SetRepeatWeekdays(v string) *ModifyAutoSnapshotPolicyRequest
	GetRepeatWeekdays() *string
	SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyRequest
	GetRetentionDays() *int32
	SetTimePoints(v string) *ModifyAutoSnapshotPolicyRequest
	GetTimePoints() *string
}

type ModifyAutoSnapshotPolicyRequest struct {
	// The ID of the automatic snapshot policy.
	//
	// You can call DescribeAutoSnapshotPolicies to query available automatic snapshot policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-extreme-233e6****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The name of the automatic snapshot policy. If this parameter is left empty, the policy name is not modified.
	//
	// Limits:
	//
	// - The name must be 2 to 128 characters in length.
	//
	// - The name must start with a letter.
	//
	// - The name can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	// The days of the week on which automatic snapshots are created.
	//
	// Cycle: week.
	//
	// Valid values: 1 to 7. For example, 1 indicates Monday. To create multiple automatic snapshots within a week, specify multiple days and separate them with commas (,). You can specify up to 7 days.
	//
	// example:
	//
	// 1,7
	RepeatWeekdays *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	// The retention period of automatic snapshots.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// - -1 (default): Automatic snapshots are permanently retained.
	//
	// - 1 to 65536: Automatic snapshots are retained for the specified number of days.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The points in time at which automatic snapshots are created.
	//
	// Unit: hours.
	//
	// Valid values: 0 to 23, which correspond to the 24 points in time from 00:00 to 23:00. For example, 1 indicates 01:00. To create multiple automatic snapshots within a day, specify multiple points in time and separate them with commas (,). You can specify up to 24 points in time.
	//
	// example:
	//
	// 0,1
	TimePoints *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s ModifyAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *ModifyAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyName() *string {
	return s.AutoSnapshotPolicyName
}

func (s *ModifyAutoSnapshotPolicyRequest) GetRepeatWeekdays() *string {
	return s.RepeatWeekdays
}

func (s *ModifyAutoSnapshotPolicyRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *ModifyAutoSnapshotPolicyRequest) GetTimePoints() *string {
	return s.TimePoints
}

func (s *ModifyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *ModifyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyName(v string) *ModifyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRepeatWeekdays(v string) *ModifyAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetTimePoints(v string) *ModifyAutoSnapshotPolicyRequest {
	s.TimePoints = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
