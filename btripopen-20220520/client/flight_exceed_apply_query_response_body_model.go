// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightExceedApplyQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightExceedApplyQueryResponseBody
	GetCode() *string
	SetMessage(v string) *FlightExceedApplyQueryResponseBody
	GetMessage() *string
	SetModule(v *FlightExceedApplyQueryResponseBodyModule) *FlightExceedApplyQueryResponseBody
	GetModule() *FlightExceedApplyQueryResponseBodyModule
	SetRequestId(v string) *FlightExceedApplyQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightExceedApplyQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightExceedApplyQueryResponseBody
	GetTraceId() *string
}

type FlightExceedApplyQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightExceedApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s FlightExceedApplyQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightExceedApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightExceedApplyQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightExceedApplyQueryResponseBody) GetModule() *FlightExceedApplyQueryResponseBodyModule {
	return s.Module
}

func (s *FlightExceedApplyQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightExceedApplyQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightExceedApplyQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightExceedApplyQueryResponseBody) SetCode(v string) *FlightExceedApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) SetMessage(v string) *FlightExceedApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) SetModule(v *FlightExceedApplyQueryResponseBodyModule) *FlightExceedApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) SetRequestId(v string) *FlightExceedApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) SetSuccess(v bool) *FlightExceedApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) SetTraceId(v string) *FlightExceedApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightExceedApplyQueryResponseBodyModule struct {
	// example:
	//
	// 1231
	ApplyId                  *int64                                                              `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ApplyIntentionInfoDo     *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo       `json:"apply_intention_info_do,omitempty" xml:"apply_intention_info_do,omitempty" type:"Struct"`
	ApplyIntentionInfoDoList []*FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList `json:"apply_intention_info_do_list,omitempty" xml:"apply_intention_info_do_list,omitempty" type:"Repeated"`
	ApplyRecommendFlights    *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights      `json:"apply_recommend_flights,omitempty" xml:"apply_recommend_flights,omitempty" type:"Struct"`
	BtripCause               *string                                                             `json:"btrip_cause,omitempty" xml:"btrip_cause,omitempty"`
	// example:
	//
	// 123
	CorpId       *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	ExceedReason *string `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	// example:
	//
	// 1
	ExceedType     *int32  `json:"exceed_type,omitempty" xml:"exceed_type,omitempty"`
	OriginStandard *string `json:"origin_standard,omitempty" xml:"origin_standard,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2022-01-12T16:47Z
	SubmitTime *string `json:"submit_time,omitempty" xml:"submit_time,omitempty"`
	// example:
	//
	// 0001A1100000007EX08O
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// test_corp
	ThirdpartCorpId *string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// example:
	//
	// user1
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s FlightExceedApplyQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightExceedApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetApplyIntentionInfoDo() *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	return s.ApplyIntentionInfoDo
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetApplyIntentionInfoDoList() []*FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	return s.ApplyIntentionInfoDoList
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetApplyRecommendFlights() *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	return s.ApplyRecommendFlights
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetBtripCause() *string {
	return s.BtripCause
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetExceedType() *int32 {
	return s.ExceedType
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetOriginStandard() *string {
	return s.OriginStandard
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetThirdpartCorpId() *string {
	return s.ThirdpartCorpId
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *FlightExceedApplyQueryResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetApplyId(v int64) *FlightExceedApplyQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetApplyIntentionInfoDo(v *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) *FlightExceedApplyQueryResponseBodyModule {
	s.ApplyIntentionInfoDo = v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetApplyIntentionInfoDoList(v []*FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) *FlightExceedApplyQueryResponseBodyModule {
	s.ApplyIntentionInfoDoList = v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetApplyRecommendFlights(v *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) *FlightExceedApplyQueryResponseBodyModule {
	s.ApplyRecommendFlights = v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetBtripCause(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.BtripCause = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetCorpId(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetExceedReason(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.ExceedReason = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetExceedType(v int32) *FlightExceedApplyQueryResponseBodyModule {
	s.ExceedType = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetOriginStandard(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.OriginStandard = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetStatus(v int32) *FlightExceedApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetSubmitTime(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.SubmitTime = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetThirdpartApplyId(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetThirdpartCorpId(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.ThirdpartCorpId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetUserId(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) SetUserName(v string) *FlightExceedApplyQueryResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo struct {
	ArrAirportName *string `json:"arr_airport_name,omitempty" xml:"arr_airport_name,omitempty"`
	// example:
	//
	// SHA
	ArrCity     *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2022-01-12T16:47Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// F
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// 1
	CabinClass     *int32  `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassStr  *string `json:"cabin_class_str,omitempty" xml:"cabin_class_str,omitempty"`
	DepAirportName *string `json:"dep_airport_name,omitempty" xml:"dep_airport_name,omitempty"`
	// example:
	//
	// SHA
	DepCity     *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2022-01-12T16:47Z
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 4
	Discount *string `json:"discount,omitempty" xml:"discount,omitempty"`
	// example:
	//
	// MU2759
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// 300
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) String() string {
	return dara.Prettify(s)
}

