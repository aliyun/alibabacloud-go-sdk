// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceMajorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowDDL(v bool) *UpgradeDBInstanceMajorVersionRequest
	GetAllowDDL() *bool
	SetCollectStatMode(v string) *UpgradeDBInstanceMajorVersionRequest
	GetCollectStatMode() *string
	SetCustomExtraInfo(v string) *UpgradeDBInstanceMajorVersionRequest
	GetCustomExtraInfo() *string
	SetDBInstanceClass(v string) *UpgradeDBInstanceMajorVersionRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *UpgradeDBInstanceMajorVersionRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *UpgradeDBInstanceMajorVersionRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *UpgradeDBInstanceMajorVersionRequest
	GetDBInstanceStorageType() *string
	SetInstanceNetworkType(v string) *UpgradeDBInstanceMajorVersionRequest
	GetInstanceNetworkType() *string
	SetPayType(v string) *UpgradeDBInstanceMajorVersionRequest
	GetPayType() *string
	SetPeriod(v string) *UpgradeDBInstanceMajorVersionRequest
	GetPeriod() *string
	SetPrivateIpAddress(v string) *UpgradeDBInstanceMajorVersionRequest
	GetPrivateIpAddress() *string
	SetResourceOwnerId(v int64) *UpgradeDBInstanceMajorVersionRequest
	GetResourceOwnerId() *int64
	SetSwitchOver(v string) *UpgradeDBInstanceMajorVersionRequest
	GetSwitchOver() *string
	SetSwitchTime(v string) *UpgradeDBInstanceMajorVersionRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *UpgradeDBInstanceMajorVersionRequest
	GetSwitchTimeMode() *string
	SetTargetMajorVersion(v string) *UpgradeDBInstanceMajorVersionRequest
	GetTargetMajorVersion() *string
	SetUpgradeMode(v string) *UpgradeDBInstanceMajorVersionRequest
	GetUpgradeMode() *string
	SetUsedTime(v string) *UpgradeDBInstanceMajorVersionRequest
	GetUsedTime() *string
	SetVPCId(v string) *UpgradeDBInstanceMajorVersionRequest
	GetVPCId() *string
	SetVSwitchId(v string) *UpgradeDBInstanceMajorVersionRequest
	GetVSwitchId() *string
	SetZoneId(v string) *UpgradeDBInstanceMajorVersionRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *UpgradeDBInstanceMajorVersionRequest
	GetZoneIdSlave1() *string
	SetZoneIdSlave2(v string) *UpgradeDBInstanceMajorVersionRequest
	GetZoneIdSlave2() *string
}

