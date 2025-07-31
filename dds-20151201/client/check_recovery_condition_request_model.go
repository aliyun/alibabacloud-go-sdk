// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRecoveryConditionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *CheckRecoveryConditionRequest
	GetBackupId() *string
	SetDatabaseNames(v string) *CheckRecoveryConditionRequest
	GetDatabaseNames() *string
	SetDestRegion(v string) *CheckRecoveryConditionRequest
	GetDestRegion() *string
	SetEngineVersion(v string) *CheckRecoveryConditionRequest
	GetEngineVersion() *string
	SetInstanceType(v string) *CheckRecoveryConditionRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *CheckRecoveryConditionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckRecoveryConditionRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CheckRecoveryConditionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CheckRecoveryConditionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckRecoveryConditionRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CheckRecoveryConditionRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *CheckRecoveryConditionRequest
	GetRestoreType() *string
	SetSourceDBInstance(v string) *CheckRecoveryConditionRequest
	GetSourceDBInstance() *string
	SetSrcRegion(v string) *CheckRecoveryConditionRequest
	GetSrcRegion() *string
}

type CheckRecoveryConditionRequest struct {
	// The backup ID.
	//
	// > 	- You can call the [DescribeBackups](https://help.aliyun.com/document_detail/62172.html) operation to query the backup ID.
	//
	// > 	- You must specify one of the **RestoreTime*	- and BackupId parameters.
	//
	// > 	- This parameter is not applicable to sharded cluster instances.
	//
	// example:
	//
	// 5664****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The name of the source database. The value is a JSON array.
	//
	// >  If you do not specify this parameter, all databases are restored by default.
	//
	// example:
	//
	// ["db1","db2"]
	DatabaseNames *string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty"`
	// The region of the backup set used for the cross-region backup and restoration.
	//
	// >  This parameter is required when you set the RestoreType parameter to 3.
	//
	// example:
	//
	// cn-hangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// The database engine version of the instance.
	//
	// 	- **6.0**
	//
	// 	- **5.0**
	//
	// 	- **4.4**
	//
	// 	- **4.2**
	//
	// 	- **4.0**
	//
	// 	- **3.4**
	//
	// example:
	//
	// 4.2
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance architecture. Valid values:
	//
	// 	- replicate
	//
	// 	- sharding
	//
	// > 	- This parameter is required when you set the RestoreType parameter to 2.
	//
	// > 	- This parameter is required when you set the RestoreType parameter to 3.
	//
	// example:
	//
	// replicate
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// sg-bp179****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which the instance is restored. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > 	- The time can be a point in time within the past seven days. The time must be earlier than the current time, but later than the time when the instance was created.
	//
	// > 	- You must specify one of the RestoreTime and **BackupId*	- parameters.
	//
	// example:
	//
	// 2022-08-22T08:00:00Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The restoration type.
	//
	// > 	- 0: The data of the source instance is restored to a new instance in the same region.
	//
	// > 	- 1: The data of the source instance is restored to an earlier point in time.
	//
	// > 	- 2: The data of a deleted instance is restored to a new instance from the backup set.
	//
	// > 	- 3: The data of a deleted instance is restored to a new instance in another region from the backup set.
	//
	// example:
	//
	// 0
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// dds-bp1378****
	SourceDBInstance *string `json:"SourceDBInstance,omitempty" xml:"SourceDBInstance,omitempty"`
	// The region where the source instance resides.
	//
	// > 	- This parameter is required when you set the RestoreType parameter to 2.
	//
	// > 	- This parameter is required when you set the RestoreType parameter to 3.
	//
	// example:
	//
	// cn-beijing
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
}

func (s CheckRecoveryConditionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckRecoveryConditionRequest) GoString() string {
	return s.String()
}

func (s *CheckRecoveryConditionRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *CheckRecoveryConditionRequest) GetDatabaseNames() *string {
	return s.DatabaseNames
}

func (s *CheckRecoveryConditionRequest) GetDestRegion() *string {
	return s.DestRegion
}

func (s *CheckRecoveryConditionRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CheckRecoveryConditionRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CheckRecoveryConditionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckRecoveryConditionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckRecoveryConditionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckRecoveryConditionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckRecoveryConditionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckRecoveryConditionRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CheckRecoveryConditionRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CheckRecoveryConditionRequest) GetSourceDBInstance() *string {
	return s.SourceDBInstance
}

func (s *CheckRecoveryConditionRequest) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *CheckRecoveryConditionRequest) SetBackupId(v string) *CheckRecoveryConditionRequest {
	s.BackupId = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetDatabaseNames(v string) *CheckRecoveryConditionRequest {
	s.DatabaseNames = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetDestRegion(v string) *CheckRecoveryConditionRequest {
	s.DestRegion = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetEngineVersion(v string) *CheckRecoveryConditionRequest {
	s.EngineVersion = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetInstanceType(v string) *CheckRecoveryConditionRequest {
	s.InstanceType = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetOwnerAccount(v string) *CheckRecoveryConditionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetOwnerId(v int64) *CheckRecoveryConditionRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetResourceGroupId(v string) *CheckRecoveryConditionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetResourceOwnerAccount(v string) *CheckRecoveryConditionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetResourceOwnerId(v int64) *CheckRecoveryConditionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetRestoreTime(v string) *CheckRecoveryConditionRequest {
	s.RestoreTime = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetRestoreType(v string) *CheckRecoveryConditionRequest {
	s.RestoreType = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetSourceDBInstance(v string) *CheckRecoveryConditionRequest {
	s.SourceDBInstance = &v
	return s
}

func (s *CheckRecoveryConditionRequest) SetSrcRegion(v string) *CheckRecoveryConditionRequest {
	s.SrcRegion = &v
	return s
}

func (s *CheckRecoveryConditionRequest) Validate() error {
	return dara.Validate(s)
}
