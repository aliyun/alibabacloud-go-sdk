// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelSuggestV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelSuggestV2ResponseBody
	GetCode() *string
	SetMessage(v string) *HotelSuggestV2ResponseBody
	GetMessage() *string
	SetModule(v *HotelSuggestV2ResponseBodyModule) *HotelSuggestV2ResponseBody
	GetModule() *HotelSuggestV2ResponseBodyModule
	SetRequestId(v string) *HotelSuggestV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelSuggestV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelSuggestV2ResponseBody
	GetTraceId() *string
}

type HotelSuggestV2ResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// None
	Message *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelSuggestV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s HotelSuggestV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelSuggestV2ResponseBody) GoString() string {
	return s.String()
}

func (s *HotelSuggestV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelSuggestV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelSuggestV2ResponseBody) GetModule() *HotelSuggestV2ResponseBodyModule {
	return s.Module
}

func (s *HotelSuggestV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelSuggestV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelSuggestV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelSuggestV2ResponseBody) SetCode(v string) *HotelSuggestV2ResponseBody {
	s.Code = &v
	return s
}

func (s *HotelSuggestV2ResponseBody) SetMessage(v string) *HotelSuggestV2ResponseBody {
	s.Message = &v
	return s
}

func (s *HotelSuggestV2ResponseBody) SetModule(v *HotelSuggestV2ResponseBodyModule) *HotelSuggestV2ResponseBody {
	s.Module = v
	return s
}

func (s *HotelSuggestV2ResponseBody) SetRequestId(v string) *HotelSuggestV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelSuggestV2ResponseBody) SetSuccess(v bool) *HotelSuggestV2ResponseBody {
	s.Success = &v
	return s
}

func (s *HotelSuggestV2ResponseBody) SetTraceId(v string) *HotelSuggestV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelSuggestV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelSuggestV2ResponseBodyModule struct {
	GuessSuggestInfos   []*HotelSuggestV2ResponseBodyModuleGuessSuggestInfos   `json:"guess_suggest_infos,omitempty" xml:"guess_suggest_infos,omitempty" type:"Repeated"`
	KeywordSuggestInfos []*HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos `json:"keyword_suggest_infos,omitempty" xml:"keyword_suggest_infos,omitempty" type:"Repeated"`
	PopularSuggestInfos []*HotelSuggestV2ResponseBodyModulePopularSuggestInfos `json:"popular_suggest_infos,omitempty" xml:"popular_suggest_infos,omitempty" type:"Repeated"`
	Tips                *string                                                `json:"tips,omitempty" xml:"tips,omitempty"`
}

func (s HotelSuggestV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelSuggestV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelSuggestV2ResponseBodyModule) GetGuessSuggestInfos() []*HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	return s.GuessSuggestInfos
}

func (s *HotelSuggestV2ResponseBodyModule) GetKeywordSuggestInfos() []*HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	return s.KeywordSuggestInfos
}

func (s *HotelSuggestV2ResponseBodyModule) GetPopularSuggestInfos() []*HotelSuggestV2ResponseBodyModulePopularSuggestInfos {
	return s.PopularSuggestInfos
}

func (s *HotelSuggestV2ResponseBodyModule) GetTips() *string {
	return s.Tips
}

func (s *HotelSuggestV2ResponseBodyModule) SetGuessSuggestInfos(v []*HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) *HotelSuggestV2ResponseBodyModule {
	s.GuessSuggestInfos = v
	return s
}

func (s *HotelSuggestV2ResponseBodyModule) SetKeywordSuggestInfos(v []*HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) *HotelSuggestV2ResponseBodyModule {
	s.KeywordSuggestInfos = v
	return s
}

func (s *HotelSuggestV2ResponseBodyModule) SetPopularSuggestInfos(v []*HotelSuggestV2ResponseBodyModulePopularSuggestInfos) *HotelSuggestV2ResponseBodyModule {
	s.PopularSuggestInfos = v
	return s
}

