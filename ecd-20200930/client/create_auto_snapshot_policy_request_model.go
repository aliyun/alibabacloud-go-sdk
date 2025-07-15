// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *CreateAutoSnapshotPolicyRequest
	GetCronExpression() *string
	SetPolicyName(v string) *CreateAutoSnapshotPolicyRequest
	GetPolicyName() *string
	SetRegionId(v string) *CreateAutoSnapshotPolicyRequest
	GetRegionId() *string
	SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest
	GetRetentionDays() *int32
}

type CreateAutoSnapshotPolicyRequest struct {
	// The CRON expression for periodic scheduling.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 0 2 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The name of the automatic snapshot policy. The name must be 2 to 128 characters in length. The name must start with a letter but cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-). This parameter is empty by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_auto_policy
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
	// This parameter is required.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateAutoSnapshotPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateAutoSnapshotPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAutoSnapshotPolicyRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *CreateAutoSnapshotPolicyRequest) SetCronExpression(v string) *CreateAutoSnapshotPolicyRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetPolicyName(v string) *CreateAutoSnapshotPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRegionId(v string) *CreateAutoSnapshotPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
