// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelSearchResponseBody
	GetCode() *string
	SetMessage(v string) *HotelSearchResponseBody
	GetMessage() *string
	SetModule(v *HotelSearchResponseBodyModule) *HotelSearchResponseBody
	GetModule() *HotelSearchResponseBodyModule
	SetRequestId(v string) *HotelSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelSearchResponseBody
	GetTraceId() *string
}

type HotelSearchResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// None
	Message *string                        `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s HotelSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelSearchResponseBody) GoString() string {
	return s.String()
}

func (s *HotelSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelSearchResponseBody) GetModule() *HotelSearchResponseBodyModule {
	return s.Module
}

func (s *HotelSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelSearchResponseBody) SetCode(v string) *HotelSearchResponseBody {
	s.Code = &v
	return s
}

func (s *HotelSearchResponseBody) SetMessage(v string) *HotelSearchResponseBody {
	s.Message = &v
	return s
}

func (s *HotelSearchResponseBody) SetModule(v *HotelSearchResponseBodyModule) *HotelSearchResponseBody {
	s.Module = v
	return s
}

func (s *HotelSearchResponseBody) SetRequestId(v string) *HotelSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelSearchResponseBody) SetSuccess(v bool) *HotelSearchResponseBody {
	s.Success = &v
	return s
}

func (s *HotelSearchResponseBody) SetTraceId(v string) *HotelSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelSearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelSearchResponseBodyModule struct {
	Count *int32                                `json:"count,omitempty" xml:"count,omitempty"`
	Items []*HotelSearchResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s HotelSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelSearchResponseBodyModule) GetCount() *int32 {
	return s.Count
}

func (s *HotelSearchResponseBodyModule) GetItems() []*HotelSearchResponseBodyModuleItems {
	return s.Items
}

func (s *HotelSearchResponseBodyModule) SetCount(v int32) *HotelSearchResponseBodyModule {
	s.Count = &v
	return s
}

func (s *HotelSearchResponseBodyModule) SetItems(v []*HotelSearchResponseBodyModuleItems) *HotelSearchResponseBodyModule {
	s.Items = v
	return s
}

func (s *HotelSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelSearchResponseBodyModuleItems struct {
	BrandName *string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// example:
	//
	// 1902
	BtandCode *string `json:"btand_code,omitempty" xml:"btand_code,omitempty"`
	// example:
	//
	// 330100
	CityCode     *string                                         `json:"city_code,omitempty" xml:"city_code,omitempty"`
	DiscountDesc *HotelSearchResponseBodyModuleItemsDiscountDesc `json:"discount_desc,omitempty" xml:"discount_desc,omitempty" type:"Struct"`
	// example:
	//
	// 100
	Distance *int32 `json:"distance,omitempty" xml:"distance,omitempty"`
	// example:
	//
	// 330100
	DistrictCode *string `json:"district_code,omitempty" xml:"district_code,omitempty"`
	HotelAddress *string `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
	// example:
	//
	// 55335212
	HotelCode *string `json:"hotel_code,omitempty" xml:"hotel_code,omitempty"`
	// example:
	//
	// Business travel hotel
	HotelEnName *string `json:"hotel_en_name,omitempty" xml:"hotel_en_name,omitempty"`
	HotelName   *string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// example:
	//
	// 3
	HotelStar *string `json:"hotel_star,omitempty" xml:"hotel_star,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/6000000000003/O1CN01xkZQR41BtPxK1PQCb_!!6000000000003-0-hotel.jpg
	ImageUrl *string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// example:
	//
	// true
	IsProtocol *bool `json:"is_protocol,omitempty" xml:"is_protocol,omitempty"`
	// example:
	//
	// 119.844005,30.054384
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// 100
	MinPrice         *float64 `json:"min_price,omitempty" xml:"min_price,omitempty"`
	OriginalMinPrice *float64 `json:"original_min_price,omitempty" xml:"original_min_price,omitempty"`
	// example:
	//
	// 3.2
	Score *string `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 0571-88886784-8888
	Tel *string `json:"tel,omitempty" xml:"tel,omitempty"`
}

func (s HotelSearchResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s HotelSearchResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *HotelSearchResponseBodyModuleItems) GetBrandName() *string {
	return s.BrandName
}

func (s *HotelSearchResponseBodyModuleItems) GetBtandCode() *string {
	return s.BtandCode
}

func (s *HotelSearchResponseBodyModuleItems) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelSearchResponseBodyModuleItems) GetDiscountDesc() *HotelSearchResponseBodyModuleItemsDiscountDesc {
	return s.DiscountDesc
}

func (s *HotelSearchResponseBodyModuleItems) GetDistance() *int32 {
	return s.Distance
}

func (s *HotelSearchResponseBodyModuleItems) GetDistrictCode() *string {
	return s.DistrictCode
}

func (s *HotelSearchResponseBodyModuleItems) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *HotelSearchResponseBodyModuleItems) GetHotelCode() *string {
	return s.HotelCode
}

func (s *HotelSearchResponseBodyModuleItems) GetHotelEnName() *string {
	return s.HotelEnName
}

func (s *HotelSearchResponseBodyModuleItems) GetHotelName() *string {
	return s.HotelName
}

func (s *HotelSearchResponseBodyModuleItems) GetHotelStar() *string {
	return s.HotelStar
}

func (s *HotelSearchResponseBodyModuleItems) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *HotelSearchResponseBodyModuleItems) GetIsProtocol() *bool {
	return s.IsProtocol
}

