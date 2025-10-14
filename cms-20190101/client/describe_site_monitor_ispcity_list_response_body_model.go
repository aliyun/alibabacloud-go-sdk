// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorISPCityListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSiteMonitorISPCityListResponseBody
	GetCode() *string
	SetIspCityList(v *DescribeSiteMonitorISPCityListResponseBodyIspCityList) *DescribeSiteMonitorISPCityListResponseBody
	GetIspCityList() *DescribeSiteMonitorISPCityListResponseBodyIspCityList
	SetMessage(v string) *DescribeSiteMonitorISPCityListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSiteMonitorISPCityListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSiteMonitorISPCityListResponseBody
	GetSuccess() *string
}

type DescribeSiteMonitorISPCityListResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried detection points.
	IspCityList *DescribeSiteMonitorISPCityListResponseBodyIspCityList `json:"IspCityList,omitempty" xml:"IspCityList,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B35D7D84-547B-4E61-B909-48A1F8A0C756
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSiteMonitorISPCityListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorISPCityListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorISPCityListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSiteMonitorISPCityListResponseBody) GetIspCityList() *DescribeSiteMonitorISPCityListResponseBodyIspCityList {
	return s.IspCityList
}

func (s *DescribeSiteMonitorISPCityListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSiteMonitorISPCityListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteMonitorISPCityListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSiteMonitorISPCityListResponseBody) SetCode(v string) *DescribeSiteMonitorISPCityListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBody) SetIspCityList(v *DescribeSiteMonitorISPCityListResponseBodyIspCityList) *DescribeSiteMonitorISPCityListResponseBody {
	s.IspCityList = v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBody) SetMessage(v string) *DescribeSiteMonitorISPCityListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBody) SetRequestId(v string) *DescribeSiteMonitorISPCityListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBody) SetSuccess(v string) *DescribeSiteMonitorISPCityListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBody) Validate() error {
	if s.IspCityList != nil {
		if err := s.IspCityList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorISPCityListResponseBodyIspCityList struct {
	IspCity []*DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity `json:"IspCity,omitempty" xml:"IspCity,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorISPCityListResponseBodyIspCityList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorISPCityListResponseBodyIspCityList) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityList) GetIspCity() []*DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	return s.IspCity
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityList) SetIspCity(v []*DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) *DescribeSiteMonitorISPCityListResponseBodyIspCityList {
	s.IspCity = v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityList) Validate() error {
	if s.IspCity != nil {
		for _, item := range s.IspCity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity struct {
	Area_en    *string `json:"Area.en,omitempty" xml:"Area.en,omitempty"`
	Area_zh_cN *string `json:"Area.zh_CN,omitempty" xml:"Area.zh_CN,omitempty"`
	// The city ID.
	//
	// example:
	//
	// 4
	City           *string `json:"City,omitempty" xml:"City,omitempty"`
	CityName_en    *string `json:"CityName.en,omitempty" xml:"CityName.en,omitempty"`
	CityName_zh_cN *string `json:"CityName.zh_CN,omitempty" xml:"CityName.zh_CN,omitempty"`
	// The country name.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// 629
	Country       *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Country_en    *string `json:"Country.en,omitempty" xml:"Country.en,omitempty"`
	Country_zh_cN *string `json:"Country.zh_CN,omitempty" xml:"Country.zh_CN,omitempty"`
	// The IP address pool.
	IPPool *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCityIPPool `json:"IPPool,omitempty" xml:"IPPool,omitempty" type:"Struct"`
	// The number of IPv4 probes.
	//
	// example:
	//
	// 4
	IPV4ProbeCount *string `json:"IPV4ProbeCount,omitempty" xml:"IPV4ProbeCount,omitempty"`
	// The number of IPv6 probes.
	//
	// example:
	//
	// 3
	IPV6ProbeCount *string `json:"IPV6ProbeCount,omitempty" xml:"IPV6ProbeCount,omitempty"`
	// The carrier ID.
	//
	// example:
	//
	// 232
	Isp           *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	IspName_en    *string `json:"IspName.en,omitempty" xml:"IspName.en,omitempty"`
	IspName_zh_cN *string `json:"IspName.zh_CN,omitempty" xml:"IspName.zh_CN,omitempty"`
	// The province name.
	//
	// example:
	//
	// 264
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Region_en    *string `json:"Region.en,omitempty" xml:"Region.en,omitempty"`
	Region_zh_cN *string `json:"Region.zh_CN,omitempty" xml:"Region.zh_CN,omitempty"`
}

func (s DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetArea_en() *string {
	return s.Area_en
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetArea_zh_cN() *string {
	return s.Area_zh_cN
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetCity() *string {
	return s.City
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetCityName_en() *string {
	return s.CityName_en
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetCityName_zh_cN() *string {
	return s.CityName_zh_cN
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetCountry() *string {
	return s.Country
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetCountry_en() *string {
	return s.Country_en
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetCountry_zh_cN() *string {
	return s.Country_zh_cN
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetIPPool() *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCityIPPool {
	return s.IPPool
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetIPV4ProbeCount() *string {
	return s.IPV4ProbeCount
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetIPV6ProbeCount() *string {
	return s.IPV6ProbeCount
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetIsp() *string {
	return s.Isp
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetIspName_en() *string {
	return s.IspName_en
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetIspName_zh_cN() *string {
	return s.IspName_zh_cN
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetRegion() *string {
	return s.Region
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetRegion_en() *string {
	return s.Region_en
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) GetRegion_zh_cN() *string {
	return s.Region_zh_cN
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetArea_en(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.Area_en = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetArea_zh_cN(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.Area_zh_cN = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetCity(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.City = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetCityName_en(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.CityName_en = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetCityName_zh_cN(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.CityName_zh_cN = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetCountry(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.Country = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetCountry_en(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.Country_en = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetCountry_zh_cN(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.Country_zh_cN = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetIPPool(v *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCityIPPool) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.IPPool = v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetIPV4ProbeCount(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.IPV4ProbeCount = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetIPV6ProbeCount(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.IPV6ProbeCount = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetIsp(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.Isp = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetIspName_en(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.IspName_en = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetIspName_zh_cN(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.IspName_zh_cN = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetRegion(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.Region = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetRegion_en(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.Region_en = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) SetRegion_zh_cN(v string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity {
	s.Region_zh_cN = &v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCity) Validate() error {
	if s.IPPool != nil {
		if err := s.IPPool.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCityIPPool struct {
	IPPool []*string `json:"IPPool,omitempty" xml:"IPPool,omitempty" type:"Repeated"`
}

func (s DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCityIPPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCityIPPool) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCityIPPool) GetIPPool() []*string {
	return s.IPPool
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCityIPPool) SetIPPool(v []*string) *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCityIPPool {
	s.IPPool = v
	return s
}

func (s *DescribeSiteMonitorISPCityListResponseBodyIspCityListIspCityIPPool) Validate() error {
	return dara.Validate(s)
}
