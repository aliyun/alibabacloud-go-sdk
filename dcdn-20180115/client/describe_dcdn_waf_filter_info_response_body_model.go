// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafFilterInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*DescribeDcdnWafFilterInfoResponseBodyContent) *DescribeDcdnWafFilterInfoResponseBody
	GetContent() []*DescribeDcdnWafFilterInfoResponseBodyContent
	SetRequestId(v string) *DescribeDcdnWafFilterInfoResponseBody
	GetRequestId() *string
}

type DescribeDcdnWafFilterInfoResponseBody struct {
	// The returned information.
	Content []*DescribeDcdnWafFilterInfoResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 153ca2cd-3c01-44be-204c-64dbc6c88630
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnWafFilterInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafFilterInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafFilterInfoResponseBody) GetContent() []*DescribeDcdnWafFilterInfoResponseBodyContent {
	return s.Content
}

func (s *DescribeDcdnWafFilterInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafFilterInfoResponseBody) SetContent(v []*DescribeDcdnWafFilterInfoResponseBodyContent) *DescribeDcdnWafFilterInfoResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBody) SetRequestId(v string) *DescribeDcdnWafFilterInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafFilterInfoResponseBodyContent struct {
	// The type of the protection policy. The value of this parameter is the same as that of the DefenseScenes parameter in the request.
	//
	// example:
	//
	// custom_acl
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The information about the match condition.
	Fields []*DescribeDcdnWafFilterInfoResponseBodyContentFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
}

func (s DescribeDcdnWafFilterInfoResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafFilterInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContent) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContent) GetFields() []*DescribeDcdnWafFilterInfoResponseBodyContentFields {
	return s.Fields
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContent) SetDefenseScene(v string) *DescribeDcdnWafFilterInfoResponseBodyContent {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContent) SetFields(v []*DescribeDcdnWafFilterInfoResponseBodyContentFields) *DescribeDcdnWafFilterInfoResponseBodyContent {
	s.Fields = v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafFilterInfoResponseBodyContentFields struct {
	// The description of the match field. This parameter is not returned or is empty if no match fields are found.
	//
	// example:
	//
	// Custom Header
	ExtendField *string `json:"ExtendField,omitempty" xml:"ExtendField,omitempty"`
	// The information about the logical symbol.
	LogicalSymbol []*DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol `json:"LogicalSymbol,omitempty" xml:"LogicalSymbol,omitempty" type:"Repeated"`
	// The match field.
	//
	// example:
	//
	// Header
	MatchField *string `json:"MatchField,omitempty" xml:"MatchField,omitempty"`
}

func (s DescribeDcdnWafFilterInfoResponseBodyContentFields) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafFilterInfoResponseBodyContentFields) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFields) GetExtendField() *string {
	return s.ExtendField
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFields) GetLogicalSymbol() []*DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol {
	return s.LogicalSymbol
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFields) GetMatchField() *string {
	return s.MatchField
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFields) SetExtendField(v string) *DescribeDcdnWafFilterInfoResponseBodyContentFields {
	s.ExtendField = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFields) SetLogicalSymbol(v []*DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) *DescribeDcdnWafFilterInfoResponseBodyContentFields {
	s.LogicalSymbol = v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFields) SetMatchField(v string) *DescribeDcdnWafFilterInfoResponseBodyContentFields {
	s.MatchField = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFields) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol struct {
	// The configurable attributes, which are bit-field variables that are shown in the following list.\\
	//
	// For example, 1(00000001) indicates that case sensitivity can be enabled and stream match cannot be enabled, and 3(00000011) indicates that case sensitivity and stream match can be enabled.
	//
	// 	- Bit (low to high) - Description
	//
	// 	- 1 - Case sensitivity
	//
	// 	- 2 - Stream match
	//
	// example:
	//
	// 1
	Attributes *int32 `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The logical symbol that is displayed in the Dynamic Content Delivery Network (DCDN) console.
	//
	// example:
	//
	// Equal to one of multiple values.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of match items that can be returned. The value of this parameter varies based on the value of the Type parameter. Valid values:
	//
	// 	- If **multi*	- is returned for the Type parameter, the value of this parameter indicates the maximum number of match items.
	//
	// 	- If **single*	- is returned for the Type parameter, the value of this parameter is 1.
	//
	// 	- If **none*	- is returned for the Type parameter, the value of this parameter is 0.
	//
	// example:
	//
	// 50
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// The information about the regular expression.
	Regexp *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp `json:"Regexp,omitempty" xml:"Regexp,omitempty" type:"Struct"`
	// The logical symbol that is passed to the backend.
	//
	// example:
	//
	// match-one
	Symbol *string `json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	// The tips that are displayed in the match item.
	//
	// example:
	//
	// You can enter up to 50 tips. Press the Enter key.
	Tip *string `json:"Tip,omitempty" xml:"Tip,omitempty"`
	// The number of match items. Valid values:
	//
	// 	- multi: You can specify multiple match items.
	//
	// 	- single: You can specify only a match item.
	//
	// 	- none: no match items.
	//
	// example:
	//
	// multi
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) GetAttributes() *int32 {
	return s.Attributes
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) GetRegexp() *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp {
	return s.Regexp
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) GetSymbol() *string {
	return s.Symbol
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) GetTip() *string {
	return s.Tip
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) SetAttributes(v int32) *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol {
	s.Attributes = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) SetDescription(v string) *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol {
	s.Description = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) SetMaxLength(v int32) *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol {
	s.MaxLength = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) SetRegexp(v *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp) *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol {
	s.Regexp = v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) SetSymbol(v string) *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol {
	s.Symbol = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) SetTip(v string) *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol {
	s.Tip = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) SetType(v string) *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol {
	s.Type = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbol) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp struct {
	// The error message returned when no items match the regular expression.
	//
	// example:
	//
	// Specify this field.
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The regular expression.
	//
	// example:
	//
	// ^\\S+$
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
}

func (s DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp) GetPattern() *string {
	return s.Pattern
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp) SetErrMsg(v string) *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp) SetPattern(v string) *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp {
	s.Pattern = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponseBodyContentFieldsLogicalSymbolRegexp) Validate() error {
	return dara.Validate(s)
}
