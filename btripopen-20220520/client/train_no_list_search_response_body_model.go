// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainNoListSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainNoListSearchResponseBody
	GetCode() *string
	SetMessage(v string) *TrainNoListSearchResponseBody
	GetMessage() *string
	SetModule(v *TrainNoListSearchResponseBodyModule) *TrainNoListSearchResponseBody
	GetModule() *TrainNoListSearchResponseBodyModule
	SetRequestId(v string) *TrainNoListSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainNoListSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainNoListSearchResponseBody
	GetTraceId() *string
}

type TrainNoListSearchResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *TrainNoListSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 210bc81a17090871660176894d008c
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

func (s TrainNoListSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchResponseBody) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainNoListSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainNoListSearchResponseBody) GetModule() *TrainNoListSearchResponseBodyModule {
	return s.Module
}

func (s *TrainNoListSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainNoListSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainNoListSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainNoListSearchResponseBody) SetCode(v string) *TrainNoListSearchResponseBody {
	s.Code = &v
	return s
}

func (s *TrainNoListSearchResponseBody) SetMessage(v string) *TrainNoListSearchResponseBody {
	s.Message = &v
	return s
}

func (s *TrainNoListSearchResponseBody) SetModule(v *TrainNoListSearchResponseBodyModule) *TrainNoListSearchResponseBody {
	s.Module = v
	return s
}

func (s *TrainNoListSearchResponseBody) SetRequestId(v string) *TrainNoListSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainNoListSearchResponseBody) SetSuccess(v bool) *TrainNoListSearchResponseBody {
	s.Success = &v
	return s
}

func (s *TrainNoListSearchResponseBody) SetTraceId(v string) *TrainNoListSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainNoListSearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainNoListSearchResponseBodyModule struct {
	TrainStationSearchVOS         []*TrainNoListSearchResponseBodyModuleTrainStationSearchVOS         `json:"train_station_search_v_o_s,omitempty" xml:"train_station_search_v_o_s,omitempty" type:"Repeated"`
	TrainTransferStationSearchVOs []*TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs `json:"train_transfer_station_search_v_os,omitempty" xml:"train_transfer_station_search_v_os,omitempty" type:"Repeated"`
}

func (s TrainNoListSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchResponseBodyModule) GetTrainStationSearchVOS() []*TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	return s.TrainStationSearchVOS
}

func (s *TrainNoListSearchResponseBodyModule) GetTrainTransferStationSearchVOs() []*TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs {
	return s.TrainTransferStationSearchVOs
}

func (s *TrainNoListSearchResponseBodyModule) SetTrainStationSearchVOS(v []*TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) *TrainNoListSearchResponseBodyModule {
	s.TrainStationSearchVOS = v
	return s
}

func (s *TrainNoListSearchResponseBodyModule) SetTrainTransferStationSearchVOs(v []*TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) *TrainNoListSearchResponseBodyModule {
	s.TrainTransferStationSearchVOs = v
	return s
}

