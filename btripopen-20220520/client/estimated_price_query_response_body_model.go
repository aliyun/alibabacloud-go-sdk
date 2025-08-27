// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedPriceQueryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EstimatedPriceQueryResponseBody
  GetCode() *int32 
  SetMessage(v string) *EstimatedPriceQueryResponseBody
  GetMessage() *string 
  SetModule(v *EstimatedPriceQueryResponseBodyModule) *EstimatedPriceQueryResponseBody
  GetModule() *EstimatedPriceQueryResponseBodyModule 
  SetRequestId(v string) *EstimatedPriceQueryResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EstimatedPriceQueryResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *EstimatedPriceQueryResponseBody
  GetTraceId() *string 
}

type EstimatedPriceQueryResponseBody struct {
  // example:
  // 
  // 0
  Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
  // example:
  // 
  // demo
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  Module *EstimatedPriceQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
  // example:
  // 
  // A5009956-1077-52FB-B520-EA8C7E91D722
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // traceId
  // 
  // example:
  // 
  // 21041ce316577904808056433edbb2
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s EstimatedPriceQueryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryResponseBody) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EstimatedPriceQueryResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EstimatedPriceQueryResponseBody) GetModule() *EstimatedPriceQueryResponseBodyModule  {
  return s.Module
}

func (s *EstimatedPriceQueryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EstimatedPriceQueryResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EstimatedPriceQueryResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *EstimatedPriceQueryResponseBody) SetCode(v int32) *EstimatedPriceQueryResponseBody {
  s.Code = &v
  return s
}

func (s *EstimatedPriceQueryResponseBody) SetMessage(v string) *EstimatedPriceQueryResponseBody {
  s.Message = &v
  return s
}

func (s *EstimatedPriceQueryResponseBody) SetModule(v *EstimatedPriceQueryResponseBodyModule) *EstimatedPriceQueryResponseBody {
  s.Module = v
  return s
}

func (s *EstimatedPriceQueryResponseBody) SetRequestId(v string) *EstimatedPriceQueryResponseBody {
  s.RequestId = &v
  return s
}

func (s *EstimatedPriceQueryResponseBody) SetSuccess(v bool) *EstimatedPriceQueryResponseBody {
  s.Success = &v
  return s
}

func (s *EstimatedPriceQueryResponseBody) SetTraceId(v string) *EstimatedPriceQueryResponseBody {
  s.TraceId = &v
  return s
}

func (s *EstimatedPriceQueryResponseBody) Validate() error {
  return dara.Validate(s)
}

type EstimatedPriceQueryResponseBodyModule struct {
  HotelFeeDetail []*EstimatedPriceQueryResponseBodyModuleHotelFeeDetail `json:"hotel_fee_detail,omitempty" xml:"hotel_fee_detail,omitempty" type:"Repeated"`
  // 酒店费用总额，单位为元
  // 
  // example:
  // 
  // 500
  TotalHotelFee *int64 `json:"total_hotel_fee,omitempty" xml:"total_hotel_fee,omitempty"`
  TrafficFee *EstimatedPriceQueryResponseBodyModuleTrafficFee `json:"traffic_fee,omitempty" xml:"traffic_fee,omitempty" type:"Struct"`
}

func (s EstimatedPriceQueryResponseBodyModule) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryResponseBodyModule) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryResponseBodyModule) GetHotelFeeDetail() []*EstimatedPriceQueryResponseBodyModuleHotelFeeDetail  {
  return s.HotelFeeDetail
}

func (s *EstimatedPriceQueryResponseBodyModule) GetTotalHotelFee() *int64  {
  return s.TotalHotelFee
}

func (s *EstimatedPriceQueryResponseBodyModule) GetTrafficFee() *EstimatedPriceQueryResponseBodyModuleTrafficFee  {
  return s.TrafficFee
}

func (s *EstimatedPriceQueryResponseBodyModule) SetHotelFeeDetail(v []*EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) *EstimatedPriceQueryResponseBodyModule {
  s.HotelFeeDetail = v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModule) SetTotalHotelFee(v int64) *EstimatedPriceQueryResponseBodyModule {
  s.TotalHotelFee = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModule) SetTrafficFee(v *EstimatedPriceQueryResponseBodyModuleTrafficFee) *EstimatedPriceQueryResponseBodyModule {
  s.TrafficFee = v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModule) Validate() error {
  return dara.Validate(s)
}

type EstimatedPriceQueryResponseBodyModuleHotelFeeDetail struct {
  City *string `json:"city,omitempty" xml:"city,omitempty"`
  // example:
  // 
  // 6
  Criterion *int64 `json:"criterion,omitempty" xml:"criterion,omitempty"`
  // example:
  // 
  // 1245
  ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
  // example:
  // 
  // 6
  Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
  // example:
  // 
  // 1
  TripDays *int32 `json:"trip_days,omitempty" xml:"trip_days,omitempty"`
}

func (s EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) GetCity() *string  {
  return s.City
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) GetCriterion() *int64  {
  return s.Criterion
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) GetItineraryId() *string  {
  return s.ItineraryId
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) GetTotal() *int64  {
  return s.Total
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) GetTripDays() *int32  {
  return s.TripDays
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) SetCity(v string) *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail {
  s.City = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) SetCriterion(v int64) *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail {
  s.Criterion = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) SetItineraryId(v string) *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail {
  s.ItineraryId = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) SetTotal(v int64) *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail {
  s.Total = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) SetTripDays(v int32) *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail {
  s.TripDays = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleHotelFeeDetail) Validate() error {
  return dara.Validate(s)
}

