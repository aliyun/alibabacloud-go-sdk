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
	BackupEnabled         *string `json:"BackupEnabled,omitempty" xml:"BackupEnabled,omitempty"`
	BackupEnabledTime     *string `json:"BackupEnabledTime,omitempty" xml:"BackupEnabledTime,omitempty"`
	CrossBackupRegion     *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	CrossBackupType       *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus      *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	LockMode              *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LogBackupEnabled      *string `json:"LogBackupEnabled,omitempty" xml:"LogBackupEnabled,omitempty"`
	LogBackupEnabledTime  *string `json:"LogBackupEnabledTime,omitempty" xml:"LogBackupEnabledTime,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetentType            *int32  `json:"RetentType,omitempty" xml:"RetentType,omitempty"`
	Retention             *int32  `json:"Retention,omitempty" xml:"Retention,omitempty"`
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
