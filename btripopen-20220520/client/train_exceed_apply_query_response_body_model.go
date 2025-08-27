// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainExceedApplyQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainExceedApplyQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TrainExceedApplyQueryResponseBody
	GetMessage() *string
	SetModule(v *TrainExceedApplyQueryResponseBodyModule) *TrainExceedApplyQueryResponseBody
	GetModule() *TrainExceedApplyQueryResponseBodyModule
	SetRequestId(v string) *TrainExceedApplyQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainExceedApplyQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainExceedApplyQueryResponseBody
	GetTraceId() *string
}

type TrainExceedApplyQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TrainExceedApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// BE3FBA15-C5E1-5B99-8120-9FB84A04FB81
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TrainExceedApplyQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainExceedApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainExceedApplyQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainExceedApplyQueryResponseBody) GetModule() *TrainExceedApplyQueryResponseBodyModule {
	return s.Module
}

func (s *TrainExceedApplyQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainExceedApplyQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainExceedApplyQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainExceedApplyQueryResponseBody) SetCode(v string) *TrainExceedApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) SetMessage(v string) *TrainExceedApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) SetModule(v *TrainExceedApplyQueryResponseBodyModule) *TrainExceedApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) SetRequestId(v string) *TrainExceedApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) SetSuccess(v bool) *TrainExceedApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) SetTraceId(v string) *TrainExceedApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainExceedApplyQueryResponseBodyModule struct {
	// example:
	//
	// apply1
	ApplyId              *int64                                                       `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ApplyIntentionInfoDO *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO `json:"apply_intention_info_d_o,omitempty" xml:"apply_intention_info_d_o,omitempty" type:"Struct"`
	BtripCause           *string                                                      `json:"btrip_cause,omitempty" xml:"btrip_cause,omitempty"`
	// example:
	//
	// corp1
	CorpId       *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	ExceedReason *string `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	// example:
	//
	// 32
	ExceedType *int32 `json:"exceed_type,omitempty" xml:"exceed_type,omitempty"`
	// example:
	//
	// 10000
	OriginStandard *string `json:"origin_standard,omitempty" xml:"origin_standard,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2022-02-24T09:55Z
	SubmitTime *string `json:"submit_time,omitempty" xml:"submit_time,omitempty"`
	// example:
	//
	// 0001A1100000007EX08O
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// testcorp1
	ThirdpartCorpId *string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// example:
	//
	// user1
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s TrainExceedApplyQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainExceedApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetApplyIntentionInfoDO() *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	return s.ApplyIntentionInfoDO
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetBtripCause() *string {
	return s.BtripCause
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetExceedType() *int32 {
	return s.ExceedType
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetOriginStandard() *string {
	return s.OriginStandard
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetThirdpartCorpId() *string {
	return s.ThirdpartCorpId
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *TrainExceedApplyQueryResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetApplyId(v int64) *TrainExceedApplyQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetApplyIntentionInfoDO(v *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) *TrainExceedApplyQueryResponseBodyModule {
	s.ApplyIntentionInfoDO = v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetBtripCause(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.BtripCause = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetCorpId(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetExceedReason(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.ExceedReason = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetExceedType(v int32) *TrainExceedApplyQueryResponseBodyModule {
	s.ExceedType = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetOriginStandard(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.OriginStandard = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetStatus(v int32) *TrainExceedApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetSubmitTime(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.SubmitTime = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetThirdpartApplyId(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetThirdpartCorpId(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.ThirdpartCorpId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetUserId(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) SetUserName(v string) *TrainExceedApplyQueryResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO struct {
	// example:
	//
	// BJS
	ArrCity     *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrStation  *string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// example:
	//
	// 2022-02-24T09:55Z
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// SHA
	DepCity     *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepStation  *string `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
	// example:
	//
	// 2022-02-24T09:55Z
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 1000
	Price    *int64  `json:"price,omitempty" xml:"price,omitempty"`
	SeatName *string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
	// example:
	//
	// G39
	TrainNo       *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	TrainTypeDesc *string `json:"train_type_desc,omitempty" xml:"train_type_desc,omitempty"`
	// example:
	//
	// 32
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) String() string {
	return dara.Prettify(s)
}

func (s TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetArrCity() *string {
	return s.ArrCity
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetArrStation() *string {
	return s.ArrStation
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetDepCity() *string {
	return s.DepCity
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetDepCityName() *string {
	return s.DepCityName
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetDepStation() *string {
	return s.DepStation
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetDepTime() *string {
	return s.DepTime
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetPrice() *int64 {
	return s.Price
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetSeatName() *string {
	return s.SeatName
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetTrainTypeDesc() *string {
	return s.TrainTypeDesc
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) GetType() *int32 {
	return s.Type
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetArrCity(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.ArrCity = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetArrCityName(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.ArrCityName = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetArrStation(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.ArrStation = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetArrTime(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.ArrTime = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetDepCity(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.DepCity = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetDepCityName(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.DepCityName = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetDepStation(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.DepStation = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetDepTime(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.DepTime = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetPrice(v int64) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.Price = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetSeatName(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.SeatName = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetTrainNo(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.TrainNo = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetTrainTypeDesc(v string) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.TrainTypeDesc = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) SetType(v int32) *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO {
	s.Type = &v
	return s
}

func (s *TrainExceedApplyQueryResponseBodyModuleApplyIntentionInfoDO) Validate() error {
	return dara.Validate(s)
}
