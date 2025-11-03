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
	// The regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4E46C22C-D3B7-4DB8-9C76-63851BE68E20
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
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegions struct {
	DdsRegion []*DescribeRegionsResponseBodyRegionsDdsRegion `json:"DdsRegion,omitempty" xml:"DdsRegion,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetDdsRegion() []*DescribeRegionsResponseBodyRegionsDdsRegion {
	return s.DdsRegion
}

func (s *DescribeRegionsResponseBodyRegions) SetDdsRegion(v []*DescribeRegionsResponseBodyRegionsDdsRegion) *DescribeRegionsResponseBodyRegions {
	s.DdsRegion = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	if s.DdsRegion != nil {
		for _, item := range s.DdsRegion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionsDdsRegion struct {
	// The public endpoint of the region.
	//
	// For example, if the value of the RegionId parameter in the response is cn-hangzhou, the following value is returned for the EndPoint parameter:
	//
	// 	- mongodb.aliyuncs.com
	//
	// example:
	//
	// mongodb.aliyuncs.com
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// The value of the LocalName parameter is in the language that is specified by the **AcceptLanguage*	- parameter. For example, if the value of the RegionId parameter in the response is **cn-hangzhou**, the following values are returned for the LocalName parameter:
	//
	// 	- If the value of the **AcceptLanguage*	- parameter is **zh**, the value **华东1（杭州）*	- is returned for the LocalName parameter.
	//
	// 	- If the value of the **AcceptLanguage*	- parameter is **en**, the value **China (Hangzhou)*	- is returned for the LocalName parameter.
	//
	// example:
	//
	// China (Hangzhou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The zones.
	Zones *DescribeRegionsResponseBodyRegionsDdsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyRegionsDdsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsDdsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) GetEndPoint() *string {
	return s.EndPoint
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) GetZones() *DescribeRegionsResponseBodyRegionsDdsRegionZones {
	return s.Zones
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) SetEndPoint(v string) *DescribeRegionsResponseBodyRegionsDdsRegion {
	s.EndPoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsDdsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) SetRegionName(v string) *DescribeRegionsResponseBodyRegionsDdsRegion {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsDdsRegionZones) *DescribeRegionsResponseBodyRegionsDdsRegion {
	s.Zones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegion) Validate() error {
	if s.Zones != nil {
		if err := s.Zones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionsDdsRegionZones struct {
	Zone []*DescribeRegionsResponseBodyRegionsDdsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsDdsRegionZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsDdsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZones) GetZone() []*DescribeRegionsResponseBodyRegionsDdsRegionZonesZone {
	return s.Zone
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZones) SetZone(v []*DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) *DescribeRegionsResponseBodyRegionsDdsRegionZones {
	s.Zone = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZones) Validate() error {
	if s.Zone != nil {
		for _, item := range s.Zone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionsDdsRegionZonesZone struct {
	// Indicates whether a virtual private cloud (VPC) is supported. Valid values:
	//
	// 	- **true**: VPC is supported.
	//
	// 	- **false**: VPC is not supported.
	//
	// example:
	//
	// true
	VpcEnabled *bool `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The name of the zone.
	//
	// The value of the ZoneName parameter is in the language that is specified by the **AcceptLanguage*	- parameter. For example, if the value of the ZoneId parameter in the response is **cn-hangzhou-h**, the following values are returned for the ZoneName parameter:
	//
	// 	- If the value of the **AcceptLanguage*	- parameter is **zh**, the value **H*	- is returned for the ZoneName parameter.
	//
	// 	- If the value of the **AcceptLanguage*	- parameter is **en**, the value **Hangzhou Zone H*	- is returned for the ZoneName parameter.
	//
	// example:
	//
	// Hangzhou Zone H
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) GetVpcEnabled() *bool {
	return s.VpcEnabled
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) SetZoneName(v string) *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone {
	s.ZoneName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsDdsRegionZonesZone) Validate() error {
	return dara.Validate(s)
}
