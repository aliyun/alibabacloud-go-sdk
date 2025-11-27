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
	// Specify the point in time at which the system collects the statistics of the instance.
	//
	// 	- **Before**: The system collects the statistics of the instance before the switchover to ensure service stability. If the instance contains a large amount of data, the upgrade may require a long period of time.
	//
	// 	- **After**: The system collects the statistics of the instance after the switchover to accelerate the upgrade. After the upgrade, if you access tables for which no statistics are generated, the query plans may be inaccurate, and your database service may be unavailable during peak hours.
	//
	// >  If you set the SwitchOver parameter to false, the value Before specifies that the system collects the statistics of the instance before the instance starts to process read and write requests, and the value After specifies that the system collects the statistics of the instance after the instance starts to process read and write requests.
	//
	// example:
	//
	// After
	CollectStatMode *string `json:"CollectStatMode,omitempty" xml:"CollectStatMode,omitempty"`
	CustomExtraInfo *string `json:"CustomExtraInfo,omitempty" xml:"CustomExtraInfo,omitempty"`
	// The new instance type of the instance. The new CPU and memory specifications of the instance must be higher than or equal to the original CPU and memory specifications. If you set the **UpgradeMode*	- parameter to **inPlaceUpgrade**, you **do not need to configure*	- this parameter.
	//
	// For example, you can upgrade the instance type from `pg.n2.small.2c` to `pg.n2.medium.2c`. The pg.n2.small.2c instance type provides 1 CPU core and 2 GB of memory. The pg.n2.medium.2c instance type provides 2 CPU cores and 4 GB of memory.
	//
	// >  For more information about the instance types of ApsaraDB RDS for PostgreSQL instances, see [Instance types for primary ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/276990.html).
	//
	// example:
	//
	// pg.n2.medium.2c
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The ID of the original instance.
	//
	// example:
	//
	// pgm-bp1gm3yh0ht1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The new storage capacity of the instance. Unit: GB If you set the **UpgradeMode*	- parameter to **inPlaceUpgrade**, you **do not need to configure*	- this parameter.
	//
	// Valid values:
	//
	// 	- **PL1 ESSD**: 20 GB to 32,000 GB
	//
	// 	- **PL2 ESSD**: 500 GB to 3,200 GB
	//
	// 	- **PL3 ESSD**: 1,500 GB to 3,200 GB
	//
	// 	- **General ESSD**: 40 GB to 2,000 GB
	//
	// >  If the original instance uses local disks, you can reduce the storage capacity of the instance when you upgrade the major engine version of the instance. For more information about the minimum storage capacity, see [Upgrade the major engine version](https://help.aliyun.com/document_detail/203309.html).
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance that runs the required major engine version.
	//
	// Valid values:
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd**: performance level 1 (PL1) Enterprise SSD (ESSD)
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// 	- **general_essd**: general ESSD
	//
	// The major engine version upgrade feature is developed based on snapshots for cloud disks. You can select a storage type after the upgrade based on the following items:
	//
	// 	- If the original instance uses standard SSDs, set this parameter to cloud_ssd.
	//
	// 	- If the original instance uses ESSDs, set this parameter to cloud_essd, cloud_essd2, cloud_essd3, or general_essd.
	//
	// 	- If the original instance uses local SSDs, set this parameter to cloud_essd, cloud_essd2, cloud_essd3, or general_essd.
	//
	// example:
	//
	// cloud_essd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The network type of the new instance. Set the value to VPC. The major engine version upgrade feature is supported only for instances that reside in VPCs.
	//
	// If the original instance resides in the classic network, you must migrate the instance to a VPC before you call this operation. For more information about how to view or change the network type of an instance, see [Change the network type of an ApsaraDB RDS for PostgreSQL instance](https://help.aliyun.com/document_detail/96761.html).
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The billing method. Set the value to Postpaid.
	//
	// >  For more information about how to change the billing method of an instance after the upgrade, see [Change the billing method of an instance from pay-as-you-go to subscription](https://help.aliyun.com/document_detail/96743.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The internal IP address of the new instance. You do not need to specify this parameter. The system automatically assigns an internal IP address based on the values of the VPCId and vSwitchId parameters.
	//
	// example:
	//
	// 172.16.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to switch your workloads over to the instance that runs the required major engine version based on your business requirements.
	//
	// Valid values:
	//
	// 	- **true**: The system automatically switches workloads over to the instance. This configuration method is used to perform an upgrade after you verify that the new major engine version is compatible with your workloads.
	//
	// 	- **false**: The system does not automatically switch your workloads over to the instance. In most cases, this configuration method is used to test whether the new major engine version is compatible with your workloads before you perform the upgrade.
	//
	// >
	//
	// 	- If you set this parameter to true, you must take note of the following items:
	//
	//     	- After the switchover is complete, you cannot roll your workloads back to the original instance. Proceed with caution.
	//
	//     	- During the switchover, the original instance processes only read requests. We recommend that you perform the switchover during off-peak hours.
	//
	//     	- If read-only instances are attached to the original instance, you can set this parameter only to false. In this case, the read-only instances that are attached to the original instance cannot be cloned. After the upgrade is complete, you must create read-only instances for the instance.
	//
	// 	- If you set this parameter to false, you must take note of the following items:
	//
	//     	- The data migration does not interrupt your workloads on the original instance.
	//
	//     	- After data is migrated to the instance that runs the required major engine version, you must update the endpoint configuration in your application. This update requires you to replace the endpoint of the original instance with the endpoint of the instance that runs the required major engine version. For more information about how to view the endpoint of an instance, see [Viewing and change of the internal and public endpoints and port numbers](https://help.aliyun.com/document_detail/96788.html).
	//
	// example:
	//
	// false
	SwitchOver *string `json:"SwitchOver,omitempty" xml:"SwitchOver,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// 2021-07-10T13:15:12Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The point in time at which the workloads are switched over. This parameter is used together with the SwitchOver parameter. This parameter is available only when you set the **SwitchOver*	- parameter to **true**.
	//
	// Valid values:
	//
	// 	- **Immediate**: The workloads are immediately switched over.
	//
	// 	- **MaintainTime**: The workloads are switched over within the maintenance window that you specify. You can call the ModifyDBInstanceMaintainTime operation to change the maintenance window of an instance.
	//
	// example:
	//
	// Immediate
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	// The major engine version of the new instance. The value of this parameter must be the major engine version on which an upgrade check is performed.
	//
	// >  You can call the UpgradeDBInstanceMajorVersionPrecheck operation to perform an upgrade check.
	//
	// example:
	//
	// 13.0
	TargetMajorVersion *string `json:"TargetMajorVersion,omitempty" xml:"TargetMajorVersion,omitempty"`
	// The upgrade mode. This parameter is required when you set the **SwitchOver*	- parameter to **true**. Valid values:
	//
	// 	- **inPlaceUpgrade**: local upgrade. The major engine version upgrade is performed on the original instance, and no new instance is created. After the upgrade, the original instance runs the required major engine version and inherits the original orders, name, tags, alert rules in CloudMonitor, and backup settings.
	//
	// 	- **blueGreenDeployment**: blue-green deployment. After the major engine version of the instance is upgraded, the original instance is retained and a new instance is created. Fees are generated for the new instance based on the billing method that you specified. However, no fees are generated for the creation of the new instance. After the upgrade is complete, fees are generated for both the original and new instances and the new instance cannot enjoy the discounts provided for the original instance.
	//
	// example:
	//
	// inPlaceUpgrade
	UpgradeMode *string `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The virtual private cloud (VPC) ID of the instance. If you set the **UpgradeMode*	- parameter to **inPlaceUpgrade**, you **do not need to configure*	- this parameter.
	//
	// You can call the DescribeDBInstanceAttribute operation to query the VPC ID of the original instance.
	//
	// example:
	//
	// vpc-bp1opxu1zkhn00gzv****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the instance that runs the required major engine version. If you set the **UpgradeMode*	- parameter to **inPlaceUpgrade**, you **do not need to configure*	- this parameter.
	//
	// 	- If the original instance runs RDS Basic Edition, configure the vSwitch ID for the instance that runs the required major engine version.
	//
	// 	- If the original instance runs RDS High-availability Edition, configure the vSwitch IDs for the instance that runs the required major engine version and its secondary instance. Separate the vSwitch IDs with commas (,).
	//
	// >  The vSwitches that you specify must reside in the same zone as the original instance. You can call the DescribeVSwitches operation to query the vSwitch IDs.
	//
	// example:
	//
	// vsw-bp10aqj6o4lclxdrm****,vsw-bp10aqj6o4lclxdrm****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone to which the primary instance that runs the required major engine version belongs. If you set the **UpgradeMode*	- parameter to **inPlaceUpgrade**, you **do not need to configure*	- this parameter.
	//
	// You can call the DescribeRegions operation to query zone IDs.
	//
	// You can select a zone that belongs to the region in which the original instance resides.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The ID of the zone to which the secondary instance runs the required major engine version belongs. This parameter is available only when the original instance runs RDS High-availability Edition. If you set the **UpgradeMode*	- parameter to **inPlaceUpgrade**, you **do not need to configure*	- this parameter.
	//
	// You can select a zone that belongs to the region in which the original instance resides.
	//
	// You can call the DescribeRegions operation to query zone IDs.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneIdSlave1 *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// cn-hangzhou-h
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
