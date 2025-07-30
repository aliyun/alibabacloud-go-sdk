// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionModelList(v []*DescribeRegionsResponseBodyRegionModelList) *DescribeRegionsResponseBody
	GetRegionModelList() []*DescribeRegionsResponseBodyRegionModelList
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// An array of regions.
	RegionModelList []*DescribeRegionsResponseBodyRegionModelList `json:"RegionModelList,omitempty" xml:"RegionModelList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F8900A96-67F7-5274-A41B-7722E1ECF8C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRegionModelList() []*DescribeRegionsResponseBodyRegionModelList {
	return s.RegionModelList
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetRegionModelList(v []*DescribeRegionsResponseBodyRegionModelList) *DescribeRegionsResponseBody {
	s.RegionModelList = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegionModelList struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// An array of zones.
	Zones []*DescribeRegionsResponseBodyRegionModelListZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionModelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionModelList) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionModelList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegionModelList) GetZones() []*DescribeRegionsResponseBodyRegionModelListZones {
	return s.Zones
}

func (s *DescribeRegionsResponseBodyRegionModelList) SetRegionId(v string) *DescribeRegionsResponseBodyRegionModelList {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModelList) SetZones(v []*DescribeRegionsResponseBodyRegionModelListZones) *DescribeRegionsResponseBodyRegionModelList {
	s.Zones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModelList) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionsResponseBodyRegionModelListZones struct {
	// The zone description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the VPC is disabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The label.
	//
	// example:
	//
	// test
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The zone name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The subdomain.
	//
	// example:
	//
	// cn-beijing-h-aliyun
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// Indicates whether the VPC is enabled.
	//
	// example:
	//
	// true
	VpcEnabled *bool `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
	// Indicates whether the virtual private cloud (VPC) is available.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionModelListZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionModelListZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) GetDescription() *string {
	return s.Description
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) GetLabel() *string {
	return s.Label
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) GetName() *string {
	return s.Name
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) GetVpcEnabled() *bool {
	return s.VpcEnabled
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) SetDescription(v string) *DescribeRegionsResponseBodyRegionModelListZones {
	s.Description = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) SetDisabled(v bool) *DescribeRegionsResponseBodyRegionModelListZones {
	s.Disabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) SetLabel(v string) *DescribeRegionsResponseBodyRegionModelListZones {
	s.Label = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) SetName(v string) *DescribeRegionsResponseBodyRegionModelListZones {
	s.Name = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) SetRegionId(v string) *DescribeRegionsResponseBodyRegionModelListZones {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) SetSubDomain(v string) *DescribeRegionsResponseBodyRegionModelListZones {
	s.SubDomain = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionModelListZones {
	s.VpcEnabled = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) SetZoneId(v string) *DescribeRegionsResponseBodyRegionModelListZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModelListZones) Validate() error {
	return dara.Validate(s)
}
