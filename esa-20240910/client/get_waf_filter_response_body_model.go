// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *GetWafFilterResponseBodyFilter) *GetWafFilterResponseBody
	GetFilter() *GetWafFilterResponseBodyFilter
	SetRequestId(v string) *GetWafFilterResponseBody
	GetRequestId() *string
}

type GetWafFilterResponseBody struct {
	// The matching engine configuration information returned.
	Filter *GetWafFilterResponseBodyFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWafFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponseBody) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBody) GetFilter() *GetWafFilterResponseBodyFilter {
	return s.Filter
}

func (s *GetWafFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWafFilterResponseBody) SetFilter(v *GetWafFilterResponseBodyFilter) *GetWafFilterResponseBody {
	s.Filter = v
	return s
}

func (s *GetWafFilterResponseBody) SetRequestId(v string) *GetWafFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWafFilterResponseBody) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWafFilterResponseBodyFilter struct {
	// The list that describes match objects and their properties.
	Fields []*GetWafFilterResponseBodyFilterFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// The phase in which WAF processes the request.
	//
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The target value of the matching engine.
	//
	// example:
	//
	// characteristics
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The rule type.
	//
	// example:
	//
	// http_custom_cc
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetWafFilterResponseBodyFilter) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponseBodyFilter) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilter) GetFields() []*GetWafFilterResponseBodyFilterFields {
	return s.Fields
}

func (s *GetWafFilterResponseBodyFilter) GetPhase() *string {
	return s.Phase
}

func (s *GetWafFilterResponseBodyFilter) GetTarget() *string {
	return s.Target
}

func (s *GetWafFilterResponseBodyFilter) GetType() *string {
	return s.Type
}

func (s *GetWafFilterResponseBodyFilter) SetFields(v []*GetWafFilterResponseBodyFilterFields) *GetWafFilterResponseBodyFilter {
	s.Fields = v
	return s
}

func (s *GetWafFilterResponseBodyFilter) SetPhase(v string) *GetWafFilterResponseBodyFilter {
	s.Phase = &v
	return s
}

func (s *GetWafFilterResponseBodyFilter) SetTarget(v string) *GetWafFilterResponseBodyFilter {
	s.Target = &v
	return s
}

func (s *GetWafFilterResponseBodyFilter) SetType(v string) *GetWafFilterResponseBodyFilter {
	s.Type = &v
	return s
}