func (s FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetArrAirportName() *string {
	return s.ArrAirportName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetCabin() *string {
	return s.Cabin
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetCabinClass() *int32 {
	return s.CabinClass
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetCabinClassStr() *string {
	return s.CabinClassStr
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetDepAirportName() *string {
	return s.DepAirportName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetDepCity() *string {
	return s.DepCity
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetDiscount() *string {
	return s.Discount
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetPrice() *int64 {
	return s.Price
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetType() *int32 {
	return s.Type
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetArrAirportName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.ArrAirportName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetArrCity(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.ArrCity = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetArrCityName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.ArrCityName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetArrTime(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.ArrTime = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCabin(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Cabin = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCabinClass(v int32) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CabinClass = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCabinClassStr(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CabinClassStr = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetDepAirportName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.DepAirportName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetDepCity(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.DepCity = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetDepCityName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.DepCityName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetDepTime(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.DepTime = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetDiscount(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Discount = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetFlightNo(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.FlightNo = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetPrice(v int64) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Price = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetType(v int32) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Type = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) Validate() error {
	return dara.Validate(s)
}

type FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList struct {
	ArrAirportName *string `json:"arr_airport_name,omitempty" xml:"arr_airport_name,omitempty"`
	ArrCity        *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityName    *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrTime        *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	Cabin          *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass     *int32  `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassStr  *string `json:"cabin_class_str,omitempty" xml:"cabin_class_str,omitempty"`
	DepAirportName *string `json:"dep_airport_name,omitempty" xml:"dep_airport_name,omitempty"`
	DepCity        *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityName    *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepTime        *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	Discount       *string `json:"discount,omitempty" xml:"discount,omitempty"`
	FlightNo       *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	Price          *int64  `json:"price,omitempty" xml:"price,omitempty"`
	Type           *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) String() string {
	return dara.Prettify(s)
}

func (s FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetArrAirportName() *string {
	return s.ArrAirportName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetCabinClass() *int32 {
	return s.CabinClass
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetCabinClassStr() *string {
	return s.CabinClassStr
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetDepAirportName() *string {
	return s.DepAirportName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetDepCity() *string {
	return s.DepCity
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetDiscount() *string {
	return s.Discount
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetPrice() *int64 {
	return s.Price
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) GetType() *int32 {
	return s.Type
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetArrAirportName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.ArrAirportName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetArrCity(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.ArrCity = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetArrCityName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.ArrCityName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetArrTime(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.ArrTime = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetCabin(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.Cabin = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetCabinClass(v int32) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.CabinClass = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetCabinClassStr(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.CabinClassStr = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetDepAirportName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.DepAirportName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetDepCity(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.DepCity = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetDepCityName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.DepCityName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetDepTime(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.DepTime = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetDiscount(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.Discount = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetFlightNo(v string) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.FlightNo = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetPrice(v int64) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.Price = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) SetType(v int32) *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList {
	s.Type = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyIntentionInfoDoList) Validate() error {
	return dara.Validate(s)
}

type FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights struct {
	ArrAirportName      *string `json:"arr_airport_name,omitempty" xml:"arr_airport_name,omitempty"`
	ArrCityName         *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrTime             *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	Cabin               *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass          *int32  `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassStr       *string `json:"cabin_class_str,omitempty" xml:"cabin_class_str,omitempty"`
	DepAirportName      *string `json:"dep_airport_name,omitempty" xml:"dep_airport_name,omitempty"`
	DepCityName         *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepTime             *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	Discount            *string `json:"discount,omitempty" xml:"discount,omitempty"`
	FlightNo            *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	Price               *int64  `json:"price,omitempty" xml:"price,omitempty"`
	TransferAirportName *string `json:"transfer_airport_name,omitempty" xml:"transfer_airport_name,omitempty"`
}

func (s FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) String() string {
	return dara.Prettify(s)
}

func (s FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GoString() string {
	return s.String()
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetArrAirportName() *string {
	return s.ArrAirportName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetCabin() *string {
	return s.Cabin
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetCabinClass() *int32 {
	return s.CabinClass
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetCabinClassStr() *string {
	return s.CabinClassStr
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetDepAirportName() *string {
	return s.DepAirportName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetDiscount() *string {
	return s.Discount
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetPrice() *int64 {
	return s.Price
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) GetTransferAirportName() *string {
	return s.TransferAirportName
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetArrAirportName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.ArrAirportName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetArrCityName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.ArrCityName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetArrTime(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.ArrTime = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetCabin(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.Cabin = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetCabinClass(v int32) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.CabinClass = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetCabinClassStr(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.CabinClassStr = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetDepAirportName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.DepAirportName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetDepCityName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.DepCityName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetDepTime(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.DepTime = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetDiscount(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.Discount = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetFlightNo(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.FlightNo = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetPrice(v int64) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.Price = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) SetTransferAirportName(v string) *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights {
	s.TransferAirportName = &v
	return s
}

func (s *FlightExceedApplyQueryResponseBodyModuleApplyRecommendFlights) Validate() error {
	return dara.Validate(s)
}
