// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardListQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TravelStandardListQueryResponseBody
	GetCode() *int32
	SetMessage(v string) *TravelStandardListQueryResponseBody
	GetMessage() *string
	SetModule(v *TravelStandardListQueryResponseBodyModule) *TravelStandardListQueryResponseBody
	GetModule() *TravelStandardListQueryResponseBodyModule
	SetRequestId(v string) *TravelStandardListQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TravelStandardListQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TravelStandardListQueryResponseBody
	GetTraceId() *string
}

type TravelStandardListQueryResponseBody struct {
	// example:
	//
	// 0
	Code    *int32                                     `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TravelStandardListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 05F9C201-1B53-5905-94AB-0D7444D8466A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041aa317070996148671005d0a0b
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TravelStandardListQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TravelStandardListQueryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TravelStandardListQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TravelStandardListQueryResponseBody) GetModule() *TravelStandardListQueryResponseBodyModule {
	return s.Module
}

func (s *TravelStandardListQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TravelStandardListQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TravelStandardListQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TravelStandardListQueryResponseBody) SetCode(v int32) *TravelStandardListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TravelStandardListQueryResponseBody) SetMessage(v string) *TravelStandardListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TravelStandardListQueryResponseBody) SetModule(v *TravelStandardListQueryResponseBodyModule) *TravelStandardListQueryResponseBody {
	s.Module = v
	return s
}

func (s *TravelStandardListQueryResponseBody) SetRequestId(v string) *TravelStandardListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TravelStandardListQueryResponseBody) SetSuccess(v bool) *TravelStandardListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TravelStandardListQueryResponseBody) SetTraceId(v string) *TravelStandardListQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TravelStandardListQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TravelStandardListQueryResponseBodyModule struct {
	Items []*TravelStandardListQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	TotalSize *int32 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s TravelStandardListQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TravelStandardListQueryResponseBodyModule) GetItems() []*TravelStandardListQueryResponseBodyModuleItems {
	return s.Items
}

func (s *TravelStandardListQueryResponseBodyModule) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *TravelStandardListQueryResponseBodyModule) SetItems(v []*TravelStandardListQueryResponseBodyModuleItems) *TravelStandardListQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *TravelStandardListQueryResponseBodyModule) SetTotalSize(v int32) *TravelStandardListQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *TravelStandardListQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TravelStandardListQueryResponseBodyModuleItems struct {
	MainReserveRule *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule   `json:"main_reserve_rule,omitempty" xml:"main_reserve_rule,omitempty" type:"Struct"`
	ReserveRuleDesc []*TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc `json:"reserve_rule_desc,omitempty" xml:"reserve_rule_desc,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Scope *int32 `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s TravelStandardListQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardListQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *TravelStandardListQueryResponseBodyModuleItems) GetMainReserveRule() *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule {
	return s.MainReserveRule
}

func (s *TravelStandardListQueryResponseBodyModuleItems) GetReserveRuleDesc() []*TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc {
	return s.ReserveRuleDesc
}

func (s *TravelStandardListQueryResponseBodyModuleItems) GetScope() *int32 {
	return s.Scope
}

func (s *TravelStandardListQueryResponseBodyModuleItems) SetMainReserveRule(v *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) *TravelStandardListQueryResponseBodyModuleItems {
	s.MainReserveRule = v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItems) SetReserveRuleDesc(v []*TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc) *TravelStandardListQueryResponseBodyModuleItems {
	s.ReserveRuleDesc = v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItems) SetScope(v int32) *TravelStandardListQueryResponseBodyModuleItems {
	s.Scope = &v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}

type TravelStandardListQueryResponseBodyModuleItemsMainReserveRule struct {
	OpenServiceTypeList []*string `json:"open_service_type_list,omitempty" xml:"open_service_type_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2006516571
	RuleCode *int64  `json:"rule_code,omitempty" xml:"rule_code,omitempty"`
	RuleDesc *string `json:"rule_desc,omitempty" xml:"rule_desc,omitempty"`
	// example:
	//
	// 6516571
	RuleId   *int64  `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
	RuleName *string `json:"rule_name,omitempty" xml:"rule_name,omitempty"`
}

func (s TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) GoString() string {
	return s.String()
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) GetOpenServiceTypeList() []*string {
	return s.OpenServiceTypeList
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) GetRuleCode() *int64 {
	return s.RuleCode
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) GetRuleName() *string {
	return s.RuleName
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) SetOpenServiceTypeList(v []*string) *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule {
	s.OpenServiceTypeList = v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) SetRuleCode(v int64) *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule {
	s.RuleCode = &v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) SetRuleDesc(v string) *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule {
	s.RuleDesc = &v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) SetRuleId(v int64) *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule {
	s.RuleId = &v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) SetRuleName(v string) *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule {
	s.RuleName = &v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItemsMainReserveRule) Validate() error {
	return dara.Validate(s)
}

type TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc struct {
	DataList []*TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
	Title    *string                                                                  `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// flight
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc) GoString() string {
	return s.String()
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc) GetDataList() []*TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList {
	return s.DataList
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc) GetTitle() *string {
	return s.Title
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc) GetType() *string {
	return s.Type
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc) SetDataList(v []*TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList) *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc {
	s.DataList = v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc) SetTitle(v string) *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc {
	s.Title = &v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc) SetType(v string) *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc {
	s.Type = &v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDesc) Validate() error {
	return dara.Validate(s)
}

type TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList) GoString() string {
	return s.String()
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList) GetKey() *string {
	return s.Key
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList) GetValue() *string {
	return s.Value
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList) SetKey(v string) *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList {
	s.Key = &v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList) SetValue(v string) *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList {
	s.Value = &v
	return s
}

func (s *TravelStandardListQueryResponseBodyModuleItemsReserveRuleDescDataList) Validate() error {
	return dara.Validate(s)
}
