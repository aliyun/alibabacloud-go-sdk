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
	// The collection of region information.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 73469C7-AA6F-4DC5-B3DB-A3DC0DE3****
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
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegions struct {
	// Indicates whether the current region supports scaling groups that reside in the classic network. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ClassicUnavailable *bool `json:"ClassicUnavailable,omitempty" xml:"ClassicUnavailable,omitempty"`
	// The region name.
	//
	// example:
	//
	// China (Beijing)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The region endpoint.
	//
	// example:
	//
	// ess.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether the current region supports scaling groups that reside in virtual private clouds (VPCs). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	VpcUnavailable *bool `json:"VpcUnavailable,omitempty" xml:"VpcUnavailable,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetClassicUnavailable() *bool {
	return s.ClassicUnavailable
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

func (s *DescribeRegionsResponseBodyRegions) GetVpcUnavailable() *bool {
	return s.VpcUnavailable
}

func (s *DescribeRegionsResponseBodyRegions) SetClassicUnavailable(v bool) *DescribeRegionsResponseBodyRegions {
	s.ClassicUnavailable = &v
	return s
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

func (s *DescribeRegionsResponseBodyRegions) SetVpcUnavailable(v bool) *DescribeRegionsResponseBodyRegions {
	s.VpcUnavailable = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
