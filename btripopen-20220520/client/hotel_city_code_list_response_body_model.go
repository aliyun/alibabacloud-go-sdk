// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelCityCodeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelCityCodeListResponseBody
	GetCode() *string
	SetMessage(v string) *HotelCityCodeListResponseBody
	GetMessage() *string
	SetModule(v []*HotelCityCodeListResponseBodyModule) *HotelCityCodeListResponseBody
	GetModule() []*HotelCityCodeListResponseBodyModule
	SetRequestId(v string) *HotelCityCodeListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelCityCodeListResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelCityCodeListResponseBody
	GetTraceId() *string
}

type HotelCityCodeListResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// response is empty.
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  []*HotelCityCodeListResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelCityCodeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelCityCodeListResponseBody) GoString() string {
	return s.String()
}

func (s *HotelCityCodeListResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelCityCodeListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelCityCodeListResponseBody) GetModule() []*HotelCityCodeListResponseBodyModule {
	return s.Module
}

func (s *HotelCityCodeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelCityCodeListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelCityCodeListResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelCityCodeListResponseBody) SetCode(v string) *HotelCityCodeListResponseBody {
	s.Code = &v
	return s
}

func (s *HotelCityCodeListResponseBody) SetMessage(v string) *HotelCityCodeListResponseBody {
	s.Message = &v
	return s
}

func (s *HotelCityCodeListResponseBody) SetModule(v []*HotelCityCodeListResponseBodyModule) *HotelCityCodeListResponseBody {
	s.Module = v
	return s
}

func (s *HotelCityCodeListResponseBody) SetRequestId(v string) *HotelCityCodeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelCityCodeListResponseBody) SetSuccess(v bool) *HotelCityCodeListResponseBody {
	s.Success = &v
	return s
}

func (s *HotelCityCodeListResponseBody) SetTraceId(v string) *HotelCityCodeListResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelCityCodeListResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelCityCodeListResponseBodyModule struct {
	Citys []*HotelCityCodeListResponseBodyModuleCitys `json:"citys,omitempty" xml:"citys,omitempty" type:"Repeated"`
	// example:
	//
	// 108800
	ProviceCode  *string `json:"provice_code,omitempty" xml:"provice_code,omitempty"`
	ProvinceName *string `json:"province_name,omitempty" xml:"province_name,omitempty"`
}

func (s HotelCityCodeListResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelCityCodeListResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelCityCodeListResponseBodyModule) GetCitys() []*HotelCityCodeListResponseBodyModuleCitys {
	return s.Citys
}

func (s *HotelCityCodeListResponseBodyModule) GetProviceCode() *string {
	return s.ProviceCode
}

func (s *HotelCityCodeListResponseBodyModule) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *HotelCityCodeListResponseBodyModule) SetCitys(v []*HotelCityCodeListResponseBodyModuleCitys) *HotelCityCodeListResponseBodyModule {
	s.Citys = v
	return s
}

func (s *HotelCityCodeListResponseBodyModule) SetProviceCode(v string) *HotelCityCodeListResponseBodyModule {
	s.ProviceCode = &v
	return s
}

func (s *HotelCityCodeListResponseBodyModule) SetProvinceName(v string) *HotelCityCodeListResponseBodyModule {
	s.ProvinceName = &v
	return s
}

func (s *HotelCityCodeListResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelCityCodeListResponseBodyModuleCitys struct {
	// example:
	//
	// 445222
	CityCode  *string                                              `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName  *string                                              `json:"city_name,omitempty" xml:"city_name,omitempty"`
	Districts []*HotelCityCodeListResponseBodyModuleCitysDistricts `json:"districts,omitempty" xml:"districts,omitempty" type:"Repeated"`
}

func (s HotelCityCodeListResponseBodyModuleCitys) String() string {
	return dara.Prettify(s)
}

func (s HotelCityCodeListResponseBodyModuleCitys) GoString() string {
	return s.String()
}

func (s *HotelCityCodeListResponseBodyModuleCitys) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelCityCodeListResponseBodyModuleCitys) GetCityName() *string {
	return s.CityName
}

func (s *HotelCityCodeListResponseBodyModuleCitys) GetDistricts() []*HotelCityCodeListResponseBodyModuleCitysDistricts {
	return s.Districts
}

func (s *HotelCityCodeListResponseBodyModuleCitys) SetCityCode(v string) *HotelCityCodeListResponseBodyModuleCitys {
	s.CityCode = &v
	return s
}

func (s *HotelCityCodeListResponseBodyModuleCitys) SetCityName(v string) *HotelCityCodeListResponseBodyModuleCitys {
	s.CityName = &v
	return s
}

func (s *HotelCityCodeListResponseBodyModuleCitys) SetDistricts(v []*HotelCityCodeListResponseBodyModuleCitysDistricts) *HotelCityCodeListResponseBodyModuleCitys {
	s.Districts = v
	return s
}

func (s *HotelCityCodeListResponseBodyModuleCitys) Validate() error {
	return dara.Validate(s)
}

type HotelCityCodeListResponseBodyModuleCitysDistricts struct {
	// example:
	//
	// 330000
	DistrictCode *string `json:"district_code,omitempty" xml:"district_code,omitempty"`
	DistrictName *string `json:"district_name,omitempty" xml:"district_name,omitempty"`
}

func (s HotelCityCodeListResponseBodyModuleCitysDistricts) String() string {
	return dara.Prettify(s)
}

func (s HotelCityCodeListResponseBodyModuleCitysDistricts) GoString() string {
	return s.String()
}

func (s *HotelCityCodeListResponseBodyModuleCitysDistricts) GetDistrictCode() *string {
	return s.DistrictCode
}

func (s *HotelCityCodeListResponseBodyModuleCitysDistricts) GetDistrictName() *string {
	return s.DistrictName
}

func (s *HotelCityCodeListResponseBodyModuleCitysDistricts) SetDistrictCode(v string) *HotelCityCodeListResponseBodyModuleCitysDistricts {
	s.DistrictCode = &v
	return s
}

func (s *HotelCityCodeListResponseBodyModuleCitysDistricts) SetDistrictName(v string) *HotelCityCodeListResponseBodyModuleCitysDistricts {
	s.DistrictName = &v
	return s
}

func (s *HotelCityCodeListResponseBodyModuleCitysDistricts) Validate() error {
	return dara.Validate(s)
}
