// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceCrossBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyRequest
	GetBackupEnabled() *string
	SetCrossBackupRegion(v string) *ModifyInstanceCrossBackupPolicyRequest
	GetCrossBackupRegion() *string
	SetCrossBackupType(v string) *ModifyInstanceCrossBackupPolicyRequest
	GetCrossBackupType() *string
	SetDBInstanceId(v string) *ModifyInstanceCrossBackupPolicyRequest
	GetDBInstanceId() *string
	SetLogBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyRequest
	GetLogBackupEnabled() *string
	SetOwnerId(v int64) *ModifyInstanceCrossBackupPolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceCrossBackupPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceCrossBackupPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceCrossBackupPolicyRequest
	GetResourceOwnerId() *int64
	SetRetentType(v int32) *ModifyInstanceCrossBackupPolicyRequest
	GetRetentType() *int32
	SetRetention(v int32) *ModifyInstanceCrossBackupPolicyRequest
	GetRetention() *int32
}

type ModifyInstanceCrossBackupPolicyRequest struct {
	BackupEnabled     *string `json:"BackupEnabled,omitempty" xml:"BackupEnabled,omitempty"`
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	CrossBackupType   *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	// This parameter is required.
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	LogBackupEnabled *string `json:"LogBackupEnabled,omitempty" xml:"LogBackupEnabled,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RetentType           *int32  `json:"RetentType,omitempty" xml:"RetentType,omitempty"`
	Retention            *int32  `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s ModifyInstanceCrossBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceCrossBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetBackupEnabled() *string {
	return s.BackupEnabled
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetCrossBackupType() *string {
	return s.CrossBackupType
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetLogBackupEnabled() *string {
	return s.LogBackupEnabled
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetRetentType() *int32 {
	return s.RetentType
}

func (s *ModifyInstanceCrossBackupPolicyRequest) GetRetention() *int32 {
	return s.Retention
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyRequest {
	s.BackupEnabled = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetCrossBackupRegion(v string) *ModifyInstanceCrossBackupPolicyRequest {
	s.CrossBackupRegion = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetCrossBackupType(v string) *ModifyInstanceCrossBackupPolicyRequest {
	s.CrossBackupType = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetDBInstanceId(v string) *ModifyInstanceCrossBackupPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetLogBackupEnabled(v string) *ModifyInstanceCrossBackupPolicyRequest {
	s.LogBackupEnabled = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetOwnerId(v int64) *ModifyInstanceCrossBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetRegionId(v string) *ModifyInstanceCrossBackupPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyInstanceCrossBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyInstanceCrossBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetRetentType(v int32) *ModifyInstanceCrossBackupPolicyRequest {
	s.RetentType = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) SetRetention(v int32) *ModifyInstanceCrossBackupPolicyRequest {
	s.Retention = &v
	return s
}

func (s *ModifyInstanceCrossBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
