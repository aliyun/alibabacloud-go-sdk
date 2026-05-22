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
	Filter *GetWafFilterResponseBodyFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
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
	Fields []*GetWafFilterResponseBodyFilterFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// http_bot
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// characteristics
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
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
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// http.request.headers
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// Header
	Label    *string                                       `json:"Label,omitempty" xml:"Label,omitempty"`
	Logics   []*GetWafFilterResponseBodyFilterFieldsLogics `json:"Logics,omitempty" xml:"Logics,omitempty" type:"Repeated"`
	MinPlan  *string                                       `json:"MinPlan,omitempty" xml:"MinPlan,omitempty"`
	Selector *GetWafFilterResponseBodyFilterFieldsSelector `json:"Selector,omitempty" xml:"Selector,omitempty" type:"Struct"`
	Sub      *bool                                         `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// example:
	//
	// e.g. Content-Type
	SubTip *string `json:"SubTip,omitempty" xml:"SubTip,omitempty"`
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
	return nil
}

type GetWafFilterResponseBodyFilterFieldsLogics struct {
	// example:
	//
	// 1
	Attributes *int32 `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	Enable     *bool  `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// input:single
	Kind     *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	MinPlan  *string `json:"MinPlan,omitempty" xml:"MinPlan,omitempty"`
	Negative *bool   `json:"Negative,omitempty" xml:"Negative,omitempty"`
	// example:
	//
	// Does not equal
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// eq
	Symbol *string `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	// example:
	//
	// e.g. image/jpeg
	Tip *string `json:"Tip,omitempty" xml:"Tip,omitempty"`
	// example:
	//
	// string
	Type      *string                                              `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// example:
	//
	// Enter a valid expression
	ErrMsg *string          `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	Length *WafQuotaInteger `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// ^example$
	Pattern *string          `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	Range   *WafQuotaInteger `json:"Range,omitempty" xml:"Range,omitempty"`
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
	Data []*GetWafFilterResponseBodyFilterFieldsSelectorData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// example:
	//
	// China
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
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
