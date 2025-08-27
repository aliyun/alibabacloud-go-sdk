// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonApplyQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CommonApplyQueryResponseBody
	GetCode() *string
	SetMessage(v string) *CommonApplyQueryResponseBody
	GetMessage() *string
	SetModule(v *CommonApplyQueryResponseBodyModule) *CommonApplyQueryResponseBody
	GetModule() *CommonApplyQueryResponseBodyModule
	SetRequestId(v string) *CommonApplyQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CommonApplyQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CommonApplyQueryResponseBody
	GetTraceId() *string
}

type CommonApplyQueryResponseBody struct {
	// example:
	//
	// 0
	Code    *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module  *CommonApplyQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s CommonApplyQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CommonApplyQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CommonApplyQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CommonApplyQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CommonApplyQueryResponseBody) GetModule() *CommonApplyQueryResponseBodyModule {
	return s.Module
}

func (s *CommonApplyQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CommonApplyQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CommonApplyQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CommonApplyQueryResponseBody) SetCode(v string) *CommonApplyQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CommonApplyQueryResponseBody) SetMessage(v string) *CommonApplyQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CommonApplyQueryResponseBody) SetModule(v *CommonApplyQueryResponseBodyModule) *CommonApplyQueryResponseBody {
	s.Module = v
	return s
}

func (s *CommonApplyQueryResponseBody) SetRequestId(v string) *CommonApplyQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CommonApplyQueryResponseBody) SetSuccess(v bool) *CommonApplyQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CommonApplyQueryResponseBody) SetTraceId(v string) *CommonApplyQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *CommonApplyQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CommonApplyQueryResponseBodyModule struct {
	// example:
	//
	// 123
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 5
	BizCategory *int32  `json:"biz_category,omitempty" xml:"biz_category,omitempty"`
	Cause       *string `json:"cause,omitempty" xml:"cause,omitempty"`
	// example:
	//
	// 0001A1100000007EX08O
	CorpId      *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	ExtendValue *string `json:"extend_value,omitempty" xml:"extend_value,omitempty"`
	// example:
	//
	// 2021-03-18T20:26Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1233
	ThirdpartCorpId *string `json:"thirdpart_corp_id,omitempty" xml:"thirdpart_corp_id,omitempty"`
	// example:
	//
	// 12344
	ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	TripCause   *string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	// example:
	//
	// user1
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CommonApplyQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CommonApplyQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CommonApplyQueryResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *CommonApplyQueryResponseBodyModule) GetBizCategory() *int32 {
	return s.BizCategory
}

func (s *CommonApplyQueryResponseBodyModule) GetCause() *string {
	return s.Cause
}

func (s *CommonApplyQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *CommonApplyQueryResponseBodyModule) GetExtendValue() *string {
	return s.ExtendValue
}

func (s *CommonApplyQueryResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CommonApplyQueryResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *CommonApplyQueryResponseBodyModule) GetThirdpartCorpId() *string {
	return s.ThirdpartCorpId
}

func (s *CommonApplyQueryResponseBodyModule) GetThirdpartId() *string {
	return s.ThirdpartId
}

func (s *CommonApplyQueryResponseBodyModule) GetTripCause() *string {
	return s.TripCause
}

func (s *CommonApplyQueryResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *CommonApplyQueryResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *CommonApplyQueryResponseBodyModule) SetApplyId(v int64) *CommonApplyQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetBizCategory(v int32) *CommonApplyQueryResponseBodyModule {
	s.BizCategory = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetCause(v string) *CommonApplyQueryResponseBodyModule {
	s.Cause = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetCorpId(v string) *CommonApplyQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetExtendValue(v string) *CommonApplyQueryResponseBodyModule {
	s.ExtendValue = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetGmtCreate(v string) *CommonApplyQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetStatus(v int32) *CommonApplyQueryResponseBodyModule {
	s.Status = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetThirdpartCorpId(v string) *CommonApplyQueryResponseBodyModule {
	s.ThirdpartCorpId = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetThirdpartId(v string) *CommonApplyQueryResponseBodyModule {
	s.ThirdpartId = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetTripCause(v string) *CommonApplyQueryResponseBodyModule {
	s.TripCause = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetUserId(v string) *CommonApplyQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) SetUserName(v string) *CommonApplyQueryResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *CommonApplyQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
