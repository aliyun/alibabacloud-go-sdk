// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelExceedApplyQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelExceedApplyQueryResponseBody
	GetCode() *string
	SetMessage(v string) *HotelExceedApplyQueryResponseBody
	GetMessage() *string
	SetModule(v *HotelExceedApplyQueryResponseBodyModule) *HotelExceedApplyQueryResponseBody
	GetModule() *HotelExceedApplyQueryResponseBodyModule
	SetRequestId(v string) *HotelExceedApplyQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelExceedApplyQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelExceedApplyQueryResponseBody
	GetTraceId() *string
}

type HotelExceedApplyQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelExceedApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// E03779E6-71C2-5082-ABE3-6315B8247EE8
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

func (s HotelExceedApplyQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelExceedApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelExceedApplyQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelExceedApplyQueryResponseBody) GetModule() *HotelExceedApplyQueryResponseBodyModule {
	return s.Module
}

func (s *HotelExceedApplyQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelExceedApplyQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelExceedApplyQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelExceedApplyQueryResponseBody) SetCode(v string) *HotelExceedApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) SetMessage(v string) *HotelExceedApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) SetModule(v *HotelExceedApplyQueryResponseBodyModule) *HotelExceedApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) SetRequestId(v string) *HotelExceedApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) SetSuccess(v bool) *HotelExceedApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) SetTraceId(v string) *HotelExceedApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelExceedApplyQueryResponseBodyModule struct {
	// example:
	//
	// 27238197
	ApplyId              *int64                                                       `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ApplyIntentionInfoDo *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo `json:"apply_intention_info_do,omitempty" xml:"apply_intention_info_do,omitempty" type:"Struct"`
	BtripCause           *string                                                      `json:"btrip_cause,omitempty" xml:"btrip_cause,omitempty"`
	// example:
	//
	// corp1
	CorpId       *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	ExceedReason *string `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	// example:
	//
	// 16
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
	// 2022-02-11T16:25Z
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

func (s HotelExceedApplyQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelExceedApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetApplyIntentionInfoDo() *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	return s.ApplyIntentionInfoDo
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetBtripCause() *string {
	return s.BtripCause
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetExceedType() *int32 {
	return s.ExceedType
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetOriginStandard() *string {
	return s.OriginStandard
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetThirdpartCorpId() *string {
	return s.ThirdpartCorpId
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *HotelExceedApplyQueryResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetApplyId(v int64) *HotelExceedApplyQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetApplyIntentionInfoDo(v *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) *HotelExceedApplyQueryResponseBodyModule {
	s.ApplyIntentionInfoDo = v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetBtripCause(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.BtripCause = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetCorpId(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetExceedReason(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.ExceedReason = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetExceedType(v int32) *HotelExceedApplyQueryResponseBodyModule {
	s.ExceedType = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetOriginStandard(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.OriginStandard = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetStatus(v int32) *HotelExceedApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetSubmitTime(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.SubmitTime = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetThirdpartApplyId(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetThirdpartCorpId(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.ThirdpartCorpId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetUserId(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) SetUserName(v string) *HotelExceedApplyQueryResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo struct {
	// example:
	//
	// 2021-07-08
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// example:
	//
	// 2021-07-08
	CheckOut *string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// example:
	//
	// SHA
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 10000
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// true
	Together *bool `json:"together,omitempty" xml:"together,omitempty"`
	// example:
	//
	// 16
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) String() string {
	return dara.Prettify(s)
}

func (s HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GoString() string {
	return s.String()
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetCityName() *string {
	return s.CityName
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetPrice() *int64 {
	return s.Price
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetTogether() *bool {
	return s.Together
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) GetType() *int32 {
	return s.Type
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCheckIn(v string) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CheckIn = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCheckOut(v string) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CheckOut = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCityCode(v string) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CityCode = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetCityName(v string) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.CityName = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetPrice(v int64) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Price = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetTogether(v bool) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Together = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) SetType(v int32) *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo {
	s.Type = &v
	return s
}

func (s *HotelExceedApplyQueryResponseBodyModuleApplyIntentionInfoDo) Validate() error {
	return dara.Validate(s)
}
