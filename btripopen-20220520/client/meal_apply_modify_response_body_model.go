// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyModifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MealApplyModifyResponseBody
	GetCode() *int32
	SetMessage(v string) *MealApplyModifyResponseBody
	GetMessage() *string
	SetModule(v *MealApplyModifyResponseBodyModule) *MealApplyModifyResponseBody
	GetModule() *MealApplyModifyResponseBodyModule
	SetRequestId(v string) *MealApplyModifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MealApplyModifyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MealApplyModifyResponseBody
	GetTraceId() *string
}

type MealApplyModifyResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Module  *MealApplyModifyResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-****-****-****-9D1AB7218B27
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

func (s MealApplyModifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MealApplyModifyResponseBody) GoString() string {
	return s.String()
}

func (s *MealApplyModifyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MealApplyModifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MealApplyModifyResponseBody) GetModule() *MealApplyModifyResponseBodyModule {
	return s.Module
}

func (s *MealApplyModifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MealApplyModifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MealApplyModifyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MealApplyModifyResponseBody) SetCode(v int32) *MealApplyModifyResponseBody {
	s.Code = &v
	return s
}

func (s *MealApplyModifyResponseBody) SetMessage(v string) *MealApplyModifyResponseBody {
	s.Message = &v
	return s
}

func (s *MealApplyModifyResponseBody) SetModule(v *MealApplyModifyResponseBodyModule) *MealApplyModifyResponseBody {
	s.Module = v
	return s
}

func (s *MealApplyModifyResponseBody) SetRequestId(v string) *MealApplyModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *MealApplyModifyResponseBody) SetSuccess(v bool) *MealApplyModifyResponseBody {
	s.Success = &v
	return s
}

func (s *MealApplyModifyResponseBody) SetTraceId(v string) *MealApplyModifyResponseBody {
	s.TraceId = &v
	return s
}

func (s *MealApplyModifyResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MealApplyModifyResponseBodyModule struct {
	// example:
	//
	// 1234
	ThirdPartyApplyId *string `json:"third_party_apply_id,omitempty" xml:"third_party_apply_id,omitempty"`
}

func (s MealApplyModifyResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s MealApplyModifyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *MealApplyModifyResponseBodyModule) GetThirdPartyApplyId() *string {
	return s.ThirdPartyApplyId
}

func (s *MealApplyModifyResponseBodyModule) SetThirdPartyApplyId(v string) *MealApplyModifyResponseBodyModule {
	s.ThirdPartyApplyId = &v
	return s
}

func (s *MealApplyModifyResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
