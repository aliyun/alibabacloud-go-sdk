// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiZoneAvailableRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v *DescribeMultiZoneAvailableRegionsResponseBodyRegions) *DescribeMultiZoneAvailableRegionsResponseBody
	GetRegions() *DescribeMultiZoneAvailableRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeMultiZoneAvailableRegionsResponseBody
	GetRequestId() *string
}

type DescribeMultiZoneAvailableRegionsResponseBody struct {
	Regions *DescribeMultiZoneAvailableRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// example:
	//
	// F03BB273-45EE-4B6C-A329-A6E6A8D15856
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBody) GetRegions() *DescribeMultiZoneAvailableRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeMultiZoneAvailableRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMultiZoneAvailableRegionsResponseBody) SetRegions(v *DescribeMultiZoneAvailableRegionsResponseBodyRegions) *DescribeMultiZoneAvailableRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBody) SetRequestId(v string) *DescribeMultiZoneAvailableRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegions struct {
	Region []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegions) GetRegion() []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	return s.Region
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegions) SetRegion(v []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) *DescribeMultiZoneAvailableRegionsResponseBodyRegions {
	s.Region = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegions) Validate() error {
	if s.Region != nil {
		for _, item := range s.Region {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion struct {
	AvailableCombines *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines `json:"AvailableCombines,omitempty" xml:"AvailableCombines,omitempty" type:"Struct"`
	LocalName         *string                                                                      `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// hbase.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) GetAvailableCombines() *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines {
	return s.AvailableCombines
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) GetRegionEndpoint() *string {
	return s.RegionEndpoint
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetAvailableCombines(v *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.AvailableCombines = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegion) Validate() error {
	if s.AvailableCombines != nil {
		if err := s.AvailableCombines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines struct {
	AvailableCombine []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine `json:"AvailableCombine,omitempty" xml:"AvailableCombine,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) GetAvailableCombine() []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine {
	return s.AvailableCombine
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) SetAvailableCombine(v []*DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines {
	s.AvailableCombine = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombines) Validate() error {
	if s.AvailableCombine != nil {
		for _, item := range s.AvailableCombine {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine struct {
	// example:
	//
	// cn-shenzhen-****-aliyun
	Id    *string                                                                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	Zones *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) GetId() *string {
	return s.Id
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) GetZones() *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones {
	return s.Zones
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) SetId(v string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine {
	s.Id = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) SetZones(v *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine {
	s.Zones = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombine) Validate() error {
	if s.Zones != nil {
		if err := s.Zones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones struct {
	Zone []*string `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) GetZone() []*string {
	return s.Zone
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) SetZone(v []*string) *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones {
	s.Zone = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponseBodyRegionsRegionAvailableCombinesAvailableCombineZones) Validate() error {
	return dara.Validate(s)
}
