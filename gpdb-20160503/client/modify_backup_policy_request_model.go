// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRetentionPeriod(v int32) *ModifyBackupPolicyRequest
	GetBackupRetentionPeriod() *int32
	SetDBInstanceId(v string) *ModifyBackupPolicyRequest
	GetDBInstanceId() *string
	SetEnableRecoveryPoint(v bool) *ModifyBackupPolicyRequest
	GetEnableRecoveryPoint() *bool
	SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest
	GetPreferredBackupTime() *string
	SetRecoveryPointPeriod(v string) *ModifyBackupPolicyRequest
	GetRecoveryPointPeriod() *string
}

type ModifyBackupPolicyRequest struct {
	// The number of days for which data backups are retained. Default value: 7. Maximum value: 7. Valid values: 1 to 7.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specified region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable automatic recovery point creation. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	EnableRecoveryPoint *bool `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
	// The data backup cycle. Separate multiple values with commas (,). Valid values:
	//
	// - Monday: Monday.
	//
	// - Tuesday: Tuesday.
	//
	// - Wednesday: Wednesday.
	//
	// - Thursday: Thursday.
	//
	// - Friday: Friday.
	//
	// - Saturday: Saturday.
	//
	// - Sunday: Sunday.
	//
	// This parameter is required.
	//
	// example:
	//
	// Tuesday,Thursday,Saturday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The data backup time. Default value: 00:00Z-01:00Z. Format: HH:mmZ-HH:mmZ (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 15:00Z-16:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The frequency of recovery points. Valid values:
	//
	// - **1**: Every hour.
	//
	// - **2**: Every two hours.
	//
	// - **4**: Every four hours.
	//
	// - **8**: Every eight hours.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 8
	RecoveryPointPeriod *string `json:"RecoveryPointPeriod,omitempty" xml:"RecoveryPointPeriod,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *ModifyBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyBackupPolicyRequest) GetEnableRecoveryPoint() *bool {
	return s.EnableRecoveryPoint
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *ModifyBackupPolicyRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *ModifyBackupPolicyRequest) GetRecoveryPointPeriod() *string {
	return s.RecoveryPointPeriod
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPeriod(v int32) *ModifyBackupPolicyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBInstanceId(v string) *ModifyBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetEnableRecoveryPoint(v bool) *ModifyBackupPolicyRequest {
	s.EnableRecoveryPoint = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetRecoveryPointPeriod(v string) *ModifyBackupPolicyRequest {
	s.RecoveryPointPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
