// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthPreBillGetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MonthPreBillGetResponseBody
	GetCode() *string
	SetMessage(v string) *MonthPreBillGetResponseBody
	GetMessage() *string
	SetModule(v []*MonthPreBillGetResponseBodyModule) *MonthPreBillGetResponseBody
	GetModule() []*MonthPreBillGetResponseBodyModule
	SetRequestId(v string) *MonthPreBillGetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MonthPreBillGetResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MonthPreBillGetResponseBody
	GetTraceId() *string
}

type MonthPreBillGetResponseBody struct {
	Code      *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module    []*MonthPreBillGetResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                              `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s MonthPreBillGetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MonthPreBillGetResponseBody) GoString() string {
	return s.String()
}

func (s *MonthPreBillGetResponseBody) GetCode() *string {
	return s.Code
}

func (s *MonthPreBillGetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MonthPreBillGetResponseBody) GetModule() []*MonthPreBillGetResponseBodyModule {
	return s.Module
}

func (s *MonthPreBillGetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MonthPreBillGetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MonthPreBillGetResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MonthPreBillGetResponseBody) SetCode(v string) *MonthPreBillGetResponseBody {
	s.Code = &v
	return s
}

func (s *MonthPreBillGetResponseBody) SetMessage(v string) *MonthPreBillGetResponseBody {
	s.Message = &v
	return s
}

func (s *MonthPreBillGetResponseBody) SetModule(v []*MonthPreBillGetResponseBodyModule) *MonthPreBillGetResponseBody {
	s.Module = v
	return s
}

func (s *MonthPreBillGetResponseBody) SetRequestId(v string) *MonthPreBillGetResponseBody {
	s.RequestId = &v
	return s
}

func (s *MonthPreBillGetResponseBody) SetSuccess(v bool) *MonthPreBillGetResponseBody {
	s.Success = &v
	return s
}

func (s *MonthPreBillGetResponseBody) SetTraceId(v string) *MonthPreBillGetResponseBody {
	s.TraceId = &v
	return s
}

func (s *MonthPreBillGetResponseBody) Validate() error {
	if s.Module != nil {
		for _, item := range s.Module {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MonthPreBillGetResponseBodyModule struct {
	EndDate                *string                                                  `json:"end_date,omitempty" xml:"end_date,omitempty"`
	MonthAccountBillDetail *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail `json:"monthAccountBillDetail,omitempty" xml:"monthAccountBillDetail,omitempty" type:"Struct"`
	StartDate              *string                                                  `json:"start_date,omitempty" xml:"start_date,omitempty"`
	Url                    *string                                                  `json:"url,omitempty" xml:"url,omitempty"`
}

func (s MonthPreBillGetResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s MonthPreBillGetResponseBodyModule) GoString() string {
	return s.String()
}

func (s *MonthPreBillGetResponseBodyModule) GetEndDate() *string {
	return s.EndDate
}

func (s *MonthPreBillGetResponseBodyModule) GetMonthAccountBillDetail() *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	return s.MonthAccountBillDetail
}

func (s *MonthPreBillGetResponseBodyModule) GetStartDate() *string {
	return s.StartDate
}

func (s *MonthPreBillGetResponseBodyModule) GetUrl() *string {
	return s.Url
}

func (s *MonthPreBillGetResponseBodyModule) SetEndDate(v string) *MonthPreBillGetResponseBodyModule {
	s.EndDate = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModule) SetMonthAccountBillDetail(v *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) *MonthPreBillGetResponseBodyModule {
	s.MonthAccountBillDetail = v
	return s
}

func (s *MonthPreBillGetResponseBodyModule) SetStartDate(v string) *MonthPreBillGetResponseBodyModule {
	s.StartDate = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModule) SetUrl(v string) *MonthPreBillGetResponseBodyModule {
	s.Url = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModule) Validate() error {
	if s.MonthAccountBillDetail != nil {
		if err := s.MonthAccountBillDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MonthPreBillGetResponseBodyModuleMonthAccountBillDetail struct {
	BillConfirmed  *int32   `json:"billConfirmed,omitempty" xml:"billConfirmed,omitempty"`
	CarAmount      *float64 `json:"carAmount,omitempty" xml:"carAmount,omitempty"`
	DamageAmount   *float64 `json:"damageAmount,omitempty" xml:"damageAmount,omitempty"`
	FlightAmount   *float64 `json:"flightAmount,omitempty" xml:"flightAmount,omitempty"`
	FuPoint        *float64 `json:"fuPoint,omitempty" xml:"fuPoint,omitempty"`
	HotelAmount    *float64 `json:"hotelAmount,omitempty" xml:"hotelAmount,omitempty"`
	IeFlightAmount *float64 `json:"ieFlightAmount,omitempty" xml:"ieFlightAmount,omitempty"`
	IeHotelAmount  *float64 `json:"ieHotelAmount,omitempty" xml:"ieHotelAmount,omitempty"`
	MailBillDate   *int64   `json:"mailBillDate,omitempty" xml:"mailBillDate,omitempty"`
	MealAmount     *float64 `json:"mealAmount,omitempty" xml:"mealAmount,omitempty"`
	ServiceAmount  *float64 `json:"serviceAmount,omitempty" xml:"serviceAmount,omitempty"`
	TrainAmount    *float64 `json:"trainAmount,omitempty" xml:"trainAmount,omitempty"`
	VasAmount      *float64 `json:"vasAmount,omitempty" xml:"vasAmount,omitempty"`
}

func (s MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) String() string {
	return dara.Prettify(s)
}

func (s MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GoString() string {
	return s.String()
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetBillConfirmed() *int32 {
	return s.BillConfirmed
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetCarAmount() *float64 {
	return s.CarAmount
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetDamageAmount() *float64 {
	return s.DamageAmount
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetFlightAmount() *float64 {
	return s.FlightAmount
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetFuPoint() *float64 {
	return s.FuPoint
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetHotelAmount() *float64 {
	return s.HotelAmount
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetIeFlightAmount() *float64 {
	return s.IeFlightAmount
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetIeHotelAmount() *float64 {
	return s.IeHotelAmount
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetMailBillDate() *int64 {
	return s.MailBillDate
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetMealAmount() *float64 {
	return s.MealAmount
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetServiceAmount() *float64 {
	return s.ServiceAmount
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetTrainAmount() *float64 {
	return s.TrainAmount
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) GetVasAmount() *float64 {
	return s.VasAmount
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetBillConfirmed(v int32) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.BillConfirmed = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetCarAmount(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.CarAmount = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetDamageAmount(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.DamageAmount = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetFlightAmount(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.FlightAmount = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetFuPoint(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.FuPoint = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetHotelAmount(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.HotelAmount = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetIeFlightAmount(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.IeFlightAmount = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetIeHotelAmount(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.IeHotelAmount = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetMailBillDate(v int64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.MailBillDate = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetMealAmount(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.MealAmount = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetServiceAmount(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.ServiceAmount = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetTrainAmount(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.TrainAmount = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) SetVasAmount(v float64) *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail {
	s.VasAmount = &v
	return s
}

func (s *MonthPreBillGetResponseBodyModuleMonthAccountBillDetail) Validate() error {
	return dara.Validate(s)
}
