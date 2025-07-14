// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartqQueryAbilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SmartqQueryAbilityResponseBody
	GetRequestId() *string
	SetResult(v *SmartqQueryAbilityResponseBodyResult) *SmartqQueryAbilityResponseBody
	GetResult() *SmartqQueryAbilityResponseBodyResult
	SetSuccess(v bool) *SmartqQueryAbilityResponseBody
	GetSuccess() *bool
}

type SmartqQueryAbilityResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// D787E1A3-A************2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	Result *SmartqQueryAbilityResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Whether the operation was successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SmartqQueryAbilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SmartqQueryAbilityResponseBody) GoString() string {
	return s.String()
}

func (s *SmartqQueryAbilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SmartqQueryAbilityResponseBody) GetResult() *SmartqQueryAbilityResponseBodyResult {
	return s.Result
}

func (s *SmartqQueryAbilityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SmartqQueryAbilityResponseBody) SetRequestId(v string) *SmartqQueryAbilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmartqQueryAbilityResponseBody) SetResult(v *SmartqQueryAbilityResponseBodyResult) *SmartqQueryAbilityResponseBody {
	s.Result = v
	return s
}

func (s *SmartqQueryAbilityResponseBody) SetSuccess(v bool) *SmartqQueryAbilityResponseBody {
	s.Success = &v
	return s
}

func (s *SmartqQueryAbilityResponseBody) Validate() error {
	return dara.Validate(s)
}

type SmartqQueryAbilityResponseBodyResult struct {
	// Suggested chart type.
	//
	// example:
	//
	// NEW_TABLE
	ChartType *string `json:"ChartType,omitempty" xml:"ChartType,omitempty"`
	// Summary information.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Schedule
	ConclusionText *string `json:"ConclusionText,omitempty" xml:"ConclusionText,omitempty"`
	// Visualized logical SQL.
	//
	// example:
	//
	// select 	- ****
	LogicSql *string `json:"LogicSql,omitempty" xml:"LogicSql,omitempty"`
	// List of column tuple types.
	MetaType []*SmartqQueryAbilityResponseBodyResultMetaType `json:"MetaType,omitempty" xml:"MetaType,omitempty" type:"Repeated"`
	// Array of data value lists.
	Values []*SmartqQueryAbilityResponseBodyResultValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s SmartqQueryAbilityResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s SmartqQueryAbilityResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SmartqQueryAbilityResponseBodyResult) GetChartType() *string {
	return s.ChartType
}

func (s *SmartqQueryAbilityResponseBodyResult) GetConclusionText() *string {
	return s.ConclusionText
}

func (s *SmartqQueryAbilityResponseBodyResult) GetLogicSql() *string {
	return s.LogicSql
}

func (s *SmartqQueryAbilityResponseBodyResult) GetMetaType() []*SmartqQueryAbilityResponseBodyResultMetaType {
	return s.MetaType
}

func (s *SmartqQueryAbilityResponseBodyResult) GetValues() []*SmartqQueryAbilityResponseBodyResultValues {
	return s.Values
}

func (s *SmartqQueryAbilityResponseBodyResult) SetChartType(v string) *SmartqQueryAbilityResponseBodyResult {
	s.ChartType = &v
	return s
}

func (s *SmartqQueryAbilityResponseBodyResult) SetConclusionText(v string) *SmartqQueryAbilityResponseBodyResult {
	s.ConclusionText = &v
	return s
}

func (s *SmartqQueryAbilityResponseBodyResult) SetLogicSql(v string) *SmartqQueryAbilityResponseBodyResult {
	s.LogicSql = &v
	return s
}

func (s *SmartqQueryAbilityResponseBodyResult) SetMetaType(v []*SmartqQueryAbilityResponseBodyResultMetaType) *SmartqQueryAbilityResponseBodyResult {
	s.MetaType = v
	return s
}

func (s *SmartqQueryAbilityResponseBodyResult) SetValues(v []*SmartqQueryAbilityResponseBodyResultValues) *SmartqQueryAbilityResponseBodyResult {
	s.Values = v
	return s
}

func (s *SmartqQueryAbilityResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type SmartqQueryAbilityResponseBodyResultMetaType struct {
	// Column tuple name.
	//
	// example:
	//
	// Polar***STPS
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// if can be null:
	// true
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Column tuple type.
	//
	// example:
	//
	// string
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SmartqQueryAbilityResponseBodyResultMetaType) String() string {
	return dara.Prettify(s)
}

func (s SmartqQueryAbilityResponseBodyResultMetaType) GoString() string {
	return s.String()
}

func (s *SmartqQueryAbilityResponseBodyResultMetaType) GetKey() *string {
	return s.Key
}

func (s *SmartqQueryAbilityResponseBodyResultMetaType) GetType() *string {
	return s.Type
}

func (s *SmartqQueryAbilityResponseBodyResultMetaType) GetValue() *string {
	return s.Value
}

func (s *SmartqQueryAbilityResponseBodyResultMetaType) SetKey(v string) *SmartqQueryAbilityResponseBodyResultMetaType {
	s.Key = &v
	return s
}

func (s *SmartqQueryAbilityResponseBodyResultMetaType) SetType(v string) *SmartqQueryAbilityResponseBodyResultMetaType {
	s.Type = &v
	return s
}

func (s *SmartqQueryAbilityResponseBodyResultMetaType) SetValue(v string) *SmartqQueryAbilityResponseBodyResultMetaType {
	s.Value = &v
	return s
}

func (s *SmartqQueryAbilityResponseBodyResultMetaType) Validate() error {
	return dara.Validate(s)
}

type SmartqQueryAbilityResponseBodyResultValues struct {
	// Data values for each row.
	//
	// if can be null:
	// true
	Row []*string `json:"Row,omitempty" xml:"Row,omitempty" type:"Repeated"`
}

func (s SmartqQueryAbilityResponseBodyResultValues) String() string {
	return dara.Prettify(s)
}

func (s SmartqQueryAbilityResponseBodyResultValues) GoString() string {
	return s.String()
}

func (s *SmartqQueryAbilityResponseBodyResultValues) GetRow() []*string {
	return s.Row
}

func (s *SmartqQueryAbilityResponseBodyResultValues) SetRow(v []*string) *SmartqQueryAbilityResponseBodyResultValues {
	s.Row = v
	return s
}

func (s *SmartqQueryAbilityResponseBodyResultValues) Validate() error {
	return dara.Validate(s)
}
