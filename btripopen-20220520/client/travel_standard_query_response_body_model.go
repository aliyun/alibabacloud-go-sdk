// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TravelStandardQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TravelStandardQueryResponseBody
	GetMessage() *string
	SetModule(v *TravelStandardQueryResponseBodyModule) *TravelStandardQueryResponseBody
	GetModule() *TravelStandardQueryResponseBodyModule
	SetRequestId(v string) *TravelStandardQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TravelStandardQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TravelStandardQueryResponseBody
	GetTraceId() *string
}

type TravelStandardQueryResponseBody struct {
	// example:
	//
	// 0
	Code    *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TravelStandardQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 210bcc3a16583004579056128d33d7
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TravelStandardQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TravelStandardQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TravelStandardQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TravelStandardQueryResponseBody) GetModule() *TravelStandardQueryResponseBodyModule {
	return s.Module
}

func (s *TravelStandardQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TravelStandardQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TravelStandardQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TravelStandardQueryResponseBody) SetCode(v string) *TravelStandardQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TravelStandardQueryResponseBody) SetMessage(v string) *TravelStandardQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TravelStandardQueryResponseBody) SetModule(v *TravelStandardQueryResponseBodyModule) *TravelStandardQueryResponseBody {
	s.Module = v
	return s
}

func (s *TravelStandardQueryResponseBody) SetRequestId(v string) *TravelStandardQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TravelStandardQueryResponseBody) SetSuccess(v bool) *TravelStandardQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TravelStandardQueryResponseBody) SetTraceId(v string) *TravelStandardQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TravelStandardQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TravelStandardQueryResponseBodyModule struct {
	ActivatedServiceTypeList []*string                                         `json:"activated_service_type_list,omitempty" xml:"activated_service_type_list,omitempty" type:"Repeated"`
	ReserveRule              *TravelStandardQueryResponseBodyModuleReserveRule `json:"reserve_rule,omitempty" xml:"reserve_rule,omitempty" type:"Struct"`
}

func (s TravelStandardQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TravelStandardQueryResponseBodyModule) GetActivatedServiceTypeList() []*string {
	return s.ActivatedServiceTypeList
}

func (s *TravelStandardQueryResponseBodyModule) GetReserveRule() *TravelStandardQueryResponseBodyModuleReserveRule {
	return s.ReserveRule
}

func (s *TravelStandardQueryResponseBodyModule) SetActivatedServiceTypeList(v []*string) *TravelStandardQueryResponseBodyModule {
	s.ActivatedServiceTypeList = v
	return s
}

func (s *TravelStandardQueryResponseBodyModule) SetReserveRule(v *TravelStandardQueryResponseBodyModuleReserveRule) *TravelStandardQueryResponseBodyModule {
	s.ReserveRule = v
	return s
}

func (s *TravelStandardQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TravelStandardQueryResponseBodyModuleReserveRule struct {
	MainReserveRule  *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule    `json:"main_reserve_rule,omitempty" xml:"main_reserve_rule,omitempty" type:"Struct"`
	ModuleConfigList []*TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList `json:"module_config_list,omitempty" xml:"module_config_list,omitempty" type:"Repeated"`
}

func (s TravelStandardQueryResponseBodyModuleReserveRule) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardQueryResponseBodyModuleReserveRule) GoString() string {
	return s.String()
}

func (s *TravelStandardQueryResponseBodyModuleReserveRule) GetMainReserveRule() *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule {
	return s.MainReserveRule
}

func (s *TravelStandardQueryResponseBodyModuleReserveRule) GetModuleConfigList() []*TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList {
	return s.ModuleConfigList
}

func (s *TravelStandardQueryResponseBodyModuleReserveRule) SetMainReserveRule(v *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) *TravelStandardQueryResponseBodyModuleReserveRule {
	s.MainReserveRule = v
	return s
}

func (s *TravelStandardQueryResponseBodyModuleReserveRule) SetModuleConfigList(v []*TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList) *TravelStandardQueryResponseBodyModuleReserveRule {
	s.ModuleConfigList = v
	return s
}

func (s *TravelStandardQueryResponseBodyModuleReserveRule) Validate() error {
	return dara.Validate(s)
}

type TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule struct {
	OpenServiceTypeList []*string `json:"open_service_type_list,omitempty" xml:"open_service_type_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2006517149
	RuleCode *int64  `json:"rule_code,omitempty" xml:"rule_code,omitempty"`
	RuleDesc *string `json:"rule_desc,omitempty" xml:"rule_desc,omitempty"`
	// example:
	//
	// 6517149
	RuleId   *int64  `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
	RuleName *string `json:"rule_name,omitempty" xml:"rule_name,omitempty"`
}

func (s TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) GoString() string {
	return s.String()
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) GetOpenServiceTypeList() []*string {
	return s.OpenServiceTypeList
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) GetRuleCode() *int64 {
	return s.RuleCode
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) GetRuleName() *string {
	return s.RuleName
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) SetOpenServiceTypeList(v []*string) *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule {
	s.OpenServiceTypeList = v
	return s
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) SetRuleCode(v int64) *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule {
	s.RuleCode = &v
	return s
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) SetRuleDesc(v string) *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule {
	s.RuleDesc = &v
	return s
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) SetRuleId(v int64) *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule {
	s.RuleId = &v
	return s
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) SetRuleName(v string) *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule {
	s.RuleName = &v
	return s
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleMainReserveRule) Validate() error {
	return dara.Validate(s)
}

type TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList struct {
	// example:
	//
	// FLIGHT_ADVANCE_BUY_DAY
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 3
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList) GoString() string {
	return s.String()
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList) GetCode() *string {
	return s.Code
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList) GetValue() *string {
	return s.Value
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList) SetCode(v string) *TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList {
	s.Code = &v
	return s
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList) SetValue(v string) *TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList {
	s.Value = &v
	return s
}

func (s *TravelStandardQueryResponseBodyModuleReserveRuleModuleConfigList) Validate() error {
	return dara.Validate(s)
}
