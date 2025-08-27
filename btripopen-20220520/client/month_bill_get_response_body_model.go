// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillGetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MonthBillGetResponseBody
	GetCode() *string
	SetMessage(v string) *MonthBillGetResponseBody
	GetMessage() *string
	SetModule(v []*MonthBillGetResponseBodyModule) *MonthBillGetResponseBody
	GetModule() []*MonthBillGetResponseBodyModule
	SetRequestId(v string) *MonthBillGetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MonthBillGetResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MonthBillGetResponseBody
	GetTraceId() *string
}

type MonthBillGetResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module  []*MonthBillGetResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
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

func (s MonthBillGetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MonthBillGetResponseBody) GoString() string {
	return s.String()
}

func (s *MonthBillGetResponseBody) GetCode() *string {
	return s.Code
}

func (s *MonthBillGetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MonthBillGetResponseBody) GetModule() []*MonthBillGetResponseBodyModule {
	return s.Module
}

func (s *MonthBillGetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MonthBillGetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MonthBillGetResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MonthBillGetResponseBody) SetCode(v string) *MonthBillGetResponseBody {
	s.Code = &v
	return s
}

func (s *MonthBillGetResponseBody) SetMessage(v string) *MonthBillGetResponseBody {
	s.Message = &v
	return s
}

func (s *MonthBillGetResponseBody) SetModule(v []*MonthBillGetResponseBodyModule) *MonthBillGetResponseBody {
	s.Module = v
	return s
}

func (s *MonthBillGetResponseBody) SetRequestId(v string) *MonthBillGetResponseBody {
	s.RequestId = &v
	return s
}

func (s *MonthBillGetResponseBody) SetSuccess(v bool) *MonthBillGetResponseBody {
	s.Success = &v
	return s
}

func (s *MonthBillGetResponseBody) SetTraceId(v string) *MonthBillGetResponseBody {
	s.TraceId = &v
	return s
}

func (s *MonthBillGetResponseBody) Validate() error {
	return dara.Validate(s)
}

