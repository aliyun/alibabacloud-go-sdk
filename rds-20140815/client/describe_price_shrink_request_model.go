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
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// This parameter is required.
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DBInstanceStorage     *int32  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// if can be null:
	// true
	DBNodeShrink *string `json:"DBNode,omitempty" xml:"DBNode,omitempty"`
	// This parameter is required.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	EngineVersion    *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InstanceUsedType *int32  `json:"InstanceUsedType,omitempty" xml:"InstanceUsedType,omitempty"`
	OrderType        *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// This parameter is required.
	Quantity               *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	TimeType               *string `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	UsedTime               *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	ZoneId                 *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
