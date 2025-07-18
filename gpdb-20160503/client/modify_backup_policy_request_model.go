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
	// The number of days for which data backup files are retained. Default value: 7. Maximum value: 7. Valid values: 1 to 7.
	//
	// example:
	//
	// 7
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable automatic point-in-time backup.
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// true
	EnableRecoveryPoint *bool `json:"EnableRecoveryPoint,omitempty" xml:"EnableRecoveryPoint,omitempty"`
	// The cycle based on which backups are performed. If more than one day of the week is specified, the days of the week are separated by commas (,). Valid values:
	//
	// 	- Monday
	//
	// 	- Tuesday
	//
	// 	- Wednesday
	//
	// 	- Thursday
	//
	// 	- Friday
	//
	// 	- Saturday
	//
	// 	- Sunday
	//
	// This parameter is required.
	//
	// example:
	//
	// Tuesday, Thursday, Saturday
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The backup window. Specify the backup window in the HH:mmZ-HH:mmZ format. The backup window must be in UTC. Default value: 00:00-01:00.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15:00Z-16:00Z
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	// The frequency of point-in-time backup.
	//
	// 	- 1: per hour
	//
	// 	- 2: per 2 hours
	//
	// 	- 4: per 4 hours
	//
	// 	- 8: per 8 hours
	//
	// Default value: 8.
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
