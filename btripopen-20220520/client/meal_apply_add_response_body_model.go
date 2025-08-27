// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyAddResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MealApplyAddResponseBody
	GetCode() *int32
	SetMessage(v string) *MealApplyAddResponseBody
	GetMessage() *string
	SetModule(v *MealApplyAddResponseBodyModule) *MealApplyAddResponseBody
	GetModule() *MealApplyAddResponseBodyModule
	SetRequestId(v string) *MealApplyAddResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MealApplyAddResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MealApplyAddResponseBody
	GetTraceId() *string
}

type MealApplyAddResponseBody struct {
	// example:
	//
	// 0
	Code    *int32                          `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                         `json:"message,omitempty" xml:"message,omitempty"`
	Module  *MealApplyAddResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210e842b16611337974412836dae27
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s MealApplyAddResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MealApplyAddResponseBody) GoString() string {
	return s.String()
}

func (s *MealApplyAddResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MealApplyAddResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MealApplyAddResponseBody) GetModule() *MealApplyAddResponseBodyModule {
	return s.Module
}

func (s *MealApplyAddResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MealApplyAddResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MealApplyAddResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MealApplyAddResponseBody) SetCode(v int32) *MealApplyAddResponseBody {
	s.Code = &v
	return s
}

func (s *MealApplyAddResponseBody) SetMessage(v string) *MealApplyAddResponseBody {
	s.Message = &v
	return s
}

func (s *MealApplyAddResponseBody) SetModule(v *MealApplyAddResponseBodyModule) *MealApplyAddResponseBody {
	s.Module = v
	return s
}

func (s *MealApplyAddResponseBody) SetRequestId(v string) *MealApplyAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *MealApplyAddResponseBody) SetSuccess(v bool) *MealApplyAddResponseBody {
	s.Success = &v
	return s
}

func (s *MealApplyAddResponseBody) SetTraceId(v string) *MealApplyAddResponseBody {
	s.TraceId = &v
	return s
}

func (s *MealApplyAddResponseBody) Validate() error {
	return dara.Validate(s)
}

type MealApplyAddResponseBodyModule struct {
	// example:
	//
	// 2024073128454753
	ThirdPartApplyId *string `json:"third_part_apply_id,omitempty" xml:"third_part_apply_id,omitempty"`
}

func (s MealApplyAddResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s MealApplyAddResponseBodyModule) GoString() string {
	return s.String()
}

func (s *MealApplyAddResponseBodyModule) GetThirdPartApplyId() *string {
	return s.ThirdPartApplyId
}

func (s *MealApplyAddResponseBodyModule) SetThirdPartApplyId(v string) *MealApplyAddResponseBodyModule {
	s.ThirdPartApplyId = &v
	return s
}

func (s *MealApplyAddResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
