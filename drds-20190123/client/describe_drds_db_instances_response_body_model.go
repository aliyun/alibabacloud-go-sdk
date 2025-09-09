// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDbInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstances(v *DescribeDrdsDbInstancesResponseBodyDbInstances) *DescribeDrdsDbInstancesResponseBody
	GetDbInstances() *DescribeDrdsDbInstancesResponseBodyDbInstances
	SetPageNumber(v string) *DescribeDrdsDbInstancesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeDrdsDbInstancesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeDrdsDbInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsDbInstancesResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeDrdsDbInstancesResponseBody
	GetTotal() *string
}

type DescribeDrdsDbInstancesResponseBody struct {
	// Indicates information about the ApsaraDB RDS for MySQL instances that are used to store the data of the specified database.
	DbInstances *DescribeDrdsDbInstancesResponseBodyDbInstances `json:"DbInstances,omitempty" xml:"DbInstances,omitempty" type:"Struct"`
	// Indicates the page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Indicates the number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates the ID of the request.
	//
	// example:
	//
	// 2F7F8080-9132-4279-85D0-B7E5C4305162
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates the number of primary ApsaraDB RDS for MySQL instances.
	//
	// example:
	//
	// 1
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsDbInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBody) GetDbInstances() *DescribeDrdsDbInstancesResponseBodyDbInstances {
	return s.DbInstances
}

func (s *DescribeDrdsDbInstancesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeDrdsDbInstancesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDrdsDbInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsDbInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsDbInstancesResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeDrdsDbInstancesResponseBody) SetDbInstances(v *DescribeDrdsDbInstancesResponseBodyDbInstances) *DescribeDrdsDbInstancesResponseBody {
	s.DbInstances = v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetPageNumber(v string) *DescribeDrdsDbInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetPageSize(v string) *DescribeDrdsDbInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetRequestId(v string) *DescribeDrdsDbInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetSuccess(v bool) *DescribeDrdsDbInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) SetTotal(v string) *DescribeDrdsDbInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDbInstancesResponseBodyDbInstances struct {
	DbInstance []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstances) GetDbInstance() []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	return s.DbInstance
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstances) SetDbInstance(v []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) *DescribeDrdsDbInstancesResponseBodyDbInstances {
	s.DbInstance = v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance struct {
	// Indicates the endpoint that is used to connect to an ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	//
	// example:
	//
	// rm-bp1t1mk5a5b******.mysql.rds.aliyuncs.com
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// Indicates the ID of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	//
	// example:
	//
	// rm-bp1t1mk5a5bdj****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Indicates the state of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database. Valid values:
	//
	// 	- **0**: The ApsaraDB RDS for MySQL instance is being created.
	//
	// 	- **1**: The ApsaraDB RDS for MySQL instance is running.
	//
	// 	- **3**: The ApsaraDB RDS for MySQL instance is being deleted.
	//
	// 	- **5**: The ApsaraDB RDS for MySQL instance is being restarted.
	//
	// 	- **6**: The ApsaraDB RDS for MySQL instance is being upgraded or downgraded.
	//
	// 	- **7**: The ApsaraDB RDS for MySQL instance is being backed up.
	//
	// 	- **8**: The network type of the ApsaraDB RDS for MySQL instance is being changed.
	//
	// 	- **9**: The ApsaraDB RDS for MySQL instance is being migrated.
	//
	// 	- **11**: The data of the ApsaraDB RDS for MySQL instance is being migrated.
	//
	// 	- **12**: A disaster-recovery instance is being generated.
	//
	// 	- **13**: Data is being imported to the ApsaraDB RDS for MySQL instance.
	//
	// 	- **14**: Data is being imported to the ApsaraDB RDS for MySQL instance from an another ApsaraDB RDS for MySQL instance.
	//
	// 	- **15**: A failover is being performed.
	//
	// 	- **16**: A temporary instance is being created.
	//
	// 	- **17**: A network is being created for the ApsaraDB RDS for MySQL instance.
	//
	// 	- **18**: The ApsaraDB RDS for MySQL instance is being cloned.
	//
	// 	- **19**: The link is being changed.
	//
	// 	- **20**: The read-only instances of the ApsaraDB RDS for MySQL instance are being migrated.
	//
	// example:
	//
	// 1
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// Indicates the type of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database. The value is set to RDS.
	//
	// example:
	//
	// RDS
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Indicates the ID of a resource.
	//
	// example:
	//
	// dm-hbgau1zp****
	DmInstanceId *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	// Indicates the engine of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates the engine version of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Indicates the point in time when the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database expires.
	//
	// example:
	//
	// 1237486127634
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates the network type of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// Indicates the billing method of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database. Valid values:
	//
	// 	- **drdsPre**: The instance uses the subscription billing method.
	//
	// 	- **drdsPost**: The instance uses the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Indicates the port that is used to connect to the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Indicates whether the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database is a primary instance or a read-only instance.
	//
	// 	- **Primary**: The instance is a primary instance.
	//
	// 	- **Readonly**: The instance is a read-only instance.
	//
	// example:
	//
	// Primary
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// Indicates information about the read-only instances of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	ReadOnlyInstances *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances `json:"ReadOnlyInstances,omitempty" xml:"ReadOnlyInstances,omitempty" type:"Struct"`
	// Indicates the read weight of the read-only instance.
	//
	// example:
	//
	// 30
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	// Indicates the number of remaining days before a subscription instance expires.
	//
	// example:
	//
	// 0
	RemainDays *int32 `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetDmInstanceId() *string {
	return s.DmInstanceId
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetRdsInstType() *string {
	return s.RdsInstType
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetReadOnlyInstances() *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances {
	return s.ReadOnlyInstances
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetReadWeight() *int32 {
	return s.ReadWeight
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) GetRemainDays() *int32 {
	return s.RemainDays
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetConnectUrl(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDbInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetEngine(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetEngineVersion(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetExpireTime(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetNetworkType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetPayType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetPort(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetRdsInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetReadOnlyInstances(v *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ReadOnlyInstances = v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetReadWeight(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) SetRemainDays(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances struct {
	ReadOnlyInstance []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance `json:"ReadOnlyInstance,omitempty" xml:"ReadOnlyInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) GetReadOnlyInstance() []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	return s.ReadOnlyInstance
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) SetReadOnlyInstance(v []*DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances {
	s.ReadOnlyInstance = v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance struct {
	// Indicates the endpoint that is used to connect to the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	//
	// example:
	//
	// rm-bp1t1mk5a5b******.mysql.rds.aliyuncs.com
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// Indicates the state of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database. Valid values:
	//
	// 	- **0**: The ApsaraDB RDS for MySQL instance is being created.
	//
	// 	- **1**: The ApsaraDB RDS for MySQL instance is running.
	//
	// 	- **3**: The ApsaraDB RDS for MySQL instance is being deleted.
	//
	// 	- **5**: The ApsaraDB RDS for MySQL instance is being restarted.
	//
	// 	- **6**: The ApsaraDB RDS for MySQL instance is being upgraded or downgraded.
	//
	// 	- **7**: The ApsaraDB RDS for MySQL instance is being backed up.
	//
	// 	- **8**: The network type of the ApsaraDB RDS for MySQL instance is being changed.
	//
	// 	- **9**: The ApsaraDB RDS for MySQL instance is being migrated.
	//
	// 	- **11**: The data of the ApsaraDB RDS for MySQL instance is being migrated.
	//
	// 	- **12**: A disaster-recovery instance is being generated.
	//
	// 	- **13**: Data is being imported to the ApsaraDB RDS for MySQL instance.
	//
	// 	- **14**: Data is being imported to the ApsaraDB RDS for MySQL instance from an another ApsaraDB RDS for MySQL instance.
	//
	// 	- **15**: A failover is being performed.
	//
	// 	- **16**: A temporary instance is being created.
	//
	// 	- **17**: A network is being created for the ApsaraDB RDS for MySQL instance.
	//
	// 	- **18**: The ApsaraDB RDS for MySQL instance is being cloned.
	//
	// 	- **19**: The link is being changed.
	//
	// 	- **20**: The read-only instances of the ApsaraDB RDS for MySQL instance are being migrated.
	//
	// example:
	//
	// 1
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// Indicates the type of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database. The value is set to RDS.
	//
	// example:
	//
	// RDS
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Indicates the ID of a resource.
	//
	// example:
	//
	// dm-hbgau1zp****
	DmInstanceId *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	// Indicates the engine of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates the engine version of the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Indicates the timestamp when the ApsaraDB RDS for MySQL instance that is used to store the data of the specified database expires.
	//
	// example:
	//
	// 123421352351234
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates the name of a read-only instance.
	//
	// example:
	//
	// **
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates the network type of the read-only instance.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// Indicates the billing method of the read-only instance.
	//
	// 	- **drdsPre**: The instance uses the subscription billing method.
	//
	// 	- **drdsPost**: The instance uses the pay-as-you-go billing method.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Indicates the port that is used to connect to the read-only instance.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Indicates the type of the read-only instance.
	//
	// example:
	//
	// RDS
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// Indicates the read weight of the read-only instance.
	//
	// example:
	//
	// 70
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	// Indicates the number of remaining days before the read-only instance expires.
	//
	// example:
	//
	// 0
	RemainDays *int32 `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetDmInstanceId() *string {
	return s.DmInstanceId
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetRdsInstType() *string {
	return s.RdsInstType
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetReadWeight() *int32 {
	return s.ReadWeight
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) GetRemainDays() *int32 {
	return s.RemainDays
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetConnectUrl(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDbInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngine(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngineVersion(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetExpireTime(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetInstanceName(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetNetworkType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetPayType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetPort(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetRdsInstType(v string) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetReadWeight(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) SetRemainDays(v int32) *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponseBodyDbInstancesDbInstanceReadOnlyInstancesReadOnlyInstance) Validate() error {
	return dara.Validate(s)
}