type EstimatedPriceQueryResponseBodyModuleTrafficFee struct {
  BtripRoutes []*EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes `json:"btrip_routes,omitempty" xml:"btrip_routes,omitempty" type:"Repeated"`
  // example:
  // 
  // demo
  ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
  // example:
  // 
  // 265000
  MaxFee *int64 `json:"max_fee,omitempty" xml:"max_fee,omitempty"`
  // example:
  // 
  // 30100
  MinFee *int64 `json:"min_fee,omitempty" xml:"min_fee,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EstimatedPriceQueryResponseBodyModuleTrafficFee) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryResponseBodyModuleTrafficFee) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) GetBtripRoutes() []*EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes  {
  return s.BtripRoutes
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) GetErrMsg() *string  {
  return s.ErrMsg
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) GetMaxFee() *int64  {
  return s.MaxFee
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) GetMinFee() *int64  {
  return s.MinFee
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) GetSuccess() *bool  {
  return s.Success
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) SetBtripRoutes(v []*EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) *EstimatedPriceQueryResponseBodyModuleTrafficFee {
  s.BtripRoutes = v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) SetErrMsg(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFee {
  s.ErrMsg = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) SetMaxFee(v int64) *EstimatedPriceQueryResponseBodyModuleTrafficFee {
  s.MaxFee = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) SetMinFee(v int64) *EstimatedPriceQueryResponseBodyModuleTrafficFee {
  s.MinFee = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) SetSuccess(v bool) *EstimatedPriceQueryResponseBodyModuleTrafficFee {
  s.Success = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFee) Validate() error {
  return dara.Validate(s)
}

type EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes struct {
  ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
  ArrDate *int64 `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
  Cheapest *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest `json:"cheapest,omitempty" xml:"cheapest,omitempty" type:"Struct"`
  DepCity *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
  DepDate *int64 `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
  // example:
  // 
  // demo
  ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
  // example:
  // 
  // 1245
  ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
  MostExpensive *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive `json:"most_expensive,omitempty" xml:"most_expensive,omitempty" type:"Struct"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) GetArrCity() *string  {
  return s.ArrCity
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) GetArrDate() *int64  {
  return s.ArrDate
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) GetCheapest() *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest  {
  return s.Cheapest
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) GetDepCity() *string  {
  return s.DepCity
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) GetDepDate() *int64  {
  return s.DepDate
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) GetErrMsg() *string  {
  return s.ErrMsg
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) GetItineraryId() *string  {
  return s.ItineraryId
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) GetMostExpensive() *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive  {
  return s.MostExpensive
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) GetSuccess() *bool  {
  return s.Success
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) SetArrCity(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes {
  s.ArrCity = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) SetArrDate(v int64) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes {
  s.ArrDate = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) SetCheapest(v *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes {
  s.Cheapest = v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) SetDepCity(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes {
  s.DepCity = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) SetDepDate(v int64) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes {
  s.DepDate = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) SetErrMsg(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes {
  s.ErrMsg = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) SetItineraryId(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes {
  s.ItineraryId = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) SetMostExpensive(v *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes {
  s.MostExpensive = v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) SetSuccess(v bool) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes {
  s.Success = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutes) Validate() error {
  return dara.Validate(s)
}

type EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest struct {
  // example:
  // 
  // 00:40
  ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
  // example:
  // 
  // 22:20
  DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
  // example:
  // 
  // 30100
  Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
  SeatGrade *string `json:"seat_grade,omitempty" xml:"seat_grade,omitempty"`
  // example:
  // 
  // MU9668
  VehicleNo *string `json:"vehicle_no,omitempty" xml:"vehicle_no,omitempty"`
}

func (s EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) GetArrTime() *string  {
  return s.ArrTime
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) GetDepTime() *string  {
  return s.DepTime
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) GetFee() *int64  {
  return s.Fee
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) GetSeatGrade() *string  {
  return s.SeatGrade
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) GetVehicleNo() *string  {
  return s.VehicleNo
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) SetArrTime(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest {
  s.ArrTime = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) SetDepTime(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest {
  s.DepTime = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) SetFee(v int64) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest {
  s.Fee = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) SetSeatGrade(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest {
  s.SeatGrade = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) SetVehicleNo(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest {
  s.VehicleNo = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesCheapest) Validate() error {
  return dara.Validate(s)
}

type EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive struct {
  // example:
  // 
  // 19:20
  ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
  // example:
  // 
  // 17:00
  DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
  // example:
  // 
  // 265000
  Fee *int64 `json:"fee,omitempty" xml:"fee,omitempty"`
  SeatGrade *string `json:"seat_grade,omitempty" xml:"seat_grade,omitempty"`
  // example:
  // 
  // CA1721
  VehicleNo *string `json:"vehicle_no,omitempty" xml:"vehicle_no,omitempty"`
}

func (s EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) GetArrTime() *string  {
  return s.ArrTime
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) GetDepTime() *string  {
  return s.DepTime
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) GetFee() *int64  {
  return s.Fee
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) GetSeatGrade() *string  {
  return s.SeatGrade
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) GetVehicleNo() *string  {
  return s.VehicleNo
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) SetArrTime(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive {
  s.ArrTime = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) SetDepTime(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive {
  s.DepTime = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) SetFee(v int64) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive {
  s.Fee = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) SetSeatGrade(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive {
  s.SeatGrade = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) SetVehicleNo(v string) *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive {
  s.VehicleNo = &v
  return s
}

func (s *EstimatedPriceQueryResponseBodyModuleTrafficFeeBtripRoutesMostExpensive) Validate() error {
  return dara.Validate(s)
}

