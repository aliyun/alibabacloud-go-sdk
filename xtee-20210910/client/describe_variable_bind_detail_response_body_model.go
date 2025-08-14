// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableBindDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVariableBindDetailResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeVariableBindDetailResponseBodyResultObject) *DescribeVariableBindDetailResponseBody
	GetResultObject() *DescribeVariableBindDetailResponseBodyResultObject
}

type DescribeVariableBindDetailResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *DescribeVariableBindDetailResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeVariableBindDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableBindDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVariableBindDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVariableBindDetailResponseBody) GetResultObject() *DescribeVariableBindDetailResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVariableBindDetailResponseBody) SetRequestId(v string) *DescribeVariableBindDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBody) SetResultObject(v *DescribeVariableBindDetailResponseBodyResultObject) *DescribeVariableBindDetailResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVariableBindDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVariableBindDetailResponseBodyResultObject struct {
	// Whether modification is allowed, default is false
	//
	// example:
	//
	// false
	AllowModify *bool `json:"allowModify,omitempty" xml:"allowModify,omitempty"`
	// Variable definition ID
	//
	// example:
	//
	// 10
	DefineId *int64 `json:"defineId,omitempty" xml:"defineId,omitempty"`
	// Variable definition title
	//
	// example:
	//
	// IP所在地_城市Code
	DefineTitle *string `json:"defineTitle,omitempty" xml:"defineTitle,omitempty"`
	// Variable description information
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Variable ID.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Bound parameters.
	Params []*DescribeVariableBindDetailResponseBodyResultObjectParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
	// List of associated policies
	RelationRules []*DescribeVariableBindDetailResponseBodyResultObjectRelationRules `json:"relationRules,omitempty" xml:"relationRules,omitempty" type:"Repeated"`
	// Title.
	//
	// example:
	//
	// 变量title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DescribeVariableBindDetailResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableBindDetailResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) GetAllowModify() *bool {
	return s.AllowModify
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) GetDefineId() *int64 {
	return s.DefineId
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) GetDefineTitle() *string {
	return s.DefineTitle
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) GetParams() []*DescribeVariableBindDetailResponseBodyResultObjectParams {
	return s.Params
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) GetRelationRules() []*DescribeVariableBindDetailResponseBodyResultObjectRelationRules {
	return s.RelationRules
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) GetTitle() *string {
	return s.Title
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) SetAllowModify(v bool) *DescribeVariableBindDetailResponseBodyResultObject {
	s.AllowModify = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) SetDefineId(v int64) *DescribeVariableBindDetailResponseBodyResultObject {
	s.DefineId = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) SetDefineTitle(v string) *DescribeVariableBindDetailResponseBodyResultObject {
	s.DefineTitle = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) SetDescription(v string) *DescribeVariableBindDetailResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) SetEventCode(v string) *DescribeVariableBindDetailResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) SetId(v int64) *DescribeVariableBindDetailResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) SetParams(v []*DescribeVariableBindDetailResponseBodyResultObjectParams) *DescribeVariableBindDetailResponseBodyResultObject {
	s.Params = v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) SetRelationRules(v []*DescribeVariableBindDetailResponseBodyResultObjectRelationRules) *DescribeVariableBindDetailResponseBodyResultObject {
	s.RelationRules = v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) SetTitle(v string) *DescribeVariableBindDetailResponseBodyResultObject {
	s.Title = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeVariableBindDetailResponseBodyResultObjectParams struct {
	// Event field name
	//
	// example:
	//
	// ip
	EventFieldName *string `json:"eventFieldName,omitempty" xml:"eventFieldName,omitempty"`
	// Whether it is required, default is false
	//
	// example:
	//
	// false
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// Bound variable name
	//
	// example:
	//
	// ip
	VariableName *string `json:"variableName,omitempty" xml:"variableName,omitempty"`
}

func (s DescribeVariableBindDetailResponseBodyResultObjectParams) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableBindDetailResponseBodyResultObjectParams) GoString() string {
	return s.String()
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectParams) GetEventFieldName() *string {
	return s.EventFieldName
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectParams) GetRequired() *bool {
	return s.Required
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectParams) GetVariableName() *string {
	return s.VariableName
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectParams) SetEventFieldName(v string) *DescribeVariableBindDetailResponseBodyResultObjectParams {
	s.EventFieldName = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectParams) SetRequired(v bool) *DescribeVariableBindDetailResponseBodyResultObjectParams {
	s.Required = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectParams) SetVariableName(v string) *DescribeVariableBindDetailResponseBodyResultObjectParams {
	s.VariableName = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectParams) Validate() error {
	return dara.Validate(s)
}

type DescribeVariableBindDetailResponseBodyResultObjectRelationRules struct {
	// Policy rule ID
	//
	// example:
	//
	// 104566
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Policy name
	//
	// example:
	//
	// 营销风险识别
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeVariableBindDetailResponseBodyResultObjectRelationRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableBindDetailResponseBodyResultObjectRelationRules) GoString() string {
	return s.String()
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectRelationRules) GetKey() *string {
	return s.Key
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectRelationRules) GetValue() *string {
	return s.Value
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectRelationRules) SetKey(v string) *DescribeVariableBindDetailResponseBodyResultObjectRelationRules {
	s.Key = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectRelationRules) SetValue(v string) *DescribeVariableBindDetailResponseBodyResultObjectRelationRules {
	s.Value = &v
	return s
}

func (s *DescribeVariableBindDetailResponseBodyResultObjectRelationRules) Validate() error {
	return dara.Validate(s)
}
