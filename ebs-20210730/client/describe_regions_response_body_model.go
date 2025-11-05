// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() []*DescribeRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// Details about the regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 17EE62D8-064E-5404-8B0D-72122478****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRegions() []*DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegions struct {
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	//
	// example:
	//
	// ebs.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Details about the zones.
	Zones []*DescribeRegionsResponseBodyRegionsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyRegions) GetRegionEndpoint() *string {
	return s.RegionEndpoint
}

func (s *DescribeRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegions) GetZones() []*DescribeRegionsResponseBodyRegionsZones {
	return s.Zones
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetZones(v []*DescribeRegionsResponseBodyRegionsZones) *DescribeRegionsResponseBodyRegions {
	s.Zones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	if s.Zones != nil {
		for _, item := range s.Zones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionsZones struct {
	// The name of the zone.
	//
	// example:
	//
	// Hangzhou Zone H
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The type of resource list.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsZones) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyRegionsZones) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *DescribeRegionsResponseBodyRegionsZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetResourceTypes(v []*string) *DescribeRegionsResponseBodyRegionsZones {
	s.ResourceTypes = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsZones) Validate() error {
	return dara.Validate(s)
}
