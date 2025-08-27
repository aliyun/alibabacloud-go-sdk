// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CostCenterQueryResponseBody
	GetCode() *string
	SetMessage(v string) *CostCenterQueryResponseBody
	GetMessage() *string
	SetModule(v []*CostCenterQueryResponseBodyModule) *CostCenterQueryResponseBody
	GetModule() []*CostCenterQueryResponseBodyModule
	SetMorePage(v bool) *CostCenterQueryResponseBody
	GetMorePage() *bool
	SetRequestId(v string) *CostCenterQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CostCenterQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CostCenterQueryResponseBody
	GetTraceId() *string
}

type CostCenterQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  []*CostCenterQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	// example:
	//
	// true
	MorePage *bool `json:"more_page,omitempty" xml:"more_page,omitempty"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
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

func (s CostCenterQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CostCenterQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CostCenterQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CostCenterQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CostCenterQueryResponseBody) GetModule() []*CostCenterQueryResponseBodyModule {
	return s.Module
}

func (s *CostCenterQueryResponseBody) GetMorePage() *bool {
	return s.MorePage
}

func (s *CostCenterQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CostCenterQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CostCenterQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CostCenterQueryResponseBody) SetCode(v string) *CostCenterQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CostCenterQueryResponseBody) SetMessage(v string) *CostCenterQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CostCenterQueryResponseBody) SetModule(v []*CostCenterQueryResponseBodyModule) *CostCenterQueryResponseBody {
	s.Module = v
	return s
}

func (s *CostCenterQueryResponseBody) SetMorePage(v bool) *CostCenterQueryResponseBody {
	s.MorePage = &v
	return s
}

func (s *CostCenterQueryResponseBody) SetRequestId(v string) *CostCenterQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CostCenterQueryResponseBody) SetSuccess(v bool) *CostCenterQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CostCenterQueryResponseBody) SetTraceId(v string) *CostCenterQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *CostCenterQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CostCenterQueryResponseBodyModule struct {
	// example:
	//
	// a@alipay.com
	AlipayNo *string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
	// example:
	//
	// ding12345678
	CorpId    *string                                       `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Disable   *int64                                        `json:"disable,omitempty" xml:"disable,omitempty"`
	EntityDOS []*CostCenterQueryResponseBodyModuleEntityDOS `json:"entity_d_o_s,omitempty" xml:"entity_d_o_s,omitempty" type:"Repeated"`
	// example:
	//
	// 7232
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 123456
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
	// rule code
	//
	// example:
	//
	// 500578154
	RuleCode *int64 `json:"rule_code,omitempty" xml:"rule_code,omitempty"`
	// example:
	//
	// 1
	Scope *int64 `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// 1
	ThirdpartId *string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CostCenterQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CostCenterQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CostCenterQueryResponseBodyModule) GetAlipayNo() *string {
	return s.AlipayNo
}

func (s *CostCenterQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *CostCenterQueryResponseBodyModule) GetDisable() *int64 {
	return s.Disable
}

func (s *CostCenterQueryResponseBodyModule) GetEntityDOS() []*CostCenterQueryResponseBodyModuleEntityDOS {
	return s.EntityDOS
}

func (s *CostCenterQueryResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *CostCenterQueryResponseBodyModule) GetNumber() *string {
	return s.Number
}

func (s *CostCenterQueryResponseBodyModule) GetRuleCode() *int64 {
	return s.RuleCode
}

func (s *CostCenterQueryResponseBodyModule) GetScope() *int64 {
	return s.Scope
}

func (s *CostCenterQueryResponseBodyModule) GetThirdpartId() *string {
	return s.ThirdpartId
}

func (s *CostCenterQueryResponseBodyModule) GetTitle() *string {
	return s.Title
}

func (s *CostCenterQueryResponseBodyModule) SetAlipayNo(v string) *CostCenterQueryResponseBodyModule {
	s.AlipayNo = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetCorpId(v string) *CostCenterQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetDisable(v int64) *CostCenterQueryResponseBodyModule {
	s.Disable = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetEntityDOS(v []*CostCenterQueryResponseBodyModuleEntityDOS) *CostCenterQueryResponseBodyModule {
	s.EntityDOS = v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetId(v int64) *CostCenterQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetNumber(v string) *CostCenterQueryResponseBodyModule {
	s.Number = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetRuleCode(v int64) *CostCenterQueryResponseBodyModule {
	s.RuleCode = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetScope(v int64) *CostCenterQueryResponseBodyModule {
	s.Scope = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetThirdpartId(v string) *CostCenterQueryResponseBodyModule {
	s.ThirdpartId = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) SetTitle(v string) *CostCenterQueryResponseBodyModule {
	s.Title = &v
	return s
}

func (s *CostCenterQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type CostCenterQueryResponseBodyModuleEntityDOS struct {
	// example:
	//
	// ding1234567
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// example:
	//
	// 12345
	EntityId *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	// example:
	//
	// 1
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
	// example:
	//
	// default_bus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 10
	UserNum *int32 `json:"user_num,omitempty" xml:"user_num,omitempty"`
}

func (s CostCenterQueryResponseBodyModuleEntityDOS) String() string {
	return dara.Prettify(s)
}

func (s CostCenterQueryResponseBodyModuleEntityDOS) GoString() string {
	return s.String()
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) GetCorpId() *string {
	return s.CorpId
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) GetEntityId() *string {
	return s.EntityId
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) GetEntityType() *string {
	return s.EntityType
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) GetName() *string {
	return s.Name
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) GetUserNum() *int32 {
	return s.UserNum
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) SetCorpId(v string) *CostCenterQueryResponseBodyModuleEntityDOS {
	s.CorpId = &v
	return s
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) SetEntityId(v string) *CostCenterQueryResponseBodyModuleEntityDOS {
	s.EntityId = &v
	return s
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) SetEntityType(v string) *CostCenterQueryResponseBodyModuleEntityDOS {
	s.EntityType = &v
	return s
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) SetName(v string) *CostCenterQueryResponseBodyModuleEntityDOS {
	s.Name = &v
	return s
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) SetUserNum(v int32) *CostCenterQueryResponseBodyModuleEntityDOS {
	s.UserNum = &v
	return s
}

func (s *CostCenterQueryResponseBodyModuleEntityDOS) Validate() error {
	return dara.Validate(s)
}
