// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribePriceShrinkRequest
	GetClientToken() *string
	SetCommodityCode(v string) *DescribePriceShrinkRequest
	GetCommodityCode() *string
	SetDBInstanceClass(v string) *DescribePriceShrinkRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *DescribePriceShrinkRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *DescribePriceShrinkRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *DescribePriceShrinkRequest
	GetDBInstanceStorageType() *string
	SetDBNodeShrink(v string) *DescribePriceShrinkRequest
	GetDBNodeShrink() *string
	SetEngine(v string) *DescribePriceShrinkRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribePriceShrinkRequest
	GetEngineVersion() *string
	SetInstanceUsedType(v int32) *DescribePriceShrinkRequest
	GetInstanceUsedType() *int32
	SetOrderType(v string) *DescribePriceShrinkRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribePriceShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePriceShrinkRequest
	GetOwnerId() *int64
	SetPayType(v string) *DescribePriceShrinkRequest
	GetPayType() *string
	SetQuantity(v int32) *DescribePriceShrinkRequest
	GetQuantity() *int32
	SetRegionId(v string) *DescribePriceShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePriceShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePriceShrinkRequest
	GetResourceOwnerId() *int64
	SetServerlessConfigShrink(v string) *DescribePriceShrinkRequest
	GetServerlessConfigShrink() *string
	SetTimeType(v string) *DescribePriceShrinkRequest
	GetTimeType() *string
	SetUsedTime(v int32) *DescribePriceShrinkRequest
	GetUsedTime() *int32
	SetZoneId(v string) *DescribePriceShrinkRequest
	GetZoneId() *string
}

type DescribePriceShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity code of the instance. Valid values:
	//
	// 	- **bards**: The instance is a pay-as-you-go primary instance. This value is available at the China site (aliyun.com).
	//
	// 	- **rds*	- (default): The instance is a subscription primary instance. This value is available at the China site (aliyun.com).
	//
	// 	- **rords**: The instance is a pay-as-you-go read-only instance. This value is available at the China site (aliyun.com).
	//
	// 	- **rds_rordspre_public_cn**: The instance is a subscription read-only instance. This value is available at the China site (aliyun.com).
	//
	// 	- **bards_intl**: The instance is a pay-as-you-go primary instance. This value is available at the international site (alibabacloud.com).
	//
	// 	- **rds_intl**: The instance is a subscription primary instance. This value is available at the international site (alibabacloud.com).
	//
	// 	- **rords_intl**: The instance is a pay-as-you-go read-only instance. This value is available at the international site (alibabacloud.com).
	//
	// 	- **rds_rordspre_public_intl**: The instance is a subscription read-only instance. This value is available at the international site (alibabacloud.com).
	//
	// >  If you want to query the price of a read-only instance, you must specify this parameter.
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The instance type of the instance. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rds.mysql.s1.small
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The ID of the instance for which you want to change the specifications or the instance that you want to renew.
	//
	// > 	- If you want to query the price of a specification change order or a renewal order, you must specify this parameter.
	//
	// > 	- If the instance is a read-only instance, you must set this parameter to the ID of its primary instance.
	//
	// example:
	//
	// rm-*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The storage capacity of the instance. Unit: GB. You can increase the storage capacity at a step size of 5 GB. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the new instance. Valid values:
	//
	// 	- **general_essd**: premium Enterprise SSD (ESSD)
	//
	// 	- **local_ssd**: premium local SSD
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd**: performance level 1 (PL1) ESSD
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The information about the node.
	//
	// >  This parameter is supported for ApsaraDB RDS for MySQL instances that run RDS Cluster Edition.
	//
	// if can be null:
	// true
	DBNodeShrink *string `json:"DBNode,omitempty" xml:"DBNode,omitempty"`
	// The database engine of the instance. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// 	- **MariaDB**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance. Valid values:
	//
	// 	- Valid values if you set Engine to **MySQL**: **5.5**, **5.6**, **5.7**, and **8.0**
	//
	// 	- Valid values if you set Engine to **SQL Server**: **08r2_ent_ha**(cloud disks, discontinued), **2008r2**(high-performance local disks, discontinued), **2012*	- (SQL Server EE Basic)**2012_ent_ha**, **2012_std_ha**, **2012_web**, **2016_ent_ha**, **2016_std_ha**, **2016_web**, **2017_ent**, **2017_std_ha**, **2017_web**, **2019_ent**, **2019_std_ha**, **2019_web**, **2022_ent**, **2022_std_ha**, and **2022_web**
	//
	// 	- Valid values if you set Engine to **PostgreSQL**: **10.0**, **11.0**, **12.0**, **13.0**, **14.0**, and **15.0**
	//
	// 	- Valid value if you set Engine to **MariaDB**: **10.3**
	//
	// >  The following information describes the valid values when you set Engine to SQLServer: `_ent` specifies SQL Server EE on RDS Cluster Edition, `_ent_ha` specifies SQL Server EE, `_std_ha` specifies SQL Server SE, and `_web` specifies SQL Server Web.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.5
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The role of the instance. Valid values:
	//
	// 	- **0**: primary instance
	//
	// 	- **3**: read-only instance
	//
	// example:
	//
	// 0
	InstanceUsedType *int32 `json:"InstanceUsedType,omitempty" xml:"InstanceUsedType,omitempty"`
	// The order type. Valid values:
	//
	// 	- **BUY**
	//
	// 	- **RENEW**
	//
	// 	- **UPGRADE**
	//
	// 	- **DOWNGRADE**
	//
	// example:
	//
	// BUY
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Prepaid**: subscription
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The number of instances that you want to purchase. Valid values: **0 to 30**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The settings of the serverless instance.
	//
	// > ApsaraDB RDS for MariaDB does not support serverless instances.
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	// The billing cycle of the subscription instance. This parameter is required when **CommodityCode*	- is set to **rds**, **rds_rordspre_public_cn**, **rds_intl**, or **rds_rordspre_public_intl**. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// example:
	//
	// Year
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The subscription duration of the instance.
	//
	// 	- If you set the **TimeType*	- parameter to **Year**, the value of the UsedTime parameter ranges from **1 to 100**.
	//
	// 	- If you set the **TimeType*	- parameter to **Month**, the value of the UsedTime parameter ranges from **1 to 999**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The zone ID of the primary instance. You can call the DescribeRegions operation to query the most recent zone list.
	//
	// >  If you specify a virtual private cloud (VPC) and a vSwitch, this parameter is required to identify the zone for the vSwitch.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribePriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribePriceShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribePriceShrinkRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribePriceShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribePriceShrinkRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *DescribePriceShrinkRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *DescribePriceShrinkRequest) GetDBNodeShrink() *string {
	return s.DBNodeShrink
}