func (s *HotelSearchResponseBodyModuleItems) GetLocation() *string {
	return s.Location
}

func (s *HotelSearchResponseBodyModuleItems) GetMinPrice() *float64 {
	return s.MinPrice
}

func (s *HotelSearchResponseBodyModuleItems) GetOriginalMinPrice() *float64 {
	return s.OriginalMinPrice
}

func (s *HotelSearchResponseBodyModuleItems) GetScore() *string {
	return s.Score
}

func (s *HotelSearchResponseBodyModuleItems) GetStatus() *int32 {
	return s.Status
}

func (s *HotelSearchResponseBodyModuleItems) GetTel() *string {
	return s.Tel
}

func (s *HotelSearchResponseBodyModuleItems) SetBrandName(v string) *HotelSearchResponseBodyModuleItems {
	s.BrandName = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetBtandCode(v string) *HotelSearchResponseBodyModuleItems {
	s.BtandCode = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetCityCode(v string) *HotelSearchResponseBodyModuleItems {
	s.CityCode = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetDiscountDesc(v *HotelSearchResponseBodyModuleItemsDiscountDesc) *HotelSearchResponseBodyModuleItems {
	s.DiscountDesc = v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetDistance(v int32) *HotelSearchResponseBodyModuleItems {
	s.Distance = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetDistrictCode(v string) *HotelSearchResponseBodyModuleItems {
	s.DistrictCode = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetHotelAddress(v string) *HotelSearchResponseBodyModuleItems {
	s.HotelAddress = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetHotelCode(v string) *HotelSearchResponseBodyModuleItems {
	s.HotelCode = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetHotelEnName(v string) *HotelSearchResponseBodyModuleItems {
	s.HotelEnName = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetHotelName(v string) *HotelSearchResponseBodyModuleItems {
	s.HotelName = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetHotelStar(v string) *HotelSearchResponseBodyModuleItems {
	s.HotelStar = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetImageUrl(v string) *HotelSearchResponseBodyModuleItems {
	s.ImageUrl = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetIsProtocol(v bool) *HotelSearchResponseBodyModuleItems {
	s.IsProtocol = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetLocation(v string) *HotelSearchResponseBodyModuleItems {
	s.Location = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetMinPrice(v float64) *HotelSearchResponseBodyModuleItems {
	s.MinPrice = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetOriginalMinPrice(v float64) *HotelSearchResponseBodyModuleItems {
	s.OriginalMinPrice = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetScore(v string) *HotelSearchResponseBodyModuleItems {
	s.Score = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetStatus(v int32) *HotelSearchResponseBodyModuleItems {
	s.Status = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) SetTel(v string) *HotelSearchResponseBodyModuleItems {
	s.Tel = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}

type HotelSearchResponseBodyModuleItemsDiscountDesc struct {
	CashReduceTotal *string                                                         `json:"cash_reduce_total,omitempty" xml:"cash_reduce_total,omitempty"`
	DinamicLabel    *string                                                         `json:"dinamic_label,omitempty" xml:"dinamic_label,omitempty"`
	DiscountDetail  []*HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail `json:"discount_detail,omitempty" xml:"discount_detail,omitempty" type:"Repeated"`
	SubTitle        *string                                                         `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	Title           *string                                                         `json:"title,omitempty" xml:"title,omitempty"`
}

func (s HotelSearchResponseBodyModuleItemsDiscountDesc) String() string {
	return dara.Prettify(s)
}

func (s HotelSearchResponseBodyModuleItemsDiscountDesc) GoString() string {
	return s.String()
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) GetCashReduceTotal() *string {
	return s.CashReduceTotal
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) GetDinamicLabel() *string {
	return s.DinamicLabel
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) GetDiscountDetail() []*HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail {
	return s.DiscountDetail
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) GetSubTitle() *string {
	return s.SubTitle
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) GetTitle() *string {
	return s.Title
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) SetCashReduceTotal(v string) *HotelSearchResponseBodyModuleItemsDiscountDesc {
	s.CashReduceTotal = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) SetDinamicLabel(v string) *HotelSearchResponseBodyModuleItemsDiscountDesc {
	s.DinamicLabel = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) SetDiscountDetail(v []*HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail) *HotelSearchResponseBodyModuleItemsDiscountDesc {
	s.DiscountDetail = v
	return s
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) SetSubTitle(v string) *HotelSearchResponseBodyModuleItemsDiscountDesc {
	s.SubTitle = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) SetTitle(v string) *HotelSearchResponseBodyModuleItemsDiscountDesc {
	s.Title = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDesc) Validate() error {
	return dara.Validate(s)
}

type HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail struct {
	LabelName []*string `json:"label_name,omitempty" xml:"label_name,omitempty" type:"Repeated"`
	MoneyDesc *string   `json:"money_desc,omitempty" xml:"money_desc,omitempty"`
}

func (s HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail) String() string {
	return dara.Prettify(s)
}

func (s HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail) GoString() string {
	return s.String()
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail) GetLabelName() []*string {
	return s.LabelName
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail) GetMoneyDesc() *string {
	return s.MoneyDesc
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail) SetLabelName(v []*string) *HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail {
	s.LabelName = v
	return s
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail) SetMoneyDesc(v string) *HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail {
	s.MoneyDesc = &v
	return s
}

func (s *HotelSearchResponseBodyModuleItemsDiscountDescDiscountDetail) Validate() error {
	return dara.Validate(s)
}
