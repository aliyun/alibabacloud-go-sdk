// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyntheticProbeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSyntheticProbeListResponseBody
	GetCode() *string
	SetIspCityList(v []*DescribeSyntheticProbeListResponseBodyIspCityList) *DescribeSyntheticProbeListResponseBody
	GetIspCityList() []*DescribeSyntheticProbeListResponseBodyIspCityList
	SetMessage(v string) *DescribeSyntheticProbeListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSyntheticProbeListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSyntheticProbeListResponseBody
	GetSuccess() *string
}

type DescribeSyntheticProbeListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried detection points.
	IspCityList []*DescribeSyntheticProbeListResponseBodyIspCityList `json:"IspCityList,omitempty" xml:"IspCityList,omitempty" type:"Repeated"`
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 87170bc7-e28a-4c93-b9bf-90a1dbe84736
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSyntheticProbeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyntheticProbeListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyntheticProbeListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSyntheticProbeListResponseBody) GetIspCityList() []*DescribeSyntheticProbeListResponseBodyIspCityList {
	return s.IspCityList
}

func (s *DescribeSyntheticProbeListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSyntheticProbeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSyntheticProbeListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSyntheticProbeListResponseBody) SetCode(v string) *DescribeSyntheticProbeListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBody) SetIspCityList(v []*DescribeSyntheticProbeListResponseBodyIspCityList) *DescribeSyntheticProbeListResponseBody {
	s.IspCityList = v
	return s
}

func (s *DescribeSyntheticProbeListResponseBody) SetMessage(v string) *DescribeSyntheticProbeListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBody) SetRequestId(v string) *DescribeSyntheticProbeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBody) SetSuccess(v string) *DescribeSyntheticProbeListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBody) Validate() error {
	if s.IspCityList != nil {
		for _, item := range s.IspCityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSyntheticProbeListResponseBodyIspCityList struct {
	AreaCn *string `json:"AreaCn,omitempty" xml:"AreaCn,omitempty"`
	// example:
	//
	// Huabei
	AreaEn *string `json:"AreaEn,omitempty" xml:"AreaEn,omitempty"`
	// example:
	//
	// 738
	City   *string `json:"City,omitempty" xml:"City,omitempty"`
	CityCn *string `json:"CityCn,omitempty" xml:"CityCn,omitempty"`
	// example:
	//
	// Beijing
	CityEn *string `json:"CityEn,omitempty" xml:"CityEn,omitempty"`
	// example:
	//
	// 629
	Country   *string `json:"Country,omitempty" xml:"Country,omitempty"`
	CountryCn *string `json:"CountryCn,omitempty" xml:"CountryCn,omitempty"`
	// example:
	//
	// China
	CountryEn *string `json:"CountryEn,omitempty" xml:"CountryEn,omitempty"`
	// example:
	//
	// 1
	IdcV4ProbeCount *int32 `json:"IdcV4ProbeCount,omitempty" xml:"IdcV4ProbeCount,omitempty"`
	// The number of IPv6 nodes in data centers.
	//
	// example:
	//
	// 1
	IdcV6ProbeCount *int32 `json:"IdcV6ProbeCount,omitempty" xml:"IdcV6ProbeCount,omitempty"`
	// The IP addresses of the monitored nodes.
	IpPool []*string `json:"IpPool,omitempty" xml:"IpPool,omitempty" type:"Repeated"`
	// example:
	//
	// 232
	Isp   *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	IspCn *string `json:"IspCn,omitempty" xml:"IspCn,omitempty"`
	// example:
	//
	// China-Unicom
	IspEn *string `json:"IspEn,omitempty" xml:"IspEn,omitempty"`
	// example:
	//
	// 1
	LmProbeCount *int32 `json:"LmProbeCount,omitempty" xml:"LmProbeCount,omitempty"`
	// example:
	//
	// 1
	MbProbeCount *int32 `json:"MbProbeCount,omitempty" xml:"MbProbeCount,omitempty"`
	// example:
	//
	// 264
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionCn *string `json:"RegionCn,omitempty" xml:"RegionCn,omitempty"`
	// example:
	//
	// Jiangxi
	RegionEn *string `json:"RegionEn,omitempty" xml:"RegionEn,omitempty"`
}

func (s DescribeSyntheticProbeListResponseBodyIspCityList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyntheticProbeListResponseBodyIspCityList) GoString() string {
	return s.String()
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetAreaCn() *string {
	return s.AreaCn
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetAreaEn() *string {
	return s.AreaEn
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetCity() *string {
	return s.City
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetCityCn() *string {
	return s.CityCn
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetCityEn() *string {
	return s.CityEn
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetCountry() *string {
	return s.Country
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetCountryCn() *string {
	return s.CountryCn
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetCountryEn() *string {
	return s.CountryEn
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetIdcV4ProbeCount() *int32 {
	return s.IdcV4ProbeCount
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetIdcV6ProbeCount() *int32 {
	return s.IdcV6ProbeCount
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetIpPool() []*string {
	return s.IpPool
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetIsp() *string {
	return s.Isp
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetIspCn() *string {
	return s.IspCn
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetIspEn() *string {
	return s.IspEn
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetLmProbeCount() *int32 {
	return s.LmProbeCount
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetMbProbeCount() *int32 {
	return s.MbProbeCount
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetRegion() *string {
	return s.Region
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetRegionCn() *string {
	return s.RegionCn
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) GetRegionEn() *string {
	return s.RegionEn
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetAreaCn(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.AreaCn = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetAreaEn(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.AreaEn = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetCity(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.City = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetCityCn(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.CityCn = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetCityEn(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.CityEn = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetCountry(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.Country = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetCountryCn(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.CountryCn = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetCountryEn(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.CountryEn = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetIdcV4ProbeCount(v int32) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.IdcV4ProbeCount = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetIdcV6ProbeCount(v int32) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.IdcV6ProbeCount = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetIpPool(v []*string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.IpPool = v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetIsp(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.Isp = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetIspCn(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.IspCn = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetIspEn(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.IspEn = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetLmProbeCount(v int32) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.LmProbeCount = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetMbProbeCount(v int32) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.MbProbeCount = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetRegion(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.Region = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetRegionCn(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.RegionCn = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) SetRegionEn(v string) *DescribeSyntheticProbeListResponseBodyIspCityList {
	s.RegionEn = &v
	return s
}

func (s *DescribeSyntheticProbeListResponseBodyIspCityList) Validate() error {
	return dara.Validate(s)
}
