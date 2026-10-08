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
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity code of the instance. Valid values:
	//
	// 	- **bards**: pay-as-you-go primary instance (China site)
	//
	// 	- **rds*	- (default): subscription primary instance (China site)
	//
	// 	- **rords**: pay-as-you-go read-only instance (China site)
	//
	// 	- **rds_rordspre_public_cn**: subscription read-only instance (China site)
	//
	// 	- **bards_intl**: pay-as-you-go primary instance (international site)
	//
	// 	- **rds_intl**: subscription primary instance (international site)
	//
	// 	- **rords_intl**: pay-as-you-go read-only instance (international site)
	//
	// 	- **rds_rordspre_public_intl**: subscription read-only instance (international site)
	//
	// > This parameter is required when you query the price of a read-only instance.
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The instance type. For more information, see [Primary instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql.x2.medium.xc
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// Instance ID of the instance for which you want to change the specifications or renew.
	//
	// > - This parameter is required when you query the price for a specification change or renewal.
	//
	// > - If the instance is a read-only instance, specify instance ID of its primary instance.
	//
	// example:
	//
	// rm-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance storage space. Unit: GB. The value increases in increments of 5 GB. For more information about the value range, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type. Valid values:
	//
	// 	- **general_essd**: Premium ESSD
	//
	// 	- **local_ssd**: Premium Local SSDs
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd**: PL1 ESSD cloud disk
	//
	// 	- **cloud_essd2**: PL2 ESSD cloud disk
	//
	// 	- **cloud_essd3**: PL3 ESSD cloud disk
	//
	// example:
	//
	// local_ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The node information.
	//
	// > This parameter is used for ApsaraDB RDS for MySQL instances in the cluster edition.
	//
	// if can be null:
	// true
	DBNode []*DescribePriceRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
	// The database engine. Valid values:
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
	// <props="china">The database engine version. Valid values:
	//
	// - **MySQL**: **5.5**, **5.6**, **5.7**, **8.0**
	//
	// - **SQL Server**: **08r2_ent_ha*	- (cloud disk, discontinued), **2008r2*	- (Premium Local SSDs, discontinued), **2012*	- (Enterprise Edition Basic), **2012_ent_ha**, **2012_std_ha**, **2012_web**, **2014_ent_ha**, **2014_std_ha**, **2016_ent_ha**, **2016_std_ha**, **2016_web**, **2017_ent**, **2017_std_ha**, **2017_web**, **2019_ent**, **2019_std_ha**, **2019_web**, **2022_ent**, **2022_std_ha**, **2022_web**
	//
	// - **PostgreSQL**: **10.0**, **11.0**, **12.0**, **13.0**, **14.0**, **15.0**
	//
	// - **MariaDB**: **10.3**
	//
	//
	//
	// <props="intl">The database engine version. Valid values:
	//
	// - **MySQL**: **5.5**, **5.6**, **5.7**, **8.0**
	//
	// - **SQL Server**: **08r2_ent_ha*	- (cloud disk, discontinued), **2008r2*	- (Premium Local SSDs, discontinued), **2012*	- (Enterprise Edition Basic), **2012_ent_ha**, **2012_std_ha**, **2012_web**, **2014_ent_ha**, **2014_std_ha**, **2016_ent_ha**, **2016_std_ha**, **2016_web**, **2017_ent**, **2017_std_ha**, **2017_web**, **2019_ent**, **2019_std_ha**, **2019_web**, **2022_ent**, **2022_std_ha**, **2022_web**
	//
	// - **PostgreSQL**: **10.0**, **11.0**, **12.0**, **13.0**, **14.0**, **15.0**
	//
	// - **MariaDB**: **10.3**
	//
	// > For SQL Server instances, `_ent` indicates Enterprise Edition (Cluster), `_ent_ha` indicates Enterprise Edition, `_std_ha` indicates Standard Edition, and `_web` indicates Web Edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance type. Valid values:
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
	// 	- **BUY**: purchase
	//
	// 	- **RENEW**: renewal
	//
	// 	- **UPGRADE**: upgrade
	//
	// 	- **DOWNGRADE**: downgrade
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
	// The number of instances to purchase. Valid values: **0 to 30**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The region ID. You can call DescribeRegions to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The settings of the serverless ApsaraDB RDS instance.
	//
	// > MariaDB does not support serverless instances.
	ServerlessConfig *DescribePriceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The subscription type. This parameter is required when **CommodityCode*	- is set to **rds**, **rds_rordspre_public_cn**, **rds_intl**, or **rds_rordspre_public_intl**. Valid values:
	//
	// 	- **Year**: yearly subscription
	//
	// 	- **Month**: monthly subscription
	//
	// example:
	//
	// Year
	TimeType *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The subscription duration. Valid values:
	//
	// 	- If **TimeType*	- is set to **Year**, the value of UsedTime ranges from **1 to 100**.
	//
	// 	- If **TimeType*	- is set to **Month**, the value of UsedTime ranges from **1 to 999**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The zone ID of the primary node. You can call DescribeRegions to query the most recent zone list.
	//
	// > If you specify a VPC and a vSwitch, this parameter is required to match the zone of the specified vSwitch.
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
	// The node specifications.
	//
	// example:
	//
	// mysql.x2.medium.xc
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
	// The maximum value of the automatic scaling range for the RDS Capacity Unit (RCU) of the instance.
	//
	// example:
	//
	// 8
	MaxCapacity *float64 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The minimum value of the automatic scaling range for the RDS Capacity Unit (RCU) of the instance.
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
