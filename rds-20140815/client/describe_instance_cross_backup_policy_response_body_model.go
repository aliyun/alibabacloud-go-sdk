// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceCrossBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupEnabled(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetBackupEnabled() *string
	SetBackupEnabledTime(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetBackupEnabledTime() *string
	SetCrossBackupRegion(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetCrossBackupRegion() *string
	SetCrossBackupType(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetCrossBackupType() *string
	SetDBInstanceDescription(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetDBInstanceId() *string
	SetDBInstanceStatus(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetDBInstanceStatus() *string
	SetEngine(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetEngine() *string
	SetEngineVersion(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetEngineVersion() *string
	SetLockMode(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetLockMode() *string
	SetLogBackupEnabled(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetLogBackupEnabled() *string
	SetLogBackupEnabledTime(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetLogBackupEnabledTime() *string
	SetRegionId(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeInstanceCrossBackupPolicyResponseBody
	GetRequestId() *string
	SetRetentType(v int32) *DescribeInstanceCrossBackupPolicyResponseBody
	GetRetentType() *int32
	SetRetention(v int32) *DescribeInstanceCrossBackupPolicyResponseBody
	GetRetention() *int32
}

type DescribeInstanceCrossBackupPolicyResponseBody struct {
	// The status of the cross-region backup feature on the instance. Valid values:
	//
	// 	- **Disable**
	//
	// 	- **Enable**
	//
	// example:
	//
	// Enable
	BackupEnabled *string `json:"BackupEnabled,omitempty" xml:"BackupEnabled,omitempty"`
	// The point in time at which the cross-region backup feature is enabled. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-06-12T05:44:21Z
	BackupEnabledTime *string `json:"BackupEnabledTime,omitempty" xml:"BackupEnabledTime,omitempty"`
	// The ID of the destination region where the cross-region backup files of the instance are stored.
	//
	// example:
	//
	// cn-shanghai
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	// The policy that is used to save the cross-region backup files of the instance. Default value: **1**. The value 1 indicates that all cross-region backup files are saved.
	//
	// example:
	//
	// 1
	CrossBackupType *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	// The name of the instance. It must be 2 to 256 characters in length. The value can contain letters, digits, underscores (_), and hyphens (-), and must start with a letter.
	//
	// >  The value cannot start with http:// or https://.
	//
	// example:
	//
	// Test database
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of the instance. For more information, see [Instance state table](https://help.aliyun.com/document_detail/26315.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The lock status of the instance. Valid values:
	//
	// 	- **Unlock**: The instance is not locked.
	//
	// 	- **ManualLock**: The instance is manually locked.
	//
	// 	- **LockByExpiration**: The instance is automatically locked due to instance expiration.
	//
	// 	- **LockByRestoration**: The instance is automatically locked before a rollback.
	//
	// 	- **LockByDiskQuota**: The instance is automatically locked because its storage capacity is exhausted and the instance is inaccessible.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The status of the cross-region log backup feature on the instance. Valid values:
	//
	// 	- **Disable**
	//
	// 	- **Enable**
	//
	// example:
	//
	// Enable
	LogBackupEnabled *string `json:"LogBackupEnabled,omitempty" xml:"LogBackupEnabled,omitempty"`
	// The time when cross-region log backup was enabled on the instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-06-12T05:44:21Z
	LogBackupEnabledTime *string `json:"LogBackupEnabledTime,omitempty" xml:"LogBackupEnabledTime,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB7667B2-72C8-497B-9BD8-3B343CEF51AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The policy that is used to retain the cross-region backup files of the instance. Default value: **1**. The value 1 indicates that the cross-region backup files of the instance are retained based on the specified retention period.
	//
	// example:
	//
	// 1
	RetentType *int32 `json:"RetentType,omitempty" xml:"RetentType,omitempty"`
	// The number of days for which the cross-region backup files of the instance are retained. Valid values: **7 to 1825**.
	//
	// example:
	//
	// 15
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s DescribeInstanceCrossBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceCrossBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetBackupEnabled() *string {
	return s.BackupEnabled
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetBackupEnabledTime() *string {
	return s.BackupEnabledTime
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetCrossBackupType() *string {
	return s.CrossBackupType
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetLogBackupEnabled() *string {
	return s.LogBackupEnabled
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetLogBackupEnabledTime() *string {
	return s.LogBackupEnabledTime
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetRetentType() *int32 {
	return s.RetentType
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) GetRetention() *int32 {
	return s.Retention
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetBackupEnabled(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.BackupEnabled = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetBackupEnabledTime(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.BackupEnabledTime = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetCrossBackupRegion(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.CrossBackupRegion = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetCrossBackupType(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.CrossBackupType = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetDBInstanceDescription(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetDBInstanceId(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetDBInstanceStatus(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetEngine(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetEngineVersion(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetLockMode(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetLogBackupEnabled(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.LogBackupEnabled = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetLogBackupEnabledTime(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.LogBackupEnabledTime = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetRegionId(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetRequestId(v string) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetRetentType(v int32) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.RetentType = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) SetRetention(v int32) *DescribeInstanceCrossBackupPolicyResponseBody {
	s.Retention = &v
	return s
}

func (s *DescribeInstanceCrossBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
