// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateToOtherZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *MigrateToOtherZoneRequest
	GetCategory() *string
	SetCustomExtraInfo(v string) *MigrateToOtherZoneRequest
	GetCustomExtraInfo() *string
	SetDBInstanceClass(v string) *MigrateToOtherZoneRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *MigrateToOtherZoneRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int64) *MigrateToOtherZoneRequest
	GetDBInstanceStorage() *int64
	SetDBInstanceStorageType(v string) *MigrateToOtherZoneRequest
	GetDBInstanceStorageType() *string
	SetEffectiveTime(v string) *MigrateToOtherZoneRequest
	GetEffectiveTime() *string
	SetIoAccelerationEnabled(v string) *MigrateToOtherZoneRequest
	GetIoAccelerationEnabled() *string
	SetIsModifySpec(v string) *MigrateToOtherZoneRequest
	GetIsModifySpec() *string
	SetOwnerAccount(v string) *MigrateToOtherZoneRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MigrateToOtherZoneRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *MigrateToOtherZoneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MigrateToOtherZoneRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *MigrateToOtherZoneRequest
	GetSwitchTime() *string
	SetVPCId(v string) *MigrateToOtherZoneRequest
	GetVPCId() *string
	SetVSwitchId(v string) *MigrateToOtherZoneRequest
	GetVSwitchId() *string
	SetZoneId(v string) *MigrateToOtherZoneRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *MigrateToOtherZoneRequest
	GetZoneIdSlave1() *string
	SetZoneIdSlave2(v string) *MigrateToOtherZoneRequest
	GetZoneIdSlave2() *string
}

