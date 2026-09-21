// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUniBackupPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateUniBackupPolicyShrinkRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateUniBackupPolicyShrinkRequest
	GetAccountPassword() *string
	SetDatabaseAddByUser(v string) *CreateUniBackupPolicyShrinkRequest
	GetDatabaseAddByUser() *string
	SetDatabaseType(v string) *CreateUniBackupPolicyShrinkRequest
	GetDatabaseType() *string
	SetFullPlanShrink(v string) *CreateUniBackupPolicyShrinkRequest
	GetFullPlanShrink() *string
	SetIncPlanShrink(v string) *CreateUniBackupPolicyShrinkRequest
	GetIncPlanShrink() *string
	SetInstanceId(v string) *CreateUniBackupPolicyShrinkRequest
	GetInstanceId() *string
	SetPolicyName(v string) *CreateUniBackupPolicyShrinkRequest
	GetPolicyName() *string
	SetRetention(v int32) *CreateUniBackupPolicyShrinkRequest
	GetRetention() *int32
	SetSpeedLimiter(v int64) *CreateUniBackupPolicyShrinkRequest
	GetSpeedLimiter() *int64
	SetUniRegionId(v string) *CreateUniBackupPolicyShrinkRequest
	GetUniRegionId() *string
	SetUuid(v string) *CreateUniBackupPolicyShrinkRequest
	GetUuid() *string
}

type CreateUniBackupPolicyShrinkRequest struct {
	// The username of the database account.
	//
	// example:
	//
	// admin
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// example:
	//
	// Pass****
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// Specifies whether the database is manually added by the user. Valid values:
	//
	// - **true**: The database is manually added.
	//
	// - **false**: The database is not manually added.
	//
	// example:
	//
	// true
	DatabaseAddByUser *string `json:"DatabaseAddByUser,omitempty" xml:"DatabaseAddByUser,omitempty"`
	// The type of the database. Valid values:
	//
	// - **MYSQL**
	//
	// - **ORACLE**
	//
	// - **MSSQL**
	//
	// This parameter is required.
	//
	// example:
	//
	// MYSQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The full backup policy. The value is in JSON format and contains the following fields:
	//
	// - **start**: the start time of the backup.
	//
	// - **interval**: the interval between backups.
	//
	// - **type**: the unit of the interval.
	//
	// - **days**: the days of the week on which backups are performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"days":[4],"interval":1,"planType":"weekly","startTime":"22:00:00"}
	FullPlanShrink *string `json:"FullPlan,omitempty" xml:"FullPlan,omitempty"`
	// The incremental backup policy. The value is in JSON format and contains the following fields:
	//
	// - **start**: the start time of the backup.
	//
	// - **interval**: the interval between backups.
	//
	// - **type**: the unit of the interval.
	//
	// - **days**: the days of the week on which backups are performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"interval":1,"planType":"daily","startTime":"23:30:00"}
	IncPlanShrink *string `json:"IncPlan,omitempty" xml:"IncPlan,omitempty"`
	// The ID of the ECS instance.
	//
	// >You can call the [DescribeUniBackupDatabase](~~DescribeUniBackupDatabase~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1fu4aqltf1huhc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the anti-ransomware backup policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql-policy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The number of days for which backup data is retained.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The network bandwidth throttling for backup network bandwidth. Unit: bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5242880
	SpeedLimiter *int64 `json:"SpeedLimiter,omitempty" xml:"SpeedLimiter,omitempty"`
	// The region in which the server protected by the backup policy resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	UniRegionId *string `json:"UniRegionId,omitempty" xml:"UniRegionId,omitempty"`
	// The UUID of the server that is backed up by the database anti-ransomware feature.
	//
	// > You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to obtain the UUID of the server.
	//
	// example:
	//
	// 045cad48-eb08-4047-a70c-713aec7b****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateUniBackupPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUniBackupPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUniBackupPolicyShrinkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateUniBackupPolicyShrinkRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateUniBackupPolicyShrinkRequest) GetDatabaseAddByUser() *string {
	return s.DatabaseAddByUser
}

func (s *CreateUniBackupPolicyShrinkRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *CreateUniBackupPolicyShrinkRequest) GetFullPlanShrink() *string {
	return s.FullPlanShrink
}

func (s *CreateUniBackupPolicyShrinkRequest) GetIncPlanShrink() *string {
	return s.IncPlanShrink
}

func (s *CreateUniBackupPolicyShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUniBackupPolicyShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateUniBackupPolicyShrinkRequest) GetRetention() *int32 {
	return s.Retention
}

func (s *CreateUniBackupPolicyShrinkRequest) GetSpeedLimiter() *int64 {
	return s.SpeedLimiter
}

func (s *CreateUniBackupPolicyShrinkRequest) GetUniRegionId() *string {
	return s.UniRegionId
}

func (s *CreateUniBackupPolicyShrinkRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CreateUniBackupPolicyShrinkRequest) SetAccountName(v string) *CreateUniBackupPolicyShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetAccountPassword(v string) *CreateUniBackupPolicyShrinkRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetDatabaseAddByUser(v string) *CreateUniBackupPolicyShrinkRequest {
	s.DatabaseAddByUser = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetDatabaseType(v string) *CreateUniBackupPolicyShrinkRequest {
	s.DatabaseType = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetFullPlanShrink(v string) *CreateUniBackupPolicyShrinkRequest {
	s.FullPlanShrink = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetIncPlanShrink(v string) *CreateUniBackupPolicyShrinkRequest {
	s.IncPlanShrink = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetInstanceId(v string) *CreateUniBackupPolicyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetPolicyName(v string) *CreateUniBackupPolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetRetention(v int32) *CreateUniBackupPolicyShrinkRequest {
	s.Retention = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetSpeedLimiter(v int64) *CreateUniBackupPolicyShrinkRequest {
	s.SpeedLimiter = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetUniRegionId(v string) *CreateUniBackupPolicyShrinkRequest {
	s.UniRegionId = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) SetUuid(v string) *CreateUniBackupPolicyShrinkRequest {
	s.Uuid = &v
	return s
}

func (s *CreateUniBackupPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
