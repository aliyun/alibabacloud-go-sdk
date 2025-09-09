// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPolicyDO(v *DescribeBackupPolicyResponseBodyBackupPolicyDO) *DescribeBackupPolicyResponseBody
	GetBackupPolicyDO() *DescribeBackupPolicyResponseBodyBackupPolicyDO
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupPolicyResponseBody
	GetSuccess() *bool
}

type DescribeBackupPolicyResponseBody struct {
	// The information about the backup policy.
	BackupPolicyDO *DescribeBackupPolicyResponseBodyBackupPolicyDO `json:"BackupPolicyDO,omitempty" xml:"BackupPolicyDO,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 8FAF3989-79CD-4A67-8FFD-97899B64****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) GetBackupPolicyDO() *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	return s.BackupPolicyDO
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupPolicyResponseBody) SetBackupPolicyDO(v *DescribeBackupPolicyResponseBodyBackupPolicyDO) *DescribeBackupPolicyResponseBody {
	s.BackupPolicyDO = v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetSuccess(v bool) *DescribeBackupPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPolicyResponseBodyBackupPolicyDO struct {
	// No value is returned.
	//
	// example:
	//
	// null
	BackupAppName *string `json:"BackupAppName,omitempty" xml:"BackupAppName,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// null
	BackupDbName *string `json:"BackupDbName,omitempty" xml:"BackupDbName,omitempty"`
	// The backup level. Valid values:
	//
	// 	- **db**: database backup
	//
	// 	- **instance**: instance backup
	//
	// example:
	//
	// instance
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// Indicates whether the log backup feature is enabled. Valid values:
	//
	// 	- **1**: The log backup feature is enabled.
	//
	// 	- **0**: The log backup feature is disabled.
	//
	// example:
	//
	// 1
	BackupLog *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **logic**: logical backup
	//
	// 	- **phy**: fast backup
	//
	// example:
	//
	// phy
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The type of the backup policy. Valid values:
	//
	// 	- **DataBackupPolicy**: a data backup policy
	//
	// 	- **LogBackupPolicy**: a log backup policy
	//
	// example:
	//
	// DataBackupPolicy
	BackupPolicyMode *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	// The retention period of backup files. Unit: days.
	//
	// example:
	//
	// 0
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// null
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The retention period of data backup files. Unit: days.
	//
	// example:
	//
	// 0
	DataBackupRetentionPeriod *int64 `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// 0
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// 0
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// 0
	HighSpaceUsageProtection *int64 `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// 0
	LocalLogRetentionHours *int64 `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// 0
	LocalLogRetentionSpace *int64 `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// The retention period of log backup files. Unit: days.
	//
	// example:
	//
	// 0
	LogBackupRetentionPeriod *int64 `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// null
	NextBackupActuallyTime *string `json:"NextBackupActuallyTime,omitempty" xml:"NextBackupActuallyTime,omitempty"`
	// The backup cycle. You can specify multiple backup cycles. Separate multiple backup cycles with commas (,). Valid values:
	//
	// 	- **0**: every Monday
	//
	// 	- **1**: every Tuesday
	//
	// 	- **2**: every Wednesday
	//
	// 	- **3**: every Thursday
	//
	// 	- **4**: every Friday
	//
	// 	- **5**: every Saturday
	//
	// 	- **6**: every Sunday
	//
	// example:
	//
	// 1,4
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The time range in which a backup is performed. The time is displayed in UTC.
	//
	// example:
	//
	// 22:00:00-23:00:00
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s DescribeBackupPolicyResponseBodyBackupPolicyDO) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyBackupPolicyDO) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetBackupAppName() *string {
	return s.BackupAppName
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetBackupDbName() *string {
	return s.BackupDbName
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetBackupLevel() *string {
	return s.BackupLevel
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetBackupLog() *string {
	return s.BackupLog
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetBackupPolicyMode() *string {
	return s.BackupPolicyMode
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetBackupRetentionPeriod() *int64 {
	return s.BackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetDataBackupRetentionPeriod() *int64 {
	return s.DataBackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetHighSpaceUsageProtection() *int64 {
	return s.HighSpaceUsageProtection
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetLocalLogRetentionHours() *int64 {
	return s.LocalLogRetentionHours
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetLocalLogRetentionSpace() *int64 {
	return s.LocalLogRetentionSpace
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetLogBackupRetentionPeriod() *int64 {
	return s.LogBackupRetentionPeriod
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetNextBackupActuallyTime() *string {
	return s.NextBackupActuallyTime
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupAppName(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupAppName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupDbName(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupDbName = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupLevel(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupLog(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupLog = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupMode(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupPolicyMode(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupPolicyMode = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetBackupType(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetDataBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetGmtCreate(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetGmtModified(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetHighSpaceUsageProtection(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLocalLogRetentionHours(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLocalLogRetentionSpace(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetLogBackupRetentionPeriod(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetNextBackupActuallyTime(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.NextBackupActuallyTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDO {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDO) Validate() error {
	return dara.Validate(s)
}
