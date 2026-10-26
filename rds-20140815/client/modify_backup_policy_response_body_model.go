// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompressType(v string) *ModifyBackupPolicyResponseBody
	GetCompressType() *string
	SetDBInstanceID(v string) *ModifyBackupPolicyResponseBody
	GetDBInstanceID() *string
	SetEnableBackupLog(v string) *ModifyBackupPolicyResponseBody
	GetEnableBackupLog() *string
	SetEnableIncrementDataBackup(v bool) *ModifyBackupPolicyResponseBody
	GetEnableIncrementDataBackup() *bool
	SetEnablePitrProtection(v bool) *ModifyBackupPolicyResponseBody
	GetEnablePitrProtection() *bool
	SetHighSpaceUsageProtection(v string) *ModifyBackupPolicyResponseBody
	GetHighSpaceUsageProtection() *string
	SetIncBackupInterval(v int32) *ModifyBackupPolicyResponseBody
	GetIncBackupInterval() *int32
	SetLocalLogRetentionHours(v int32) *ModifyBackupPolicyResponseBody
	GetLocalLogRetentionHours() *int32
	SetLocalLogRetentionSpace(v string) *ModifyBackupPolicyResponseBody
	GetLocalLogRetentionSpace() *string
	SetLogBackupLocalRetentionNumber(v int32) *ModifyBackupPolicyResponseBody
	GetLogBackupLocalRetentionNumber() *int32
	SetRequestId(v string) *ModifyBackupPolicyResponseBody
	GetRequestId() *string
}

type ModifyBackupPolicyResponseBody struct {
	// The backup compression method. Valid values:
	//
	// 	- **0**: no compression.
	//
	// 	- **1**: zlib compression.
	//
	// 	- **2**: parallel zlib compression.
	//
	// 	- **4**: quicklz compression. Database and table restoration is enabled.
	//
	// 	- **8**: quicklz compression for MySQL 8.0. Database and table restoration is not supported.
	//
	// example:
	//
	// 4
	CompressType *string `json:"CompressType,omitempty" xml:"CompressType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceID *string `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// Indicates whether instance log backup is enabled. Valid values:
	//
	// 	- **1**: enabled.
	//
	// 	- **0**: shutdown.
	//
	//
	// > Instance log backup is enabled by default and cannot be shut down for ApsaraDB RDS for SQL Server instances.
	//
	// example:
	//
	// 1
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Indicates whether incremental backup is enabled.
	EnableIncrementDataBackup *bool `json:"EnableIncrementDataBackup,omitempty" xml:"EnableIncrementDataBackup,omitempty"`
	// Indicates whether point-in-time recovery (PITR) is enabled for the ApsaraDB RDS for MySQL instance.
	EnablePitrProtection *bool `json:"EnablePitrProtection,omitempty" xml:"EnablePitrProtection,omitempty"`
	// Indicates whether binary log files are unconditionally deleted when the storage usage of the **ApsaraDB RDS for MySQL*	- instance exceeds 80% or the remaining storage space is less than 5 GB.
	//
	// example:
	//
	// Disable
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// The high-frequency incremental backup interval. Unit: minutes. Valid values: 60, 120, 240, 360, and 720.
	//
	// example:
	//
	// 120
	IncBackupInterval *int32 `json:"IncBackupInterval,omitempty" xml:"IncBackupInterval,omitempty"`
	// The number of hours for which instance log backups are retained locally on the **MySQL*	- instance.
	//
	// example:
	//
	// 18
	LocalLogRetentionHours *int32 `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// The maximum loop space usage of the local binary log on the **MySQL*	- instance.
	//
	// example:
	//
	// 30
	LocalLogRetentionSpace *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// The number of binary log files retained locally on the **ApsaraDB RDS for MySQL*	- instance.
	//
	// example:
	//
	// 60
	LogBackupLocalRetentionNumber *int32 `json:"LogBackupLocalRetentionNumber,omitempty" xml:"LogBackupLocalRetentionNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DA147739-AEAD-4417-9089-65E9B1D8240D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBody) GetCompressType() *string {
	return s.CompressType
}

func (s *ModifyBackupPolicyResponseBody) GetDBInstanceID() *string {
	return s.DBInstanceID
}

func (s *ModifyBackupPolicyResponseBody) GetEnableBackupLog() *string {
	return s.EnableBackupLog
}

func (s *ModifyBackupPolicyResponseBody) GetEnableIncrementDataBackup() *bool {
	return s.EnableIncrementDataBackup
}

func (s *ModifyBackupPolicyResponseBody) GetEnablePitrProtection() *bool {
	return s.EnablePitrProtection
}

func (s *ModifyBackupPolicyResponseBody) GetHighSpaceUsageProtection() *string {
	return s.HighSpaceUsageProtection
}

func (s *ModifyBackupPolicyResponseBody) GetIncBackupInterval() *int32 {
	return s.IncBackupInterval
}

func (s *ModifyBackupPolicyResponseBody) GetLocalLogRetentionHours() *int32 {
	return s.LocalLogRetentionHours
}

func (s *ModifyBackupPolicyResponseBody) GetLocalLogRetentionSpace() *string {
	return s.LocalLogRetentionSpace
}

func (s *ModifyBackupPolicyResponseBody) GetLogBackupLocalRetentionNumber() *int32 {
	return s.LogBackupLocalRetentionNumber
}

func (s *ModifyBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupPolicyResponseBody) SetCompressType(v string) *ModifyBackupPolicyResponseBody {
	s.CompressType = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetDBInstanceID(v string) *ModifyBackupPolicyResponseBody {
	s.DBInstanceID = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetEnableBackupLog(v string) *ModifyBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetEnableIncrementDataBackup(v bool) *ModifyBackupPolicyResponseBody {
	s.EnableIncrementDataBackup = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetEnablePitrProtection(v bool) *ModifyBackupPolicyResponseBody {
	s.EnablePitrProtection = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetHighSpaceUsageProtection(v string) *ModifyBackupPolicyResponseBody {
	s.HighSpaceUsageProtection = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetIncBackupInterval(v int32) *ModifyBackupPolicyResponseBody {
	s.IncBackupInterval = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetLocalLogRetentionHours(v int32) *ModifyBackupPolicyResponseBody {
	s.LocalLogRetentionHours = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetLocalLogRetentionSpace(v string) *ModifyBackupPolicyResponseBody {
	s.LocalLogRetentionSpace = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetLogBackupLocalRetentionNumber(v int32) *ModifyBackupPolicyResponseBody {
	s.LogBackupLocalRetentionNumber = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) SetRequestId(v string) *ModifyBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
