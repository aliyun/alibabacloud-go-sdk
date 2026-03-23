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
	DBNode []*DescribePriceRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
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
	Quantity             *int32                                `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	RegionId             *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerlessConfig     *DescribePriceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	TimeType             *string                               `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	UsedTime             *int32                                `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	ZoneId               *string                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	MaxCapacity *float64 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
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
