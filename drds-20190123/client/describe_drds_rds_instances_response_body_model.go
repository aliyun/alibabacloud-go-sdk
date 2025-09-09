// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsRdsInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstances(v *DescribeDrdsRdsInstancesResponseBodyDbInstances) *DescribeDrdsRdsInstancesResponseBody
	GetDbInstances() *DescribeDrdsRdsInstancesResponseBodyDbInstances
	SetRequestId(v string) *DescribeDrdsRdsInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsRdsInstancesResponseBody
	GetSuccess() *bool
}

type DescribeDrdsRdsInstancesResponseBody struct {
	// The information about the custom ApsaraDB RDS for MySQL instances at the storage layer.
	DbInstances *DescribeDrdsRdsInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 123DB16B-02F2-45F7-A571-843991******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsRdsInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBody) GetDbInstances() *DescribeDrdsRdsInstancesResponseBodyDbInstances {
	return s.DbInstances
}

func (s *DescribeDrdsRdsInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsRdsInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetDbInstances(v *DescribeDrdsRdsInstancesResponseBodyDbInstances) *DescribeDrdsRdsInstancesResponseBody {
	s.DbInstances = v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetRequestId(v string) *DescribeDrdsRdsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBody) SetSuccess(v bool) *DescribeDrdsRdsInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsRdsInstancesResponseBodyDbInstances struct {
	DbInstance []*DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstances) GetDbInstance() []*DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	return s.DbInstance
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstances) SetDbInstance(v []*DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) *DescribeDrdsRdsInstancesResponseBodyDbInstances {
	s.DbInstance = v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance struct {
	// The internal endpoint of the custom ApsaraDB RDS for MySQL instance at the storage layer.
	//
	// example:
	//
	// rm-***************.mysql.rds.aliyuncs.com
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The number of CPU cores of the custom ApsaraDB RDS for MySQL instance at the storage layer.
	//
	// example:
	//
	// 8
	DBInstanceCPU *string `json:"DBInstanceCPU,omitempty" xml:"DBInstanceCPU,omitempty"`
	// The instance family of the custom ApsaraDB RDS for MySQL instance at the storage layer. Valid values:
	//
	// 	- **x**: general-purpose instance family
	//
	// 	- **d**: dedicated instance family
	//
	// 	- **h**: dedicated host instance family
	//
	// example:
	//
	// x
	DBInstanceClassType *string `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	// The ID of the custom ApsaraDB RDS for MySQL instance at the storage layer.
	//
	// example:
	//
	// rm-*****************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The memory size of the custom ApsaraDB RDS for MySQL instance at the storage layer. Unit: MB.
	//
	// example:
	//
	// 8192
	DBInstanceMemory *int64 `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty"`
	// The status of the custom ApsaraDB RDS for MySQL instance at the storage layer. Valid values:
	//
	// 	- 0: The instance is being created.
	//
	// 	- 1: The instance is running.
	//
	// 	- 3: The instance is being deleted.
	//
	// 	- 5: The instance is being restarted.
	//
	// 	- 6: The instance is being upgraded or downgraded.
	//
	// 	- 7: The instance is being backed up.
	//
	// 	- 8: The network type of the instance is being changed.
	//
	// 	- 9: The instance is being migrated.
	//
	// 	- 11: The data stored on the instance is being migrated.
	//
	// 	- 12: A disaster recovery instance is being generated.
	//
	// 	- 13: Data is being imported to the instance.
	//
	// 	- 14: Data is being imported from another RDS instance to the instance.
	//
	// 	- 15: A switchover is being performed.
	//
	// 	- 16: A temporary instance is being created.
	//
	// 	- 17: The network of the instance is being created.
	//
	// 	- 18: The instance is being cloned.
	//
	// 	- 19: The link is being changed.
	//
	// 	- 20: The read-only RDS instances of the instance are being migrated.
	//
	// example:
	//
	// 1
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The storage space of the custom ApsaraDB RDS for MySQL instance at the storage layer. Unit: GB.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int64 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The type of the instance at the storage layer. The value is RDS.
	//
	// example:
	//
	// RDS
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// dm-*************
	DmInstanceId *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	// The engine type of the custom ApsaraDB RDS for MySQL instance at the storage layer. The value is MySQL.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version of the custom ApsaraDB RDS for MySQL instance at the storage layer. The value is 8.0.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The lock mode of the RDS instance. Valid values:
	//
	// 0: The instance is not locked.
	//
	// 1: The instance is manually locked.
	//
	// 2: The instance is automatically locked if the instance expires.
	//
	// 3: The instance is automatically locked if the instance is rolled back.
	//
	// 4: The instance is automatically locked if the storage space of the instance reaches the upper limit.
	//
	// 5: The instance is automatically locked if the storage space of the read-only instances reaches the upper limit.
	//
	// example:
	//
	// 1
	LockMode *int32 `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The reason why the RDS instance is locked.
	//
	// example:
	//
	// Manually Locked
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The network type of the custom ApsaraDB RDS for MySQL instance at the storage layer. The value is VPC.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the custom ApsaraDB RDS for MySQL instance at the storage layer. Valid values:
	//
	// 	- Postpaid: pay-as-you-go
	//
	// 	- Prepaid: subscription
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port used to connect to the instance over an internal network.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the custom ApsaraDB RDS for MySQL instance at the storage layer. Valid values:
	//
	// 	- Primary: primary instance
	//
	// 	- Readonly: read-only instance
	//
	// example:
	//
	// Primary
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// The read and write weights of the custom ApsaraDB RDS for MySQL instance at the storage layer.
	//
	// example:
	//
	// 0
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetDBInstanceCPU() *string {
	return s.DBInstanceCPU
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetDBInstanceClassType() *string {
	return s.DBInstanceClassType
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetDBInstanceMemory() *int64 {
	return s.DBInstanceMemory
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetDBInstanceStorage() *int64 {
	return s.DBInstanceStorage
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetDmInstanceId() *string {
	return s.DmInstanceId
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetLockMode() *int32 {
	return s.LockMode
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetRdsInstType() *string {
	return s.RdsInstType
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) GetReadWeight() *int32 {
	return s.ReadWeight
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetConnectUrl(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceCPU(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceCPU = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceClassType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceClassType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceId(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceMemory(v int64) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceMemory = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStorage(v int64) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDbInstType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetDmInstanceId(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetEngine(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetEngineVersion(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetLockMode(v int32) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetLockReason(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.LockReason = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetNetworkType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetPayType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetPort(v int32) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetRdsInstType(v string) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) SetReadWeight(v int32) *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponseBodyDbInstancesDbInstance) Validate() error {
	return dara.Validate(s)
}
