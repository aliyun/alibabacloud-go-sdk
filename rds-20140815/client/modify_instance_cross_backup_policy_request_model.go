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
	// Specifies whether to enable the cross-region backup feature, which includes data backup and log backup. Valid values:
	//
	// 	- **0**: Disabled.
	//
	// 	- **1**: Enabled.
	//
	// >When you enable the cross-region backup feature, you must specify the destination region ID.
	//
	// example:
	//
	// 1
	BackupEnabled *string `json:"BackupEnabled,omitempty" xml:"BackupEnabled,omitempty"`
	// The ID of the destination region for cross-region backup.
	//
	// example:
	//
	// cn-shanghai
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	// The type of cross-region backup retention. The only valid value is **1**, which indicates that all backups are retained.
	//
	// example:
	//
	// 1
	CrossBackupType *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable cross-region log backup. Valid values:
	//
	// 	- **0**: Disabled.
	//
	// 	- **1**: Enabled.
	//
	// >You can enable cross-region log backup only when the cross-region backup feature is enabled.
	//
	// example:
	//
	// 1
	LogBackupEnabled *string `json:"LogBackupEnabled,omitempty" xml:"LogBackupEnabled,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the source instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The cross-region backup retention method. The only valid value is **1**, which indicates retention by duration.
	//
	// example:
	//
	// 1
	RetentType *int32 `json:"RetentType,omitempty" xml:"RetentType,omitempty"`
	// The number of days for which cross-region backups are retained. Valid values: **7 to 1825**.
	//
	// example:
	//
	// 7
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
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