func (s *HotelSuggestV2ResponseBodyModule) SetTips(v string) *HotelSuggestV2ResponseBodyModule {
	s.Tips = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelSuggestV2ResponseBodyModuleGuessSuggestInfos struct {
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 300100
	CityCode    *int32  `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName    *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	DisplayName *string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// example:
	//
	// 53853318
	HotelId *string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i3/O1CN01qKg25r1rKLOKxT3vB_!!6000000005612-2-tps-32-32.png
	Icon  *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Point *string `json:"point,omitempty" xml:"point,omitempty"`
	// example:
	//
	// 524
	Price *string `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 0
	Region *int32 `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 10
	Type     *int32  `json:"type,omitempty" xml:"type,omitempty"`
	TypeDesc *string `json:"type_desc,omitempty" xml:"type_desc,omitempty"`
}

func (s HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GoString() string {
	return s.String()
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetAddress() *string {
	return s.Address
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetCityCode() *int32 {
	return s.CityCode
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetCityName() *string {
	return s.CityName
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetDisplayName() *string {
	return s.DisplayName
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetHotelId() *string {
	return s.HotelId
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetIcon() *string {
	return s.Icon
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetPoint() *string {
	return s.Point
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetPrice() *string {
	return s.Price
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetRegion() *int32 {
	return s.Region
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetType() *int32 {
	return s.Type
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) GetTypeDesc() *string {
	return s.TypeDesc
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetAddress(v string) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.Address = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetCityCode(v int32) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.CityCode = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetCityName(v string) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.CityName = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetDisplayName(v string) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.DisplayName = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetHotelId(v string) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.HotelId = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetIcon(v string) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.Icon = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetPoint(v string) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.Point = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetPrice(v string) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.Price = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetRegion(v int32) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.Region = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetType(v int32) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.Type = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) SetTypeDesc(v string) *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos {
	s.TypeDesc = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleGuessSuggestInfos) Validate() error {
	return dara.Validate(s)
}

type HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos struct {
	Address              *string                                                                    `json:"address,omitempty" xml:"address,omitempty"`
	BusinessAreaWithCity []*HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity `json:"business_area_with_city,omitempty" xml:"business_area_with_city,omitempty" type:"Repeated"`
	// example:
	//
	// 300100
	CityCode    *int32  `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName    *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	DisplayName *string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// example:
	//
	// 53853318
	HotelId *string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i3/O1CN01qKg25r1rKLOKxT3vB_!!6000000005612-2-tps-32-32.png
	Icon  *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Point *string `json:"point,omitempty" xml:"point,omitempty"`
	// example:
	//
	// 524
	Price *string `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 0
	Region *int32 `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 10
	Type     *int32  `json:"type,omitempty" xml:"type,omitempty"`
	TypeDesc *string `json:"type_desc,omitempty" xml:"type_desc,omitempty"`
}

func (s HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GoString() string {
	return s.String()
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetAddress() *string {
	return s.Address
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetBusinessAreaWithCity() []*HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	return s.BusinessAreaWithCity
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetCityCode() *int32 {
	return s.CityCode
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetCityName() *string {
	return s.CityName
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetDisplayName() *string {
	return s.DisplayName
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetHotelId() *string {
	return s.HotelId
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetIcon() *string {
	return s.Icon
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetPoint() *string {
	return s.Point
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetPrice() *string {
	return s.Price
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetRegion() *int32 {
	return s.Region
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetType() *int32 {
	return s.Type
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) GetTypeDesc() *string {
	return s.TypeDesc
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetAddress(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.Address = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetBusinessAreaWithCity(v []*HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.BusinessAreaWithCity = v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetCityCode(v int32) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.CityCode = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetCityName(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.CityName = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetDisplayName(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.DisplayName = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetHotelId(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.HotelId = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetIcon(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.Icon = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetPoint(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.Point = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetPrice(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.Price = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetRegion(v int32) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.Region = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetType(v int32) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.Type = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) SetTypeDesc(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos {
	s.TypeDesc = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfos) Validate() error {
	return dara.Validate(s)
}

type HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity struct {
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 300100
	CityCode    *int32  `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName    *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	DisplayName *string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// example:
	//
	// 57140953
	HotelId *string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/imgextra/i3/O1CN01qKg25r1rKLOKxT3vB_!!6000000005612-2-tps-32-32.png
	Icon  *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Point *string `json:"point,omitempty" xml:"point,omitempty"`
	// example:
	//
	// 524
	Price *string `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 0
	Region *int32 `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 5
	Type     *int32  `json:"type,omitempty" xml:"type,omitempty"`
	TypeDesc *string `json:"type_desc,omitempty" xml:"type_desc,omitempty"`
}

func (s HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) String() string {
	return dara.Prettify(s)
}

func (s HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GoString() string {
	return s.String()
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetAddress() *string {
	return s.Address
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetCityCode() *int32 {
	return s.CityCode
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetCityName() *string {
	return s.CityName
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetDisplayName() *string {
	return s.DisplayName
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetHotelId() *string {
	return s.HotelId
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetIcon() *string {
	return s.Icon
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetPoint() *string {
	return s.Point
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetPrice() *string {
	return s.Price
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetRegion() *int32 {
	return s.Region
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetType() *int32 {
	return s.Type
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) GetTypeDesc() *string {
	return s.TypeDesc
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetAddress(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.Address = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetCityCode(v int32) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.CityCode = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetCityName(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.CityName = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetDisplayName(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.DisplayName = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetHotelId(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.HotelId = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetIcon(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.Icon = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetPoint(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.Point = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetPrice(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.Price = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetRegion(v int32) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.Region = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetType(v int32) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.Type = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) SetTypeDesc(v string) *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity {
	s.TypeDesc = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModuleKeywordSuggestInfosBusinessAreaWithCity) Validate() error {
	return dara.Validate(s)
}

type HotelSuggestV2ResponseBodyModulePopularSuggestInfos struct {
	// example:
	//
	// https://gw.alicdn.com/imgextra/i1/O1CN01x0q19E1QZSqLHVVNh_!!6000000001990-2-tps-54-54.png
	Icon         *string                                                            `json:"icon,omitempty" xml:"icon,omitempty"`
	PopularInfos []*HotelSuggestV2ResponseBodyModulePopularSuggestInfosPopularInfos `json:"popular_infos,omitempty" xml:"popular_infos,omitempty" type:"Repeated"`
	Title        *string                                                            `json:"title,omitempty" xml:"title,omitempty"`
}

func (s HotelSuggestV2ResponseBodyModulePopularSuggestInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelSuggestV2ResponseBodyModulePopularSuggestInfos) GoString() string {
	return s.String()
}

func (s *HotelSuggestV2ResponseBodyModulePopularSuggestInfos) GetIcon() *string {
	return s.Icon
}

func (s *HotelSuggestV2ResponseBodyModulePopularSuggestInfos) GetPopularInfos() []*HotelSuggestV2ResponseBodyModulePopularSuggestInfosPopularInfos {
	return s.PopularInfos
}

func (s *HotelSuggestV2ResponseBodyModulePopularSuggestInfos) GetTitle() *string {
	return s.Title
}

func (s *HotelSuggestV2ResponseBodyModulePopularSuggestInfos) SetIcon(v string) *HotelSuggestV2ResponseBodyModulePopularSuggestInfos {
	s.Icon = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModulePopularSuggestInfos) SetPopularInfos(v []*HotelSuggestV2ResponseBodyModulePopularSuggestInfosPopularInfos) *HotelSuggestV2ResponseBodyModulePopularSuggestInfos {
	s.PopularInfos = v
	return s
}

func (s *HotelSuggestV2ResponseBodyModulePopularSuggestInfos) SetTitle(v string) *HotelSuggestV2ResponseBodyModulePopularSuggestInfos {
	s.Title = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModulePopularSuggestInfos) Validate() error {
	return dara.Validate(s)
}

type HotelSuggestV2ResponseBodyModulePopularSuggestInfosPopularInfos struct {
	DisplayName *string `json:"display_name,omitempty" xml:"display_name,omitempty"`
}

func (s HotelSuggestV2ResponseBodyModulePopularSuggestInfosPopularInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelSuggestV2ResponseBodyModulePopularSuggestInfosPopularInfos) GoString() string {
	return s.String()
}

func (s *HotelSuggestV2ResponseBodyModulePopularSuggestInfosPopularInfos) GetDisplayName() *string {
	return s.DisplayName
}

func (s *HotelSuggestV2ResponseBodyModulePopularSuggestInfosPopularInfos) SetDisplayName(v string) *HotelSuggestV2ResponseBodyModulePopularSuggestInfosPopularInfos {
	s.DisplayName = &v
	return s
}

func (s *HotelSuggestV2ResponseBodyModulePopularSuggestInfosPopularInfos) Validate() error {
	return dara.Validate(s)
}
