// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelStaticInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelStaticInfoResponseBody
	GetCode() *string
	SetMessage(v string) *HotelStaticInfoResponseBody
	GetMessage() *string
	SetModule(v *HotelStaticInfoResponseBodyModule) *HotelStaticInfoResponseBody
	GetModule() *HotelStaticInfoResponseBodyModule
	SetRequestId(v string) *HotelStaticInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelStaticInfoResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelStaticInfoResponseBody
	GetTraceId() *string
}

type HotelStaticInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// operation success.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *HotelStaticInfoResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
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

func (s HotelStaticInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoResponseBody) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelStaticInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelStaticInfoResponseBody) GetModule() *HotelStaticInfoResponseBodyModule {
	return s.Module
}

func (s *HotelStaticInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelStaticInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelStaticInfoResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelStaticInfoResponseBody) SetCode(v string) *HotelStaticInfoResponseBody {
	s.Code = &v
	return s
}

func (s *HotelStaticInfoResponseBody) SetMessage(v string) *HotelStaticInfoResponseBody {
	s.Message = &v
	return s
}

func (s *HotelStaticInfoResponseBody) SetModule(v *HotelStaticInfoResponseBodyModule) *HotelStaticInfoResponseBody {
	s.Module = v
	return s
}

func (s *HotelStaticInfoResponseBody) SetRequestId(v string) *HotelStaticInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelStaticInfoResponseBody) SetSuccess(v bool) *HotelStaticInfoResponseBody {
	s.Success = &v
	return s
}

func (s *HotelStaticInfoResponseBody) SetTraceId(v string) *HotelStaticInfoResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelStaticInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelStaticInfoResponseBodyModule struct {
	HotelStaticInfos []*HotelStaticInfoResponseBodyModuleHotelStaticInfos `json:"hotel_static_infos,omitempty" xml:"hotel_static_infos,omitempty" type:"Repeated"`
}

func (s HotelStaticInfoResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoResponseBodyModule) GetHotelStaticInfos() []*HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	return s.HotelStaticInfos
}

func (s *HotelStaticInfoResponseBodyModule) SetHotelStaticInfos(v []*HotelStaticInfoResponseBodyModuleHotelStaticInfos) *HotelStaticInfoResponseBodyModule {
	s.HotelStaticInfos = v
	return s
}

