// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *ModifyAutoSnapshotPolicyRequest
	GetCronExpression() *string
	SetDiskType(v string) *ModifyAutoSnapshotPolicyRequest
	GetDiskType() *string
	SetPolicyId(v string) *ModifyAutoSnapshotPolicyRequest
	GetPolicyId() *string
	SetPolicyName(v string) *ModifyAutoSnapshotPolicyRequest
	GetPolicyName() *string
	SetRegionId(v string) *ModifyAutoSnapshotPolicyRequest
	GetRegionId() *string
	SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyRequest
	GetRetentionDays() *int32
}

type ModifyAutoSnapshotPolicyRequest struct {
	// The CRON expression.
	//
	// example:
	//
	// 0 20 16 ? 	- 1,2,3,4,5,6,7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	DiskType       *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the automatic snapshot policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-itcmrhqt01tdo****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the automatic snapshot policy. The name must be 2 to 128 characters in length. The name must start with a letter but cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-). This parameter is empty by default.
	//
	// example:
	//
	// Automatic system snapshot
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The retention period of the automatic snapshots. Unit: days. Valid values: 1 to 180.
	//
	// example:
	//
	// 5
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
}

func (s ModifyAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyAutoSnapshotPolicyRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *ModifyAutoSnapshotPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ModifyAutoSnapshotPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ModifyAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAutoSnapshotPolicyRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *ModifyAutoSnapshotPolicyRequest) SetCronExpression(v string) *ModifyAutoSnapshotPolicyRequest {
	s.CronExpression = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetDiskType(v string) *ModifyAutoSnapshotPolicyRequest {
	s.DiskType = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetPolicyId(v string) *ModifyAutoSnapshotPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetPolicyName(v string) *ModifyAutoSnapshotPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRegionId(v string) *ModifyAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
