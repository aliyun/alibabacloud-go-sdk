// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableClassesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeAvailableClassesRequest
	GetCategory() *string
	SetCommodityCode(v string) *DescribeAvailableClassesRequest
	GetCommodityCode() *string
	SetDBInstanceId(v string) *DescribeAvailableClassesRequest
	GetDBInstanceId() *string
	SetDBInstanceStorageType(v string) *DescribeAvailableClassesRequest
	GetDBInstanceStorageType() *string
	SetEngine(v string) *DescribeAvailableClassesRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeAvailableClassesRequest
	GetEngineVersion() *string
	SetInstanceChargeType(v string) *DescribeAvailableClassesRequest
	GetInstanceChargeType() *string
	SetOrderType(v string) *DescribeAvailableClassesRequest
	GetOrderType() *string
	SetRegionId(v string) *DescribeAvailableClassesRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeAvailableClassesRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeAvailableClassesRequest
	GetZoneId() *string
}

type DescribeAvailableClassesRequest struct {
	// This parameter is required.
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// This parameter is required.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	EngineVersion      *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	OrderType          *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// This parameter is required.
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableClassesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableClassesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableClassesRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeAvailableClassesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeAvailableClassesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAvailableClassesRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *DescribeAvailableClassesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAvailableClassesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAvailableClassesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAvailableClassesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeAvailableClassesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableClassesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableClassesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableClassesRequest) SetCategory(v string) *DescribeAvailableClassesRequest {
	s.Category = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetCommodityCode(v string) *DescribeAvailableClassesRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetDBInstanceId(v string) *DescribeAvailableClassesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetDBInstanceStorageType(v string) *DescribeAvailableClassesRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetEngine(v string) *DescribeAvailableClassesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetEngineVersion(v string) *DescribeAvailableClassesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetInstanceChargeType(v string) *DescribeAvailableClassesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetOrderType(v string) *DescribeAvailableClassesRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetRegionId(v string) *DescribeAvailableClassesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetResourceOwnerId(v int64) *DescribeAvailableClassesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableClassesRequest) SetZoneId(v string) *DescribeAvailableClassesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableClassesRequest) Validate() error {
	return dara.Validate(s)
}
