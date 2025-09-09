// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupLocalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPolicyDO(v *DescribeBackupLocalResponseBodyBackupPolicyDO) *DescribeBackupLocalResponseBody
	GetBackupPolicyDO() *DescribeBackupLocalResponseBodyBackupPolicyDO
	SetRequestId(v string) *DescribeBackupLocalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupLocalResponseBody
	GetSuccess() *bool
}

type DescribeBackupLocalResponseBody struct {
	// The information about the backup policy.
	BackupPolicyDO *DescribeBackupLocalResponseBodyBackupPolicyDO `json:"BackupPolicyDO,omitempty" xml:"BackupPolicyDO,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FDC9CFD5-306D-4A23-9D8C-057274C6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupLocalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLocalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponseBody) GetBackupPolicyDO() *DescribeBackupLocalResponseBodyBackupPolicyDO {
	return s.BackupPolicyDO
}

func (s *DescribeBackupLocalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupLocalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupLocalResponseBody) SetBackupPolicyDO(v *DescribeBackupLocalResponseBodyBackupPolicyDO) *DescribeBackupLocalResponseBody {
	s.BackupPolicyDO = v
	return s
}

func (s *DescribeBackupLocalResponseBody) SetRequestId(v string) *DescribeBackupLocalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupLocalResponseBody) SetSuccess(v bool) *DescribeBackupLocalResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupLocalResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupLocalResponseBodyBackupPolicyDO struct {
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
	// No value is returned.
	//
	// example:
	//
	// null
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// null
	BackupLog *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// null
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// null
	BackupPolicyMode *string `json:"BackupPolicyMode,omitempty" xml:"BackupPolicyMode,omitempty"`
	// No value is returned.
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
	// No value is returned.
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
	// Indicates whether the feature is enabled to forcibly delete binary log files if the used storage space of the instance exceeds 90% of the total storage space or the remaining storage space is less than 5 GB. Valid values:
	//
	// 	- 1: The feature is enabled.
	//
	// 	- 0: The feature is disabled.
	//
	// example:
	//
	// 1
	HighSpaceUsageProtection *int64 `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// The number of hours for which log backup files are retained on the instance. Valid values: 0 to 168. Default value: **18**. The value **0*	- indicates that log backup files are not retained.
	//
	// example:
	//
	// 18
	LocalLogRetentionHours *int64 `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// The maximum storage usage that is allowed for local log files. Valid values: 0 to 50. Default value: 30.
	//
	// example:
	//
	// 30
	LocalLogRetentionSpace *int64 `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// No value is returned.
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
	// No value is returned.
	//
	// example:
	//
	// null
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// No value is returned.
	//
	// example:
	//
	// null
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s DescribeBackupLocalResponseBodyBackupPolicyDO) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLocalResponseBodyBackupPolicyDO) GoString() string {
	return s.String()
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetBackupAppName() *string {
	return s.BackupAppName
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetBackupDbName() *string {
	return s.BackupDbName
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetBackupLevel() *string {
	return s.BackupLevel
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetBackupLog() *string {
	return s.BackupLog
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetBackupPolicyMode() *string {
	return s.BackupPolicyMode
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetBackupRetentionPeriod() *int64 {
	return s.BackupRetentionPeriod
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetDataBackupRetentionPeriod() *int64 {
	return s.DataBackupRetentionPeriod
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetHighSpaceUsageProtection() *int64 {
	return s.HighSpaceUsageProtection
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetLocalLogRetentionHours() *int64 {
	return s.LocalLogRetentionHours
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetLocalLogRetentionSpace() *int64 {
	return s.LocalLogRetentionSpace
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetLogBackupRetentionPeriod() *int64 {
	return s.LogBackupRetentionPeriod
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetNextBackupActuallyTime() *string {
	return s.NextBackupActuallyTime
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupAppName(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupAppName = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupDbName(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupDbName = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupLevel(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupLog(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupLog = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupMode(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupPolicyMode(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupPolicyMode = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetBackupType(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetDataBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetGmtCreate(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.GmtCreate = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetGmtModified(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.GmtModified = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetHighSpaceUsageProtection(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLocalLogRetentionHours(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLocalLogRetentionSpace(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetLogBackupRetentionPeriod(v int64) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetNextBackupActuallyTime(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.NextBackupActuallyTime = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetPreferredBackupPeriod(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) SetPreferredBackupTime(v string) *DescribeBackupLocalResponseBodyBackupPolicyDO {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupLocalResponseBodyBackupPolicyDO) Validate() error {
	return dara.Validate(s)
}
