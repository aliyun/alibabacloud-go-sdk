// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreateDdrDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSetId(v string) *CheckCreateDdrDBInstanceRequest
	GetBackupSetId() *string
	SetDBInstanceClass(v string) *CheckCreateDdrDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceStorage(v int32) *CheckCreateDdrDBInstanceRequest
	GetDBInstanceStorage() *int32
	SetEngine(v string) *CheckCreateDdrDBInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CheckCreateDdrDBInstanceRequest
	GetEngineVersion() *string
	SetOwnerId(v int64) *CheckCreateDdrDBInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CheckCreateDdrDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CheckCreateDdrDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CheckCreateDdrDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckCreateDdrDBInstanceRequest
	GetResourceOwnerId() *int64
	SetRestoreTime(v string) *CheckCreateDdrDBInstanceRequest
	GetRestoreTime() *string
	SetRestoreType(v string) *CheckCreateDdrDBInstanceRequest
	GetRestoreType() *string
	SetSourceDBInstanceName(v string) *CheckCreateDdrDBInstanceRequest
	GetSourceDBInstanceName() *string
	SetSourceRegion(v string) *CheckCreateDdrDBInstanceRequest
	GetSourceRegion() *string
}

type CheckCreateDdrDBInstanceRequest struct {
	// The ID of the backup set that is used for the restoration. You can call the DescribeCrossRegionBackups operation to query the backup set ID.
	//
	// >  This parameter must be specified when the **RestoreType*	- parameter is set to **0**.
	//
	// example:
	//
	// 14358
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The instance type of the destination instance. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rds.mysql.s1.small
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The storage capacity of the destination instance. Valid values: **5 to 2000**. Unit: GB. You can increase the storage capacity in increments of 5 GB. For more information, see [Primary instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The database engine of the destination instance. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The major engine version of the destination instance. The value of this parameter varies based on the value of **Engine**.
	//
	// 	- Valid values when Engine is set to MySQL: **5.5, 5.6, 5.7, and 8.0**
	//
	// 	- Valid values when Engine is set to SQLServer: **2008r2, 08r2_ent_ha, 2012, 2012_ent_ha, 2012_std_ha, 2012_web, 2014_std_ha, 2016_ent_ha, 2016_std_ha, 2016_web, 2017_std_ha, 2017_ent, 2019_std_ha, and 2019_ent**
	//
	// 	- PostgreSQL: **10.0, 11.0, 12.0, 13.0, 14.0, and 15.0**
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the destination instance. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time to which you want to restore data. The point in time that you specify must be earlier than the current time. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > If you set **RestoreType*	- to **1**, you must also specify this parameter.
	//
	// example:
	//
	// 2019-05-30T03:29:10Z
	RestoreTime *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The method that is used to restore data. Valid values:
	//
	// 	- **0**: restores data from a backup set. If you set this parameter to 0, you must also specify the **BackupSetId*	- parameter.
	//
	// 	- **1**: restores data to a point in time. If you set this parameter to 1, you must also specify the **RestoreTime**, **SourceRegion**, and **SourceDBInstanceName*	- parameters.
	//
	// Default value: **0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The ID of the source instance if you want to restore data to a point in time.
	//
	// >  This parameter must be specified when the **RestoreType*	- parameter is set to **1**.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	// The region ID of the source instance if you want to restore data to a point in time.
	//
	// > If you set **RestoreType*	- to **1**, you must also specify this parameter.
	//
	// example:
	//
	// cn-hangzhou
	SourceRegion *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
}

func (s CheckCreateDdrDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateDdrDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CheckCreateDdrDBInstanceRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CheckCreateDdrDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CheckCreateDdrDBInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *CheckCreateDdrDBInstanceRequest) GetEngine() *string {
	return s.Engine
}

func (s *CheckCreateDdrDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CheckCreateDdrDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckCreateDdrDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckCreateDdrDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckCreateDdrDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckCreateDdrDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckCreateDdrDBInstanceRequest) GetRestoreTime() *string {
	return s.RestoreTime
}

func (s *CheckCreateDdrDBInstanceRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CheckCreateDdrDBInstanceRequest) GetSourceDBInstanceName() *string {
	return s.SourceDBInstanceName
}

func (s *CheckCreateDdrDBInstanceRequest) GetSourceRegion() *string {
	return s.SourceRegion
}

func (s *CheckCreateDdrDBInstanceRequest) SetBackupSetId(v string) *CheckCreateDdrDBInstanceRequest {
	s.BackupSetId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetDBInstanceClass(v string) *CheckCreateDdrDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetDBInstanceStorage(v int32) *CheckCreateDdrDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetEngine(v string) *CheckCreateDdrDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetEngineVersion(v string) *CheckCreateDdrDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetOwnerId(v int64) *CheckCreateDdrDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetRegionId(v string) *CheckCreateDdrDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetResourceGroupId(v string) *CheckCreateDdrDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetResourceOwnerAccount(v string) *CheckCreateDdrDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetResourceOwnerId(v int64) *CheckCreateDdrDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetRestoreTime(v string) *CheckCreateDdrDBInstanceRequest {
	s.RestoreTime = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetRestoreType(v string) *CheckCreateDdrDBInstanceRequest {
	s.RestoreType = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetSourceDBInstanceName(v string) *CheckCreateDdrDBInstanceRequest {
	s.SourceDBInstanceName = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) SetSourceRegion(v string) *CheckCreateDdrDBInstanceRequest {
	s.SourceRegion = &v
	return s
}

func (s *CheckCreateDdrDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
