// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() *DescribeRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// Details of the regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FF8EB261-5447-4B1B-9F14-294CEA008A9F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRegions() *DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetRegion() []*DescribeRegionsResponseBodyRegionsRegion {
	return s.Region
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Details of the zones.
	Zones *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetZones() *DescribeRegionsResponseBodyRegionsRegionZones {
	return s.Zones
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegionsRegionZones struct {
	Zone []*DescribeRegionsResponseBodyRegionsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZones) GetZone() []*DescribeRegionsResponseBodyRegionsRegionZonesZone {
	return s.Zone
}

func (s *DescribeRegionsResponseBodyRegionsRegionZones) SetZone(v []*DescribeRegionsResponseBodyRegionsRegionZonesZone) *DescribeRegionsResponseBodyRegionsRegionZones {
	s.Zone = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZones) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegionsRegionZonesZone struct {
	// Indicates whether Virtual Private Cloud (VPC) is available.
	//
	// 	- **true**: VPC is available.
	//
	// 	- **false**: VPC is unavailable.
	//
	// example:
	//
	// true
	VpcEnabled *bool `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) GetVpcEnabled() *bool {
	return s.VpcEnabled
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) Validate() error {
	return dara.Validate(s)
}
