// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelCorpCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChannelCorpCreateResponseBody
	GetCode() *string
	SetMessage(v string) *ChannelCorpCreateResponseBody
	GetMessage() *string
	SetModule(v *ChannelCorpCreateResponseBodyModule) *ChannelCorpCreateResponseBody
	GetModule() *ChannelCorpCreateResponseBodyModule
	SetRequestId(v string) *ChannelCorpCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChannelCorpCreateResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ChannelCorpCreateResponseBody
	GetTraceId() *string
}

type ChannelCorpCreateResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  *ChannelCorpCreateResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s ChannelCorpCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChannelCorpCreateResponseBody) GoString() string {
	return s.String()
}

func (s *ChannelCorpCreateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChannelCorpCreateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChannelCorpCreateResponseBody) GetModule() *ChannelCorpCreateResponseBodyModule {
	return s.Module
}

func (s *ChannelCorpCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChannelCorpCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChannelCorpCreateResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ChannelCorpCreateResponseBody) SetCode(v string) *ChannelCorpCreateResponseBody {
	s.Code = &v
	return s
}

func (s *ChannelCorpCreateResponseBody) SetMessage(v string) *ChannelCorpCreateResponseBody {
	s.Message = &v
	return s
}

func (s *ChannelCorpCreateResponseBody) SetModule(v *ChannelCorpCreateResponseBodyModule) *ChannelCorpCreateResponseBody {
	s.Module = v
	return s
}

func (s *ChannelCorpCreateResponseBody) SetRequestId(v string) *ChannelCorpCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChannelCorpCreateResponseBody) SetSuccess(v bool) *ChannelCorpCreateResponseBody {
	s.Success = &v
	return s
}

func (s *ChannelCorpCreateResponseBody) SetTraceId(v string) *ChannelCorpCreateResponseBody {
	s.TraceId = &v
	return s
}

func (s *ChannelCorpCreateResponseBody) Validate() error {
	return dara.Validate(s)
}

type ChannelCorpCreateResponseBodyModule struct {
	AdministratorName *string `json:"administrator_name,omitempty" xml:"administrator_name,omitempty"`
	// example:
	//
	// 18378889782
	AdministratorPhone *string `json:"administrator_phone,omitempty" xml:"administrator_phone,omitempty"`
	// example:
	//
	// dingaa15ca45cba9ee744a5
	CorpId   *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// example:
	//
	// open12mplkebufu8701012jq
	OpenAgentId *string `json:"open_agent_id,omitempty" xml:"open_agent_id,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s ChannelCorpCreateResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ChannelCorpCreateResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ChannelCorpCreateResponseBodyModule) GetAdministratorName() *string {
	return s.AdministratorName
}

func (s *ChannelCorpCreateResponseBodyModule) GetAdministratorPhone() *string {
	return s.AdministratorPhone
}

func (s *ChannelCorpCreateResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *ChannelCorpCreateResponseBodyModule) GetCorpName() *string {
	return s.CorpName
}

func (s *ChannelCorpCreateResponseBodyModule) GetOpenAgentId() *string {
	return s.OpenAgentId
}

func (s *ChannelCorpCreateResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *ChannelCorpCreateResponseBodyModule) SetAdministratorName(v string) *ChannelCorpCreateResponseBodyModule {
	s.AdministratorName = &v
	return s
}

func (s *ChannelCorpCreateResponseBodyModule) SetAdministratorPhone(v string) *ChannelCorpCreateResponseBodyModule {
	s.AdministratorPhone = &v
	return s
}

func (s *ChannelCorpCreateResponseBodyModule) SetCorpId(v string) *ChannelCorpCreateResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *ChannelCorpCreateResponseBodyModule) SetCorpName(v string) *ChannelCorpCreateResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *ChannelCorpCreateResponseBodyModule) SetOpenAgentId(v string) *ChannelCorpCreateResponseBodyModule {
	s.OpenAgentId = &v
	return s
}

func (s *ChannelCorpCreateResponseBodyModule) SetUserId(v string) *ChannelCorpCreateResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *ChannelCorpCreateResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
