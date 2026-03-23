// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceCrossBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetBackupEnabled() *string
	SetCrossBackupRegion(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetCrossBackupRegion() *string
	SetCrossBackupType(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetCrossBackupType() *string
	SetDBInstanceId(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetDBInstanceId() *string
	SetLogBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetLogBackupEnabled() *string
	SetRegionId(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetRegionId() *string
	SetRequestId(v string) *ModifyInstanceCrossBackupPolicyResponseBody
	GetRequestId() *string
	SetRetentType(v int32) *ModifyInstanceCrossBackupPolicyResponseBody
	GetRetentType() *int32
	SetRetention(v int32) *ModifyInstanceCrossBackupPolicyResponseBody
	GetRetention() *int32
}

type ModifyInstanceCrossBackupPolicyResponseBody struct {
	BackupEnabled     *string `json:"BackupEnabled,omitempty" xml:"BackupEnabled,omitempty"`
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	CrossBackupType   *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	LogBackupEnabled  *string `json:"LogBackupEnabled,omitempty" xml:"LogBackupEnabled,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetentType        *int32  `json:"RetentType,omitempty" xml:"RetentType,omitempty"`
	Retention         *int32  `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s ModifyInstanceCrossBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceCrossBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetBackupEnabled() *string {
	return s.BackupEnabled
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetCrossBackupType() *string {
	return s.CrossBackupType
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetLogBackupEnabled() *string {
	return s.LogBackupEnabled
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetRetentType() *int32 {
	return s.RetentType
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) GetRetention() *int32 {
	return s.Retention
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.BackupEnabled = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetCrossBackupRegion(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.CrossBackupRegion = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetCrossBackupType(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.CrossBackupType = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetDBInstanceId(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetLogBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.LogBackupEnabled = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetRegionId(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetRequestId(v string) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetRetentType(v int32) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.RetentType = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) SetRetention(v int32) *ModifyInstanceCrossBackupPolicyResponseBody {
	s.Retention = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
