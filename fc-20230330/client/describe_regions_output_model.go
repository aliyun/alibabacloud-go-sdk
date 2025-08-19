// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v *DescribeRegionsOutputRegions) *DescribeRegionsOutput
	GetRegions() *DescribeRegionsOutputRegions
}

type DescribeRegionsOutput struct {
	Regions *DescribeRegionsOutputRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsOutput) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsOutput) GoString() string {
	return s.String()
}

func (s *DescribeRegionsOutput) GetRegions() *DescribeRegionsOutputRegions {
	return s.Regions
}

func (s *DescribeRegionsOutput) SetRegions(v *DescribeRegionsOutputRegions) *DescribeRegionsOutput {
	s.Regions = v
	return s
}

func (s *DescribeRegionsOutput) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsOutputRegions struct {
	Region []*DescribeRegionsOutputRegionsRegion `json:"Region" xml:"Region" type:"Repeated"`
}

func (s DescribeRegionsOutputRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsOutputRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsOutputRegions) GetRegion() []*DescribeRegionsOutputRegionsRegion {
	return s.Region
}

func (s *DescribeRegionsOutputRegions) SetRegion(v []*DescribeRegionsOutputRegionsRegion) *DescribeRegionsOutputRegions {
	s.Region = v
	return s
}

func (s *DescribeRegionsOutputRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsOutputRegionsRegion struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsOutputRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsOutputRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsOutputRegionsRegion) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsOutputRegionsRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsOutputRegionsRegion) SetLocalName(v string) *DescribeRegionsOutputRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsOutputRegionsRegion) SetRegionId(v string) *DescribeRegionsOutputRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsOutputRegionsRegion) Validate() error {
	return dara.Validate(s)
}
