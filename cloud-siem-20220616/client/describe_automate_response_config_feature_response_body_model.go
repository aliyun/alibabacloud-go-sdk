// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutomateResponseConfigFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAutomateResponseConfigFeatureResponseBody
	GetCode() *int32
	SetData(v []*DescribeAutomateResponseConfigFeatureResponseBodyData) *DescribeAutomateResponseConfigFeatureResponseBody
	GetData() []*DescribeAutomateResponseConfigFeatureResponseBodyData
	SetMessage(v string) *DescribeAutomateResponseConfigFeatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAutomateResponseConfigFeatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAutomateResponseConfigFeatureResponseBody
	GetSuccess() *bool
}

type DescribeAutomateResponseConfigFeatureResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAutomateResponseConfigFeatureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) GetData() []*DescribeAutomateResponseConfigFeatureResponseBodyData {
	return s.Data
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetCode(v int32) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetData(v []*DescribeAutomateResponseConfigFeatureResponseBodyData) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetMessage(v string) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetRequestId(v string) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) SetSuccess(v bool) *DescribeAutomateResponseConfigFeatureResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBody) Validate() error {
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

type DescribeAutomateResponseConfigFeatureResponseBodyData struct {
	// The data type of the condition field in the automated response rule.
	//
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The name of the condition field in the automated response rule.
	//
	// example:
	//
	// alert_desc
	Feature *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// The enumerated values of the right operand for the field.
	RightValueEnums []*DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums `json:"RightValueEnums,omitempty" xml:"RightValueEnums,omitempty" type:"Repeated"`
	// The operators that are supported for the condition field.
	SupportOperators []*DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators `json:"SupportOperators,omitempty" xml:"SupportOperators,omitempty" type:"Repeated"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) GetDataType() *string {
	return s.DataType
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) GetFeature() *string {
	return s.Feature
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) GetRightValueEnums() []*DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums {
	return s.RightValueEnums
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) GetSupportOperators() []*DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	return s.SupportOperators
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetDataType(v string) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.DataType = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetFeature(v string) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.Feature = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetRightValueEnums(v []*DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.RightValueEnums = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) SetSupportOperators(v []*DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) *DescribeAutomateResponseConfigFeatureResponseBodyData {
	s.SupportOperators = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyData) Validate() error {
	if s.RightValueEnums != nil {
		for _, item := range s.RightValueEnums {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SupportOperators != nil {
		for _, item := range s.SupportOperators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums struct {
	// The enumerated value of the right operand.
	//
	// example:
	//
	// serious
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The internal code of the enumerated value.
	//
	// example:
	//
	// aliyun.siem.automate.feature.alert_level.serious
	ValueMds *string `json:"ValueMds,omitempty" xml:"ValueMds,omitempty"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) GetValue() *string {
	return s.Value
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) GetValueMds() *string {
	return s.ValueMds
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) SetValue(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums {
	s.Value = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) SetValueMds(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums {
	s.ValueMds = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataRightValueEnums) Validate() error {
	return dara.Validate(s)
}

type DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators struct {
	// Indicates whether the right operand is required. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	HasRightValue *bool `json:"HasRightValue,omitempty" xml:"HasRightValue,omitempty"`
	// The position of the operator in the operator list.
	//
	// example:
	//
	// 3
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The operator.
	//
	// example:
	//
	// <=
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The description of the operator in Chinese.
	//
	// example:
	//
	// larger than or equal to
	OperatorDescCn *string `json:"OperatorDescCn,omitempty" xml:"OperatorDescCn,omitempty"`
	// The description of the operator in English.
	//
	// example:
	//
	// larger than or equal to
	OperatorDescEn *string `json:"OperatorDescEn,omitempty" xml:"OperatorDescEn,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// <=
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The data types that are supported by the operator. The data types are separated by commas (,).
	//
	// example:
	//
	// varchar
	SupportDataType *string `json:"SupportDataType,omitempty" xml:"SupportDataType,omitempty"`
	// The scenarios that are supported by the operator. Multiple scenarios are separated by commas (,), such as aggregation scenarios. By default, this parameter is empty.
	//
	// example:
	//
	// [AGGREGATE]
	SupportTag []*string `json:"SupportTag,omitempty" xml:"SupportTag,omitempty" type:"Repeated"`
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GetHasRightValue() *bool {
	return s.HasRightValue
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GetOperator() *string {
	return s.Operator
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GetOperatorDescCn() *string {
	return s.OperatorDescCn
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GetOperatorDescEn() *string {
	return s.OperatorDescEn
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GetOperatorName() *string {
	return s.OperatorName
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GetSupportDataType() *string {
	return s.SupportDataType
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) GetSupportTag() []*string {
	return s.SupportTag
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetHasRightValue(v bool) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.HasRightValue = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetIndex(v int32) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.Index = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperator(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.Operator = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperatorDescCn(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.OperatorDescCn = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperatorDescEn(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.OperatorDescEn = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetOperatorName(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.OperatorName = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetSupportDataType(v string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.SupportDataType = &v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) SetSupportTag(v []*string) *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators {
	s.SupportTag = v
	return s
}

func (s *DescribeAutomateResponseConfigFeatureResponseBodyDataSupportOperators) Validate() error {
	return dara.Validate(s)
}
