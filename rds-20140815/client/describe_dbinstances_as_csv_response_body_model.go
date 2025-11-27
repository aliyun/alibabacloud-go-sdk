// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesAsCsvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBInstancesAsCsvResponseBodyItems) *DescribeDBInstancesAsCsvResponseBody
	GetItems() *DescribeDBInstancesAsCsvResponseBodyItems
	SetRequestId(v string) *DescribeDBInstancesAsCsvResponseBody
	GetRequestId() *string
}

type DescribeDBInstancesAsCsvResponseBody struct {
	// An array that consists of the fields in **DBInstanceAttribute**.
	Items *DescribeDBInstancesAsCsvResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A444291****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstancesAsCsvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesAsCsvResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesAsCsvResponseBody) GetItems() *DescribeDBInstancesAsCsvResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstancesAsCsvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancesAsCsvResponseBody) SetItems(v *DescribeDBInstancesAsCsvResponseBodyItems) *DescribeDBInstancesAsCsvResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBody) SetRequestId(v string) *DescribeDBInstancesAsCsvResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesAsCsvResponseBodyItems struct {
	DBInstanceAttribute []*DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesAsCsvResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesAsCsvResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesAsCsvResponseBodyItems) GetDBInstanceAttribute() []*DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	return s.DBInstanceAttribute
}

