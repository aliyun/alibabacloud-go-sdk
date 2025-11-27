// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesForCloneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBInstancesForCloneResponseBodyItems) *DescribeDBInstancesForCloneResponseBody
	GetItems() *DescribeDBInstancesForCloneResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstancesForCloneResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBInstancesForCloneResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBInstancesForCloneResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBInstancesForCloneResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBInstancesForCloneResponseBody struct {
	// An array that consists of the details about the instances.
	Items *DescribeDBInstancesForCloneResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 12
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 120
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstancesForCloneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesForCloneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForCloneResponseBody) GetItems() *DescribeDBInstancesForCloneResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstancesForCloneResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesForCloneResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBInstancesForCloneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancesForCloneResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBInstancesForCloneResponseBody) SetItems(v *DescribeDBInstancesForCloneResponseBodyItems) *DescribeDBInstancesForCloneResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBody) SetPageNumber(v int32) *DescribeDBInstancesForCloneResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBody) SetPageRecordCount(v int32) *DescribeDBInstancesForCloneResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBody) SetRequestId(v string) *DescribeDBInstancesForCloneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstancesForCloneResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesForCloneResponseBodyItems struct {
	DBInstance []*DescribeDBInstancesForCloneResponseBodyItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesForCloneResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesForCloneResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForCloneResponseBodyItems) GetDBInstance() []*DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	return s.DBInstance
}

