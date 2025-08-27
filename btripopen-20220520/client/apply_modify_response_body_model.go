// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyModifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyModifyResponseBody
	GetCode() *string
	SetMessage(v string) *ApplyModifyResponseBody
	GetMessage() *string
	SetModule(v *ApplyModifyResponseBodyModule) *ApplyModifyResponseBody
	GetModule() *ApplyModifyResponseBodyModule
	SetRequestId(v string) *ApplyModifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyModifyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ApplyModifyResponseBody
	GetTraceId() *string
}

type ApplyModifyResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module。
	Module *ApplyModifyResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
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

func (s ApplyModifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyModifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyModifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyModifyResponseBody) GetModule() *ApplyModifyResponseBodyModule {
	return s.Module
}

func (s *ApplyModifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyModifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyModifyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ApplyModifyResponseBody) SetCode(v string) *ApplyModifyResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyModifyResponseBody) SetMessage(v string) *ApplyModifyResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyModifyResponseBody) SetModule(v *ApplyModifyResponseBodyModule) *ApplyModifyResponseBody {
	s.Module = v
	return s
}

func (s *ApplyModifyResponseBody) SetRequestId(v string) *ApplyModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyModifyResponseBody) SetSuccess(v bool) *ApplyModifyResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyModifyResponseBody) SetTraceId(v string) *ApplyModifyResponseBody {
	s.TraceId = &v
	return s
}

func (s *ApplyModifyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ApplyModifyResponseBodyModule struct {
	// example:
	//
	// 118526587
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// thirdpart12132
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// 20220702001
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
}

func (s ApplyModifyResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ApplyModifyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ApplyModifyResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *ApplyModifyResponseBodyModule) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *ApplyModifyResponseBodyModule) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *ApplyModifyResponseBodyModule) SetApplyId(v int64) *ApplyModifyResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *ApplyModifyResponseBodyModule) SetThirdpartApplyId(v string) *ApplyModifyResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyModifyResponseBodyModule) SetThirdpartBusinessId(v string) *ApplyModifyResponseBodyModule {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyModifyResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
