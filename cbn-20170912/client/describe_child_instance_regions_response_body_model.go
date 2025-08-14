// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChildInstanceRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v *DescribeChildInstanceRegionsResponseBodyRegions) *DescribeChildInstanceRegionsResponseBody
	GetRegions() *DescribeChildInstanceRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeChildInstanceRegionsResponseBody
	GetRequestId() *string
}

type DescribeChildInstanceRegionsResponseBody struct {
	// A list of regions.
	Regions *DescribeChildInstanceRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// D5CEED59-36AA-47CC-9D81-16F71C46BD80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChildInstanceRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChildInstanceRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsResponseBody) GetRegions() *DescribeChildInstanceRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeChildInstanceRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChildInstanceRegionsResponseBody) SetRegions(v *DescribeChildInstanceRegionsResponseBodyRegions) *DescribeChildInstanceRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeChildInstanceRegionsResponseBody) SetRequestId(v string) *DescribeChildInstanceRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChildInstanceRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeChildInstanceRegionsResponseBodyRegions struct {
	Region []*DescribeChildInstanceRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeChildInstanceRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeChildInstanceRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsResponseBodyRegions) GetRegion() []*DescribeChildInstanceRegionsResponseBodyRegionsRegion {
	return s.Region
}

func (s *DescribeChildInstanceRegionsResponseBodyRegions) SetRegion(v []*DescribeChildInstanceRegionsResponseBodyRegionsRegion) *DescribeChildInstanceRegionsResponseBodyRegions {
	s.Region = v
	return s
}

func (s *DescribeChildInstanceRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeChildInstanceRegionsResponseBodyRegionsRegion struct {
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeChildInstanceRegionsResponseBodyRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeChildInstanceRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeChildInstanceRegionsResponseBodyRegionsRegion) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeChildInstanceRegionsResponseBodyRegionsRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeChildInstanceRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeChildInstanceRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeChildInstanceRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeChildInstanceRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeChildInstanceRegionsResponseBodyRegionsRegion) Validate() error {
	return dara.Validate(s)
}
