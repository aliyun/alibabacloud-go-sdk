// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectManagerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddProjectManagerResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddProjectManagerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddProjectManagerResponseBody
	GetMessage() *string
	SetModule(v *AddProjectManagerResponseBodyModule) *AddProjectManagerResponseBody
	GetModule() *AddProjectManagerResponseBodyModule
	SetRequestId(v string) *AddProjectManagerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddProjectManagerResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AddProjectManagerResponseBody
	GetTraceId() *string
}

type AddProjectManagerResponseBody struct {
	Code           *string                              `json:"code,omitempty" xml:"code,omitempty"`
	HttpStatusCode *int32                               `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module         *AddProjectManagerResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId      *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                                `json:"success,omitempty" xml:"success,omitempty"`
	TraceId        *string                              `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AddProjectManagerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddProjectManagerResponseBody) GoString() string {
	return s.String()
}

func (s *AddProjectManagerResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddProjectManagerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddProjectManagerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddProjectManagerResponseBody) GetModule() *AddProjectManagerResponseBodyModule {
	return s.Module
}

func (s *AddProjectManagerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddProjectManagerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddProjectManagerResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AddProjectManagerResponseBody) SetCode(v string) *AddProjectManagerResponseBody {
	s.Code = &v
	return s
}

func (s *AddProjectManagerResponseBody) SetHttpStatusCode(v int32) *AddProjectManagerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddProjectManagerResponseBody) SetMessage(v string) *AddProjectManagerResponseBody {
	s.Message = &v
	return s
}

func (s *AddProjectManagerResponseBody) SetModule(v *AddProjectManagerResponseBodyModule) *AddProjectManagerResponseBody {
	s.Module = v
	return s
}

func (s *AddProjectManagerResponseBody) SetRequestId(v string) *AddProjectManagerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProjectManagerResponseBody) SetSuccess(v bool) *AddProjectManagerResponseBody {
	s.Success = &v
	return s
}

func (s *AddProjectManagerResponseBody) SetTraceId(v string) *AddProjectManagerResponseBody {
	s.TraceId = &v
	return s
}

func (s *AddProjectManagerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddProjectManagerResponseBodyModule struct {
	AddNum    *int32 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	RemoveNum *int32 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
	RuleCode  *int64 `json:"rule_code,omitempty" xml:"rule_code,omitempty"`
}

func (s AddProjectManagerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s AddProjectManagerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *AddProjectManagerResponseBodyModule) GetAddNum() *int32 {
	return s.AddNum
}

func (s *AddProjectManagerResponseBodyModule) GetRemoveNum() *int32 {
	return s.RemoveNum
}

func (s *AddProjectManagerResponseBodyModule) GetRuleCode() *int64 {
	return s.RuleCode
}

func (s *AddProjectManagerResponseBodyModule) SetAddNum(v int32) *AddProjectManagerResponseBodyModule {
	s.AddNum = &v
	return s
}

func (s *AddProjectManagerResponseBodyModule) SetRemoveNum(v int32) *AddProjectManagerResponseBodyModule {
	s.RemoveNum = &v
	return s
}

func (s *AddProjectManagerResponseBodyModule) SetRuleCode(v int64) *AddProjectManagerResponseBodyModule {
	s.RuleCode = &v
	return s
}

func (s *AddProjectManagerResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
