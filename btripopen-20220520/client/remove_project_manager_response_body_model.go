// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectManagerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveProjectManagerResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemoveProjectManagerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveProjectManagerResponseBody
	GetMessage() *string
	SetModule(v *RemoveProjectManagerResponseBodyModule) *RemoveProjectManagerResponseBody
	GetModule() *RemoveProjectManagerResponseBodyModule
	SetRequestId(v string) *RemoveProjectManagerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveProjectManagerResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *RemoveProjectManagerResponseBody
	GetTraceId() *string
}

type RemoveProjectManagerResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Module  *RemoveProjectManagerResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s RemoveProjectManagerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectManagerResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveProjectManagerResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveProjectManagerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveProjectManagerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveProjectManagerResponseBody) GetModule() *RemoveProjectManagerResponseBodyModule {
	return s.Module
}

func (s *RemoveProjectManagerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveProjectManagerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveProjectManagerResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *RemoveProjectManagerResponseBody) SetCode(v string) *RemoveProjectManagerResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveProjectManagerResponseBody) SetHttpStatusCode(v int32) *RemoveProjectManagerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveProjectManagerResponseBody) SetMessage(v string) *RemoveProjectManagerResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveProjectManagerResponseBody) SetModule(v *RemoveProjectManagerResponseBodyModule) *RemoveProjectManagerResponseBody {
	s.Module = v
	return s
}

func (s *RemoveProjectManagerResponseBody) SetRequestId(v string) *RemoveProjectManagerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveProjectManagerResponseBody) SetSuccess(v bool) *RemoveProjectManagerResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveProjectManagerResponseBody) SetTraceId(v string) *RemoveProjectManagerResponseBody {
	s.TraceId = &v
	return s
}

func (s *RemoveProjectManagerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveProjectManagerResponseBodyModule struct {
	// example:
	//
	// 1
	AddNum *int32 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	// example:
	//
	// 1
	RemoveNum *int32 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
	// example:
	//
	// 500578154
	RuleCode *int64 `json:"rule_code,omitempty" xml:"rule_code,omitempty"`
}

func (s RemoveProjectManagerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectManagerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *RemoveProjectManagerResponseBodyModule) GetAddNum() *int32 {
	return s.AddNum
}

func (s *RemoveProjectManagerResponseBodyModule) GetRemoveNum() *int32 {
	return s.RemoveNum
}

func (s *RemoveProjectManagerResponseBodyModule) GetRuleCode() *int64 {
	return s.RuleCode
}

func (s *RemoveProjectManagerResponseBodyModule) SetAddNum(v int32) *RemoveProjectManagerResponseBodyModule {
	s.AddNum = &v
	return s
}

func (s *RemoveProjectManagerResponseBodyModule) SetRemoveNum(v int32) *RemoveProjectManagerResponseBodyModule {
	s.RemoveNum = &v
	return s
}

func (s *RemoveProjectManagerResponseBodyModule) SetRuleCode(v int64) *RemoveProjectManagerResponseBodyModule {
	s.RuleCode = &v
	return s
}

func (s *RemoveProjectManagerResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
