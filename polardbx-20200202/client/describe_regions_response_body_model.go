// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeRegionsResponseBody
	GetCode() *int32
	SetErrorCode(v int32) *DescribeRegionsResponseBody
	GetErrorCode() *int32
	SetMessage(v string) *DescribeRegionsResponseBody
	GetMessage() *string
	SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() *DescribeRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRegionsResponseBody
	GetSuccess() *bool
}

type DescribeRegionsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeRegionsResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DescribeRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRegionsResponseBody) GetRegions() *DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRegionsResponseBody) SetCode(v int32) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrorCode(v int32) *DescribeRegionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetMessage(v string) *DescribeRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
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

type DescribeRegionsResponseBodyRegionsRegion struct {
	// example:
	//
	// ch-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// true
	SupportPolarx10 *bool `json:"SupportPolarx10,omitempty" xml:"SupportPolarx10,omitempty"`
	// example:
	//
	// true
	SupportPolarx20 *bool                                          `json:"SupportPolarx20,omitempty" xml:"SupportPolarx20,omitempty"`
	Zones           *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
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

func (s *DescribeRegionsResponseBodyRegionsRegion) GetSupportPolarx10() *bool {
	return s.SupportPolarx10
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetSupportPolarx20() *bool {
	return s.SupportPolarx20
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetZones() *DescribeRegionsResponseBodyRegionsRegionZones {
	return s.Zones
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetSupportPolarx10(v bool) *DescribeRegionsResponseBodyRegionsRegion {
	s.SupportPolarx10 = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetSupportPolarx20(v bool) *DescribeRegionsResponseBodyRegionsRegion {
	s.SupportPolarx20 = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) Validate() error {
	if s.Zones != nil {
		if err := s.Zones.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type DescribeRegionsResponseBodyRegionsRegionZonesZone struct {
	// example:
	//
	// true
	VpcEnabled *bool `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	// example:
	//
	// cn-hangzhou-a
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