func (s *GetWafFilterResponseBodyFilter) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWafFilterResponseBodyFilterFields struct {
	// Indicates whether the current plan supports this match object.
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The parameter of the match object used internally by the system.
	//
	// example:
	//
	// http.request.headers
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The display label of the match object.
	//
	// example:
	//
	// Header
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The list of logical operator properties that define the logical conditions used for matching.
	Logics []*GetWafFilterResponseBodyFilterFieldsLogics `json:"Logics,omitempty" xml:"Logics,omitempty" type:"Repeated"`
	// The minimum plan that supports this match object, displayed when the current plan does not support it.
	//
	// example:
	//
	// high
	MinPlan *string `json:"MinPlan,omitempty" xml:"MinPlan,omitempty"`
	// The selector object that defines how to select the match object.
	Selector *GetWafFilterResponseBodyFilterFieldsSelector `json:"Selector,omitempty" xml:"Selector,omitempty" type:"Struct"`
	// Indicates whether the match object contains subfields.
	//
	// example:
	//
	// true
	Sub *bool `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// The hint provided to users about how to enter subfields.
	//
	// example:
	//
	// e.g. Content-Type
	SubTip *string `json:"SubTip,omitempty" xml:"SubTip,omitempty"`
	// The enumerated sub-item list (dropdown subfields for grouped fields such as ali.websdk). Top-level match objects populate this list. Sub-items that are flat fields can be used directly as the left-hand side of an expression.
	Subs []*GetWafFilterResponseBodyFilterFieldsSubs `json:"Subs,omitempty" xml:"Subs,omitempty" type:"Repeated"`
}

func (s GetWafFilterResponseBodyFilterFields) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFields) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFields) GetEnable() *bool {
	return s.Enable
}

func (s *GetWafFilterResponseBodyFilterFields) GetKey() *string {
	return s.Key
}

func (s *GetWafFilterResponseBodyFilterFields) GetLabel() *string {
	return s.Label
}

func (s *GetWafFilterResponseBodyFilterFields) GetLogics() []*GetWafFilterResponseBodyFilterFieldsLogics {
	return s.Logics
}

func (s *GetWafFilterResponseBodyFilterFields) GetMinPlan() *string {
	return s.MinPlan
}

func (s *GetWafFilterResponseBodyFilterFields) GetSelector() *GetWafFilterResponseBodyFilterFieldsSelector {
	return s.Selector
}

func (s *GetWafFilterResponseBodyFilterFields) GetSub() *bool {
	return s.Sub
}

func (s *GetWafFilterResponseBodyFilterFields) GetSubTip() *string {
	return s.SubTip
}

func (s *GetWafFilterResponseBodyFilterFields) GetSubs() []*GetWafFilterResponseBodyFilterFieldsSubs {
	return s.Subs
}

func (s *GetWafFilterResponseBodyFilterFields) SetEnable(v bool) *GetWafFilterResponseBodyFilterFields {
	s.Enable = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetKey(v string) *GetWafFilterResponseBodyFilterFields {
	s.Key = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetLabel(v string) *GetWafFilterResponseBodyFilterFields {
	s.Label = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetLogics(v []*GetWafFilterResponseBodyFilterFieldsLogics) *GetWafFilterResponseBodyFilterFields {
	s.Logics = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetMinPlan(v string) *GetWafFilterResponseBodyFilterFields {
	s.MinPlan = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetSelector(v *GetWafFilterResponseBodyFilterFieldsSelector) *GetWafFilterResponseBodyFilterFields {
	s.Selector = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetSub(v bool) *GetWafFilterResponseBodyFilterFields {
	s.Sub = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetSubTip(v string) *GetWafFilterResponseBodyFilterFields {
	s.SubTip = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) SetSubs(v []*GetWafFilterResponseBodyFilterFieldsSubs) *GetWafFilterResponseBodyFilterFields {
	s.Subs = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFields) Validate() error {
	if s.Logics != nil {
		for _, item := range s.Logics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Selector != nil {
		if err := s.Selector.Validate(); err != nil {
			return err
		}
	}
	if s.Subs != nil {
		for _, item := range s.Subs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWafFilterResponseBodyFilterFieldsLogics struct {
	// The configurable attributes, such as whether the match is case-sensitive.
	//
	// example:
	//
	// 1
	Attributes *int32 `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// Indicates whether the current plan supports this match operator.
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the value input field. Valid values:
	//
	// 	- select:single: single-select input field
	//
	// 	- select:multi: multi-select input field
	//
	// 	- input:single: single input field
	//
	// 	- input:multi: multi input field
	//
	// example:
	//
	// input:single
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The minimum plan that supports this match operator, displayed when the current plan does not support it.
	//
	// example:
	//
	// high
	MinPlan *string `json:"MinPlan,omitempty" xml:"MinPlan,omitempty"`
	// Indicates whether the match result is negated.
	Negative *bool `json:"Negative,omitempty" xml:"Negative,omitempty"`
	// The display label of the match operator.
	//
	// example:
	//
	// Does not equal
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The parameter of the match operator used internally by the system.
	//
	// example:
	//
	// eq
	Symbol *string `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	// The input hint that helps users provide valid values required by the rule.
	//
	// example:
	//
	// e.g. image/jpeg
	Tip *string `json:"Tip,omitempty" xml:"Tip,omitempty"`
	// The type of the value. Valid values:
	//
	// 	- integer: integer
	//
	// 	- integer_slice: integer array
	//
	// 	- string: string
	//
	// 	- string_slice: string array
	//
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The validator object that defines the validation rules for values.
	Validator *GetWafFilterResponseBodyFilterFieldsLogicsValidator `json:"Validator,omitempty" xml:"Validator,omitempty" type:"Struct"`
}

func (s GetWafFilterResponseBodyFilterFieldsLogics) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsLogics) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) GetAttributes() *int32 {
	return s.Attributes
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) GetEnable() *bool {
	return s.Enable
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) GetKind() *string {
	return s.Kind
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) GetMinPlan() *string {
	return s.MinPlan
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) GetNegative() *bool {
	return s.Negative
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) GetOperator() *string {
	return s.Operator
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) GetSymbol() *string {
	return s.Symbol
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) GetTip() *string {
	return s.Tip
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) GetType() *string {
	return s.Type
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) GetValidator() *GetWafFilterResponseBodyFilterFieldsLogicsValidator {
	return s.Validator
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetAttributes(v int32) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Attributes = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetEnable(v bool) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Enable = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetKind(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Kind = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetMinPlan(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.MinPlan = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetNegative(v bool) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Negative = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetOperator(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Operator = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetSymbol(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Symbol = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetTip(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Tip = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetType(v string) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Type = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) SetValidator(v *GetWafFilterResponseBodyFilterFieldsLogicsValidator) *GetWafFilterResponseBodyFilterFieldsLogics {
	s.Validator = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogics) Validate() error {
	if s.Validator != nil {
		if err := s.Validator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWafFilterResponseBodyFilterFieldsLogicsValidator struct {
	// The error message returned when validation fails.
	//
	// example:
	//
	// Enter a valid expression
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The length limit of the value.
	Length *WafQuotaInteger `json:"Length,omitempty" xml:"Length,omitempty"`
	// The regular expression pattern for the value, used for string validation.
	//
	// example:
	//
	// ^example$
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The numeric range of the value, used for number validation.
	Range *WafQuotaInteger `json:"Range,omitempty" xml:"Range,omitempty"`
}

func (s GetWafFilterResponseBodyFilterFieldsLogicsValidator) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsLogicsValidator) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) GetLength() *WafQuotaInteger {
	return s.Length
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) GetPattern() *string {
	return s.Pattern
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) GetRange() *WafQuotaInteger {
	return s.Range
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) SetErrMsg(v string) *GetWafFilterResponseBodyFilterFieldsLogicsValidator {
	s.ErrMsg = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) SetLength(v *WafQuotaInteger) *GetWafFilterResponseBodyFilterFieldsLogicsValidator {
	s.Length = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) SetPattern(v string) *GetWafFilterResponseBodyFilterFieldsLogicsValidator {
	s.Pattern = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) SetRange(v *WafQuotaInteger) *GetWafFilterResponseBodyFilterFieldsLogicsValidator {
	s.Range = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsLogicsValidator) Validate() error {
	if s.Length != nil {
		if err := s.Length.Validate(); err != nil {
			return err
		}
	}
	if s.Range != nil {
		if err := s.Range.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWafFilterResponseBodyFilterFieldsSelector struct {
	// The list of available data when the selector kind is data.
	Data []*GetWafFilterResponseBodyFilterFieldsSelectorData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The kind of the selector, such as whether it is used for selecting data items or other purposes.
	//
	// example:
	//
	// data
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
}

func (s GetWafFilterResponseBodyFilterFieldsSelector) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsSelector) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsSelector) GetData() []*GetWafFilterResponseBodyFilterFieldsSelectorData {
	return s.Data
}

func (s *GetWafFilterResponseBodyFilterFieldsSelector) GetKind() *string {
	return s.Kind
}

func (s *GetWafFilterResponseBodyFilterFieldsSelector) SetData(v []*GetWafFilterResponseBodyFilterFieldsSelectorData) *GetWafFilterResponseBodyFilterFieldsSelector {
	s.Data = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSelector) SetKind(v string) *GetWafFilterResponseBodyFilterFieldsSelector {
	s.Kind = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSelector) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWafFilterResponseBodyFilterFieldsSelectorData struct {
	// The display label of the available data.
	//
	// example:
	//
	// China
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The parameter value of the available data.
	//
	// example:
	//
	// CN
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetWafFilterResponseBodyFilterFieldsSelectorData) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsSelectorData) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsSelectorData) GetLabel() *string {
	return s.Label
}

func (s *GetWafFilterResponseBodyFilterFieldsSelectorData) GetValue() *string {
	return s.Value
}

func (s *GetWafFilterResponseBodyFilterFieldsSelectorData) SetLabel(v string) *GetWafFilterResponseBodyFilterFieldsSelectorData {
	s.Label = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSelectorData) SetValue(v string) *GetWafFilterResponseBodyFilterFieldsSelectorData {
	s.Value = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSelectorData) Validate() error {
	return dara.Validate(s)
}

type GetWafFilterResponseBodyFilterFieldsSubs struct {
	// Indicates whether the current plan supports this match object.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The parameter of the sub-item match object.
	//
	// example:
	//
	// ali.websdk.umid
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The display label of the sub-item match object.
	//
	// example:
	//
	// Web UMID
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The list of logical operator properties applicable to the sub-item (same structure as the parent Logics).
	Logics []*GetWafFilterResponseBodyFilterFieldsSubsLogics `json:"Logics,omitempty" xml:"Logics,omitempty" type:"Repeated"`
	// The minimum plan that supports this match object, displayed when the current plan does not support it.
	//
	// example:
	//
	// high
	MinPlan *string `json:"MinPlan,omitempty" xml:"MinPlan,omitempty"`
}

func (s GetWafFilterResponseBodyFilterFieldsSubs) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsSubs) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) GetEnable() *bool {
	return s.Enable
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) GetKey() *string {
	return s.Key
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) GetLabel() *string {
	return s.Label
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) GetLogics() []*GetWafFilterResponseBodyFilterFieldsSubsLogics {
	return s.Logics
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) GetMinPlan() *string {
	return s.MinPlan
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) SetEnable(v bool) *GetWafFilterResponseBodyFilterFieldsSubs {
	s.Enable = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) SetKey(v string) *GetWafFilterResponseBodyFilterFieldsSubs {
	s.Key = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) SetLabel(v string) *GetWafFilterResponseBodyFilterFieldsSubs {
	s.Label = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) SetLogics(v []*GetWafFilterResponseBodyFilterFieldsSubsLogics) *GetWafFilterResponseBodyFilterFieldsSubs {
	s.Logics = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) SetMinPlan(v string) *GetWafFilterResponseBodyFilterFieldsSubs {
	s.MinPlan = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubs) Validate() error {
	if s.Logics != nil {
		for _, item := range s.Logics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWafFilterResponseBodyFilterFieldsSubsLogics struct {
	// The field attributes.
	//
	// example:
	//
	// 0
	Attributes *int32 `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// Indicates whether the current plan supports this match operator.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the value input field. Valid values:
	//
	// 	- select:single: single-select input field
	//
	// 	- select:multi: multi-select input field
	//
	// 	- input:single: single input field
	//
	// 	- input:multi: multi input field
	//
	// example:
	//
	// select:single
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The minimum plan that supports this match operator, displayed when the current plan does not support it.
	//
	// example:
	//
	// high
	MinPlan *string `json:"MinPlan,omitempty" xml:"MinPlan,omitempty"`
	// Indicates whether the match result is negated.
	//
	// example:
	//
	// false
	Negative *bool `json:"Negative,omitempty" xml:"Negative,omitempty"`
	// The display label of the match operator.
	//
	// example:
	//
	// Equal
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The parameter of the match operator used internally by the system.
	//
	// example:
	//
	// eq
	Symbol *string `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	// The input hint that helps users provide valid values required by the rule.
	//
	// example:
	//
	// e.g. image/jpeg
	Tip *string `json:"Tip,omitempty" xml:"Tip,omitempty"`
	// The type of the value. Valid values:
	//
	// 	- integer: integer
	//
	// 	- integer_slice: integer array
	//
	// 	- string: string
	//
	// 	- string_slice: string array
	//
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The validator object that defines the validation rules for values.
	Validator *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator `json:"Validator,omitempty" xml:"Validator,omitempty" type:"Struct"`
}

func (s GetWafFilterResponseBodyFilterFieldsSubsLogics) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsSubsLogics) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) GetAttributes() *int32 {
	return s.Attributes
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) GetEnable() *bool {
	return s.Enable
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) GetKind() *string {
	return s.Kind
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) GetMinPlan() *string {
	return s.MinPlan
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) GetNegative() *bool {
	return s.Negative
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) GetOperator() *string {
	return s.Operator
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) GetSymbol() *string {
	return s.Symbol
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) GetTip() *string {
	return s.Tip
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) GetType() *string {
	return s.Type
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) GetValidator() *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator {
	return s.Validator
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) SetAttributes(v int32) *GetWafFilterResponseBodyFilterFieldsSubsLogics {
	s.Attributes = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) SetEnable(v bool) *GetWafFilterResponseBodyFilterFieldsSubsLogics {
	s.Enable = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) SetKind(v string) *GetWafFilterResponseBodyFilterFieldsSubsLogics {
	s.Kind = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) SetMinPlan(v string) *GetWafFilterResponseBodyFilterFieldsSubsLogics {
	s.MinPlan = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) SetNegative(v bool) *GetWafFilterResponseBodyFilterFieldsSubsLogics {
	s.Negative = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) SetOperator(v string) *GetWafFilterResponseBodyFilterFieldsSubsLogics {
	s.Operator = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) SetSymbol(v string) *GetWafFilterResponseBodyFilterFieldsSubsLogics {
	s.Symbol = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) SetTip(v string) *GetWafFilterResponseBodyFilterFieldsSubsLogics {
	s.Tip = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) SetType(v string) *GetWafFilterResponseBodyFilterFieldsSubsLogics {
	s.Type = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) SetValidator(v *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) *GetWafFilterResponseBodyFilterFieldsSubsLogics {
	s.Validator = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogics) Validate() error {
	if s.Validator != nil {
		if err := s.Validator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator struct {
	// The error message returned when validation fails.
	//
	// example:
	//
	// Enter a valid expression
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The length limit of the value.
	Length *WafQuotaInteger `json:"Length,omitempty" xml:"Length,omitempty"`
	// The regular expression pattern for the value, used for string validation.
	//
	// example:
	//
	// ^example$
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The numeric range of the value, used for number validation.
	Range *WafQuotaInteger `json:"Range,omitempty" xml:"Range,omitempty"`
}

func (s GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) String() string {
	return dara.Prettify(s)
}

func (s GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) GoString() string {
	return s.String()
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) GetLength() *WafQuotaInteger {
	return s.Length
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) GetPattern() *string {
	return s.Pattern
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) GetRange() *WafQuotaInteger {
	return s.Range
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) SetErrMsg(v string) *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator {
	s.ErrMsg = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) SetLength(v *WafQuotaInteger) *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator {
	s.Length = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) SetPattern(v string) *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator {
	s.Pattern = &v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) SetRange(v *WafQuotaInteger) *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator {
	s.Range = v
	return s
}

func (s *GetWafFilterResponseBodyFilterFieldsSubsLogicsValidator) Validate() error {
	if s.Length != nil {
		if err := s.Length.Validate(); err != nil {
			return err
		}
	}
	if s.Range != nil {
		if err := s.Range.Validate(); err != nil {
			return err
		}
	}
	return nil
}
