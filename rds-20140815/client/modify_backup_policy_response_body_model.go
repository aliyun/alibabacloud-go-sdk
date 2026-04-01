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
	SetHighSpaceUsageProtection(v string) *ModifyBackupPolicyResponseBody
	GetHighSpaceUsageProtection() *string
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
	// The method that is used to compress backups. Valid values:
	//
	// 	- **0:*	- Backups are not compressed.
	//
	// 	- **1**: Backups are compressed by using the zlib tool.
	//
	// 	- **2**: Backups are compressed in parallel by using the zlib tool.
	//
	// 	- **4**: Backups are compressed by using the QuickLZ tool and can be used to restore individual databases and tables.
	//
	// 	- **8**: Backups are compressed by using the QuickLZ tool but cannot be used to restore individual databases or tables. This value is supported only for instances that run MySQL 8.0.
	//
	// example:
	//
	// 4
	CompressType *string `json:"CompressType,omitempty" xml:"CompressType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceID *string `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// Indicates whether the log backup feature is enabled. Valid values:
	//
	// 	- **1**: The feature is enabled.
	//
	// 	- **0**: The feature is disabled.
	//
	// example:
	//
	// 1
	EnableBackupLog *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Specifies whether to forcefully delete log backup files from the instance when the storage usage of the instance exceeds 80% or the amount of remaining storage on the instance is less than 5 GB.
	//
	// example:
	//
	// Disable
	HighSpaceUsageProtection *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	// The number of hours for which log backup files are retained on the instance.
	//
	// example:
	//
	// 18
	LocalLogRetentionHours *int32 `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	// The maximum storage usage that is allowed for log backup files on the instance.
	//
	// example:
	//
	// 30
	LocalLogRetentionSpace *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	// The number of binary log files on the instance.
	//
	// example:
	//
	// 60
	LogBackupLocalRetentionNumber *int32 `json:"LogBackupLocalRetentionNumber,omitempty" xml:"LogBackupLocalRetentionNumber,omitempty"`
	// The ID of the request.
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

func (s *ModifyBackupPolicyResponseBody) GetHighSpaceUsageProtection() *string {
	return s.HighSpaceUsageProtection
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

func (s *ModifyBackupPolicyResponseBody) SetHighSpaceUsageProtection(v string) *ModifyBackupPolicyResponseBody {
	s.HighSpaceUsageProtection = &v
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
