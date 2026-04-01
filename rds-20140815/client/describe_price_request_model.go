// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribePriceRequest
	GetClientToken() *string
	SetCommodityCode(v string) *DescribePriceRequest
	GetCommodityCode() *string
	SetDBInstanceClass(v string) *DescribePriceRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *DescribePriceRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *DescribePriceRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *DescribePriceRequest
	GetDBInstanceStorageType() *string
	SetDBNode(v []*DescribePriceRequestDBNode) *DescribePriceRequest
	GetDBNode() []*DescribePriceRequestDBNode
	SetEngine(v string) *DescribePriceRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribePriceRequest
	GetEngineVersion() *string
	SetInstanceUsedType(v int32) *DescribePriceRequest
	GetInstanceUsedType() *int32
	SetOrderType(v string) *DescribePriceRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribePriceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePriceRequest
	GetOwnerId() *int64
	SetPayType(v string) *DescribePriceRequest
	GetPayType() *string
	SetQuantity(v int32) *DescribePriceRequest
	GetQuantity() *int32
	SetRegionId(v string) *DescribePriceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePriceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePriceRequest
	GetResourceOwnerId() *int64
	SetServerlessConfig(v *DescribePriceRequestServerlessConfig) *DescribePriceRequest
	GetServerlessConfig() *DescribePriceRequestServerlessConfig
	SetTimeType(v string) *DescribePriceRequest
	GetTimeType() *string
	SetUsedTime(v int32) *DescribePriceRequest
	GetUsedTime() *int32
	SetZoneId(v string) *DescribePriceRequest
	GetZoneId() *string
}

type DescribePriceRequest struct {
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
	DBNode []*DescribePriceRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
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
	ServerlessConfig *DescribePriceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
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

func (s DescribePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribePriceRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribePriceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribePriceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribePriceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *DescribePriceRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *DescribePriceRequest) GetDBNode() []*DescribePriceRequestDBNode {
	return s.DBNode
}

func (s *DescribePriceRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribePriceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribePriceRequest) GetInstanceUsedType() *int32 {
	return s.InstanceUsedType
}

func (s *DescribePriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePriceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePriceRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribePriceRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *DescribePriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePriceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePriceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePriceRequest) GetServerlessConfig() *DescribePriceRequestServerlessConfig {
	return s.ServerlessConfig
}

func (s *DescribePriceRequest) GetTimeType() *string {
	return s.TimeType
}

func (s *DescribePriceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *DescribePriceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePriceRequest) SetClientToken(v string) *DescribePriceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribePriceRequest) SetCommodityCode(v string) *DescribePriceRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribePriceRequest) SetDBInstanceClass(v string) *DescribePriceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribePriceRequest) SetDBInstanceId(v string) *DescribePriceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribePriceRequest) SetDBInstanceStorage(v int32) *DescribePriceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribePriceRequest) SetDBInstanceStorageType(v string) *DescribePriceRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *DescribePriceRequest) SetDBNode(v []*DescribePriceRequestDBNode) *DescribePriceRequest {
	s.DBNode = v
	return s
}

func (s *DescribePriceRequest) SetEngine(v string) *DescribePriceRequest {
	s.Engine = &v
	return s
}

func (s *DescribePriceRequest) SetEngineVersion(v string) *DescribePriceRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceUsedType(v int32) *DescribePriceRequest {
	s.InstanceUsedType = &v
	return s
}

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerAccount(v string) *DescribePriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerId(v int64) *DescribePriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetPayType(v string) *DescribePriceRequest {
	s.PayType = &v
	return s
}

func (s *DescribePriceRequest) SetQuantity(v int32) *DescribePriceRequest {
	s.Quantity = &v
	return s
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerAccount(v string) *DescribePriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerId(v int64) *DescribePriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetServerlessConfig(v *DescribePriceRequestServerlessConfig) *DescribePriceRequest {
	s.ServerlessConfig = v
	return s
}

func (s *DescribePriceRequest) SetTimeType(v string) *DescribePriceRequest {
	s.TimeType = &v
	return s
}

func (s *DescribePriceRequest) SetUsedTime(v int32) *DescribePriceRequest {
	s.UsedTime = &v
	return s
}

func (s *DescribePriceRequest) SetZoneId(v string) *DescribePriceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribePriceRequest) Validate() error {
	if s.DBNode != nil {
		for _, item := range s.DBNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServerlessConfig != nil {
		if err := s.ServerlessConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceRequestDBNode struct {
	// The instance type of the node.
	//
	// example:
	//
	// mysql.n2.small.xc
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The zone ID of the node.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribePriceRequestDBNode) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestDBNode) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestDBNode) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribePriceRequestDBNode) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePriceRequestDBNode) SetClassCode(v string) *DescribePriceRequestDBNode {
	s.ClassCode = &v
	return s
}

func (s *DescribePriceRequestDBNode) SetZoneId(v string) *DescribePriceRequestDBNode {
	s.ZoneId = &v
	return s
}

func (s *DescribePriceRequestDBNode) Validate() error {
	return dara.Validate(s)
}

type DescribePriceRequestServerlessConfig struct {
	// The maximum number of RDS Capacity Units (RCUs).
	//
	// example:
	//
	// 8
	MaxCapacity *float64 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The minimum number of RCUs.
	//
	// example:
	//
	// 0.5
	MinCapacity *float64 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s DescribePriceRequestServerlessConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestServerlessConfig) GetMaxCapacity() *float64 {
	return s.MaxCapacity
}

func (s *DescribePriceRequestServerlessConfig) GetMinCapacity() *float64 {
	return s.MinCapacity
}

func (s *DescribePriceRequestServerlessConfig) SetMaxCapacity(v float64) *DescribePriceRequestServerlessConfig {
	s.MaxCapacity = &v
	return s
}

func (s *DescribePriceRequestServerlessConfig) SetMinCapacity(v float64) *DescribePriceRequestServerlessConfig {
	s.MinCapacity = &v
	return s
}

func (s *DescribePriceRequestServerlessConfig) Validate() error {
	return dara.Validate(s)
}