func (s *TrainNoListSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainNoListSearchResponseBodyModuleTrainStationSearchVOS struct {
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
	// 2024-05-07 14:46:06
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// 30
	CostTime *string `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// example:
	//
	// BTC
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	DepStationName *string `json:"dep_station_name,omitempty" xml:"dep_station_name,omitempty"`
	// example:
	//
	// 2024-05-06 14:46:06
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
	SaleFlag    *string `json:"sale_flag,omitempty" xml:"sale_flag,omitempty"`
	SaleFlagMsg *string `json:"sale_flag_msg,omitempty" xml:"sale_flag_msg,omitempty"`
	// example:
	//
	// 0
	SeagmentIndex *string                                                              `json:"seagment_index,omitempty" xml:"seagment_index,omitempty"`
	SeatInfos     []*TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos `json:"seat_infos,omitempty" xml:"seat_infos,omitempty" type:"Repeated"`
	// example:
	//
	// D2345
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// example:
	//
	// 1
	TrainType *string `json:"train_type,omitempty" xml:"train_type,omitempty"`
}

func (s TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetArrDayTag() *string {
	return s.ArrDayTag
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetArrStationName() *string {
	return s.ArrStationName
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetCostTime() *string {
	return s.CostTime
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetDepStationName() *string {
	return s.DepStationName
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetIsEndStation() *int32 {
	return s.IsEndStation
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetIsStartStation() *int32 {
	return s.IsStartStation
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetPrice() *string {
	return s.Price
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetSaleFlag() *string {
	return s.SaleFlag
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetSaleFlagMsg() *string {
	return s.SaleFlagMsg
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetSeagmentIndex() *string {
	return s.SeagmentIndex
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetSeatInfos() []*TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos {
	return s.SeatInfos
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) GetTrainType() *string {
	return s.TrainType
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetArrDayTag(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.ArrDayTag = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetArrStationCode(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.ArrStationCode = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetArrStationName(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.ArrStationName = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetArrTime(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.ArrTime = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetCostTime(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.CostTime = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetDepStationCode(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.DepStationCode = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetDepStationName(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.DepStationName = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetDepTime(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.DepTime = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetIsEndStation(v int32) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.IsEndStation = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetIsStartStation(v int32) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.IsStartStation = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetPrice(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.Price = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetSaleFlag(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.SaleFlag = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetSaleFlagMsg(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.SaleFlagMsg = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetSeagmentIndex(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.SeagmentIndex = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetSeatInfos(v []*TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.SeatInfos = v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetTrainNo(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.TrainNo = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) SetTrainType(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS {
	s.TrainType = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOS) Validate() error {
	return dara.Validate(s)
}

type TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos struct {
	// example:
	//
	// 40900
	Price    *int64  `json:"price,omitempty" xml:"price,omitempty"`
	SeatName *string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
	// example:
	//
	// 14
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	Stock    *string `json:"stock,omitempty" xml:"stock,omitempty"`
}

func (s TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) GetPrice() *int64 {
	return s.Price
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) GetSeatName() *string {
	return s.SeatName
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) GetStock() *string {
	return s.Stock
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) SetPrice(v int64) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos {
	s.Price = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) SetSeatName(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos {
	s.SeatName = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) SetSeatType(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos {
	s.SeatType = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) SetStock(v string) *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos {
	s.Stock = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainStationSearchVOSSeatInfos) Validate() error {
	return dara.Validate(s)
}

type TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs struct {
	ArrStation         *string                                                                               `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	DepStation         *string                                                                               `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
	LineKey            *string                                                                               `json:"line_key,omitempty" xml:"line_key,omitempty"`
	MiddleStation      *string                                                                               `json:"middle_station,omitempty" xml:"middle_station,omitempty"`
	TransferDetailList []*TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList `json:"transfer_detail_list,omitempty" xml:"transfer_detail_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TransferType *string `json:"transfer_type,omitempty" xml:"transfer_type,omitempty"`
}

func (s TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) GetArrStation() *string {
	return s.ArrStation
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) GetDepStation() *string {
	return s.DepStation
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) GetLineKey() *string {
	return s.LineKey
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) GetMiddleStation() *string {
	return s.MiddleStation
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) GetTransferDetailList() []*TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	return s.TransferDetailList
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) GetTransferType() *string {
	return s.TransferType
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) SetArrStation(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs {
	s.ArrStation = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) SetDepStation(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs {
	s.DepStation = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) SetLineKey(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs {
	s.LineKey = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) SetMiddleStation(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs {
	s.MiddleStation = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) SetTransferDetailList(v []*TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs {
	s.TransferDetailList = v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) SetTransferType(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs {
	s.TransferType = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOs) Validate() error {
	return dara.Validate(s)
}

type TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList struct {
	// example:
	//
	// 1
	ArrDayTag *string `json:"arr_day_tag,omitempty" xml:"arr_day_tag,omitempty"`
	// example:
	//
	// 123344
	ArrStationCode *string `json:"arr_station_code,omitempty" xml:"arr_station_code,omitempty"`
	ArrStationName *string `json:"arr_station_name,omitempty" xml:"arr_station_name,omitempty"`
	// example:
	//
	// 2024-05-07 14:46:06
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// 234567
	CostTime *string `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// example:
	//
	// 12334
	DepStationCode *string `json:"dep_station_code,omitempty" xml:"dep_station_code,omitempty"`
	DepStationName *string `json:"dep_station_name,omitempty" xml:"dep_station_name,omitempty"`
	// example:
	//
	// 2024-05-06 14:46:06
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
	// 105000
	Price *string `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 1
	SaleFlag    *string                                                                                        `json:"sale_flag,omitempty" xml:"sale_flag,omitempty"`
	SaleFlagMsg *string                                                                                        `json:"sale_flag_msg,omitempty" xml:"sale_flag_msg,omitempty"`
	SeatInfos   []*TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos `json:"seat_infos,omitempty" xml:"seat_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	SegmentIndex *string `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// example:
	//
	// K2345
	TrainNo *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// example:
	//
	// 1
	TrainType *string `json:"train_type,omitempty" xml:"train_type,omitempty"`
}

func (s TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetArrDayTag() *string {
	return s.ArrDayTag
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetArrStationCode() *string {
	return s.ArrStationCode
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetArrStationName() *string {
	return s.ArrStationName
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetCostTime() *string {
	return s.CostTime
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetDepStationCode() *string {
	return s.DepStationCode
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetDepStationName() *string {
	return s.DepStationName
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetIsEndStation() *int32 {
	return s.IsEndStation
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetIsStartStation() *int32 {
	return s.IsStartStation
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetPrice() *string {
	return s.Price
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetSaleFlag() *string {
	return s.SaleFlag
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetSaleFlagMsg() *string {
	return s.SaleFlagMsg
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetSeatInfos() []*TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos {
	return s.SeatInfos
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetSegmentIndex() *string {
	return s.SegmentIndex
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) GetTrainType() *string {
	return s.TrainType
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetArrDayTag(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.ArrDayTag = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetArrStationCode(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.ArrStationCode = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetArrStationName(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.ArrStationName = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetArrTime(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.ArrTime = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetCostTime(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.CostTime = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetDepStationCode(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.DepStationCode = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetDepStationName(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.DepStationName = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetDepTime(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.DepTime = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetIsEndStation(v int32) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.IsEndStation = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetIsStartStation(v int32) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.IsStartStation = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetPrice(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.Price = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetSaleFlag(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.SaleFlag = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetSaleFlagMsg(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.SaleFlagMsg = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetSeatInfos(v []*TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.SeatInfos = v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetSegmentIndex(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.SegmentIndex = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetTrainNo(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.TrainNo = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) SetTrainType(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList {
	s.TrainType = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailList) Validate() error {
	return dara.Validate(s)
}

type TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos struct {
	// example:
	//
	// 40900
	Price    *int64  `json:"price,omitempty" xml:"price,omitempty"`
	SeatName *string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
	// example:
	//
	// 14
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// example:
	//
	// 33
	Stock *string `json:"stock,omitempty" xml:"stock,omitempty"`
}

func (s TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) String() string {
	return dara.Prettify(s)
}

func (s TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) GoString() string {
	return s.String()
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) GetPrice() *int64 {
	return s.Price
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) GetSeatName() *string {
	return s.SeatName
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) GetStock() *string {
	return s.Stock
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) SetPrice(v int64) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos {
	s.Price = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) SetSeatName(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos {
	s.SeatName = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) SetSeatType(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos {
	s.SeatType = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) SetStock(v string) *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos {
	s.Stock = &v
	return s
}

func (s *TrainNoListSearchResponseBodyModuleTrainTransferStationSearchVOsTransferDetailListSeatInfos) Validate() error {
	return dara.Validate(s)
}
