// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRegionAndIspResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsps(v *DescribeDcdnRegionAndIspResponseBodyIsps) *DescribeDcdnRegionAndIspResponseBody
	GetIsps() *DescribeDcdnRegionAndIspResponseBodyIsps
	SetRegions(v *DescribeDcdnRegionAndIspResponseBodyRegions) *DescribeDcdnRegionAndIspResponseBody
	GetRegions() *DescribeDcdnRegionAndIspResponseBodyRegions
	SetRequestId(v string) *DescribeDcdnRegionAndIspResponseBody
	GetRequestId() *string
}

type DescribeDcdnRegionAndIspResponseBody struct {
	// The list of ISPs.
	Isps *DescribeDcdnRegionAndIspResponseBodyIsps `json:"Isps,omitempty" xml:"Isps,omitempty" type:"Struct"`
	// The list of regions.
	Regions *DescribeDcdnRegionAndIspResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2387C335-932C-4E1E-862C-1C4363B6DE72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnRegionAndIspResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponseBody) GetIsps() *DescribeDcdnRegionAndIspResponseBodyIsps {
	return s.Isps
}

func (s *DescribeDcdnRegionAndIspResponseBody) GetRegions() *DescribeDcdnRegionAndIspResponseBodyRegions {
	return s.Regions
}

func (s *DescribeDcdnRegionAndIspResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnRegionAndIspResponseBody) SetIsps(v *DescribeDcdnRegionAndIspResponseBodyIsps) *DescribeDcdnRegionAndIspResponseBody {
	s.Isps = v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBody) SetRegions(v *DescribeDcdnRegionAndIspResponseBodyRegions) *DescribeDcdnRegionAndIspResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBody) SetRequestId(v string) *DescribeDcdnRegionAndIspResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnRegionAndIspResponseBodyIsps struct {
	Isp []*DescribeDcdnRegionAndIspResponseBodyIspsIsp `json:"Isp,omitempty" xml:"Isp,omitempty" type:"Repeated"`
}

func (s DescribeDcdnRegionAndIspResponseBodyIsps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponseBodyIsps) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponseBodyIsps) GetIsp() []*DescribeDcdnRegionAndIspResponseBodyIspsIsp {
	return s.Isp
}

func (s *DescribeDcdnRegionAndIspResponseBodyIsps) SetIsp(v []*DescribeDcdnRegionAndIspResponseBodyIspsIsp) *DescribeDcdnRegionAndIspResponseBodyIsps {
	s.Isp = v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBodyIsps) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnRegionAndIspResponseBodyIspsIsp struct {
	// The English name of the region.
	//
	// example:
	//
	// unicom
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	// The Chinese name of the ISP.
	//
	// example:
	//
	// 联通
	NameZh *string `json:"NameZh,omitempty" xml:"NameZh,omitempty"`
}

func (s DescribeDcdnRegionAndIspResponseBodyIspsIsp) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponseBodyIspsIsp) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponseBodyIspsIsp) GetNameEn() *string {
	return s.NameEn
}

func (s *DescribeDcdnRegionAndIspResponseBodyIspsIsp) GetNameZh() *string {
	return s.NameZh
}

func (s *DescribeDcdnRegionAndIspResponseBodyIspsIsp) SetNameEn(v string) *DescribeDcdnRegionAndIspResponseBodyIspsIsp {
	s.NameEn = &v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBodyIspsIsp) SetNameZh(v string) *DescribeDcdnRegionAndIspResponseBodyIspsIsp {
	s.NameZh = &v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBodyIspsIsp) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnRegionAndIspResponseBodyRegions struct {
	Region []*DescribeDcdnRegionAndIspResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeDcdnRegionAndIspResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegions) GetRegion() []*DescribeDcdnRegionAndIspResponseBodyRegionsRegion {
	return s.Region
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegions) SetRegion(v []*DescribeDcdnRegionAndIspResponseBodyRegionsRegion) *DescribeDcdnRegionAndIspResponseBodyRegions {
	s.Region = v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnRegionAndIspResponseBodyRegionsRegion struct {
	// The English name of the region.
	//
	// example:
	//
	// liaoning
	NameEn *string `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	// The Chinese name of the region.
	//
	// example:
	//
	// 辽宁省
	NameZh *string `json:"NameZh,omitempty" xml:"NameZh,omitempty"`
}

func (s DescribeDcdnRegionAndIspResponseBodyRegionsRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegionsRegion) GetNameEn() *string {
	return s.NameEn
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegionsRegion) GetNameZh() *string {
	return s.NameZh
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegionsRegion) SetNameEn(v string) *DescribeDcdnRegionAndIspResponseBodyRegionsRegion {
	s.NameEn = &v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegionsRegion) SetNameZh(v string) *DescribeDcdnRegionAndIspResponseBodyRegionsRegion {
	s.NameZh = &v
	return s
}

func (s *DescribeDcdnRegionAndIspResponseBodyRegionsRegion) Validate() error {
	return dara.Validate(s)
}
