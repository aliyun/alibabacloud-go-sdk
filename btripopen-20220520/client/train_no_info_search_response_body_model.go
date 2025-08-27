// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainNoInfoSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainNoInfoSearchResponseBody
	GetCode() *string
	SetMessage(v string) *TrainNoInfoSearchResponseBody
	GetMessage() *string
	SetModule(v *TrainNoInfoSearchResponseBodyModule) *TrainNoInfoSearchResponseBody
	GetModule() *TrainNoInfoSearchResponseBodyModule
	SetRequestId(v string) *TrainNoInfoSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainNoInfoSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainNoInfoSearchResponseBody
	GetTraceId() *string
}

type TrainNoInfoSearchResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainNoInfoSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// requestId
	//
	// example:
	//
	// 2136019116915615924561621e06ee
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainNoInfoSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainNoInfoSearchResponseBody) GoString() string {
	return s.String()
}

func (s *TrainNoInfoSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainNoInfoSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainNoInfoSearchResponseBody) GetModule() *TrainNoInfoSearchResponseBodyModule {
	return s.Module
}

func (s *TrainNoInfoSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainNoInfoSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainNoInfoSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainNoInfoSearchResponseBody) SetCode(v string) *TrainNoInfoSearchResponseBody {
	s.Code = &v
	return s
}

func (s *TrainNoInfoSearchResponseBody) SetMessage(v string) *TrainNoInfoSearchResponseBody {
	s.Message = &v
	return s
}

func (s *TrainNoInfoSearchResponseBody) SetModule(v *TrainNoInfoSearchResponseBodyModule) *TrainNoInfoSearchResponseBody {
	s.Module = v
	return s
}

func (s *TrainNoInfoSearchResponseBody) SetRequestId(v string) *TrainNoInfoSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainNoInfoSearchResponseBody) SetSuccess(v bool) *TrainNoInfoSearchResponseBody {
	s.Success = &v
	return s
}

func (s *TrainNoInfoSearchResponseBody) SetTraceId(v string) *TrainNoInfoSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainNoInfoSearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainNoInfoSearchResponseBodyModule struct {
	TrainStationSearchVO         *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO         `json:"train_station_search_v_o,omitempty" xml:"train_station_search_v_o,omitempty" type:"Struct"`
	TrainTransferStationSearchVO *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO `json:"train_transfer_station_search_v_o,omitempty" xml:"train_transfer_station_search_v_o,omitempty" type:"Struct"`
}

func (s TrainNoInfoSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainNoInfoSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainNoInfoSearchResponseBodyModule) GetTrainStationSearchVO() *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	return s.TrainStationSearchVO
}

func (s *TrainNoInfoSearchResponseBodyModule) GetTrainTransferStationSearchVO() *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO {
	return s.TrainTransferStationSearchVO
}