type MonthBillGetResponseBodyModule struct {
	EndDate *string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// CorpMonthAccountBillFeeDetail
	MonthAccountBillDetail *MonthBillGetResponseBodyModuleMonthAccountBillDetail `json:"monthAccountBillDetail,omitempty" xml:"monthAccountBillDetail,omitempty" type:"Struct"`
	StartDate              *string                                               `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// example:
	//
	// https://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s MonthBillGetResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s MonthBillGetResponseBodyModule) GoString() string {
	return s.String()
}

func (s *MonthBillGetResponseBodyModule) GetEndDate() *string {
	return s.EndDate
}

func (s *MonthBillGetResponseBodyModule) GetMonthAccountBillDetail() *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	return s.MonthAccountBillDetail
}

func (s *MonthBillGetResponseBodyModule) GetStartDate() *string {
	return s.StartDate
}

func (s *MonthBillGetResponseBodyModule) GetUrl() *string {
	return s.Url
}

func (s *MonthBillGetResponseBodyModule) SetEndDate(v string) *MonthBillGetResponseBodyModule {
	s.EndDate = &v
	return s
}

func (s *MonthBillGetResponseBodyModule) SetMonthAccountBillDetail(v *MonthBillGetResponseBodyModuleMonthAccountBillDetail) *MonthBillGetResponseBodyModule {
	s.MonthAccountBillDetail = v
	return s
}

func (s *MonthBillGetResponseBodyModule) SetStartDate(v string) *MonthBillGetResponseBodyModule {
	s.StartDate = &v
	return s
}

func (s *MonthBillGetResponseBodyModule) SetUrl(v string) *MonthBillGetResponseBodyModule {
	s.Url = &v
	return s
}

func (s *MonthBillGetResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type MonthBillGetResponseBodyModuleMonthAccountBillDetail struct {
	BillConfirmed *int32 `json:"billConfirmed,omitempty" xml:"billConfirmed,omitempty"`
	// 用车金额（单位：元）
	//
	// example:
	//
	// xx.xx
	CarAmount *float64 `json:"carAmount,omitempty" xml:"carAmount,omitempty"`
	// 违约金金额（单位：元）
	//
	// example:
	//
	// xx.xx
	DamageAmount *float64 `json:"damageAmount,omitempty" xml:"damageAmount,omitempty"`
	// 机票金额（单位：元）
	//
	// example:
	//
	// xx.xx
	FlightAmount *float64 `json:"flightAmount,omitempty" xml:"flightAmount,omitempty"`
	// 福豆金额（单位：元）
	//
	// example:
	//
	// xx.xx
	FuPoint *float64 `json:"fuPoint,omitempty" xml:"fuPoint,omitempty"`
	// 酒店金额（单位：元）
	//
	// example:
	//
	// xx.xx
	HotelAmount *float64 `json:"hotelAmount,omitempty" xml:"hotelAmount,omitempty"`
	// 国际机票金额（单位：元）
	//
	// example:
	//
	// xx.xx
	IeFlightAmount *float64 `json:"ieFlightAmount,omitempty" xml:"ieFlightAmount,omitempty"`
	IeHotelAmount  *float64 `json:"ieHotelAmount,omitempty" xml:"ieHotelAmount,omitempty"`
	// 账期日：YYYYMMDD
	//
	// example:
	//
	// 20200501
	MailBillDate *int64   `json:"mailBillDate,omitempty" xml:"mailBillDate,omitempty"`
	MealAmount   *float64 `json:"mealAmount,omitempty" xml:"mealAmount,omitempty"`
	// 服务费金额（单位：元）
	//
	// example:
	//
	// xx.xx
	ServiceAmount *float64 `json:"serviceAmount,omitempty" xml:"serviceAmount,omitempty"`
	// 火车票金额（单位：元）
	//
	// example:
	//
	// xx.xx
	TrainAmount *float64 `json:"trainAmount,omitempty" xml:"trainAmount,omitempty"`
	VasAmount   *float64 `json:"vasAmount,omitempty" xml:"vasAmount,omitempty"`
}

func (s MonthBillGetResponseBodyModuleMonthAccountBillDetail) String() string {
	return dara.Prettify(s)
}

func (s MonthBillGetResponseBodyModuleMonthAccountBillDetail) GoString() string {
	return s.String()
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetBillConfirmed() *int32 {
	return s.BillConfirmed
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetCarAmount() *float64 {
	return s.CarAmount
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetDamageAmount() *float64 {
	return s.DamageAmount
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetFlightAmount() *float64 {
	return s.FlightAmount
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetFuPoint() *float64 {
	return s.FuPoint
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetHotelAmount() *float64 {
	return s.HotelAmount
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetIeFlightAmount() *float64 {
	return s.IeFlightAmount
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetIeHotelAmount() *float64 {
	return s.IeHotelAmount
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetMailBillDate() *int64 {
	return s.MailBillDate
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetMealAmount() *float64 {
	return s.MealAmount
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetServiceAmount() *float64 {
	return s.ServiceAmount
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetTrainAmount() *float64 {
	return s.TrainAmount
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) GetVasAmount() *float64 {
	return s.VasAmount
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetBillConfirmed(v int32) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.BillConfirmed = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetCarAmount(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.CarAmount = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetDamageAmount(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.DamageAmount = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetFlightAmount(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.FlightAmount = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetFuPoint(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.FuPoint = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetHotelAmount(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.HotelAmount = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetIeFlightAmount(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.IeFlightAmount = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetIeHotelAmount(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.IeHotelAmount = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetMailBillDate(v int64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.MailBillDate = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetMealAmount(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.MealAmount = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetServiceAmount(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.ServiceAmount = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetTrainAmount(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.TrainAmount = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) SetVasAmount(v float64) *MonthBillGetResponseBodyModuleMonthAccountBillDetail {
	s.VasAmount = &v
	return s
}

func (s *MonthBillGetResponseBodyModuleMonthAccountBillDetail) Validate() error {
	return dara.Validate(s)
}
