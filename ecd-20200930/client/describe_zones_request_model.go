// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeZonesRequest
	GetRegionId() *string
	SetVpcAccessZone(v bool) *DescribeZonesRequest
	GetVpcAccessZone() *bool
	SetZoneType(v string) *DescribeZonesRequest
	GetZoneType() *string
}

type DescribeZonesRequest struct {
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcAccessZone *bool   `json:"VpcAccessZone,omitempty" xml:"VpcAccessZone,omitempty"`
	// The type of the zone. Default value: `AvailabilityZone`. This value indicates Alibaba Cloud zones.
	//
	// example:
	//
	// AvailabilityZone
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesRequest) GetVpcAccessZone() *bool {
	return s.VpcAccessZone
}

func (s *DescribeZonesRequest) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetVpcAccessZone(v bool) *DescribeZonesRequest {
	s.VpcAccessZone = &v
	return s
}

func (s *DescribeZonesRequest) SetZoneType(v string) *DescribeZonesRequest {
	s.ZoneType = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