func (s *DescribePriceShrinkRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribePriceShrinkRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribePriceShrinkRequest) GetInstanceUsedType() *int32 {
	return s.InstanceUsedType
}

func (s *DescribePriceShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePriceShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePriceShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePriceShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribePriceShrinkRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *DescribePriceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePriceShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePriceShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePriceShrinkRequest) GetServerlessConfigShrink() *string {
	return s.ServerlessConfigShrink
}

func (s *DescribePriceShrinkRequest) GetTimeType() *string {
	return s.TimeType
}

func (s *DescribePriceShrinkRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *DescribePriceShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePriceShrinkRequest) SetClientToken(v string) *DescribePriceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetCommodityCode(v string) *DescribePriceShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetDBInstanceClass(v string) *DescribePriceShrinkRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetDBInstanceId(v string) *DescribePriceShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetDBInstanceStorage(v int32) *DescribePriceShrinkRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetDBInstanceStorageType(v string) *DescribePriceShrinkRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetDBNodeShrink(v string) *DescribePriceShrinkRequest {
	s.DBNodeShrink = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetEngine(v string) *DescribePriceShrinkRequest {
	s.Engine = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetEngineVersion(v string) *DescribePriceShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetInstanceUsedType(v int32) *DescribePriceShrinkRequest {
	s.InstanceUsedType = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetOrderType(v string) *DescribePriceShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetOwnerAccount(v string) *DescribePriceShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetOwnerId(v int64) *DescribePriceShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetPayType(v string) *DescribePriceShrinkRequest {
	s.PayType = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetQuantity(v int32) *DescribePriceShrinkRequest {
	s.Quantity = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetRegionId(v string) *DescribePriceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetResourceOwnerAccount(v string) *DescribePriceShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetResourceOwnerId(v int64) *DescribePriceShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetServerlessConfigShrink(v string) *DescribePriceShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetTimeType(v string) *DescribePriceShrinkRequest {
	s.TimeType = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetUsedTime(v int32) *DescribePriceShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetZoneId(v string) *DescribePriceShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribePriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
