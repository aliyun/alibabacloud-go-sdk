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
	CompressType                  *string `json:"CompressType,omitempty" xml:"CompressType,omitempty"`
	DBInstanceID                  *string `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	EnableBackupLog               *string `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	HighSpaceUsageProtection      *string `json:"HighSpaceUsageProtection,omitempty" xml:"HighSpaceUsageProtection,omitempty"`
	LocalLogRetentionHours        *int32  `json:"LocalLogRetentionHours,omitempty" xml:"LocalLogRetentionHours,omitempty"`
	LocalLogRetentionSpace        *string `json:"LocalLogRetentionSpace,omitempty" xml:"LocalLogRetentionSpace,omitempty"`
	LogBackupLocalRetentionNumber *int32  `json:"LogBackupLocalRetentionNumber,omitempty" xml:"LogBackupLocalRetentionNumber,omitempty"`
	RequestId                     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