func (s *DescribeDBInstancesForCloneResponseBodyItems) SetDBInstance(v []*DescribeDBInstancesForCloneResponseBodyItemsDBInstance) *DescribeDBInstancesForCloneResponseBodyItems {
	s.DBInstance = v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItems) Validate() error {
	if s.DBInstance != nil {
		for _, item := range s.DBInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesForCloneResponseBodyItemsDBInstance struct {
	// The RDS edition of the instance. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition
	//
	// 	- **HighAvailability**: RDS High-availability Edition
	//
	// 	- **Finance**: RDS Enterprise Edition
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The connection mode of the instance. Valid values:
	//
	// 	- **Standard**: standard mode
	//
	// 	- **Safe**: database proxy mode
	//
	// example:
	//
	// Standard
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The time when the instance was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-11-05T11:26:02Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance type of the instance. For more information, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// mysql.n1.micro.1
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The name of the instance. It must be 2 to 256 characters in length. The value can contain letters, digits, underscores (_), and hyphens (-). The value must start with a letter.
	//
	// > The value cannot start with http:// or https://.
	//
	// example:
	//
	// Test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network connection type of the instance. Valid values:
	//
	// 	- **Internet**
	//
	// 	- **Intranet**
	//
	// example:
	//
	// Intranet
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The status of the instance. For more information, see [Instance state table](https://help.aliyun.com/document_detail/26315.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd/ephemeral_ssd**: local SSD
	//
	// 	- **cloud_ssd**: standard SSD.
	//
	// 	- **cloud_essd**: enhanced SSD (ESSD)
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The role of the instance. Valid values:
	//
	// 	- **Primary**: primary instance
	//
	// 	- **Readonly**: read-only instance
	//
	// 	- **Guard**: disaster recovery instance
	//
	// 	- **Temp**: temporary instance
	//
	// example:
	//
	// Primary
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The time when the instance was destroyed. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-05T11:26:02Z
	DestroyTime *string `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	// The database engine of the instance. Valid values:
	//
	// 	- MySQL
	//
	// 	- SQLServer
	//
	// 	- PostgreSQL
	//
	// 	- PPAS
	//
	// 	- MariaDB
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the instance expired. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-11-28T11:26:02Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the disaster recovery instance. This parameter is returned only when the instance is a primary instance and has a disaster recovery instance.
	//
	// example:
	//
	// rm-uf64zsuxxxxxxxxxx
	GuardDBInstanceId *string `json:"GuardDBInstanceId,omitempty" xml:"GuardDBInstanceId,omitempty"`
	// The ID of the instance role.
	//
	// example:
	//
	// 1
	InsId *int32 `json:"InsId,omitempty" xml:"InsId,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **Classic**
	//
	// 	- **VPC**
	//
	// example:
	//
	// Classic
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The lock method of the instance. Valid values:
	//
	// 	- **Unlock**: The instance is not locked.
	//
	// 	- **ManualLock**: The instance is manually locked.
	//
	// 	- **LockByExpiration**: The instance is automatically locked after it expires.
	//
	// 	- **LockByRestoration**: The instance is automatically locked before a rollback.
	//
	// 	- **LockByDiskQuota**: The instance is automatically locked because its storage capacity is exhausted and the instance is inaccessible.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the instance was locked.
	//
	// example:
	//
	// instance_expired
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The ID of the primary instance. If the value of this parameter is null, the instance is a primary instance.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	// Indicates whether multi-region deployment is used. Valid values:
	//
	// 	- **true**: Multi-region deployment is used.
	//
	// 	- **false**: Multi-region deployment is not used.
	//
	// example:
	//
	// false
	MutriORsignle *bool `json:"MutriORsignle,omitempty" xml:"MutriORsignle,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// An array consisting of the IDs of the read-only instances that are attached to the primary instance.
	ReadOnlyDBInstanceIds *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIds `json:"ReadOnlyDBInstanceIds,omitempty" xml:"ReadOnlyDBInstanceIds,omitempty" type:"Struct"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// None.
	//
	// example:
	//
	// N/A
	ReplicateId *string `json:"ReplicateId,omitempty" xml:"ReplicateId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the temporary instance.
	//
	// example:
	//
	// sub138xxxxx_rm-xxxxx
	TempDBInstanceId *string `json:"TempDBInstanceId,omitempty" xml:"TempDBInstanceId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-uf6adz52c2pxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the instance in the VPC.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1opxu1zkhxxxxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesForCloneResponseBodyItemsDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetDestroyTime() *string {
	return s.DestroyTime
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetGuardDBInstanceId() *string {
	return s.GuardDBInstanceId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetInsId() *int32 {
	return s.InsId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetMasterInstanceId() *string {
	return s.MasterInstanceId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetMutriORsignle() *bool {
	return s.MutriORsignle
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetReadOnlyDBInstanceIds() *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIds {
	return s.ReadOnlyDBInstanceIds
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetReplicateId() *string {
	return s.ReplicateId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetTempDBInstanceId() *string {
	return s.TempDBInstanceId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetVpcCloudInstanceId() *string {
	return s.VpcCloudInstanceId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetCategory(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.Category = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetConnectionMode(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetCreateTime(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetDBInstanceClass(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetDBInstanceDescription(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetDBInstanceId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetDBInstanceNetType(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetDBInstanceStorageType(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.DBInstanceStorageType = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetDBInstanceType(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetDestroyTime(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.DestroyTime = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetEngine(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetEngineVersion(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetExpireTime(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetGuardDBInstanceId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.GuardDBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetInsId(v int32) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.InsId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetLockMode(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetLockReason(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetMasterInstanceId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetMutriORsignle(v bool) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.MutriORsignle = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetPayType(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetReadOnlyDBInstanceIds(v *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.ReadOnlyDBInstanceIds = v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetRegionId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetReplicateId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.ReplicateId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetResourceGroupId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetTempDBInstanceId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.TempDBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetVSwitchId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetVpcCloudInstanceId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetVpcId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) SetZoneId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstance) Validate() error {
	if s.ReadOnlyDBInstanceIds != nil {
		if err := s.ReadOnlyDBInstanceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIds struct {
	ReadOnlyDBInstanceId []*DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId `json:"ReadOnlyDBInstanceId,omitempty" xml:"ReadOnlyDBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) GetReadOnlyDBInstanceId() []*DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId {
	return s.ReadOnlyDBInstanceId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) SetReadOnlyDBInstanceId(v []*DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIds {
	s.ReadOnlyDBInstanceId = v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIds) Validate() error {
	if s.ReadOnlyDBInstanceId != nil {
		for _, item := range s.ReadOnlyDBInstanceId {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId struct {
	// The ID of the read-only instance.
	//
	// example:
	//
	// rm-bpxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) SetDBInstanceId(v string) *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForCloneResponseBodyItemsDBInstanceReadOnlyDBInstanceIdsReadOnlyDBInstanceId) Validate() error {
	return dara.Validate(s)
}
