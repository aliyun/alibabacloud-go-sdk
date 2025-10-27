// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsMixedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v *DescribeRegionsMixedResponseBodyRegions) *DescribeRegionsMixedResponseBody
	GetRegions() *DescribeRegionsMixedResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsMixedResponseBody
	GetRequestId() *string
}

type DescribeRegionsMixedResponseBody struct {
	// The queried regions.
	Regions *DescribeRegionsMixedResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2E9450FF-F8AD-54C6-B3C3-009FBD7C0700
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsMixedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsMixedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsMixedResponseBody) GetRegions() *DescribeRegionsMixedResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsMixedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsMixedResponseBody) SetRegions(v *DescribeRegionsMixedResponseBodyRegions) *DescribeRegionsMixedResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsMixedResponseBody) SetRequestId(v string) *DescribeRegionsMixedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsMixedResponseBody) Validate() error {
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsMixedResponseBodyRegions struct {
	Region []*DescribeRegionsMixedResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsMixedResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsMixedResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsMixedResponseBodyRegions) GetRegion() []*DescribeRegionsMixedResponseBodyRegionsRegion {
	return s.Region
}

func (s *DescribeRegionsMixedResponseBodyRegions) SetRegion(v []*DescribeRegionsMixedResponseBodyRegionsRegion) *DescribeRegionsMixedResponseBodyRegions {
	s.Region = v
	return s
}

func (s *DescribeRegionsMixedResponseBodyRegions) Validate() error {
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

type DescribeRegionsMixedResponseBodyRegionsRegion struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The version of the cluster.
	//
	// Valid values:
	//
	// 	- 2.0
	//
	// 	- 3.0
	//
	// example:
	//
	// 3.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeRegionsMixedResponseBodyRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsMixedResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsMixedResponseBodyRegionsRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsMixedResponseBodyRegionsRegion) GetVersion() *string {
	return s.Version
}

func (s *DescribeRegionsMixedResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsMixedResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsMixedResponseBodyRegionsRegion) SetVersion(v string) *DescribeRegionsMixedResponseBodyRegionsRegion {
	s.Version = &v
	return s
}

func (s *DescribeRegionsMixedResponseBodyRegionsRegion) Validate() error {
	return dara.Validate(s)
}