func (s *DescribeDBInstancesAsCsvResponseBodyItems) SetDBInstanceAttribute(v []*DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) *DescribeDBInstancesAsCsvResponseBodyItems {
	s.DBInstanceAttribute = v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItems) Validate() error {
	if s.DBInstanceAttribute != nil {
		for _, item := range s.DBInstanceAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute struct {
	// The maximum number of accounts.
	//
	// example:
	//
	// 500
	AccountMaxQuantity *int32 `json:"AccountMaxQuantity,omitempty" xml:"AccountMaxQuantity,omitempty"`
	// The type of the account.
	//
	// example:
	//
	// super
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The service availability of the instance in percentage.
	//
	// example:
	//
	// 100
	AvailabilityValue *string `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty"`
	// The category of the instance.
	//
	// example:
	//
	// 0
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The connection mode of the instance. Valid values:
	//
	// 	- **Performance**: standard mode.
	//
	// 	- **Safety**: enhanced mode
	//
	// example:
	//
	// Performance
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2011-05-30T12:11:04Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 2
	DBInstanceCPU *string `json:"DBInstanceCPU,omitempty" xml:"DBInstanceCPU,omitempty"`
	// The instance type of the instance.
	//
	// example:
	//
	// rds.mys2.small
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance family.
	//
	// example:
	//
	// s
	DBInstanceClassType *string `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	// The instance description.
	//
	// example:
	//
	// 0
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The memory capacity of the instance. Unit: MB.
	//
	// example:
	//
	// 4096
	DBInstanceMemory *int64 `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **Internet**
	//
	// 	- **Intranet**
	//
	// example:
	//
	// Internet
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The storage capacity of the instance. Unit: GB.
	//
	// example:
	//
	// 10
	DBInstanceStorage     *int32  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The instance type. Valid values:
	//
	// 	- **Primary**: primary instance
	//
	// 	- **ReadOnly**: read-only instance
	//
	// 	- **Guard**: disaster recovery instance
	//
	// 	- **Temp**: temporary instance
	//
	// example:
	//
	// Primary
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The maximum number of databases that can be created on the instance.
	//
	// example:
	//
	// 200
	DBMaxQuantity *int32 `json:"DBMaxQuantity,omitempty" xml:"DBMaxQuantity,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2019-03-27T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// API
	ExportKey *string `json:"ExportKey,omitempty" xml:"ExportKey,omitempty"`
	// The ID of the disaster recovery instance that is attached to the primary instance.
	//
	// example:
	//
	// rm-uf64zsuxxxxxxxxxx
	GuardDBInstanceId *string `json:"GuardDBInstanceId,omitempty" xml:"GuardDBInstanceId,omitempty"`
	// The ID of the instance from which incremental data comes. The incremental data of a disaster recovery instance comes from its primary instance. The incremental data of a read-only instance comes from its primary instance. If this parameter is not returned, the instance is a primary instance.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	IncrementSourceDBInstanceId *string `json:"IncrementSourceDBInstanceId,omitempty" xml:"IncrementSourceDBInstanceId,omitempty"`
	// The network type.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The lock mode of the instance.
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
	// The maintenance window of the instance. The time follows the ISO 8601 standard and is displayed in UTC. In the ApsaraDB RDS console, the maintenance window is displayed in UTC+8.
	//
	// example:
	//
	// 00:00Z-02:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// The primary instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	// The maximum number of concurrent connections.
	//
	// example:
	//
	// 60
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum number of I/O requests per second.
	//
	// example:
	//
	// 150
	MaxIOPS *int32 `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// The billing method of the instance.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port that is used to connect to the instance over an internal network.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The latency of data replication from the primary instance to the read-only instance. This parameter is valid for read-only instances.
	//
	// example:
	//
	// 0
	ReadDelayTime *string `json:"ReadDelayTime,omitempty" xml:"ReadDelayTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP addresses in the whitelist.
	//
	// example:
	//
	// 42.xx.xx.xx
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// A deprecated parameter. You do not need to specify this parameter.
	SlaveZones *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttributeSlaveZones `json:"SlaveZones,omitempty" xml:"SlaveZones,omitempty" type:"Struct"`
	// N/A.
	//
	// example:
	//
	// No
	SupportUpgradeAccountType *string `json:"SupportUpgradeAccountType,omitempty" xml:"SupportUpgradeAccountType,omitempty"`
	// The tags.
	//
	// example:
	//
	// 0
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the temporary instance that is attached to the primary instance.
	//
	// example:
	//
	// rm-uf64zsuxxxxxxxxxx
	TempDBInstanceId *string `json:"TempDBInstanceId,omitempty" xml:"TempDBInstanceId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6adz52c2pxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-uf6f7l4fg90xxxxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetAccountMaxQuantity() *int32 {
	return s.AccountMaxQuantity
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetAvailabilityValue() *string {
	return s.AvailabilityValue
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceCPU() *string {
	return s.DBInstanceCPU
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceClassType() *string {
	return s.DBInstanceClassType
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceMemory() *int64 {
	return s.DBInstanceMemory
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetDBMaxQuantity() *int32 {
	return s.DBMaxQuantity
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetExportKey() *string {
	return s.ExportKey
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetGuardDBInstanceId() *string {
	return s.GuardDBInstanceId
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetIncrementSourceDBInstanceId() *string {
	return s.IncrementSourceDBInstanceId
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetMaintainTime() *string {
	return s.MaintainTime
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetMasterInstanceId() *string {
	return s.MasterInstanceId
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetMaxIOPS() *int32 {
	return s.MaxIOPS
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetReadDelayTime() *string {
	return s.ReadDelayTime
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetSlaveZones() *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttributeSlaveZones {
	return s.SlaveZones
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetSupportUpgradeAccountType() *string {
	return s.SupportUpgradeAccountType
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetTags() *string {
	return s.Tags
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetTempDBInstanceId() *string {
	return s.TempDBInstanceId
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetAccountMaxQuantity(v int32) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.AccountMaxQuantity = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetAccountType(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.AccountType = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetAvailabilityValue(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.AvailabilityValue = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetCategory(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.Category = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetConnectionMode(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetConnectionString(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetCreationTime(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceCPU(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceCPU = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceClass(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceClassType(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceClassType = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceDescription(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceId(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceMemory(v int64) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceMemory = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceNetType(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceStatus(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceStorage(v int32) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceStorageType(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceStorageType = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBInstanceType(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetDBMaxQuantity(v int32) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.DBMaxQuantity = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetEngine(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetEngineVersion(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetExpireTime(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetExportKey(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.ExportKey = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetGuardDBInstanceId(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.GuardDBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetIncrementSourceDBInstanceId(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.IncrementSourceDBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetInstanceNetworkType(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetLockMode(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetLockReason(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetMaintainTime(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetMasterInstanceId(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.MasterInstanceId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetMaxConnections(v int32) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetMaxIOPS(v int32) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetPayType(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetPort(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetReadDelayTime(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.ReadDelayTime = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetRegionId(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetSecurityIPList(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetSlaveZones(v *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttributeSlaveZones) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.SlaveZones = v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetSupportUpgradeAccountType(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.SupportUpgradeAccountType = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetTags(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.Tags = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetTempDBInstanceId(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.TempDBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetVSwitchId(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetVpcId(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) SetZoneId(v string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttribute) Validate() error {
	if s.SlaveZones != nil {
		if err := s.SlaveZones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttributeSlaveZones struct {
	SlaveRegion []*string `json:"slaveRegion,omitempty" xml:"slaveRegion,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttributeSlaveZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttributeSlaveZones) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttributeSlaveZones) GetSlaveRegion() []*string {
	return s.SlaveRegion
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttributeSlaveZones) SetSlaveRegion(v []*string) *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttributeSlaveZones {
	s.SlaveRegion = v
	return s
}

func (s *DescribeDBInstancesAsCsvResponseBodyItemsDBInstanceAttributeSlaveZones) Validate() error {
	return dara.Validate(s)
}