func (s *HotelStaticInfoResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelStaticInfoResponseBodyModuleHotelStaticInfos struct {
	BlockRoomTypeInformation map[string]*string `json:"block_room_type_information,omitempty" xml:"block_room_type_information,omitempty"`
	// example:
	//
	// 19039
	Brand     *string `json:"brand,omitempty" xml:"brand,omitempty"`
	BrandName *string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// CN
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// example:
	//
	// CN
	CountryCode *string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 330183
	District     *string                                                      `json:"district,omitempty" xml:"district,omitempty"`
	DistrictName *string                                                      `json:"district_name,omitempty" xml:"district_name,omitempty"`
	ExpandInfo   *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo `json:"expand_info,omitempty" xml:"expand_info,omitempty" type:"Struct"`
	HotelAddress *string                                                      `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
	// example:
	//
	// Building 5, Haichuang Building, 998 Wenyi West Road
	HotelEnAddress *string `json:"hotel_en_address,omitempty" xml:"hotel_en_address,omitempty"`
	// example:
	//
	// Business travel hotel
	HotelEnName *string `json:"hotel_en_name,omitempty" xml:"hotel_en_name,omitempty"`
	// example:
	//
	// 55335212
	HotelId   *string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	HotelName *string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// example:
	//
	// 2020
	HotelOpenTime *string `json:"hotel_open_time,omitempty" xml:"hotel_open_time,omitempty"`
	// example:
	//
	// 0086#0571#28350117
	HotelPhones *string `json:"hotel_phones,omitempty" xml:"hotel_phones,omitempty"`
	// example:
	//
	// demo
	HotelPolicies *string `json:"hotel_policies,omitempty" xml:"hotel_policies,omitempty"`
	// example:
	//
	// 0086#0571#28350117
	Hotelfax   *string                                                        `json:"hotelfax,omitempty" xml:"hotelfax,omitempty"`
	Hotelpics  *string                                                        `json:"hotelpics,omitempty" xml:"hotelpics,omitempty"`
	Imageinfos []*HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos `json:"imageinfos,omitempty" xml:"imageinfos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	InvoiceProviderMethod *string  `json:"invoice_provider_method,omitempty" xml:"invoice_provider_method,omitempty"`
	InvoiceTypes          []*int32 `json:"invoice_types,omitempty" xml:"invoice_types,omitempty" type:"Repeated"`
	// example:
	//
	// 119.844005,30.054384
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// 330000
	Province     *string `json:"province,omitempty" xml:"province,omitempty"`
	ProvinceName *string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// example:
	//
	// 2
	RatingAverage *string                                                       `json:"rating_average,omitempty" xml:"rating_average,omitempty"`
	RoomInfos     []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos `json:"room_infos,omitempty" xml:"room_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Star *string `json:"star,omitempty" xml:"star,omitempty"`
	// example:
	//
	// 2
	StarRate *string `json:"star_rate,omitempty" xml:"star_rate,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// demo
	Themes *string `json:"themes,omitempty" xml:"themes,omitempty"`
	// example:
	//
	// demo
	VisaReminding *bool `json:"visa_reminding,omitempty" xml:"visa_reminding,omitempty"`
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfos) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetBlockRoomTypeInformation() map[string]*string {
	return s.BlockRoomTypeInformation
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetBrand() *string {
	return s.Brand
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetBrandName() *string {
	return s.BrandName
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetCityName() *string {
	return s.CityName
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetCountry() *string {
	return s.Country
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetCountryCode() *string {
	return s.CountryCode
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetDescription() *string {
	return s.Description
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetDistrict() *string {
	return s.District
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetDistrictName() *string {
	return s.DistrictName
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetExpandInfo() *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	return s.ExpandInfo
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetHotelEnAddress() *string {
	return s.HotelEnAddress
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetHotelEnName() *string {
	return s.HotelEnName
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetHotelId() *string {
	return s.HotelId
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetHotelName() *string {
	return s.HotelName
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetHotelOpenTime() *string {
	return s.HotelOpenTime
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetHotelPhones() *string {
	return s.HotelPhones
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetHotelPolicies() *string {
	return s.HotelPolicies
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetHotelfax() *string {
	return s.Hotelfax
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetHotelpics() *string {
	return s.Hotelpics
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetImageinfos() []*HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos {
	return s.Imageinfos
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetInvoiceProviderMethod() *string {
	return s.InvoiceProviderMethod
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetInvoiceTypes() []*int32 {
	return s.InvoiceTypes
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetLocation() *string {
	return s.Location
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetProvince() *string {
	return s.Province
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetRatingAverage() *string {
	return s.RatingAverage
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetRoomInfos() []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	return s.RoomInfos
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetStar() *string {
	return s.Star
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetStarRate() *string {
	return s.StarRate
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetStatus() *string {
	return s.Status
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetThemes() *string {
	return s.Themes
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) GetVisaReminding() *bool {
	return s.VisaReminding
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetBlockRoomTypeInformation(v map[string]*string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.BlockRoomTypeInformation = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetBrand(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Brand = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetBrandName(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.BrandName = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetCityCode(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.CityCode = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetCityName(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.CityName = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetCountry(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Country = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetCountryCode(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.CountryCode = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetDescription(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Description = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetDistrict(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.District = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetDistrictName(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.DistrictName = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetExpandInfo(v *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.ExpandInfo = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetHotelAddress(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.HotelAddress = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetHotelEnAddress(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.HotelEnAddress = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetHotelEnName(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.HotelEnName = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetHotelId(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.HotelId = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetHotelName(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.HotelName = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetHotelOpenTime(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.HotelOpenTime = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetHotelPhones(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.HotelPhones = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetHotelPolicies(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.HotelPolicies = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetHotelfax(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Hotelfax = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetHotelpics(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Hotelpics = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetImageinfos(v []*HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Imageinfos = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetInvoiceProviderMethod(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.InvoiceProviderMethod = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetInvoiceTypes(v []*int32) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.InvoiceTypes = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetLocation(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Location = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetProvince(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Province = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetProvinceName(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.ProvinceName = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetRatingAverage(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.RatingAverage = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetRoomInfos(v []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.RoomInfos = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetStar(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Star = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetStarRate(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.StarRate = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetStatus(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Status = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetThemes(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.Themes = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) SetVisaReminding(v bool) *HotelStaticInfoResponseBodyModuleHotelStaticInfos {
	s.VisaReminding = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfos) Validate() error {
	return dara.Validate(s)
}

type HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo struct {
	// example:
	//
	// 17:00
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// example:
	//
	// 12:00
	CheckOut *string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// example:
	//
	// 2020
	DecorateTime    *string   `json:"decorate_time,omitempty" xml:"decorate_time,omitempty"`
	Floors          *string   `json:"floors,omitempty" xml:"floors,omitempty"`
	HotelFacilities []*string `json:"hotel_facilities,omitempty" xml:"hotel_facilities,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	HotelType *int32 `json:"hotel_type,omitempty" xml:"hotel_type,omitempty"`
	// example:
	//
	// 2020
	OpeningTime    *string   `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	RoomFacilities []*string `json:"room_facilities,omitempty" xml:"room_facilities,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Rooms   *int32    `json:"rooms,omitempty" xml:"rooms,omitempty"`
	Service []*string `json:"service,omitempty" xml:"service,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ThemeTag *string `json:"theme_tag,omitempty" xml:"theme_tag,omitempty"`
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetDecorateTime() *string {
	return s.DecorateTime
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetFloors() *string {
	return s.Floors
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetHotelFacilities() []*string {
	return s.HotelFacilities
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetHotelType() *int32 {
	return s.HotelType
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetOpeningTime() *string {
	return s.OpeningTime
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetRoomFacilities() []*string {
	return s.RoomFacilities
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetRooms() *int32 {
	return s.Rooms
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetService() []*string {
	return s.Service
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) GetThemeTag() *string {
	return s.ThemeTag
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetCheckIn(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.CheckIn = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetCheckOut(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.CheckOut = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetDecorateTime(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.DecorateTime = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetFloors(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.Floors = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetHotelFacilities(v []*string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.HotelFacilities = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetHotelType(v int32) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.HotelType = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetOpeningTime(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.OpeningTime = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetRoomFacilities(v []*string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.RoomFacilities = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetRooms(v int32) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.Rooms = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetService(v []*string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.Service = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) SetThemeTag(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo {
	s.ThemeTag = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosExpandInfo) Validate() error {
	return dara.Validate(s)
}

type HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Tag  *int32  `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/6000000000003/O1CN01xkZQR41BtPxK1PQCb_!!6000000000003-0-hotel.jpg
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos) GetDesc() *string {
	return s.Desc
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos) GetTag() *int32 {
	return s.Tag
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos) GetUrl() *string {
	return s.Url
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos) SetDesc(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos {
	s.Desc = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos) SetTag(v int32) *HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos {
	s.Tag = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos) SetUrl(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos {
	s.Url = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosImageinfos) Validate() error {
	return dara.Validate(s)
}

type HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos struct {
	BedInfoGroupList []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupList `json:"bed_info_group_list,omitempty" xml:"bed_info_group_list,omitempty" type:"Repeated"`
	BedInfos         []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos         `json:"bed_infos,omitempty" xml:"bed_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ExtraBed *int32 `json:"extra_bed,omitempty" xml:"extra_bed,omitempty"`
	// example:
	//
	// demo
	ExtraBedDesc *string `json:"extra_bed_desc,omitempty" xml:"extra_bed_desc,omitempty"`
	// example:
	//
	// 1,2,3,4,5,6
	Floor *string `json:"floor,omitempty" xml:"floor,omitempty"`
	// example:
	//
	// 0
	InternetWay *string `json:"internet_way,omitempty" xml:"internet_way,omitempty"`
	// example:
	//
	// 1
	MaxOccupancy     *int32    `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	RoomFacilities   *string   `json:"room_facilities,omitempty" xml:"room_facilities,omitempty"`
	RoomFacilityList []*string `json:"room_facility_list,omitempty" xml:"room_facility_list,omitempty" type:"Repeated"`
	// example:
	//
	// 71652158
	RoomId *string `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// example:
	//
	// //img.alicdn.com/imgextra/i3/6000000000003/O1CN01xkZQR41BtPxK1PQCb_!!6000000000003-0-hotel.jpg
	RoomImage  *string   `json:"room_image,omitempty" xml:"room_image,omitempty"`
	RoomImages []*string `json:"room_images,omitempty" xml:"room_images,omitempty" type:"Repeated"`
	RoomName   *string   `json:"room_name,omitempty" xml:"room_name,omitempty"`
	RoomType   *int32    `json:"room_type,omitempty" xml:"room_type,omitempty"`
	Roomarea   *string   `json:"roomarea,omitempty" xml:"roomarea,omitempty"`
	// example:
	//
	// 1
	Rooms *int32 `json:"rooms,omitempty" xml:"rooms,omitempty"`
	// example:
	//
	// 0
	Window     *string `json:"window,omitempty" xml:"window,omitempty"`
	WindowBad  *string `json:"window_bad,omitempty" xml:"window_bad,omitempty"`
	WindowView *string `json:"window_view,omitempty" xml:"window_view,omitempty"`
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetBedInfoGroupList() []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupList {
	return s.BedInfoGroupList
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetBedInfos() []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos {
	return s.BedInfos
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetExtraBed() *int32 {
	return s.ExtraBed
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetExtraBedDesc() *string {
	return s.ExtraBedDesc
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetFloor() *string {
	return s.Floor
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetInternetWay() *string {
	return s.InternetWay
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetMaxOccupancy() *int32 {
	return s.MaxOccupancy
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetRoomFacilities() *string {
	return s.RoomFacilities
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetRoomFacilityList() []*string {
	return s.RoomFacilityList
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetRoomId() *string {
	return s.RoomId
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetRoomImage() *string {
	return s.RoomImage
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetRoomImages() []*string {
	return s.RoomImages
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetRoomName() *string {
	return s.RoomName
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetRoomType() *int32 {
	return s.RoomType
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetRoomarea() *string {
	return s.Roomarea
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetRooms() *int32 {
	return s.Rooms
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetWindow() *string {
	return s.Window
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetWindowBad() *string {
	return s.WindowBad
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) GetWindowView() *string {
	return s.WindowView
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetBedInfoGroupList(v []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupList) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.BedInfoGroupList = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetBedInfos(v []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.BedInfos = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetExtraBed(v int32) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.ExtraBed = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetExtraBedDesc(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.ExtraBedDesc = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetFloor(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.Floor = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetInternetWay(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.InternetWay = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetMaxOccupancy(v int32) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.MaxOccupancy = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetRoomFacilities(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.RoomFacilities = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetRoomFacilityList(v []*string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.RoomFacilityList = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetRoomId(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.RoomId = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetRoomImage(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.RoomImage = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetRoomImages(v []*string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.RoomImages = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetRoomName(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.RoomName = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetRoomType(v int32) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.RoomType = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetRoomarea(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.Roomarea = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetRooms(v int32) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.Rooms = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetWindow(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.Window = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetWindowBad(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.WindowBad = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) SetWindowView(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos {
	s.WindowView = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfos) Validate() error {
	return dara.Validate(s)
}

type HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupList struct {
	BedInfos []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos `json:"bed_Infos,omitempty" xml:"bed_Infos,omitempty" type:"Repeated"`
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupList) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupList) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupList) GetBedInfos() []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos {
	return s.BedInfos
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupList) SetBedInfos(v []*HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupList {
	s.BedInfos = v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupList) Validate() error {
	return dara.Validate(s)
}

type HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos struct {
	BedDesc *string `json:"bed_desc,omitempty" xml:"bed_desc,omitempty"`
	BedNum  *int32  `json:"bed_num,omitempty" xml:"bed_num,omitempty"`
	BedSize *string `json:"bed_size,omitempty" xml:"bed_size,omitempty"`
	BedType *string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	Length  *string `json:"length,omitempty" xml:"length,omitempty"`
	Width   *string `json:"width,omitempty" xml:"width,omitempty"`
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) GetBedDesc() *string {
	return s.BedDesc
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) GetBedNum() *int32 {
	return s.BedNum
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) GetBedSize() *string {
	return s.BedSize
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) GetBedType() *string {
	return s.BedType
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) GetLength() *string {
	return s.Length
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) GetWidth() *string {
	return s.Width
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) SetBedDesc(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos {
	s.BedDesc = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) SetBedNum(v int32) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos {
	s.BedNum = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) SetBedSize(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos {
	s.BedSize = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) SetBedType(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos {
	s.BedType = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) SetLength(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos {
	s.Length = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) SetWidth(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos {
	s.Width = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfoGroupListBedInfos) Validate() error {
	return dara.Validate(s)
}

type HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos struct {
	BedDesc *string `json:"bed_desc,omitempty" xml:"bed_desc,omitempty"`
	// example:
	//
	// 1
	BedNum *int32 `json:"bed_num,omitempty" xml:"bed_num,omitempty"`
	// example:
	//
	// 1.8m
	BedSize *string `json:"bed_size,omitempty" xml:"bed_size,omitempty"`
	// example:
	//
	// 0
	BedType *string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	Length  *string `json:"length,omitempty" xml:"length,omitempty"`
	Width   *string `json:"width,omitempty" xml:"width,omitempty"`
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) GetBedDesc() *string {
	return s.BedDesc
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) GetBedNum() *int32 {
	return s.BedNum
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) GetBedSize() *string {
	return s.BedSize
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) GetBedType() *string {
	return s.BedType
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) GetLength() *string {
	return s.Length
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) GetWidth() *string {
	return s.Width
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) SetBedDesc(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos {
	s.BedDesc = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) SetBedNum(v int32) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos {
	s.BedNum = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) SetBedSize(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos {
	s.BedSize = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) SetBedType(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos {
	s.BedType = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) SetLength(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos {
	s.Length = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) SetWidth(v string) *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos {
	s.Width = &v
	return s
}

func (s *HotelStaticInfoResponseBodyModuleHotelStaticInfosRoomInfosBedInfos) Validate() error {
	return dara.Validate(s)
}