func (s *TrainNoInfoSearchResponseBodyModule) SetTrainStationSearchVO(v *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) *TrainNoInfoSearchResponseBodyModule {
	s.TrainStationSearchVO = v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModule) SetTrainTransferStationSearchVO(v *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO) *TrainNoInfoSearchResponseBodyModule {
	s.TrainTransferStationSearchVO = v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO struct {
	// example:
	//
	// 0
	ArrDayTag *string `json:"arr_day_tag,omitempty" xml:"arr_day_tag,omitempty"`
	// example:
	//
	// BDC
	ArrStationCode *string `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	ArrStationName *string `json:"arr_station_name,omitempty" xml:"arr_station_name,omitempty"`
	// example:
	//
	// 2024-05-07 15:19:01
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// 31
	CostTime *string `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// example:
	//
	// BTC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	DepStationName *string `json:"dep_station_name,omitempty" xml:"dep_station_name,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 1
	IsEndStation *int32 `json:"is_end_station,omitempty" xml:"is_end_station,omitempty"`
	// example:
	//
	// 0
	IsStartStation *int32 `json:"is_start_station,omitempty" xml:"is_start_station,omitempty"`
	// example:
	//
	// 54000
	Price *string `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 0
	SaleFlag  *string                                                             `json:"sale_flag,omitempty" xml:"sale_flag,omitempty"`
	SeatInfos []*TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos `json:"seat_infos,omitempty" xml:"seat_infos,omitempty" type:"Repeated"`
	// example:
	//
	// k2345
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// example:
	//
	// 1
	TrainType *string `json:"train_type,omitempty" xml:"train_type,omitempty"`
}

func (s TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) String() string {
	return dara.Prettify(s)
}

func (s TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GoString() string {
	return s.String()
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetArrDayTag() *string {
	return s.ArrDayTag
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetArrStationName() *string {
	return s.ArrStationName
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetCostTime() *string {
	return s.CostTime
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetDepStationName() *string {
	return s.DepStationName
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetIsEndStation() *int32 {
	return s.IsEndStation
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetIsStartStation() *int32 {
	return s.IsStartStation
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetPrice() *string {
	return s.Price
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetSaleFlag() *string {
	return s.SaleFlag
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetSeatInfos() []*TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos {
	return s.SeatInfos
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) GetTrainType() *string {
	return s.TrainType
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetArrDayTag(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.ArrDayTag = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetArrStationCode(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.ArrStationCode = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetArrStationName(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.ArrStationName = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetArrTime(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.ArrTime = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetCostTime(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.CostTime = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetDepStationCode(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.DepStationCode = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetDepStationName(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.DepStationName = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetDepTime(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.DepTime = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetIsEndStation(v int32) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.IsEndStation = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetIsStartStation(v int32) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.IsStartStation = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetPrice(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.Price = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetSaleFlag(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.SaleFlag = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetSeatInfos(v []*TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.SeatInfos = v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetTrainNo(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.TrainNo = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) SetTrainType(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO {
	s.TrainType = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVO) Validate() error {
	return dara.Validate(s)
}

type TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos struct {
	// example:
	//
	// 10000
	Price    *int64  `json:"price,omitempty" xml:"price,omitempty"`
	SeatName *string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
	// example:
	//
	// 14
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	Stock    *string `json:"stock,omitempty" xml:"stock,omitempty"`
}

func (s TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) GoString() string {
	return s.String()
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) GetPrice() *int64 {
	return s.Price
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) GetSeatName() *string {
	return s.SeatName
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) GetStock() *string {
	return s.Stock
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) SetPrice(v int64) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos {
	s.Price = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) SetSeatName(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos {
	s.SeatName = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) SetSeatType(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos {
	s.SeatType = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) SetStock(v string) *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos {
	s.Stock = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainStationSearchVOSeatInfos) Validate() error {
	return dara.Validate(s)
}

type TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO struct {
	TransferDetailList []*TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList `json:"transfer_detail_list,omitempty" xml:"transfer_detail_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TransferType *string `json:"transfer_type,omitempty" xml:"transfer_type,omitempty"`
}

func (s TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO) String() string {
	return dara.Prettify(s)
}

func (s TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO) GoString() string {
	return s.String()
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO) GetTransferDetailList() []*TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	return s.TransferDetailList
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO) GetTransferType() *string {
	return s.TransferType
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO) SetTransferDetailList(v []*TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO {
	s.TransferDetailList = v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO) SetTransferType(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO {
	s.TransferType = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVO) Validate() error {
	return dara.Validate(s)
}

type TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList struct {
	// example:
	//
	// 1
	ArrDayTag *string `json:"arr_day_tag,omitempty" xml:"arr_day_tag,omitempty"`
	// example:
	//
	// BDC
	ArrStationCode *string `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	ArrStationName *string `json:"arr_station_name,omitempty" xml:"arr_station_name,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// 111
	CostTime *string `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// example:
	//
	// BTC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	DepStationName *string `json:"dep_station_name,omitempty" xml:"dep_station_name,omitempty"`
	// example:
	//
	// 2024-05-06 15:19:01
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 1
	IsEndStation *int32 `json:"is_end_station,omitempty" xml:"is_end_station,omitempty"`
	// example:
	//
	// 1
	IsStartStation *int32 `json:"is_start_station,omitempty" xml:"is_start_station,omitempty"`
	// example:
	//
	// 54000
	Price *string `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 1
	SaleFlag  *string                                                                                       `json:"sale_flag,omitempty" xml:"sale_flag,omitempty"`
	SeatInfos []*TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos `json:"seat_infos,omitempty" xml:"seat_infos,omitempty" type:"Repeated"`
	// example:
	//
	// D1234
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// example:
	//
	// 1
	TrainType *string `json:"train_type,omitempty" xml:"train_type,omitempty"`
}

func (s TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) String() string {
	return dara.Prettify(s)
}

func (s TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GoString() string {
	return s.String()
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetArrDayTag() *string {
	return s.ArrDayTag
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetArrStationName() *string {
	return s.ArrStationName
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetCostTime() *string {
	return s.CostTime
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetDepStationName() *string {
	return s.DepStationName
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetIsEndStation() *int32 {
	return s.IsEndStation
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetIsStartStation() *int32 {
	return s.IsStartStation
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetPrice() *string {
	return s.Price
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetSaleFlag() *string {
	return s.SaleFlag
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetSeatInfos() []*TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos {
	return s.SeatInfos
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) GetTrainType() *string {
	return s.TrainType
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetArrDayTag(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.ArrDayTag = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetArrStationCode(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.ArrStationCode = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetArrStationName(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.ArrStationName = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetArrTime(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.ArrTime = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetCostTime(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.CostTime = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetDepStationCode(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.DepStationCode = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetDepStationName(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.DepStationName = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetDepTime(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.DepTime = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetIsEndStation(v int32) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.IsEndStation = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetIsStartStation(v int32) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.IsStartStation = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetPrice(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.Price = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetSaleFlag(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.SaleFlag = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetSeatInfos(v []*TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.SeatInfos = v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetTrainNo(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.TrainNo = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) SetTrainType(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList {
	s.TrainType = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailList) Validate() error {
	return dara.Validate(s)
}

type TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos struct {
	// example:
	//
	// 67000
	Price    *int64  `json:"price,omitempty" xml:"price,omitempty"`
	SeatName *string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
	// example:
	//
	// 14
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	Stock    *string `json:"stock,omitempty" xml:"stock,omitempty"`
}

func (s TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) GoString() string {
	return s.String()
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) GetPrice() *int64 {
	return s.Price
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) GetSeatName() *string {
	return s.SeatName
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) GetStock() *string {
	return s.Stock
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) SetPrice(v int64) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos {
	s.Price = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) SetSeatName(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos {
	s.SeatName = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) SetSeatType(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos {
	s.SeatType = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) SetStock(v string) *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos {
	s.Stock = &v
	return s
}

func (s *TrainNoInfoSearchResponseBodyModuleTrainTransferStationSearchVOTransferDetailListSeatInfos) Validate() error {
	return dara.Validate(s)
}
