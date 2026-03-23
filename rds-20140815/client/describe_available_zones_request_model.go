// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeAvailableZonesRequest
	GetCategory() *string
	SetCommodityCode(v string) *DescribeAvailableZonesRequest
	GetCommodityCode() *string
	SetDBInstanceName(v string) *DescribeAvailableZonesRequest
	GetDBInstanceName() *string
	SetDispenseMode(v string) *DescribeAvailableZonesRequest
	GetDispenseMode() *string
	SetEngine(v string) *DescribeAvailableZonesRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeAvailableZonesRequest
	GetEngineVersion() *string
	SetRegionId(v string) *DescribeAvailableZonesRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeAvailableZonesRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeAvailableZonesRequest
	GetZoneId() *string
}

type DescribeAvailableZonesRequest struct {
	Category       *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CommodityCode  *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DispenseMode   *string `json:"DispenseMode,omitempty" xml:"DispenseMode,omitempty"`
	// This parameter is required.
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// This parameter is required.
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId          *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeAvailableZonesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeAvailableZonesRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAvailableZonesRequest) GetDispenseMode() *string {
	return s.DispenseMode
}

func (s *DescribeAvailableZonesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAvailableZonesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAvailableZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableZonesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAvailableZonesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableZonesRequest) SetCategory(v string) *DescribeAvailableZonesRequest {
	s.Category = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetCommodityCode(v string) *DescribeAvailableZonesRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetDBInstanceName(v string) *DescribeAvailableZonesRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetDispenseMode(v string) *DescribeAvailableZonesRequest {
	s.DispenseMode = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetEngine(v string) *DescribeAvailableZonesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetEngineVersion(v string) *DescribeAvailableZonesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetRegionId(v string) *DescribeAvailableZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetResourceOwnerId(v int64) *DescribeAvailableZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableZonesRequest) SetZoneId(v string) *DescribeAvailableZonesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableZonesRequest) Validate() error {
	return dara.Validate(s)
}
