// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePrivateLineAreasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveAreasList(v *DescribeLivePrivateLineAreasResponseBodyLiveAreasList) *DescribeLivePrivateLineAreasResponseBody
	GetLiveAreasList() *DescribeLivePrivateLineAreasResponseBodyLiveAreasList
	SetRequestId(v string) *DescribeLivePrivateLineAreasResponseBody
	GetRequestId() *string
}

type DescribeLivePrivateLineAreasResponseBody struct {
	// Details about the regions.
	LiveAreasList *DescribeLivePrivateLineAreasResponseBodyLiveAreasList `json:"LiveAreasList,omitempty" xml:"LiveAreasList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C4865B85-664B-19D3-BB16-C62FB83C8226
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLivePrivateLineAreasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAreasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAreasResponseBody) GetLiveAreasList() *DescribeLivePrivateLineAreasResponseBodyLiveAreasList {
	return s.LiveAreasList
}

func (s *DescribeLivePrivateLineAreasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLivePrivateLineAreasResponseBody) SetLiveAreasList(v *DescribeLivePrivateLineAreasResponseBodyLiveAreasList) *DescribeLivePrivateLineAreasResponseBody {
	s.LiveAreasList = v
	return s
}

func (s *DescribeLivePrivateLineAreasResponseBody) SetRequestId(v string) *DescribeLivePrivateLineAreasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLivePrivateLineAreasResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLivePrivateLineAreasResponseBodyLiveAreasList struct {
	LiveArea []*DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea `json:"LiveArea,omitempty" xml:"LiveArea,omitempty" type:"Repeated"`
}

func (s DescribeLivePrivateLineAreasResponseBodyLiveAreasList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAreasResponseBodyLiveAreasList) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasList) GetLiveArea() []*DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea {
	return s.LiveArea
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasList) SetLiveArea(v []*DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea) *DescribeLivePrivateLineAreasResponseBodyLiveAreasList {
	s.LiveArea = v
	return s
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasList) Validate() error {
	return dara.Validate(s)
}

type DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea struct {
	// The region type. Valid values:
	//
	// 	- domestic: in the Chinese mainland
	//
	// 	- overseas: outside the Chinese mainland
	//
	// example:
	//
	// domestic
	RegionType *string `json:"RegionType,omitempty" xml:"RegionType,omitempty"`
	// The regions.
	Regions *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea) GetRegionType() *string {
	return s.RegionType
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea) GetRegions() *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegions {
	return s.Regions
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea) SetRegionType(v string) *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea {
	s.RegionType = &v
	return s
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea) SetRegions(v *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegions) *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea {
	s.Regions = v
	return s
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveArea) Validate() error {
	return dara.Validate(s)
}

type DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegions struct {
	Region []*DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegions) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegions) GetRegion() []*DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion {
	return s.Region
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegions) SetRegion(v []*DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion) *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegions {
	s.Region = v
	return s
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion struct {
	// The region name.
	//
	// example:
	//
	// cn-shenzhen
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion) SetLocalName(v string) *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion) SetRegionId(v string) *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeLivePrivateLineAreasResponseBodyLiveAreasListLiveAreaRegionsRegion) Validate() error {
	return dara.Validate(s)
}
