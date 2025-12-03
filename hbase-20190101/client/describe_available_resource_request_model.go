// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeAvailableResourceRequest
	GetChargeType() *string
	SetCoreInstanceType(v string) *DescribeAvailableResourceRequest
	GetCoreInstanceType() *string
	SetDiskType(v string) *DescribeAvailableResourceRequest
	GetDiskType() *string
	SetEngine(v string) *DescribeAvailableResourceRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeAvailableResourceRequest
	GetEngineVersion() *string
	SetRegionId(v string) *DescribeAvailableResourceRequest
	GetRegionId() *string
	SetZoneId(v string) *DescribeAvailableResourceRequest
	GetZoneId() *string
}

type DescribeAvailableResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeAvailableResourceRequest) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *DescribeAvailableResourceRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeAvailableResourceRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeAvailableResourceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeAvailableResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableResourceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAvailableResourceRequest) SetChargeType(v string) *DescribeAvailableResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetCoreInstanceType(v string) *DescribeAvailableResourceRequest {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetDiskType(v string) *DescribeAvailableResourceRequest {
	s.DiskType = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetEngine(v string) *DescribeAvailableResourceRequest {
	s.Engine = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetEngineVersion(v string) *DescribeAvailableResourceRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) Validate() error {
	return dara.Validate(s)
}
