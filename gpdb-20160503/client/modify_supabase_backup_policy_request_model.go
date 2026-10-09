// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRetentionPeriod(v int32) *ModifySupabaseBackupPolicyRequest
	GetBackupRetentionPeriod() *int32
	SetEnableRecoveryPoint(v bool) *ModifySupabaseBackupPolicyRequest
	GetEnableRecoveryPoint() *bool
	SetPreferredBackupPeriod(v string) *ModifySupabaseBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *ModifySupabaseBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetProjectId(v string) *ModifySupabaseBackupPolicyRequest
	GetProjectId() *string
	SetRecoveryPointPeriod(v string) *ModifySupabaseBackupPolicyRequest
	GetRecoveryPointPeriod() *string
	SetRegionId(v string) *ModifySupabaseBackupPolicyRequest
	GetRegionId() *string
}

type ModifySupabaseBackupPolicyRequest struct {
	// The data backup retention period. Unit: days. Valid values: 1 to 7.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// Specifies whether to enable automatic recovery points. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// If this parameter is not specified, false is used.
	EnableRecoveryPoint *bool `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
	// The data backup cycle. Separate multiple values with commas (,). Valid values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, and Sunday.
	//
	// This parameter is required.
	//
	// example:
	//
	// Wednesday,Friday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The start time of the data backup. The time is in UTC and follows the HH:mmZ format, such as 01:00Z. The HH:mmZ-HH:mmZ time range format is also supported, and the server uses the start time of the range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// Instance ID of the Supabase instance. You can obtain instance ID on the Supabase page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbp-263****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The interval for the automatic creation of recovery points. Unit: hours. Valid values: 1/6 (10 minutes), 1/2 (30 minutes), 1, 2, 4, and 8. This parameter takes effect only when EnableRecoveryPoint is set to true. If this parameter is not specified, the default value 1 is used. If EnableRecoveryPoint is set to false, this parameter is ignored.
	//
	// example:
	//
	// 1
	RecoveryPointPeriod *string `json:"RecoveryPointPeriod,omitempty" xml:"RecoveryPointPeriod,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifySupabaseBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifySupabaseBackupPolicyRequest) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *ModifySupabaseBackupPolicyRequest) GetEnableRecoveryPoint() *bool {
	return s.EnableRecoveryPoint
}

func (s *ModifySupabaseBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifySupabaseBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifySupabaseBackupPolicyRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifySupabaseBackupPolicyRequest) GetRecoveryPointPeriod() *string {
	return s.RecoveryPointPeriod
}

func (s *ModifySupabaseBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySupabaseBackupPolicyRequest) SetBackupRetentionPeriod(v int32) *ModifySupabaseBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifySupabaseBackupPolicyRequest) SetEnableRecoveryPoint(v bool) *ModifySupabaseBackupPolicyRequest {
	s.EnableRecoveryPoint = &v
	return s
}

func (s *ModifySupabaseBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifySupabaseBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifySupabaseBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifySupabaseBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifySupabaseBackupPolicyRequest) SetProjectId(v string) *ModifySupabaseBackupPolicyRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifySupabaseBackupPolicyRequest) SetRecoveryPointPeriod(v string) *ModifySupabaseBackupPolicyRequest {
	s.RecoveryPointPeriod = &v
	return s
}

func (s *ModifySupabaseBackupPolicyRequest) SetRegionId(v string) *ModifySupabaseBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySupabaseBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
