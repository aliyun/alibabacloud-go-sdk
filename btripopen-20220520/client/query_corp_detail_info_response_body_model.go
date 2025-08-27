// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCorpDetailInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCorpDetailInfoResponseBody
	GetCode() *string
	SetMessage(v string) *QueryCorpDetailInfoResponseBody
	GetMessage() *string
	SetModule(v *QueryCorpDetailInfoResponseBodyModule) *QueryCorpDetailInfoResponseBody
	GetModule() *QueryCorpDetailInfoResponseBodyModule
	SetRequestId(v string) *QueryCorpDetailInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCorpDetailInfoResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *QueryCorpDetailInfoResponseBody
	GetTraceId() *string
}

type QueryCorpDetailInfoResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  *QueryCorpDetailInfoResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-****-****-****-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce********056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s QueryCorpDetailInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCorpDetailInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCorpDetailInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCorpDetailInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCorpDetailInfoResponseBody) GetModule() *QueryCorpDetailInfoResponseBodyModule {
	return s.Module
}

func (s *QueryCorpDetailInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCorpDetailInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCorpDetailInfoResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *QueryCorpDetailInfoResponseBody) SetCode(v string) *QueryCorpDetailInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBody) SetMessage(v string) *QueryCorpDetailInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBody) SetModule(v *QueryCorpDetailInfoResponseBodyModule) *QueryCorpDetailInfoResponseBody {
	s.Module = v
	return s
}

func (s *QueryCorpDetailInfoResponseBody) SetRequestId(v string) *QueryCorpDetailInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBody) SetSuccess(v bool) *QueryCorpDetailInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBody) SetTraceId(v string) *QueryCorpDetailInfoResponseBody {
	s.TraceId = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCorpDetailInfoResponseBodyModule struct {
	// example:
	//
	// btrip5txxxxxxx
	CorpId   *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// example:
	//
	// open12********012jq
	OpenAgentId    *string `json:"open_agent_id,omitempty" xml:"open_agent_id,omitempty"`
	SuperAdminName *string `json:"super_admin_name,omitempty" xml:"super_admin_name,omitempty"`
	// example:
	//
	// 138xxxx0001
	SuperAdminPhone *string `json:"super_admin_phone,omitempty" xml:"super_admin_phone,omitempty"`
	// example:
	//
	// user1234
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s QueryCorpDetailInfoResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryCorpDetailInfoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryCorpDetailInfoResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *QueryCorpDetailInfoResponseBodyModule) GetCorpName() *string {
	return s.CorpName
}

func (s *QueryCorpDetailInfoResponseBodyModule) GetOpenAgentId() *string {
	return s.OpenAgentId
}

func (s *QueryCorpDetailInfoResponseBodyModule) GetSuperAdminName() *string {
	return s.SuperAdminName
}

func (s *QueryCorpDetailInfoResponseBodyModule) GetSuperAdminPhone() *string {
	return s.SuperAdminPhone
}

func (s *QueryCorpDetailInfoResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *QueryCorpDetailInfoResponseBodyModule) SetCorpId(v string) *QueryCorpDetailInfoResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBodyModule) SetCorpName(v string) *QueryCorpDetailInfoResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBodyModule) SetOpenAgentId(v string) *QueryCorpDetailInfoResponseBodyModule {
	s.OpenAgentId = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBodyModule) SetSuperAdminName(v string) *QueryCorpDetailInfoResponseBodyModule {
	s.SuperAdminName = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBodyModule) SetSuperAdminPhone(v string) *QueryCorpDetailInfoResponseBodyModule {
	s.SuperAdminPhone = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBodyModule) SetUserId(v string) *QueryCorpDetailInfoResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *QueryCorpDetailInfoResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
