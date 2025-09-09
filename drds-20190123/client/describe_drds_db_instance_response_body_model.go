// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDbInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstance(v *DescribeDrdsDbInstanceResponseBodyDbInstance) *DescribeDrdsDbInstanceResponseBody
	GetDbInstance() *DescribeDrdsDbInstanceResponseBodyDbInstance
	SetRequestId(v string) *DescribeDrdsDbInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsDbInstanceResponseBody
	GetSuccess() *bool
}

type DescribeDrdsDbInstanceResponseBody struct {
	// The detailed information about the returned custom ApsaraDB RDS for MySQL instance.
	DbInstance *DescribeDrdsDbInstanceResponseBodyDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 4FE09970-CA69-4144-88CA-67FB4BTY56G3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDbInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBody) GetDbInstance() *DescribeDrdsDbInstanceResponseBodyDbInstance {
	return s.DbInstance
}

func (s *DescribeDrdsDbInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsDbInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsDbInstanceResponseBody) SetDbInstance(v *DescribeDrdsDbInstanceResponseBodyDbInstance) *DescribeDrdsDbInstanceResponseBody {
	s.DbInstance = v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBody) SetRequestId(v string) *DescribeDrdsDbInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBody) SetSuccess(v bool) *DescribeDrdsDbInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDbInstanceResponseBodyDbInstance struct {
	// The URL used to connect to the custom ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// rm-***************.mysql.rds.aliyuncs.com
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The ID of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// rm-***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The state of the instance.
	//
	// example:
	//
	// 1
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The role of the instance. Valid values:
	//
	// 	- **Primary**: The instance is a primary instance.
	//
	// 	- **ReadOnly**: The instance is a read-only instance.
	//
	// example:
	//
	// Primary
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// dm-*************
	DmInstanceId *string `json:"DmInstanceId,omitempty" xml:"DmInstanceId,omitempty"`
	// The engine of the database that is run on the instance. Valid value: **MySQL**.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version of the database that is run on the instance. Valid values: **5.7**.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the custom ApsaraDB RDS for MySQL instance expires. The value of this parameter is a UNIX timestamp. Unit: seconds.
	//
	// >  This parameter is returned only when the custom ApsaraDB RDS for MySQL instance is a subscription instance.
	//
	// example:
	//
	// 12341434315
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The type of the network. Valid values: **VPC**.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the custom ApsaraDB RDS for MySQL instance. Valid values:
	//
	// 	- **Prepaid**: subscription
	//
	// 	- **Postaid**: pay-as-you-go
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port used to connect to the custom ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// RDS
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// The list of read-only ApsaraDB RDS for MySQL instances.
	ReadOnlyInstances *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances `json:"ReadOnlyInstances,omitempty" xml:"ReadOnlyInstances,omitempty" type:"Struct"`
	// The read ratio of the instance.
	//
	// example:
	//
	// 70
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	// The number of remaining days before the instance expires.
	//
	// example:
	//
	// 0
	RemainDays *string `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetDmInstanceId() *string {
	return s.DmInstanceId
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetRdsInstType() *string {
	return s.RdsInstType
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetReadOnlyInstances() *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances {
	return s.ReadOnlyInstances
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetReadWeight() *int32 {
	return s.ReadWeight
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) GetRemainDays() *string {
	return s.RemainDays
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetConnectUrl(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDbInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetEngine(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetEngineVersion(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetExpireTime(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetNetworkType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetPayType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetPort(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetRdsInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetReadOnlyInstances(v *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ReadOnlyInstances = v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetReadWeight(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) SetRemainDays(v string) *DescribeDrdsDbInstanceResponseBodyDbInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances struct {
	ReadOnlyInstance []*DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance `json:"ReadOnlyInstance,omitempty" xml:"ReadOnlyInstance,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) GetReadOnlyInstance() []*DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	return s.ReadOnlyInstance
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) SetReadOnlyInstance(v []*DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances {
	s.ReadOnlyInstance = v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance struct {
	// The URL used to connect to the read-only instance.
	//
	// example:
	//
	// rm-bp1ub71ct9skc3yxx.mysql.rds.aliyuncs.com
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The ID of the read-only instance.
	//
	// example:
	//
	// rm-bp1ub71ct9skc****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The state of the read-only instance.
	//
	// example:
	//
	// 1
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The role of the read-only instance.
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
	// The engine of the database that is run on the read-only instance.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The engine version of the database that is run on the read-only instance.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The timestamp that indicates when the read-only instance expires.
	//
	// example:
	//
	// 1823487328173
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The network type of the read-only instance.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the read-only instance.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The port used to connect to the read-only instance.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// RDS
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// The read ratio of the read-only instance.
	//
	// example:
	//
	// 30
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
	// The number of remaining days before the read-only instance expires.
	//
	// example:
	//
	// 0
	RemainDays *string `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	// This parameter is unavailable for read-only instances.
	//
	// example:
	//
	// 0
	VersionAction *int32 `json:"VersionAction,omitempty" xml:"VersionAction,omitempty"`
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetDmInstanceId() *string {
	return s.DmInstanceId
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetRdsInstType() *string {
	return s.RdsInstType
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetReadWeight() *int32 {
	return s.ReadWeight
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetRemainDays() *string {
	return s.RemainDays
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) GetVersionAction() *int32 {
	return s.VersionAction
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetConnectUrl(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDBInstanceStatus(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDbInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetDmInstanceId(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.DmInstanceId = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngine(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetEngineVersion(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetExpireTime(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetNetworkType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetPayType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetPort(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetRdsInstType(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetReadWeight(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetRemainDays(v string) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) SetVersionAction(v int32) *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance {
	s.VersionAction = &v
	return s
}

func (s *DescribeDrdsDbInstanceResponseBodyDbInstanceReadOnlyInstancesReadOnlyInstance) Validate() error {
	return dara.Validate(s)
}
