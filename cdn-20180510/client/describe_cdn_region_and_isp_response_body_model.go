// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnRegionAndIspResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsps(v *DescribeCdnRegionAndIspResponseBodyIsps) *DescribeCdnRegionAndIspResponseBody
	GetIsps() *DescribeCdnRegionAndIspResponseBodyIsps
	SetRegions(v *DescribeCdnRegionAndIspResponseBodyRegions) *DescribeCdnRegionAndIspResponseBody
	GetRegions() *DescribeCdnRegionAndIspResponseBodyRegions
	SetRequestId(v string) *DescribeCdnRegionAndIspResponseBody
	GetRequestId() *string
}

type DescribeCdnRegionAndIspResponseBody struct {
	// The list of ISPs.
	Isps *DescribeCdnRegionAndIspResponseBodyIsps `json:"Isps,omitempty" xml:"Isps,omitempty" type:"Struct"`
	// The list of regions.
	Regions *DescribeCdnRegionAndIspResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2387C335-932C-4E1E-862C-1C4363B6DE72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnRegionAndIspResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponseBody) GetIsps() *DescribeCdnRegionAndIspResponseBodyIsps {
	return s.Isps
}

func (s *DescribeCdnRegionAndIspResponseBody) GetRegions() *DescribeCdnRegionAndIspResponseBodyRegions {
	return s.Regions
}

func (s *DescribeCdnRegionAndIspResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnRegionAndIspResponseBody) SetIsps(v *DescribeCdnRegionAndIspResponseBodyIsps) *DescribeCdnRegionAndIspResponseBody {
	s.Isps = v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBody) SetRegions(v *DescribeCdnRegionAndIspResponseBodyRegions) *DescribeCdnRegionAndIspResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBody) SetRequestId(v string) *DescribeCdnRegionAndIspResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnRegionAndIspResponseBodyIsps struct {
	Isp []*DescribeCdnRegionAndIspResponseBodyIspsIsp `json:"Isp,omitempty" xml:"Isp,omitempty" type:"Repeated"`
}

func (s DescribeCdnRegionAndIspResponseBodyIsps) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponseBodyIsps) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponseBodyIsps) GetIsp() []*DescribeCdnRegionAndIspResponseBodyIspsIsp {
	return s.Isp
}

func (s *DescribeCdnRegionAndIspResponseBodyIsps) SetIsp(v []*DescribeCdnRegionAndIspResponseBodyIspsIsp) *DescribeCdnRegionAndIspResponseBodyIsps {
	s.Isp = v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBodyIsps) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnRegionAndIspResponseBodyIspsIsp struct {
	// The English name of the ISP.
	//
	// example:
	//
	// unicom
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	// The Chinese name of the ISP.
	NameZh *string `json:"NameZh,omitempty" xml:"NameZh,omitempty"`
}

func (s DescribeCdnRegionAndIspResponseBodyIspsIsp) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponseBodyIspsIsp) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponseBodyIspsIsp) GetNameEn() *string {
	return s.NameEn
}

func (s *DescribeCdnRegionAndIspResponseBodyIspsIsp) GetNameZh() *string {
	return s.NameZh
}

func (s *DescribeCdnRegionAndIspResponseBodyIspsIsp) SetNameEn(v string) *DescribeCdnRegionAndIspResponseBodyIspsIsp {
	s.NameEn = &v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBodyIspsIsp) SetNameZh(v string) *DescribeCdnRegionAndIspResponseBodyIspsIsp {
	s.NameZh = &v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBodyIspsIsp) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnRegionAndIspResponseBodyRegions struct {
	Region []*DescribeCdnRegionAndIspResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeCdnRegionAndIspResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponseBodyRegions) GetRegion() []*DescribeCdnRegionAndIspResponseBodyRegionsRegion {
	return s.Region
}

func (s *DescribeCdnRegionAndIspResponseBodyRegions) SetRegion(v []*DescribeCdnRegionAndIspResponseBodyRegionsRegion) *DescribeCdnRegionAndIspResponseBodyRegions {
	s.Region = v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnRegionAndIspResponseBodyRegionsRegion struct {
	// The English name of the region.
	//
	// example:
	//
	// liaoning
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	// The Chinese name of the region.
	NameZh *string `json:"NameZh,omitempty" xml:"NameZh,omitempty"`
}

func (s DescribeCdnRegionAndIspResponseBodyRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnRegionAndIspResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspResponseBodyRegionsRegion) GetNameEn() *string {
	return s.NameEn
}

func (s *DescribeCdnRegionAndIspResponseBodyRegionsRegion) GetNameZh() *string {
	return s.NameZh
}

func (s *DescribeCdnRegionAndIspResponseBodyRegionsRegion) SetNameEn(v string) *DescribeCdnRegionAndIspResponseBodyRegionsRegion {
	s.NameEn = &v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBodyRegionsRegion) SetNameZh(v string) *DescribeCdnRegionAndIspResponseBodyRegionsRegion {
	s.NameZh = &v
	return s
}

func (s *DescribeCdnRegionAndIspResponseBodyRegionsRegion) Validate() error {
	return dara.Validate(s)
}
