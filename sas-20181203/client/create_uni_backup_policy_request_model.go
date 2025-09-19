// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUniBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateUniBackupPolicyRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateUniBackupPolicyRequest
	GetAccountPassword() *string
	SetDatabaseAddByUser(v string) *CreateUniBackupPolicyRequest
	GetDatabaseAddByUser() *string
	SetDatabaseType(v string) *CreateUniBackupPolicyRequest
	GetDatabaseType() *string
	SetFullPlan(v map[string]interface{}) *CreateUniBackupPolicyRequest
	GetFullPlan() map[string]interface{}
	SetIncPlan(v map[string]interface{}) *CreateUniBackupPolicyRequest
	GetIncPlan() map[string]interface{}
	SetInstanceId(v string) *CreateUniBackupPolicyRequest
	GetInstanceId() *string
	SetPolicyName(v string) *CreateUniBackupPolicyRequest
	GetPolicyName() *string
	SetRetention(v int32) *CreateUniBackupPolicyRequest
	GetRetention() *int32
	SetSpeedLimiter(v int64) *CreateUniBackupPolicyRequest
	GetSpeedLimiter() *int64
	SetUniRegionId(v string) *CreateUniBackupPolicyRequest
	GetUniRegionId() *string
	SetUuid(v string) *CreateUniBackupPolicyRequest
	GetUuid() *string
}

type CreateUniBackupPolicyRequest struct {
	// The name of the database account.
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
	// Specifies whether the database is manually added. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DatabaseAddByUser *string `json:"DatabaseAddByUser,omitempty" xml:"DatabaseAddByUser,omitempty"`
	// The type of the database. Valid values:
	//
	// 	- **MYSQL**
	//
	// 	- **ORACLE**
	//
	// 	- **MSSQL**
	//
	// This parameter is required.
	//
	// example:
	//
	// MYSQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The policy for full data backup. The value of this parameter is a JSON string. The JSON string contains the following fields:
	//
	// 	- **start**: the start time of a backup task.
	//
	// 	- **interval**: the interval of backup tasks.
	//
	// 	- **type**: the unit of the interval.
	//
	// 	- **days**: the days of a week on which a backup task is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"days":[4],"interval":1,"planType":"weekly","startTime":"22:00:00"}
	FullPlan map[string]interface{} `json:"FullPlan,omitempty" xml:"FullPlan,omitempty"`
	// The policy for incremental data backup. The value of this parameter is a JSON string. The JSON string contains the following fields:
	//
	// 	- **start**: the start time of a backup task.
	//
	// 	- **interval**: the interval of backup tasks.
	//
	// 	- **type**: the unit of the interval.
	//
	// 	- **days**: the days of a week on which a backup task is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"interval":1,"planType":"daily","startTime":"23:30:00"}
	IncPlan map[string]interface{} `json:"IncPlan,omitempty" xml:"IncPlan,omitempty"`
	// The ID of the Elastic Compute Service (ECS) instance.
	//
	// >  You can call the [DescribeUniBackupDatabase](~~DescribeUniBackupDatabase~~) operation to query the IDs of ECS instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1fu4aqltf1huhc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the anti-ransomware policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql-policy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The retention period of backup data.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The maximum network bandwidth that is allowed during data backup. Unit: bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5242880
	SpeedLimiter *int64 `json:"SpeedLimiter,omitempty" xml:"SpeedLimiter,omitempty"`
	// The region in which the server resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	UniRegionId *string `json:"UniRegionId,omitempty" xml:"UniRegionId,omitempty"`
	// The UUID of the server whose data is backed up based on the anti-ransomware policy.
	//
	// >  You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 045cad48-eb08-4047-a70c-713aec7b****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateUniBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUniBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateUniBackupPolicyRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateUniBackupPolicyRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateUniBackupPolicyRequest) GetDatabaseAddByUser() *string {
	return s.DatabaseAddByUser
}

func (s *CreateUniBackupPolicyRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *CreateUniBackupPolicyRequest) GetFullPlan() map[string]interface{} {
	return s.FullPlan
}

func (s *CreateUniBackupPolicyRequest) GetIncPlan() map[string]interface{} {
	return s.IncPlan
}

func (s *CreateUniBackupPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUniBackupPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateUniBackupPolicyRequest) GetRetention() *int32 {
	return s.Retention
}

func (s *CreateUniBackupPolicyRequest) GetSpeedLimiter() *int64 {
	return s.SpeedLimiter
}

func (s *CreateUniBackupPolicyRequest) GetUniRegionId() *string {
	return s.UniRegionId
}

func (s *CreateUniBackupPolicyRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CreateUniBackupPolicyRequest) SetAccountName(v string) *CreateUniBackupPolicyRequest {
	s.AccountName = &v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetAccountPassword(v string) *CreateUniBackupPolicyRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetDatabaseAddByUser(v string) *CreateUniBackupPolicyRequest {
	s.DatabaseAddByUser = &v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetDatabaseType(v string) *CreateUniBackupPolicyRequest {
	s.DatabaseType = &v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetFullPlan(v map[string]interface{}) *CreateUniBackupPolicyRequest {
	s.FullPlan = v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetIncPlan(v map[string]interface{}) *CreateUniBackupPolicyRequest {
	s.IncPlan = v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetInstanceId(v string) *CreateUniBackupPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetPolicyName(v string) *CreateUniBackupPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetRetention(v int32) *CreateUniBackupPolicyRequest {
	s.Retention = &v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetSpeedLimiter(v int64) *CreateUniBackupPolicyRequest {
	s.SpeedLimiter = &v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetUniRegionId(v string) *CreateUniBackupPolicyRequest {
	s.UniRegionId = &v
	return s
}

func (s *CreateUniBackupPolicyRequest) SetUuid(v string) *CreateUniBackupPolicyRequest {
	s.Uuid = &v
	return s
}

func (s *CreateUniBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
