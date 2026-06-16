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
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5093C7EE-8E27-5FC9-9B88-40626BA540C0
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
	RegionEnName  *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	RegionName    *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	RegionNo      *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	RegionNoAlias *string `json:"RegionNoAlias,omitempty" xml:"RegionNoAlias,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetRegionEnName() *string {
	return s.RegionEnName
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeRegionsResponseBodyRegionsRegion) GetRegionNoAlias() *string {
	return s.RegionNoAlias
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEnName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionNo(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionNo = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionNoAlias(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionNoAlias = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) Validate() error {
	return dara.Validate(s)
}
