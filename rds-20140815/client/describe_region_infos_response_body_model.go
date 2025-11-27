// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v *DescribeRegionInfosResponseBodyRegions) *DescribeRegionInfosResponseBody
	GetRegions() *DescribeRegionInfosResponseBodyRegions
	SetRequestId(v string) *DescribeRegionInfosResponseBody
	GetRequestId() *string
}

type DescribeRegionInfosResponseBody struct {
	// A list of regions.
	Regions *DescribeRegionInfosResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5414A4E5-4C36-4461-95FC-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionInfosResponseBody) GetRegions() *DescribeRegionInfosResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionInfosResponseBody) SetRegions(v *DescribeRegionInfosResponseBodyRegions) *DescribeRegionInfosResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionInfosResponseBody) SetRequestId(v string) *DescribeRegionInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionInfosResponseBody) Validate() error {
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionInfosResponseBodyRegions struct {
	RDSRegion []*DescribeRegionInfosResponseBodyRegionsRDSRegion `json:"RDSRegion,omitempty" xml:"RDSRegion,omitempty" type:"Repeated"`
}

func (s DescribeRegionInfosResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionInfosResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionInfosResponseBodyRegions) GetRDSRegion() []*DescribeRegionInfosResponseBodyRegionsRDSRegion {
	return s.RDSRegion
}

func (s *DescribeRegionInfosResponseBodyRegions) SetRDSRegion(v []*DescribeRegionInfosResponseBodyRegionsRDSRegion) *DescribeRegionInfosResponseBodyRegions {
	s.RDSRegion = v
	return s
}

func (s *DescribeRegionInfosResponseBodyRegions) Validate() error {
	if s.RDSRegion != nil {
		for _, item := range s.RDSRegion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionInfosResponseBodyRegionsRDSRegion struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionInfosResponseBodyRegionsRDSRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionInfosResponseBodyRegionsRDSRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionInfosResponseBodyRegionsRDSRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionInfosResponseBodyRegionsRDSRegion) SetRegionId(v string) *DescribeRegionInfosResponseBodyRegionsRDSRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionInfosResponseBodyRegionsRDSRegion) Validate() error {
	return dara.Validate(s)
}