type MigrateToOtherZoneRequest struct {
	// The RDS edition of the instance. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition
	//
	// 	- **HighAvailability**: RDS High-availability Edition
	//
	// 	- **AlwaysOn**: SQL Server on RDS Cluster Edition
	//
	// 	- **cluster**: MySQL on RDS Cluster Edition
	//
	// 	- **Finance**: RDS Enterprise Edition
	//
	// example:
	//
	// HighAvailability
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The new instance type of the instance. You can change the instance type of the instance. You cannot change the storage type of the instance. If you set **IsModifySpec*	- to **true**, you must specify at least one of DBInstanceClass and **DBInstanceStorage**.
	//
	// For more information about instance types, see [Primary ApsaraDB RDS for MySQL instance types](https://help.aliyun.com/document_detail/276975.html).
	//
	// example:
	//
	// mysql.x4.xlarge.2
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The new storage capacity of the instance. If you set **IsModifySpec*	- to **true**, you must specify at least one of DBInstanceStorage and **DBInstanceClass**.
	//
	// Unit: GB. The available storage capacity range varies based on the instance type of the instance. For more information, see [Primary ApsaraDB RDS for MySQL instance types](https://help.aliyun.com/document_detail/276975.html).
	//
	// example:
	//
	// 500
	DBInstanceStorage *int64 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: local SSD. This is the recommended storage type.
	//
	// 	- **general_essd**: general Enterprise SSD (ESSD). This is the recommended storage type.
	//
	// 	- **cloud_essd**: PL1 ESSD
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// 	- **cloud_ssd**: standard SSD. This storage type is not recommended. Standard SSDs are no longer available for purchase in some Alibaba Cloud regions.
	//
	// The default value of this parameter is determined by the instance type specified by the **DBInstanceClass*	- parameter.
	//
	// 	- If the instance type specifies the local SSD storage type, the default value of this parameter is **local_ssd**.
	//
	// 	- If the instance type specifies the standard SSD or ESSD storage type, the default value of this parameter is **cloud_essd**.
	//
	// >  Serverless instances support only PL1 ESSDs and general ESSDs.
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The time when you want the change to take effect. Valid values:
	//
	// 	- **Immediately*	- (default): The change immediately takes effect.
	//
	// 	- **MaintainTime**: The change takes effect during the maintenance window. For more information, see ModifyDBInstanceMaintainTime.
	//
	// 	- **ScheduleTime**: The change takes effect at the point in time that you specify.
	//
	// >  If you set this parameter to **ScheduleTime**, you must specify the **SwitchTime*	- parameter.
	//
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	// Specifies whether to change the specifications of the instance during the cross-zone migration. Valid values:
	//
	// 	- **true**: You want to change the specifications of the instance during the cross-zone migration. If you set this parameter to **true**, you must specify at least one of **DBInstanceClass*	- and **DBInstanceStorage**.
	//
	// 	- **false*	- (default): You do not want to change the specifications of the instance during the cross-zone migration.
	//
	// > This parameter applies only to instances that run MySQL.
	//
	// example:
	//
	// true
	IsModifySpec         *string `json:"IsModifySpec,omitempty" xml:"IsModifySpec,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The migration time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > This parameter is used with **EffectiveTime**. You must specify this parameter only when **EffectiveTime*	- is set to **ScheduleTime**.
	//
	// example:
	//
	// 2021-12-14T15:15:15Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The ID of the virtual private cloud (VPC). Do not change the VPC of the instance when you migrate the instance across zones.
	//
	// 	- This parameter must be specified when the instance resides in a VPC.
	//
	// 	- If the instance runs SQL Server, you can change the VPC of the instance.
	//
	// example:
	//
	// vpc-xxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// 	- This parameter must be specified when the instance resides in a VPC. You can call the DescribeVSwitches operation to query existing vSwitches.
	//
	// 	- If the instance runs PostgreSQL or SQL Server and a secondary zone is specified for the instance, you can specify multiple vSwitch IDs, each of which corresponds to a zone. Separate the vSwitch IDs with commas (,).
	//
	// example:
	//
	// vsw-uf6adz52c2pxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the destination zone. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The secondary zone 1 of the instance.
	//
	// >  This parameter must be configured if the instance runs RDS editions other than RDS Basic Edition.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// The secondary zone 2 of the instance.
	//
	// >  You can specify this parameter only for instances that run RDS Enterprise Edition.
	//
	// example:
	//
	// cn-hangzhou-d
	ZoneIdSlave2 *string `json:"ZoneIdSlave2,omitempty" xml:"ZoneIdSlave2,omitempty"`
}

func (s MigrateToOtherZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateToOtherZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneRequest) GetCategory() *string {
	return s.Category
}

func (s *MigrateToOtherZoneRequest) GetCustomExtraInfo() *string {
	return s.CustomExtraInfo
}

func (s *MigrateToOtherZoneRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *MigrateToOtherZoneRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateToOtherZoneRequest) GetDBInstanceStorage() *int64 {
	return s.DBInstanceStorage
}

func (s *MigrateToOtherZoneRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *MigrateToOtherZoneRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *MigrateToOtherZoneRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *MigrateToOtherZoneRequest) GetIsModifySpec() *string {
	return s.IsModifySpec
}

func (s *MigrateToOtherZoneRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MigrateToOtherZoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateToOtherZoneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MigrateToOtherZoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateToOtherZoneRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *MigrateToOtherZoneRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *MigrateToOtherZoneRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *MigrateToOtherZoneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *MigrateToOtherZoneRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *MigrateToOtherZoneRequest) GetZoneIdSlave2() *string {
	return s.ZoneIdSlave2
}

func (s *MigrateToOtherZoneRequest) SetCategory(v string) *MigrateToOtherZoneRequest {
	s.Category = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetCustomExtraInfo(v string) *MigrateToOtherZoneRequest {
	s.CustomExtraInfo = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetDBInstanceClass(v string) *MigrateToOtherZoneRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetDBInstanceId(v string) *MigrateToOtherZoneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetDBInstanceStorage(v int64) *MigrateToOtherZoneRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetDBInstanceStorageType(v string) *MigrateToOtherZoneRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetEffectiveTime(v string) *MigrateToOtherZoneRequest {
	s.EffectiveTime = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetIoAccelerationEnabled(v string) *MigrateToOtherZoneRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetIsModifySpec(v string) *MigrateToOtherZoneRequest {
	s.IsModifySpec = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetOwnerAccount(v string) *MigrateToOtherZoneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetOwnerId(v int64) *MigrateToOtherZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetResourceOwnerAccount(v string) *MigrateToOtherZoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetResourceOwnerId(v int64) *MigrateToOtherZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetSwitchTime(v string) *MigrateToOtherZoneRequest {
	s.SwitchTime = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetVPCId(v string) *MigrateToOtherZoneRequest {
	s.VPCId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetVSwitchId(v string) *MigrateToOtherZoneRequest {
	s.VSwitchId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetZoneId(v string) *MigrateToOtherZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetZoneIdSlave1(v string) *MigrateToOtherZoneRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetZoneIdSlave2(v string) *MigrateToOtherZoneRequest {
	s.ZoneIdSlave2 = &v
	return s
}

func (s *MigrateToOtherZoneRequest) Validate() error {
	return dara.Validate(s)
}
