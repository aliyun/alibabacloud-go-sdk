// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCorpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSubCorpResponseBody
	GetCode() *string
	SetMessage(v string) *CreateSubCorpResponseBody
	GetMessage() *string
	SetModule(v *CreateSubCorpResponseBodyModule) *CreateSubCorpResponseBody
	GetModule() *CreateSubCorpResponseBodyModule
	SetRequestId(v string) *CreateSubCorpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSubCorpResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateSubCorpResponseBody
	GetTraceId() *string
}

type CreateSubCorpResponseBody struct {
	Code      *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                          `json:"message,omitempty" xml:"message,omitempty"`
	Module    *CreateSubCorpResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                            `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                          `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CreateSubCorpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCorpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubCorpResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSubCorpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSubCorpResponseBody) GetModule() *CreateSubCorpResponseBodyModule {
	return s.Module
}

func (s *CreateSubCorpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSubCorpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSubCorpResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateSubCorpResponseBody) SetCode(v string) *CreateSubCorpResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSubCorpResponseBody) SetMessage(v string) *CreateSubCorpResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSubCorpResponseBody) SetModule(v *CreateSubCorpResponseBodyModule) *CreateSubCorpResponseBody {
	s.Module = v
	return s
}

func (s *CreateSubCorpResponseBody) SetRequestId(v string) *CreateSubCorpResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubCorpResponseBody) SetSuccess(v bool) *CreateSubCorpResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSubCorpResponseBody) SetTraceId(v string) *CreateSubCorpResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateSubCorpResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSubCorpResponseBodyModule struct {
	SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
}

func (s CreateSubCorpResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCorpResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateSubCorpResponseBodyModule) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *CreateSubCorpResponseBodyModule) SetSubCorpId(v string) *CreateSubCorpResponseBodyModule {
	s.SubCorpId = &v
	return s
}

func (s *CreateSubCorpResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