type UpgradeDBInstanceMajorVersionRequest struct {
	AllowDDL *bool `json:"AllowDDL,omitempty" xml:"AllowDDL,omitempty"`
	// Specifies when to execute statistics information collection on the database.
	//
	// - **Before**: Execute collection before the switchover. This ensures business stability. If the instance has a large data volume, the upgrade may take a long time.
	//
	// - **After**: Execute collection after the switchover. The upgrade is faster. Accessing tables without generated statistics information after the upgrade may cause inaccurate execution plans. During peak hours, this may cause the database to break down.
	//
	// > For non-switchover scenarios, "before switchover" means statistics information is collected before the new instance is opened for read/write, and "after switchover" means statistics information is collected after the new instance is opened for read/write.
	//
	// example:
	//
	// After
	CollectStatMode *string `json:"CollectStatMode,omitempty" xml:"CollectStatMode,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The instance type after the upgrade. The CPU and memory configurations must be greater than or equal to those of the original instance type. If **UpgradeMode*	- is set to **inPlaceUpgrade*	- or **zeroDownTimeUpgrade**, **you do not need to configure*	- this parameter.
	//
	// For example, if the original instance type is `pg.n2.small.2c` with 1 CPU core and 2 GB of memory, you can upgrade it to `pg.n2.medium.2c` with 2 CPU cores and 4 GB of memory.
	//
	// > For the instance type codes of ApsaraDB RDS for PostgreSQL, refer to [Primary ApsaraDB RDS for PostgreSQL instance types](https://help.aliyun.com/document_detail/276990.html).
	//
	// example:
	//
	// pg.n2.medium.2c
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID of the original instance.
	//
	// example:
	//
	// pgm-bp1gm3yh0ht1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance storage capacity after the upgrade. Unit: GB. If **UpgradeMode*	- (upgrade pattern) is set to **inPlaceUpgrade*	- or **zeroDownTimeUpgrade**, **you do not need to configure*	- this parameter.
	//
	// Valid values:
	//
	// - **PL1 ESSD cloud disk**: 20 GB to 3200 GB
	//
	// - **PL2 ESSD cloud disk**: 500 GB to 3200 GB
	//
	// - **PL3 ESSD cloud disk**: 1500 GB to 3200 GB
	//
	// - **Premium performance disk**: 40 GB to 2000 GB
	//
	// > When upgrading the major engine version of an instance with Premium Local SSDs, storage capacity reduction is supported. For the minimum storage capacity, refer to [Upgrade the major engine version of a database](https://help.aliyun.com/document_detail/203309.html).
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance after the upgrade.
	//
	// Valid values:
	//
	// - **cloud_ssd**: standard SSD
	//
	// - **cloud_essd**: PL1 ESSD
	//
	// - **cloud_essd2**: PL2 ESSD
	//
	// - **cloud_essd3**: PL3 ESSD
	//
	// - **general_essd**: premium performance disk
	//
	//
	// The major engine version upgrade feature is based on cloud disk snapshots. The supported storage types after the upgrade are as follows:
	//
	// - If the original instance uses a standard SSD, you can select standard SSD.
	//
	// - If the original instance uses an ESSD cloud disk, you can select PL1 ESSD, PL2 ESSD, PL3 ESSD, or premium performance disk.
	//
	// - If the original instance uses Premium Local SSDs, you can select PL1 ESSD, PL2 ESSD, PL3 ESSD, or premium performance disk.
	//
	// example:
	//
	// cloud_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The network type of the instance after the upgrade. Set this parameter to VPC. Only VPC-connected instances support major engine version upgrades.
	//
	// If the network type is classic network, switch to VPC first. For information about how to view or switch the network type, refer to [Switch the network type](https://help.aliyun.com/document_detail/96761.html).
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The billing method of the instance. Set this parameter to Postpaid for pay-as-you-go billing.
	//
	// > If you want to change the billing method after the upgrade, refer to [Switch from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/96743.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Reserved parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// You do not need to configure this parameter. It specifies the internal IP address of the target instance. The system automatically assigns an IP address based on VPCId and vSwitchId by default.
	//
	// example:
	//
	// 172.16.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The switchover configuration. Specifies whether to switch traffic to the new version instance based on your business requirements.
	//
	// Valid values:
	//
	// - **true**: Switchover is performed and automatic switchover is enabled. This option is typically used to execute the formal upgrade after confirming that your business can run stably on the new version.
	//
	// - **false**: Switchover is not performed and automatic switchover is not enabled. This option is typically used to test the compatibility of your application with the new version before the formal upgrade.
	//
	// > - If you select switchover:
	//
	// >     - Switchover cannot be rolled back after execution. Proceed with caution.
	//
	// >     - During the switchover procedure, the original instance becomes read-only and writes are not allowed. Execute the switchover during off-peak hours.
	//
	// >     - If read-only instances are created for the original instance, you cannot select switchover. You can only upgrade the instance without switchover, and the original read-only instances are not cloned. After the upgrade, create new PostgreSQL read-only instances for the new version instance.
	//
	// > - If you do not select switchover:
	//
	// >     - The business on the original instance is not affected during migration.
	//
	// >     - To upgrade the instance without switchover, change the database connection address in your application to the database connection address of the new instance after migration is complete. For information about how to view the connection address, refer to [View or modify the internal and public endpoints and port numbers](https://help.aliyun.com/document_detail/96788.html).
	//
	// example:
	//
	// false
	SwitchOver *string `json:"SwitchOver,omitempty" xml:"SwitchOver,omitempty"`
	// Reserved parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// 2021-07-10T13:15:12Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// This parameter is used together with SwitchOver and takes effect only when **SwitchOver*	- is set to **true**. Specifies the switchover time.
	//
	// Valid values:
	//
	// - **Immediate**: The switchover takes effect immediately.
	//
	// - **MaintainTime**: The switchover takes effect during the maintenance window. You can call the ModifyDBInstanceMaintainTime operation to modify the maintenance window.
	//
	// example:
	//
	// Immediate
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	// The target major engine version of the instance after the upgrade. This value must be the same as the target version specified during the pre-upgrade check.
	//
	// > You can call the UpgradeDBInstanceMajorVersionPrecheck operation to perform a pre-upgrade check for the major engine version upgrade.
	//
	// example:
	//
	// 13.0
	TargetMajorVersion *string `json:"TargetMajorVersion,omitempty" xml:"TargetMajorVersion,omitempty"`
	// The upgrade pattern. Configure this parameter when **SwitchOver*	- is set to **true**. Valid values:
	//
	// - **inPlaceUpgrade**: In-place upgrade. The major engine version upgrade task is executed on the original instance without creating a new version instance. After the upgrade, the original instance inherits the existing order, instance name, tags, CloudMonitor alert rules, and backup rules.
	//
	// - **blueGreenDeployment**: Blue-green deployment. The major engine version upgrade retains the original instance and creates a new version instance. The new instance is free of charge during creation. After the new instance is created, fees are incurred and the billing method may change. After the upgrade, both the original and new instances incur fees, and the new instance does not inherit the discounts of the original instance.
	//
	// - **zeroDownTimeUpgrade**: Zero-downtime upgrade. The system uses pg_upgrade to upgrade the original instance to the target version and uses native logical replication for incremental updates. Active switchover is supported during the upgrade procedure, and you can validate the higher version instance before the switchover. From the start of the upgrade until the active switchover, the instance maintains normal read/write operations. During the switchover, the read-only duration is at the second level.
	//
	// example:
	//
	// inPlaceUpgrade
	UpgradeMode *string `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
	// Reserved parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID. If **UpgradeMode*	- is set to **inPlaceUpgrade*	- or **zeroDownTimeUpgrade**, **you do not need to configure*	- this parameter.
	//
	// You can call the DescribeDBInstanceAttribute operation to query the VPC ID of the original instance.
	//
	// example:
	//
	// vpc-bp1opxu1zkhn00gzv****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the target instance. If **UpgradeMode*	- (upgrade pattern) is set to **inPlaceUpgrade*	- or **zeroDownTimeUpgrade**, **you do not need to configure*	- this parameter.
	//
	// - If the original instance is a Basic Edition instance, specify the vSwitch ID of the target instance.
	//
	// - If the original instance is a high-availability series instance, you can specify the vSwitch IDs of the target primary and secondary instances, separated by commas (,).
	//
	// > The target vSwitch must be in the same zone as the original instance. You can call the DescribeVSwitches operation to query vSwitches.
	//
	// example:
	//
	// vsw-bp10aqj6o4lclxdrm****,vsw-bp10aqj6o4lclxdrm****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The primary zone ID of the target instance. If **UpgradeMode*	- is set to **inPlaceUpgrade*	- or **zeroDownTimeUpgrade**, **you do not need to configure*	- this parameter.
	//
	// You can call the DescribeRegions operation to query zone IDs.
	//
	// ApsaraDB RDS for PostgreSQL allows you to deploy the new instance in a different zone within the same region as the original instance after the upgrade.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// This parameter can be configured only when the original instance is a high-availability series instance. Specifies the secondary zone ID of the target instance. If **UpgradeMode*	- (upgrade pattern) is set to **inPlaceUpgrade*	- or **zeroDownTimeUpgrade**, **you do not need to configure*	- this parameter.
	//
	// ApsaraDB RDS for PostgreSQL allows you to deploy the new secondary instance in a different zone within the same region as the original instance after the upgrade.
	//
	// You can call the DescribeRegions operation to query zone IDs.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// Reserved parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneIdSlave2 *string `json:"ZoneIdSlave2,omitempty" xml:"ZoneIdSlave2,omitempty"`
}

func (s UpgradeDBInstanceMajorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceMajorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetAllowDDL() *bool {
	return s.AllowDDL
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetCollectStatMode() *string {
	return s.CollectStatMode
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetCustomExtraInfo() *string {
	return s.CustomExtraInfo
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetPayType() *string {
	return s.PayType
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetPeriod() *string {
	return s.Period
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetSwitchOver() *string {
	return s.SwitchOver
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetTargetMajorVersion() *string {
	return s.TargetMajorVersion
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetUpgradeMode() *string {
	return s.UpgradeMode
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *UpgradeDBInstanceMajorVersionRequest) GetZoneIdSlave2() *string {
	return s.ZoneIdSlave2
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetAllowDDL(v bool) *UpgradeDBInstanceMajorVersionRequest {
	s.AllowDDL = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetCollectStatMode(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.CollectStatMode = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetCustomExtraInfo(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.CustomExtraInfo = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetDBInstanceClass(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetDBInstanceId(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetDBInstanceStorage(v int32) *UpgradeDBInstanceMajorVersionRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetDBInstanceStorageType(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetInstanceNetworkType(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetPayType(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.PayType = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetPeriod(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.Period = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetPrivateIpAddress(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceMajorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetSwitchOver(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.SwitchOver = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetSwitchTime(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetSwitchTimeMode(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetTargetMajorVersion(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.TargetMajorVersion = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetUpgradeMode(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.UpgradeMode = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetUsedTime(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.UsedTime = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetVPCId(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.VPCId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetVSwitchId(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetZoneId(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.ZoneId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetZoneIdSlave1(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) SetZoneIdSlave2(v string) *UpgradeDBInstanceMajorVersionRequest {
	s.ZoneIdSlave2 = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionRequest) Validate() error {
	return dara.Validate(s)
}
