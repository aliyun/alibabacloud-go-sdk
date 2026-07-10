// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAddResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyAddResponseBody
	GetCode() *string
	SetMessage(v string) *ApplyAddResponseBody
	GetMessage() *string
	SetModule(v *ApplyAddResponseBodyModule) *ApplyAddResponseBody
	GetModule() *ApplyAddResponseBodyModule
	SetRequestId(v string) *ApplyAddResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyAddResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ApplyAddResponseBody
	GetTraceId() *string
}

type ApplyAddResponseBody struct {
	Code      *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                     `json:"message,omitempty" xml:"message,omitempty"`
	Module    *ApplyAddResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                       `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                     `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ApplyAddResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyAddResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyAddResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyAddResponseBody) GetModule() *ApplyAddResponseBodyModule {
	return s.Module
}

func (s *ApplyAddResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyAddResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyAddResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ApplyAddResponseBody) SetCode(v string) *ApplyAddResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyAddResponseBody) SetMessage(v string) *ApplyAddResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyAddResponseBody) SetModule(v *ApplyAddResponseBodyModule) *ApplyAddResponseBody {
	s.Module = v
	return s
}

func (s *ApplyAddResponseBody) SetRequestId(v string) *ApplyAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyAddResponseBody) SetSuccess(v bool) *ApplyAddResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyAddResponseBody) SetTraceId(v string) *ApplyAddResponseBody {
	s.TraceId = &v
	return s
}

func (s *ApplyAddResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyAddResponseBodyModule struct {
	ApplyId             *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ThirdpartApplyId    *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
}

func (s ApplyAddResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ApplyAddResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ApplyAddResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *ApplyAddResponseBodyModule) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *ApplyAddResponseBodyModule) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *ApplyAddResponseBodyModule) SetApplyId(v int64) *ApplyAddResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *ApplyAddResponseBodyModule) SetThirdpartApplyId(v string) *ApplyAddResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *ApplyAddResponseBodyModule) SetThirdpartBusinessId(v string) *ApplyAddResponseBodyModule {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *ApplyAddResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
