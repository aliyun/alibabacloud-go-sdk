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
	// Specifies whether to enable the cross-region backup feature on the instance. This parameter specifies whether you can back up data and logs. Valid values:
	//
	// 	- **0**: disables the feature.
	//
	// 	- **1:*	- enables the feature.
	//
	// > Before you enable the cross-region backup feature, you must configure the CrossBackupRegion parameter.
	//
	// example:
	//
	// 1
	BackupEnabled *string `json:"BackupEnabled,omitempty" xml:"BackupEnabled,omitempty"`
	// The ID of the region in which the cross-region backup files of the instance are stored.
	//
	// example:
	//
	// cn-shanghai
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	// The policy that is used to save the cross-region backup files of the instance. Set the value to **1**. The value 1 specifies that all cross-region backup files are saved.
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
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable the cross-region log backup feature on the instance. Valid values:
	//
	// 	- **0**: disables the feature.
	//
	// 	- **1:*	- enables the feature.
	//
	// > You can enable the cross-region log backup feature only when the cross-region backup feature is enabled.
	//
	// example:
	//
	// 1
	LogBackupEnabled *string `json:"LogBackupEnabled,omitempty" xml:"LogBackupEnabled,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the source instance. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The policy that is used to retain the cross-region backup files of the instance. Set the value to 1. The value **1*	- specifies that the cross-region backup files of the instance are retained based on the specified retention period.
	//
	// example:
	//
	// 1
	RetentType *int32 `json:"RetentType,omitempty" xml:"RetentType,omitempty"`
	// The number of days for which the cross-region backup files of the instance are retained. Valid values: **7 to 1825**.
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
